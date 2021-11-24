package poolinterface

import (
	"github.com/artworkk/balancer-factory-logs/enums"
	"github.com/artworkk/balancer-factory-logs/types"
	"github.com/ethereum/go-ethereum/common"
)

type PoolBalancerV2 struct {
	// Key
	Chain     enums.ChainType
	LpAddress common.Address
	// Info
	Dex               string
	Token0            common.Address
	Token1            common.Address
	Token0Index       int
	Token1Index       int
	SwapFeePercentage *types.UInt256
	BalanceToken0     *types.UInt256
	BalanceToken1     *types.UInt256
	Token0Decimal     int
	Token1Decimal     int
	ScalingFactory0   *types.UInt256
	ScalingFactory1   *types.UInt256
	NormalizedWeight0 *types.UInt256
	NormalizedWeight1 *types.UInt256
	UpdatedAt         int64
	UpdatedBlock      uint64
}
