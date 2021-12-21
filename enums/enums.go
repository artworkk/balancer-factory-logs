package enums

import "github.com/ethereum/go-ethereum/common"

type ChainType string
type DexType string
type NodeAddress string

// ChainType constants are used as key in the maps
const (
	Ethereum ChainType = "ethereum"
	Polygon  ChainType = "polygon"
	Arbitrum ChainType = "arbitrum"
)

// Pool types
const (
	Wp         DexType = "wp"
	Wp2Tokens  DexType = "wp2tokens"
	StablePool DexType = "stablepool"
)

// NodeAddress, factory addresses, and factory genesis blocks are used as values
const (
	EthereumNode NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/eth/mainnet"
	PolygonNode  NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/polygon/mainnet"
	ArbitrumNode NodeAddress = "https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/arbitrum/mainnet"
)

const (
	WPFactoryGenesisEthereum         int64 = 12272147
	WPFactoryGenesisPolygon          int64 = 15832998
	WPFactoryGenesisArbitrum         int64 = 222863
	Wp2TokensFactoryGenesisEthereum  int64 = 12272147
	Wp2TokensFactoryGenesisPolygon   int64 = 15832998
	Wp2TokensFactoryGenesisArbitrum  int64 = 222864
	StablePoolFactoryGenesisEthereum int64 = 12703127
	StablePoolFactoryGenesisPolygon  int64 = 16138680
	StablePoolFactoryGenesisArbitrum int64 = 222866
)

var (
	WPFactoryEthereum         common.Address = common.HexToAddress("0x8E9aa87E45e92bad84D5F8DD1bff34Fb92637dE9")
	WPFactoryPolygon          common.Address = common.HexToAddress("0x8E9aa87E45e92bad84D5F8DD1bff34Fb92637dE9")
	WPFactoryArbitrum         common.Address = common.HexToAddress("0x7dFdEF5f355096603419239CE743BfaF1120312B")
	Wp2TokensFactoryEthereum  common.Address = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
	Wp2TokensFactoryPolygon   common.Address = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
	Wp2TokensFactoryArbitrum  common.Address = common.HexToAddress("0xCF0a32Bbef8F064969F21f7e02328FB577382018")
	StablePoolFactoryEthereum common.Address = common.HexToAddress("0xc66Ba2B6595D3613CCab350C886aCE23866EDe24")
	StablePoolFactoryPolygon  common.Address = common.HexToAddress("0xc66Ba2B6595D3613CCab350C886aCE23866EDe24")
	StablePoolFactoryArbitrum common.Address = common.HexToAddress("0x2433477A10FC5d31B9513C638F19eE85CaED53Fd")
)

// Maps
var (
	DexNode = map[ChainType]NodeAddress{
		Ethereum: EthereumNode,
		Polygon:  PolygonNode,
		Arbitrum: ArbitrumNode,
	}
	Wp2TokensFactoryAddr = map[ChainType]common.Address{
		Ethereum: Wp2TokensFactoryEthereum,
		Polygon:  Wp2TokensFactoryPolygon,
		Arbitrum: Wp2TokensFactoryArbitrum,
	}
	StablePoolFactoryAddr = map[ChainType]common.Address{
		Ethereum: StablePoolFactoryEthereum,
		Polygon:  StablePoolFactoryPolygon,
		Arbitrum: StablePoolFactoryArbitrum,
	}
	WPFactoryAddr = map[ChainType]common.Address{
		Ethereum: WPFactoryEthereum,
		Polygon:  WPFactoryPolygon,
		Arbitrum: WPFactoryArbitrum,
	}
	Wp2TokensGenesis = map[ChainType]int64{
		Ethereum: Wp2TokensFactoryGenesisEthereum,
		Polygon:  Wp2TokensFactoryGenesisPolygon,
		Arbitrum: Wp2TokensFactoryGenesisArbitrum,
	}
	StablePoolGenesis = map[ChainType]int64{
		Ethereum: StablePoolFactoryGenesisEthereum,
		Polygon:  StablePoolFactoryGenesisPolygon,
		Arbitrum: StablePoolFactoryGenesisArbitrum,
	}
	WPGenesis = map[ChainType]int64{
		Ethereum: WPFactoryGenesisEthereum,
		Polygon:  WPFactoryGenesisPolygon,
		Arbitrum: WPFactoryGenesisArbitrum,
	}
)
