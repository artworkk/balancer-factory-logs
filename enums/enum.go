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
	EthereumNode NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/eth/mainnet"
	PolygonNode  NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/polygon/mainnet"
)

const (
	Wp2TokensFactoryGenesisEthereum  int64 = 12272147
	Wp2TokensFactoryGenesisPolygon   int64 = 15832998
	StablePoolFactoryGenesisEthereum int64 = 12703127
	StablePoolFactoryGenesisPolygon  int64 = 16138680
)

var (
	Wp2tokensFactoryEthereum  common.Address = common.HexToAddress("0x8E9aa87E45e92bad84D5F8DD1bff34Fb92637dE9")
	Wp2TokensFactoryPolygon   common.Address = common.HexToAddress("0x8E9aa87E45e92bad84D5F8DD1bff34Fb92637dE9")
	StablePoolFactoryEthereum common.Address = common.HexToAddress("0xc66Ba2B6595D3613CCab350C886aCE23866EDe24")
	StablePoolFactoryPolygon  common.Address = common.HexToAddress("0xc66Ba2B6595D3613CCab350C886aCE23866EDe24")
)

// Maps
var (
	DexNode = map[ChainType]NodeAddress{
		Ethereum: EthereumNode,
		Polygon:  PolygonNode,
	}
	Wp2TokensFactoryAddr = map[ChainType]common.Address{
		Ethereum: Wp2tokensFactoryEthereum,
		Polygon:  Wp2TokensFactoryPolygon,
	}
	StablePoolFactoryAddr = map[ChainType]common.Address{
		Ethereum: StablePoolFactoryEthereum,
		Polygon:  StablePoolFactoryPolygon,
	}
	Wp2TokensGenesis = map[ChainType]int64{
		Ethereum: Wp2TokensFactoryGenesisEthereum,
		Polygon:  Wp2TokensFactoryGenesisPolygon,
	}
	StablePoolGenesis = map[ChainType]int64{
		Ethereum: StablePoolFactoryGenesisEthereum,
		Polygon:  StablePoolFactoryGenesisPolygon,
	}
)
