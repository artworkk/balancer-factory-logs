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

	"github.com/artworkk/balancer-factory-logs/contract"
	"github.com/artworkk/balancer-factory-logs/enums"
	"github.com/artworkk/balancer-factory-logs/poolinterface"
	libtypes "github.com/artworkk/balancer-factory-logs/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
)

type param struct {
	chain enums.ChainType
}

type poolData struct {
	PoolId        string          `json:"poolId"`
	LpAddress     common.Address  `json:"lpAddress"`
	Chain         enums.ChainType `json:"chain"`
	Token0        common.Address  `json:"token0"`
	Token1        common.Address  `json:"token1"`
	BalanceToken0 *big.Int        `json:"balanceToken0"`
	BalanceToken1 *big.Int        `json:"balanceToken1"`
	WeightToken0  *big.Int        `json:"weightToken0"`
	WeightToken1  *big.Int        `json:"weightToken1"`
	SwapFeePerc   *big.Int        `json:"swapFeePercentage"`
	IsPaused      bool            `json:"isPaused"`
}

var (
	pools = make(chan *poolinterface.PoolBalancerV2)
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

	// Balancer factory genesis
	fromBlock := enums.DexFactGenesis[chain]
	// Balancer factory addr
	factoryAddress := enums.DexFactAddr[chain]
	// Balancer factory ABI
	balancerFactoryAbi, err := abi.JSON(strings.NewReader(contract.Balancerv2wp2tokenfactoryABI))
	if err != nil {
		log.Fatalf("[%s]\tfailed to parse balancer abi", chain)
	}
	eventPoolCreated := balancerFactoryAbi.Events["PoolCreated"].ID

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
			Addresses: []common.Address{factoryAddress},
			Topics: [][]common.Hash{{
				eventPoolCreated,
			}},
		})
		if err != nil {
			log.Fatalf("[%s]\tfailed to filter logs\n%s", chain, err.Error())
		}
		for _, _log := range logs {
			if _log.Address != factoryAddress {
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
				case eventPoolCreated:
					lpAddress := common.HexToAddress(l.Topics[1].String())
					vaultAddress := common.HexToAddress("0xba12222222228d8ba445958a75a0704d566bf2c8")
					log.Printf("[%s]\tFound PoolCreated on block %d at address %s\n",
						chain, l.BlockNumber, lpAddress)

					callOpts := &bind.CallOpts{
						BlockNumber: big.NewInt(currentBlock),
						Context:     ctx,
					}
					balancerV2Wp2Token, err := contract.NewBalancerv2wp2token(lpAddress, client)
					if err != nil {
						log.Fatalf("[%s]\tfailed to create balancerV2Wp2Token for pool %v\n%v\n", chain, lpAddress, err.Error())
					}
					poolId, err := balancerV2Wp2Token.GetPoolId(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool ID from LP address %v weights\n%v\n", chain, lpAddress, err.Error())
					}
					balancerV2TheVault, err := contract.NewBalancerv2thevault(vaultAddress, client)
					if err != nil {
						log.Fatalf("[%s]\tfailed to create balancerV2TheVault for pool %v\n%s\n", chain, lpAddress, err.Error())
					}
					poolTokens, err := balancerV2TheVault.GetPoolTokens(callOpts, poolId)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool tokens for pool %v\n%v\n", chain, fmt.Sprintf("0x%x", poolId), err.Error())
					}
					weights, err := balancerV2Wp2Token.GetNormalizedWeights(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool %v weights\n", chain, lpAddress)
					}
					swapFeePerc, err := balancerV2Wp2Token.GetSwapFeePercentage(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool %v swap fee percentage\n", chain, lpAddress)
					}
					header, err := client.HeaderByNumber(ctx, big.NewInt(int64(currentBlock)))
					if err != nil {
						log.Fatalf("[%s]\tfailed to get block header\n", chain)
					}
					balToken0, _ := uint256.FromBig(poolTokens.Balances[0])
					balToken1, _ := uint256.FromBig(poolTokens.Balances[1])
					swapFeePercentage, _ := uint256.FromBig(swapFeePerc)
					weightToken0, _ := uint256.FromBig(weights[0])
					weightToken1, _ := uint256.FromBig(weights[1])

					pools <- &poolinterface.PoolBalancerV2{
						Chain:             chain,
						LpAddress:         lpAddress,
						Dex:               "BalancerV2Wp2Token",
						Token0:            poolTokens.Tokens[0],
						Token1:            poolTokens.Tokens[1],
						BalanceToken0:     &libtypes.UInt256{Int: balToken0},
						BalanceToken1:     &libtypes.UInt256{Int: balToken1},
						SwapFeePercentage: &libtypes.UInt256{Int: swapFeePercentage},
						NormalizedWeight0: &libtypes.UInt256{Int: weightToken0},
						NormalizedWeight1: &libtypes.UInt256{Int: weightToken1},
						UpdatedAt:         int64(header.Time),
						UpdatedBlock:      uint64(currentBlock),
					}
				}
			}(_log)
			wg.Wait()
		}
	}
}
