// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wp2tokens

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPoolSwapStructsSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IPoolSwapStructsSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// IPriceOracleOracleAccumulatorQuery is an auto generated low-level Go binding around an user-defined struct.
type IPriceOracleOracleAccumulatorQuery struct {
	Variable uint8
	Ago      *big.Int
}

// IPriceOracleOracleAverageQuery is an auto generated low-level Go binding around an user-defined struct.
type IPriceOracleOracleAverageQuery struct {
	Variable uint8
	Secs     *big.Int
	Ago      *big.Int
}

// BalancerV2Wp2TokensMetaData contains all meta data concerning the BalancerV2Wp2Tokens contract.
var BalancerV2Wp2TokensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"normalizedWeight0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"normalizedWeight1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structWeightedPool2Tokens.NewPoolParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OracleEnabledChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLargestSafeQueryWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPriceOracle.Variable\",\"name\":\"variable\",\"type\":\"uint8\"}],\"name\":\"getLatest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiscData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"logInvariant\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"logTotalSupply\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"oracleSampleCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNormalizedWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPriceOracle.Variable\",\"name\":\"variable\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ago\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.OracleAccumulatorQuery[]\",\"name\":\"queries\",\"type\":\"tuple[]\"}],\"name\":\"getPastAccumulators\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"results\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"logPairPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accLogPairPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"logBptPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accLogBptPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"logInvariant\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"accLogInvariant\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPriceOracle.Variable\",\"name\":\"variable\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ago\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.OracleAverageQuery[]\",\"name\":\"queries\",\"type\":\"tuple[]\"}],\"name\":\"getTimeWeightedAverage\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"results\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSamples\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"dueProtocolFeeAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancerV2Wp2TokensABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2Wp2TokensMetaData.ABI instead.
var BalancerV2Wp2TokensABI = BalancerV2Wp2TokensMetaData.ABI

// BalancerV2Wp2Tokens is an auto generated Go binding around an Ethereum contract.
type BalancerV2Wp2Tokens struct {
	BalancerV2Wp2TokensCaller     // Read-only binding to the contract
	BalancerV2Wp2TokensTransactor // Write-only binding to the contract
	BalancerV2Wp2TokensFilterer   // Log filterer for contract events
}

// BalancerV2Wp2TokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2Wp2TokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2Wp2TokensSession struct {
	Contract     *BalancerV2Wp2Tokens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BalancerV2Wp2TokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2Wp2TokensCallerSession struct {
	Contract *BalancerV2Wp2TokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BalancerV2Wp2TokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2Wp2TokensTransactorSession struct {
	Contract     *BalancerV2Wp2TokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BalancerV2Wp2TokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2Wp2TokensRaw struct {
	Contract *BalancerV2Wp2Tokens // Generic contract binding to access the raw methods on
}

// BalancerV2Wp2TokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensCallerRaw struct {
	Contract *BalancerV2Wp2TokensCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2Wp2TokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensTransactorRaw struct {
	Contract *BalancerV2Wp2TokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2Wp2Tokens creates a new instance of BalancerV2Wp2Tokens, bound to a specific deployed contract.
func NewBalancerV2Wp2Tokens(address common.Address, backend bind.ContractBackend) (*BalancerV2Wp2Tokens, error) {
	contract, err := bindBalancerV2Wp2Tokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2Tokens{BalancerV2Wp2TokensCaller: BalancerV2Wp2TokensCaller{contract: contract}, BalancerV2Wp2TokensTransactor: BalancerV2Wp2TokensTransactor{contract: contract}, BalancerV2Wp2TokensFilterer: BalancerV2Wp2TokensFilterer{contract: contract}}, nil
}

// NewBalancerV2Wp2TokensCaller creates a new read-only instance of BalancerV2Wp2Tokens, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2Wp2TokensCaller, error) {
	contract, err := bindBalancerV2Wp2Tokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensCaller{contract: contract}, nil
}

// NewBalancerV2Wp2TokensTransactor creates a new write-only instance of BalancerV2Wp2Tokens, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2Wp2TokensTransactor, error) {
	contract, err := bindBalancerV2Wp2Tokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensTransactor{contract: contract}, nil
}

// NewBalancerV2Wp2TokensFilterer creates a new log filterer instance of BalancerV2Wp2Tokens, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2Wp2TokensFilterer, error) {
	contract, err := bindBalancerV2Wp2Tokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFilterer{contract: contract}, nil
}

// bindBalancerV2Wp2Tokens binds a generic wrapper to an already deployed contract.
func bindBalancerV2Wp2Tokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerV2Wp2TokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Wp2Tokens.Contract.BalancerV2Wp2TokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.BalancerV2Wp2TokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.BalancerV2Wp2TokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Wp2Tokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.DOMAINSEPARATOR(&_BalancerV2Wp2Tokens.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.DOMAINSEPARATOR(&_BalancerV2Wp2Tokens.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.Allowance(&_BalancerV2Wp2Tokens.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.Allowance(&_BalancerV2Wp2Tokens.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.BalanceOf(&_BalancerV2Wp2Tokens.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.BalanceOf(&_BalancerV2Wp2Tokens.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Decimals() (uint8, error) {
	return _BalancerV2Wp2Tokens.Contract.Decimals(&_BalancerV2Wp2Tokens.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) Decimals() (uint8, error) {
	return _BalancerV2Wp2Tokens.Contract.Decimals(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.GetActionId(&_BalancerV2Wp2Tokens.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.GetActionId(&_BalancerV2Wp2Tokens.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetAuthorizer() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetAuthorizer(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetAuthorizer() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetAuthorizer(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetInvariant() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetInvariant(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetInvariant() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetInvariant(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetLargestSafeQueryWindow is a free data retrieval call binding the contract method 0xffd088eb.
//
// Solidity: function getLargestSafeQueryWindow() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetLargestSafeQueryWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getLargestSafeQueryWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLargestSafeQueryWindow is a free data retrieval call binding the contract method 0xffd088eb.
//
// Solidity: function getLargestSafeQueryWindow() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetLargestSafeQueryWindow() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLargestSafeQueryWindow(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetLargestSafeQueryWindow is a free data retrieval call binding the contract method 0xffd088eb.
//
// Solidity: function getLargestSafeQueryWindow() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetLargestSafeQueryWindow() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLargestSafeQueryWindow(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetLastInvariant is a free data retrieval call binding the contract method 0x9b02cdde.
//
// Solidity: function getLastInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetLastInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getLastInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastInvariant is a free data retrieval call binding the contract method 0x9b02cdde.
//
// Solidity: function getLastInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetLastInvariant() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLastInvariant(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetLastInvariant is a free data retrieval call binding the contract method 0x9b02cdde.
//
// Solidity: function getLastInvariant() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetLastInvariant() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLastInvariant(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetLatest is a free data retrieval call binding the contract method 0xb10be739.
//
// Solidity: function getLatest(uint8 variable) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetLatest(opts *bind.CallOpts, variable uint8) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getLatest", variable)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatest is a free data retrieval call binding the contract method 0xb10be739.
//
// Solidity: function getLatest(uint8 variable) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetLatest(variable uint8) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLatest(&_BalancerV2Wp2Tokens.CallOpts, variable)
}

// GetLatest is a free data retrieval call binding the contract method 0xb10be739.
//
// Solidity: function getLatest(uint8 variable) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetLatest(variable uint8) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetLatest(&_BalancerV2Wp2Tokens.CallOpts, variable)
}

// GetMiscData is a free data retrieval call binding the contract method 0x4a6b0b15.
//
// Solidity: function getMiscData() view returns(int256 logInvariant, int256 logTotalSupply, uint256 oracleSampleCreationTimestamp, uint256 oracleIndex, bool oracleEnabled, uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetMiscData(opts *bind.CallOpts) (struct {
	LogInvariant                  *big.Int
	LogTotalSupply                *big.Int
	OracleSampleCreationTimestamp *big.Int
	OracleIndex                   *big.Int
	OracleEnabled                 bool
	SwapFeePercentage             *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getMiscData")

	outstruct := new(struct {
		LogInvariant                  *big.Int
		LogTotalSupply                *big.Int
		OracleSampleCreationTimestamp *big.Int
		OracleIndex                   *big.Int
		OracleEnabled                 bool
		SwapFeePercentage             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LogInvariant = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LogTotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OracleSampleCreationTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OracleIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OracleEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.SwapFeePercentage = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMiscData is a free data retrieval call binding the contract method 0x4a6b0b15.
//
// Solidity: function getMiscData() view returns(int256 logInvariant, int256 logTotalSupply, uint256 oracleSampleCreationTimestamp, uint256 oracleIndex, bool oracleEnabled, uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetMiscData() (struct {
	LogInvariant                  *big.Int
	LogTotalSupply                *big.Int
	OracleSampleCreationTimestamp *big.Int
	OracleIndex                   *big.Int
	OracleEnabled                 bool
	SwapFeePercentage             *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetMiscData(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetMiscData is a free data retrieval call binding the contract method 0x4a6b0b15.
//
// Solidity: function getMiscData() view returns(int256 logInvariant, int256 logTotalSupply, uint256 oracleSampleCreationTimestamp, uint256 oracleIndex, bool oracleEnabled, uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetMiscData() (struct {
	LogInvariant                  *big.Int
	LogTotalSupply                *big.Int
	OracleSampleCreationTimestamp *big.Int
	OracleIndex                   *big.Int
	OracleEnabled                 bool
	SwapFeePercentage             *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetMiscData(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetNormalizedWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getNormalizedWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetNormalizedWeights(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetNormalizedWeights(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetOwner() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetOwner(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetOwner() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetOwner(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetPastAccumulators is a free data retrieval call binding the contract method 0x6b843239.
//
// Solidity: function getPastAccumulators((uint8,uint256)[] queries) view returns(int256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetPastAccumulators(opts *bind.CallOpts, queries []IPriceOracleOracleAccumulatorQuery) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getPastAccumulators", queries)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPastAccumulators is a free data retrieval call binding the contract method 0x6b843239.
//
// Solidity: function getPastAccumulators((uint8,uint256)[] queries) view returns(int256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetPastAccumulators(queries []IPriceOracleOracleAccumulatorQuery) ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPastAccumulators(&_BalancerV2Wp2Tokens.CallOpts, queries)
}

// GetPastAccumulators is a free data retrieval call binding the contract method 0x6b843239.
//
// Solidity: function getPastAccumulators((uint8,uint256)[] queries) view returns(int256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetPastAccumulators(queries []IPriceOracleOracleAccumulatorQuery) ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPastAccumulators(&_BalancerV2Wp2Tokens.CallOpts, queries)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPausedState(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPausedState(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetPoolId() ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPoolId(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetPoolId() ([32]byte, error) {
	return _BalancerV2Wp2Tokens.Contract.GetPoolId(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetRate() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetRate(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetRate() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetRate(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetSample is a free data retrieval call binding the contract method 0x60d1507c.
//
// Solidity: function getSample(uint256 index) view returns(int256 logPairPrice, int256 accLogPairPrice, int256 logBptPrice, int256 accLogBptPrice, int256 logInvariant, int256 accLogInvariant, uint256 timestamp)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetSample(opts *bind.CallOpts, index *big.Int) (struct {
	LogPairPrice    *big.Int
	AccLogPairPrice *big.Int
	LogBptPrice     *big.Int
	AccLogBptPrice  *big.Int
	LogInvariant    *big.Int
	AccLogInvariant *big.Int
	Timestamp       *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getSample", index)

	outstruct := new(struct {
		LogPairPrice    *big.Int
		AccLogPairPrice *big.Int
		LogBptPrice     *big.Int
		AccLogBptPrice  *big.Int
		LogInvariant    *big.Int
		AccLogInvariant *big.Int
		Timestamp       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LogPairPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccLogPairPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LogBptPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccLogBptPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LogInvariant = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccLogInvariant = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSample is a free data retrieval call binding the contract method 0x60d1507c.
//
// Solidity: function getSample(uint256 index) view returns(int256 logPairPrice, int256 accLogPairPrice, int256 logBptPrice, int256 accLogBptPrice, int256 logInvariant, int256 accLogInvariant, uint256 timestamp)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetSample(index *big.Int) (struct {
	LogPairPrice    *big.Int
	AccLogPairPrice *big.Int
	LogBptPrice     *big.Int
	AccLogBptPrice  *big.Int
	LogInvariant    *big.Int
	AccLogInvariant *big.Int
	Timestamp       *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetSample(&_BalancerV2Wp2Tokens.CallOpts, index)
}

// GetSample is a free data retrieval call binding the contract method 0x60d1507c.
//
// Solidity: function getSample(uint256 index) view returns(int256 logPairPrice, int256 accLogPairPrice, int256 logBptPrice, int256 accLogBptPrice, int256 logInvariant, int256 accLogInvariant, uint256 timestamp)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetSample(index *big.Int) (struct {
	LogPairPrice    *big.Int
	AccLogPairPrice *big.Int
	LogBptPrice     *big.Int
	AccLogBptPrice  *big.Int
	LogInvariant    *big.Int
	AccLogInvariant *big.Int
	Timestamp       *big.Int
}, error) {
	return _BalancerV2Wp2Tokens.Contract.GetSample(&_BalancerV2Wp2Tokens.CallOpts, index)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetSwapFeePercentage(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetSwapFeePercentage(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetTimeWeightedAverage is a free data retrieval call binding the contract method 0x1dccd830.
//
// Solidity: function getTimeWeightedAverage((uint8,uint256,uint256)[] queries) view returns(uint256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetTimeWeightedAverage(opts *bind.CallOpts, queries []IPriceOracleOracleAverageQuery) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getTimeWeightedAverage", queries)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTimeWeightedAverage is a free data retrieval call binding the contract method 0x1dccd830.
//
// Solidity: function getTimeWeightedAverage((uint8,uint256,uint256)[] queries) view returns(uint256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetTimeWeightedAverage(queries []IPriceOracleOracleAverageQuery) ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetTimeWeightedAverage(&_BalancerV2Wp2Tokens.CallOpts, queries)
}

// GetTimeWeightedAverage is a free data retrieval call binding the contract method 0x1dccd830.
//
// Solidity: function getTimeWeightedAverage((uint8,uint256,uint256)[] queries) view returns(uint256[] results)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetTimeWeightedAverage(queries []IPriceOracleOracleAverageQuery) ([]*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetTimeWeightedAverage(&_BalancerV2Wp2Tokens.CallOpts, queries)
}

// GetTotalSamples is a free data retrieval call binding the contract method 0xb48b5b40.
//
// Solidity: function getTotalSamples() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetTotalSamples(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getTotalSamples")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSamples is a free data retrieval call binding the contract method 0xb48b5b40.
//
// Solidity: function getTotalSamples() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetTotalSamples() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetTotalSamples(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetTotalSamples is a free data retrieval call binding the contract method 0xb48b5b40.
//
// Solidity: function getTotalSamples() pure returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetTotalSamples() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.GetTotalSamples(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) GetVault() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetVault(&_BalancerV2Wp2Tokens.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) GetVault() (common.Address, error) {
	return _BalancerV2Wp2Tokens.Contract.GetVault(&_BalancerV2Wp2Tokens.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Name() (string, error) {
	return _BalancerV2Wp2Tokens.Contract.Name(&_BalancerV2Wp2Tokens.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) Name() (string, error) {
	return _BalancerV2Wp2Tokens.Contract.Name(&_BalancerV2Wp2Tokens.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.Nonces(&_BalancerV2Wp2Tokens.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.Nonces(&_BalancerV2Wp2Tokens.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Symbol() (string, error) {
	return _BalancerV2Wp2Tokens.Contract.Symbol(&_BalancerV2Wp2Tokens.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) Symbol() (string, error) {
	return _BalancerV2Wp2Tokens.Contract.Symbol(&_BalancerV2Wp2Tokens.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2Wp2Tokens.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) TotalSupply() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.TotalSupply(&_BalancerV2Wp2Tokens.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerV2Wp2Tokens.Contract.TotalSupply(&_BalancerV2Wp2Tokens.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Approve(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Approve(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "decreaseApproval", spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.DecreaseApproval(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.DecreaseApproval(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// EnableOracle is a paid mutator transaction binding the contract method 0x292c914a.
//
// Solidity: function enableOracle() returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) EnableOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "enableOracle")
}

// EnableOracle is a paid mutator transaction binding the contract method 0x292c914a.
//
// Solidity: function enableOracle() returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) EnableOracle() (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.EnableOracle(&_BalancerV2Wp2Tokens.TransactOpts)
}

// EnableOracle is a paid mutator transaction binding the contract method 0x292c914a.
//
// Solidity: function enableOracle() returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) EnableOracle() (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.EnableOracle(&_BalancerV2Wp2Tokens.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "increaseApproval", spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.IncreaseApproval(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.IncreaseApproval(&_BalancerV2Wp2Tokens.TransactOpts, spender, amount)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnExitPool(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnExitPool(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[] amountsIn, uint256[] dueProtocolFeeAmounts)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[] amountsIn, uint256[] dueProtocolFeeAmounts)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnJoinPool(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[] amountsIn, uint256[] dueProtocolFeeAmounts)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnJoinPool(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) OnSwap(opts *bind.TransactOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "onSwap", request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnSwap(&_BalancerV2Wp2Tokens.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.OnSwap(&_BalancerV2Wp2Tokens.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Permit(&_BalancerV2Wp2Tokens.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Permit(&_BalancerV2Wp2Tokens.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) QueryExit(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.QueryExit(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.QueryExit(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) QueryJoin(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.QueryJoin(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.QueryJoin(&_BalancerV2Wp2Tokens.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.SetPaused(&_BalancerV2Wp2Tokens.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.SetPaused(&_BalancerV2Wp2Tokens.TransactOpts, paused)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.SetSwapFeePercentage(&_BalancerV2Wp2Tokens.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.SetSwapFeePercentage(&_BalancerV2Wp2Tokens.TransactOpts, swapFeePercentage)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Transfer(&_BalancerV2Wp2Tokens.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.Transfer(&_BalancerV2Wp2Tokens.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.TransferFrom(&_BalancerV2Wp2Tokens.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2Wp2Tokens.Contract.TransferFrom(&_BalancerV2Wp2Tokens.TransactOpts, sender, recipient, amount)
}

// BalancerV2Wp2TokensApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensApprovalIterator struct {
	Event *BalancerV2Wp2TokensApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerV2Wp2TokensApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerV2Wp2TokensApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerV2Wp2TokensApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensApproval represents a Approval event raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BalancerV2Wp2TokensApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerV2Wp2Tokens.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensApprovalIterator{contract: _BalancerV2Wp2Tokens.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerV2Wp2Tokens.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensApproval)
				if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) ParseApproval(log types.Log) (*BalancerV2Wp2TokensApproval, error) {
	event := new(BalancerV2Wp2TokensApproval)
	if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2Wp2TokensOracleEnabledChangedIterator is returned from FilterOracleEnabledChanged and is used to iterate over the raw logs and unpacked data for OracleEnabledChanged events raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensOracleEnabledChangedIterator struct {
	Event *BalancerV2Wp2TokensOracleEnabledChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerV2Wp2TokensOracleEnabledChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensOracleEnabledChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerV2Wp2TokensOracleEnabledChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerV2Wp2TokensOracleEnabledChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensOracleEnabledChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensOracleEnabledChanged represents a OracleEnabledChanged event raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensOracleEnabledChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOracleEnabledChanged is a free log retrieval operation binding the contract event 0x3e350b41e86a8e10f804ade6d35340d620be35569cc75ac943e8bb14ab80ead1.
//
// Solidity: event OracleEnabledChanged(bool enabled)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) FilterOracleEnabledChanged(opts *bind.FilterOpts) (*BalancerV2Wp2TokensOracleEnabledChangedIterator, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.FilterLogs(opts, "OracleEnabledChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensOracleEnabledChangedIterator{contract: _BalancerV2Wp2Tokens.contract, event: "OracleEnabledChanged", logs: logs, sub: sub}, nil
}

// WatchOracleEnabledChanged is a free log subscription operation binding the contract event 0x3e350b41e86a8e10f804ade6d35340d620be35569cc75ac943e8bb14ab80ead1.
//
// Solidity: event OracleEnabledChanged(bool enabled)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) WatchOracleEnabledChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensOracleEnabledChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.WatchLogs(opts, "OracleEnabledChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensOracleEnabledChanged)
				if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "OracleEnabledChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleEnabledChanged is a log parse operation binding the contract event 0x3e350b41e86a8e10f804ade6d35340d620be35569cc75ac943e8bb14ab80ead1.
//
// Solidity: event OracleEnabledChanged(bool enabled)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) ParseOracleEnabledChanged(log types.Log) (*BalancerV2Wp2TokensOracleEnabledChanged, error) {
	event := new(BalancerV2Wp2TokensOracleEnabledChanged)
	if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "OracleEnabledChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2Wp2TokensPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensPausedStateChangedIterator struct {
	Event *BalancerV2Wp2TokensPausedStateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerV2Wp2TokensPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensPausedStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerV2Wp2TokensPausedStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerV2Wp2TokensPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensPausedStateChanged represents a PausedStateChanged event raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BalancerV2Wp2TokensPausedStateChangedIterator, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensPausedStateChangedIterator{contract: _BalancerV2Wp2Tokens.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensPausedStateChanged)
				if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) ParsePausedStateChanged(log types.Log) (*BalancerV2Wp2TokensPausedStateChanged, error) {
	event := new(BalancerV2Wp2TokensPausedStateChanged)
	if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2Wp2TokensSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensSwapFeePercentageChangedIterator struct {
	Event *BalancerV2Wp2TokensSwapFeePercentageChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerV2Wp2TokensSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensSwapFeePercentageChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerV2Wp2TokensSwapFeePercentageChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerV2Wp2TokensSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BalancerV2Wp2TokensSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensSwapFeePercentageChangedIterator{contract: _BalancerV2Wp2Tokens.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2Wp2Tokens.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensSwapFeePercentageChanged)
				if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BalancerV2Wp2TokensSwapFeePercentageChanged, error) {
	event := new(BalancerV2Wp2TokensSwapFeePercentageChanged)
	if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2Wp2TokensTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensTransferIterator struct {
	Event *BalancerV2Wp2TokensTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerV2Wp2TokensTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerV2Wp2TokensTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerV2Wp2TokensTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensTransfer represents a Transfer event raised by the BalancerV2Wp2Tokens contract.
type BalancerV2Wp2TokensTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BalancerV2Wp2TokensTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerV2Wp2Tokens.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensTransferIterator{contract: _BalancerV2Wp2Tokens.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerV2Wp2Tokens.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensTransfer)
				if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerV2Wp2Tokens *BalancerV2Wp2TokensFilterer) ParseTransfer(log types.Log) (*BalancerV2Wp2TokensTransfer, error) {
	event := new(BalancerV2Wp2TokensTransfer)
	if err := _BalancerV2Wp2Tokens.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
