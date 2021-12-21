package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	contract "github.com/artworkk/balancer-factory-logs/contract"
	stablepool "github.com/artworkk/balancer-factory-logs/contract/balancerv2_stablepool"
	weightedpool "github.com/artworkk/balancer-factory-logs/contract/balancerv2_wp"
	wp2tokens "github.com/artworkk/balancer-factory-logs/contract/balancerv2_wp2tokens"
	"github.com/artworkk/balancer-factory-logs/enums"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type param struct {
	chain enums.ChainType
}

type poolData struct {
	PoolId    string          `json:"poolId"`
	LpAddress common.Address  `json:"lpAddress"`
	Chain     enums.ChainType `json:"chain"`
	Token0    common.Address  `json:"token0"`
	Token1    common.Address  `json:"token1"`
	// WeightToken0 *big.Int       `json:"weightToken0"`
	// WeightToken1 *big.Int       `json:"weightToken1"`
	// SwapFeePerc  *big.Int       `json:"swapFeePercentage"`
	// IsPaused     bool           `json:"isPaused"`
}

var (
	pools = make(chan *poolData)
)

func main() {
	ctx := context.Background()
	params := []param{
		{chain: enums.Ethereum},
		{chain: enums.Polygon},
	}
	var wg sync.WaitGroup
	for _, _param := range params {
		go func(c context.Context, p param) {
			defer func() {
				wg.Done()
			}()
			wg.Add(1)
			getCreatedPool(c, p)
		}(ctx, _param)
	}
	go func() {
		for newPool := range pools {
			_json, _ := json.Marshal(newPool)
			fmt.Println(string(_json))
		}
	}()
	time.Sleep(5 * time.Millisecond)
	wg.Wait()
}

func getCreatedPool(ctx context.Context, param param) {
	chain := param.chain
	nodeAddr := enums.DexNode[chain]
	client, err := ethclient.Dial(string(nodeAddr))
	if err != nil {
		log.Fatalf("[%s]\tfailed to init client", chain)
	}
	log.Printf("[%s]\tInitialized client", chain)
	currentBlockUint, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("[%s]\tfailed to get current block number", chain)
	}
	currentBlock := int64(currentBlockUint)

	// BalancerV2 WeightedPool2Tokens
	wp2tokensFactoryAddress := enums.Wp2TokensFactoryAddr[chain]
	wp2tokensFactory, err := abi.JSON(strings.NewReader(wp2tokens.BalancerV2Wp2TokensFactoryABI))
	if err != nil {
		log.Fatalf("[%s]\tfailed to parse BalancerV2Wp2TokensFactoryABI", chain)
	}
	// BalancerV2 StablePool
	stablepoolFactoryAddress := enums.StablePoolFactoryAddr[chain]
	stablepoolFactory, err := abi.JSON(strings.NewReader(stablepool.BalancerV2StablePoolFactoryABI))
	if err != nil {
		log.Fatalf("[%s]\tfailed to parse BalancerV2StablePoolFactoryABI", chain)
	}
	wpFactoryAddress := enums.WPFactoryAddr[chain]
	wpFactory, err := abi.JSON(strings.NewReader(weightedpool.WpFactoryABI))
	if err != nil {
		log.Fatalf("[%s]\tfailed to parse BalancerV2WpFactoryABI", "kuy")
	}
	wp2tokensPoolCreated := wp2tokensFactory.Events["PoolCreated"].ID
	stablepoolPoolCreated := stablepoolFactory.Events["PoolCreated"].ID
	wpPoolcreated := wpFactory.Events["PoolCreated"].ID

	var fromBlock int64
	wp2tokensGenesis := enums.Wp2TokensGenesis[chain]
	stablepoolGenesis := enums.StablePoolGenesis[chain]
	wpGenesis := enums.WPGenesis[chain]
	if stablepoolGenesis < wp2tokensGenesis {
		fromBlock = stablepoolGenesis
	} else {
		fromBlock = wp2tokensGenesis
	}
	if fromBlock > wpGenesis {
		fromBlock = wpGenesis
	}

	var wg sync.WaitGroup
	step := int64(1024)
	guard := make(chan struct{}, step)

	log.Printf("[%s]\tFiltering logs..", chain)
	for from := fromBlock; from < currentBlock; from += step {
		to := from + step + 1
		if to > currentBlock {
			to = currentBlock
		}
		logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(from),
			ToBlock:   big.NewInt(to),
			Addresses: []common.Address{
				// wp2tokensFactoryAddress,
				// stablepoolFactoryAddress,
				wpFactoryAddress,
			},
			Topics: [][]common.Hash{{
				// wp2tokensPoolCreated,
				// stablepoolPoolCreated,
				wpPoolcreated,
			}},
		})
		if err != nil {
			log.Fatalf("[%s]\tfailed to filter logs\n%s", chain, err.Error())
		}
		for _, _log := range logs {
			switch _log.Address {
			case
				wp2tokensFactoryAddress,
				stablepoolFactoryAddress,
				wpFactoryAddress:
			default:
				continue
			}
			guard <- struct{}{}
			wg.Add(1)
			go func(l types.Log) {
				defer func() {
					<-guard
					wg.Done()
				}()
				switch l.Topics[0] {
				case
					wp2tokensPoolCreated,
					stablepoolPoolCreated,
					wpPoolcreated:
					switch l.Address {
					case wpFactoryAddress:
						lpAddress := common.HexToAddress(l.Topics[1].String())
						log.Printf("[%s]\tFound WeightedPool PoolCreated on block %d at address %s\n",
							chain, l.BlockNumber, lpAddress)
						balancerv2wp, err := weightedpool.NewWp(lpAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerv2wp for pool %v\n%v\n", chain, lpAddress, err.Error())
						}
						callOpts := &bind.CallOpts{
							BlockNumber: big.NewInt(currentBlock),
							Context:     ctx,
						}
						poolId, err := balancerv2wp.GetPoolId(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool ID from LP address %v weights\n%v\n", chain, lpAddress, err.Error())
						}
						vaultAddress, err := balancerv2wp.GetVault(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get vault address: %v\n", chain, err.Error())
						}
						balancerV2TheVault, err := contract.NewBalancerv2thevault(vaultAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerV2TheVault: %v\n", chain, err.Error())
						}
						poolTokens, err := balancerV2TheVault.GetPoolTokens(callOpts, poolId)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool tokens for pool %v\n%v\n", chain, fmt.Sprintf("0x%x", poolId), err.Error())
						}
						pools <- &poolData{
							PoolId:    fmt.Sprintf("0x%x", poolId),
							LpAddress: lpAddress,
							Chain:     chain,
							Token0:    poolTokens.Tokens[0],
							Token1:    poolTokens.Tokens[1],
						}
					case stablepoolFactoryAddress:
						lpAddress := common.HexToAddress(l.Topics[1].String())
						log.Printf("[%s]\tFound StablePool PoolCreated on block %d at address %s\n",
							chain, l.BlockNumber, lpAddress)
						balancerv2stablepool, err := stablepool.NewBalancerV2StablePool(lpAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerv2stablepool for pool %v\n%v\n", chain, lpAddress, err.Error())
						}
						callOpts := &bind.CallOpts{
							BlockNumber: big.NewInt(currentBlock),
							Context:     ctx,
						}
						poolId, err := balancerv2stablepool.GetPoolId(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool ID from LP address %v weights\n%v\n", chain, lpAddress, err.Error())
						}
						vaultAddress, err := balancerv2stablepool.GetVault(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get vault address: %v\n", chain, err.Error())
						}
						balancerV2TheVault, err := contract.NewBalancerv2thevault(vaultAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerV2TheVault: %v\n", chain, err.Error())
						}
						poolTokens, err := balancerV2TheVault.GetPoolTokens(callOpts, poolId)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool tokens for pool %v\n%v\n", chain, fmt.Sprintf("0x%x", poolId), err.Error())
						}
						pools <- &poolData{
							PoolId:    fmt.Sprintf("0x%x", poolId),
							LpAddress: lpAddress,
							Chain:     chain,
							Token0:    poolTokens.Tokens[0],
							Token1:    poolTokens.Tokens[1],
						}
					case wp2tokensFactoryAddress:
						lpAddress := common.HexToAddress(l.Topics[1].String())
						log.Printf("[%s]\tFound Wp2Tokens PoolCreated on block %d at address %s\n",
							chain, l.BlockNumber, lpAddress)
						balancerv2wp2tokens, err := wp2tokens.NewBalancerV2Wp2Tokens(lpAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerv2wp2tokens for pool %v\n%v\n", chain, lpAddress, err.Error())
						}
						callOpts := &bind.CallOpts{
							BlockNumber: big.NewInt(currentBlock),
							Context:     ctx,
						}
						poolId, err := balancerv2wp2tokens.GetPoolId(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool ID from LP address %v weights\n%v\n", chain, lpAddress, err.Error())
						}
						vaultAddress, err := balancerv2wp2tokens.GetVault(callOpts)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get vault address: %v\n", chain, err.Error())
						}
						balancerV2TheVault, err := contract.NewBalancerv2thevault(vaultAddress, client)
						if err != nil {
							log.Fatalf("[%s]\tfailed to create balancerV2TheVault: %v\n", chain, err.Error())
						}
						poolTokens, err := balancerV2TheVault.GetPoolTokens(callOpts, poolId)
						if err != nil {
							log.Fatalf("[%s]\tfailed to get pool tokens for pool %v\n%v\n", chain, fmt.Sprintf("0x%x", poolId), err.Error())
						}
						// weights, err := balancerv2wp2tokens.GetNormalizedWeights(callOpts)
						// if err != nil {
						// 	log.Fatalf("[%s]\tfailed to get pool %v weights\n", chain, lpAddress)
						// }
						// swapFeePerc, err := balancerv2wp2tokens.GetSwapFeePercentage(callOpts)
						// if err != nil {
						// 	log.Fatalf("[%s]\tfailed to get pool %v swap fee percentage\n", chain, lpAddress)
						// }
						// paused, err := balancerV2TheVault.GetPausedState(callOpts)
						// if err != nil {
						// 	log.Fatalf("[%s]\tFailed to get the vault paused state", chain)
						// }
						// isPaused := paused.Paused
						pools <- &poolData{
							PoolId:    fmt.Sprintf("0x%x", poolId),
							LpAddress: lpAddress,
							Chain:     chain,
							Token0:    poolTokens.Tokens[0],
							Token1:    poolTokens.Tokens[1],
							// WeightToken0: weights[0],
							// WeightToken1: weights[1],
							// SwapFeePerc:  swapFeePerc,
							// IsPaused:     isPaused,
						}
					}
				}
			}(_log)
			wg.Wait()
		}
	}
}
