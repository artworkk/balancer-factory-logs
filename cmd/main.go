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

	factoryContract "github.com/artworkk/balancer-factory-logs/contracts-go/balancer/factory"
	vaultContract "github.com/artworkk/balancer-factory-logs/contracts-go/balancer/thevault"
	poolContract "github.com/artworkk/balancer-factory-logs/contracts-go/balancer/weightedpool"
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
	PoolId       string          `json:"poolId"`
	LpAddress    common.Address  `json:"lpAddress"`
	Chain        enums.ChainType `json:"chain"`
	// thevault.GetPoolTokens(poolId) doesn't work yet
	Token0       common.Address  `json:"token0"`
	Token1       common.Address  `json:"token1"`
	WeightToken0 *big.Int        `json:"weightToken0"`
	WeightToken1 *big.Int        `json:"weightToken1"`
	SwapFeePerc  *big.Int        `json:"swapFeePercentage"`
	IsPaused     bool            `json:"isPaused"`
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

	// Balancer factory genesis
	fromBlock := enums.DexFactGenesis[chain]
	// Balancer factory addr
	factoryAddress := enums.DexFactAddr[chain]
	// Balancer factory ABI
	balancerFactoryAbi, err := abi.JSON(strings.NewReader(factoryContract.FactoryABI))
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
					chain := chain
					lpAddress := common.HexToAddress(l.Topics[1].String())
					log.Printf("[%s]\tFound PoolCreated on block %d at address %s\n",
						chain, l.BlockNumber, lpAddress)

					callOpts := &bind.CallOpts{
						BlockNumber: big.NewInt(currentBlock),
						Context:     ctx,
					}
					poolCaller, err := poolContract.NewContract(lpAddress, client)
					if err != nil {
						log.Fatalf("[%s]\tfailed to create poolCaller for pool %v\n%v\n", chain, lpAddress, err.Error())
					}
					poolId, err := poolCaller.GetPoolId(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool ID from LP address %v weights\n%v\n", chain, lpAddress, err.Error())
					}
					vaultCaller, err := vaultContract.NewThevaultCaller(lpAddress, client)
					if err != nil {
						log.Fatalf("[%s]\tfailed to create vaultCaller for pool %v\n%s\n", chain, lpAddress, err.Error())
					}
					poolTokens, err := vaultCaller.GetPoolTokens(callOpts, poolId)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool tokens for pool %v\n%v\n", chain, fmt.Sprintf("0x%x", poolId), err.Error())
					}
					weights, err := poolCaller.GetNormalizedWeights(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool %v weights\n", chain, lpAddress)
					}
					swapFeePerc, err := poolCaller.GetSwapFeePercentage(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tfailed to get pool %v swap fee percentage\n", chain, lpAddress)
					}
					paused, err := vaultCaller.GetPausedState(callOpts)
					if err != nil {
						log.Fatalf("[%s]\tFailed to get the vault paused state", chain)
					}
					isPaused := paused.Paused
					pools <- &poolData{
						PoolId:       fmt.Sprintf("0x%x", poolId),
						LpAddress:    lpAddress,
						Chain:        chain,
						Token0:       poolTokens.Tokens[0],
						Token1:       poolTokens.Tokens[1],
						WeightToken0: weights[0],
						WeightToken1: weights[1],
						SwapFeePerc:  swapFeePerc,
						IsPaused:     isPaused,
					}
				}
			}(_log)
			wg.Wait()
		}
	}
}