package enums

import "github.com/ethereum/go-ethereum/common"

type ChainType string
type NodeAddress string

// ChainType constants are used as key in the maps
const (
	Ethereum ChainType = "ethereum"
	Polygon  ChainType = "polygon"
)

// NodeAddress, factory addresses, and factory genesis blocks are used as values
const (
	EthereumNode           NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/eth/mainnet"
	PolygonNode            NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/polygon/mainnet"
	EthereumFactoryGenesis int64       = 12349891
	PolygonFactoryGenesis  int64       = 15869090
)

var (
	EthereumFactoryAddress common.Address = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
	PolygonFactoryAddress  common.Address = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
)

// Maps
var (
	DexNode = map[ChainType]NodeAddress{
		Ethereum: EthereumNode,
		Polygon:  PolygonNode,
	}
	DexFactAddr = map[ChainType]common.Address{
		Ethereum: EthereumFactoryAddress,
		Polygon:  PolygonFactoryAddress,
	}
	DexFactGenesis = map[ChainType]int64{
		Ethereum: EthereumFactoryGenesis,
		Polygon:  PolygonFactoryGenesis,
	}
)
