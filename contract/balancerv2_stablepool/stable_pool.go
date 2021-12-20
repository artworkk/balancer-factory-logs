// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stablepool

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

// BalancerV2StablePoolMetaData contains all meta data concerning the BalancerV2StablePool contract.
var BalancerV2StablePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amplificationParameter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AmpUpdateStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentValue\",\"type\":\"uint256\"}],\"name\":\"AmpUpdateStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmplificationParameter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"swapRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"indexIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"poolConfig\",\"type\":\"bytes\"}],\"name\":\"setAssetManagerPoolConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rawEndValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"startAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancerV2StablePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2StablePoolMetaData.ABI instead.
var BalancerV2StablePoolABI = BalancerV2StablePoolMetaData.ABI

// BalancerV2StablePool is an auto generated Go binding around an Ethereum contract.
type BalancerV2StablePool struct {
	BalancerV2StablePoolCaller     // Read-only binding to the contract
	BalancerV2StablePoolTransactor // Write-only binding to the contract
	BalancerV2StablePoolFilterer   // Log filterer for contract events
}

// BalancerV2StablePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2StablePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2StablePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2StablePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2StablePoolSession struct {
	Contract     *BalancerV2StablePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalancerV2StablePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2StablePoolCallerSession struct {
	Contract *BalancerV2StablePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BalancerV2StablePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2StablePoolTransactorSession struct {
	Contract     *BalancerV2StablePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BalancerV2StablePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2StablePoolRaw struct {
	Contract *BalancerV2StablePool // Generic contract binding to access the raw methods on
}

// BalancerV2StablePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2StablePoolCallerRaw struct {
	Contract *BalancerV2StablePoolCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2StablePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2StablePoolTransactorRaw struct {
	Contract *BalancerV2StablePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2StablePool creates a new instance of BalancerV2StablePool, bound to a specific deployed contract.
func NewBalancerV2StablePool(address common.Address, backend bind.ContractBackend) (*BalancerV2StablePool, error) {
	contract, err := bindBalancerV2StablePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePool{BalancerV2StablePoolCaller: BalancerV2StablePoolCaller{contract: contract}, BalancerV2StablePoolTransactor: BalancerV2StablePoolTransactor{contract: contract}, BalancerV2StablePoolFilterer: BalancerV2StablePoolFilterer{contract: contract}}, nil
}

// NewBalancerV2StablePoolCaller creates a new read-only instance of BalancerV2StablePool, bound to a specific deployed contract.
func NewBalancerV2StablePoolCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2StablePoolCaller, error) {
	contract, err := bindBalancerV2StablePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolCaller{contract: contract}, nil
}

// NewBalancerV2StablePoolTransactor creates a new write-only instance of BalancerV2StablePool, bound to a specific deployed contract.
func NewBalancerV2StablePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2StablePoolTransactor, error) {
	contract, err := bindBalancerV2StablePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolTransactor{contract: contract}, nil
}

// NewBalancerV2StablePoolFilterer creates a new log filterer instance of BalancerV2StablePool, bound to a specific deployed contract.
func NewBalancerV2StablePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2StablePoolFilterer, error) {
	contract, err := bindBalancerV2StablePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFilterer{contract: contract}, nil
}

// bindBalancerV2StablePool binds a generic wrapper to an already deployed contract.
func bindBalancerV2StablePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerV2StablePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2StablePool *BalancerV2StablePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2StablePool.Contract.BalancerV2StablePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2StablePool *BalancerV2StablePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.BalancerV2StablePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2StablePool *BalancerV2StablePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.BalancerV2StablePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2StablePool *BalancerV2StablePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2StablePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerV2StablePool.Contract.DOMAINSEPARATOR(&_BalancerV2StablePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerV2StablePool.Contract.DOMAINSEPARATOR(&_BalancerV2StablePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.Allowance(&_BalancerV2StablePool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.Allowance(&_BalancerV2StablePool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.BalanceOf(&_BalancerV2StablePool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.BalanceOf(&_BalancerV2StablePool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Decimals() (uint8, error) {
	return _BalancerV2StablePool.Contract.Decimals(&_BalancerV2StablePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) Decimals() (uint8, error) {
	return _BalancerV2StablePool.Contract.Decimals(&_BalancerV2StablePool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerV2StablePool.Contract.GetActionId(&_BalancerV2StablePool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerV2StablePool.Contract.GetActionId(&_BalancerV2StablePool.CallOpts, selector)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetAmplificationParameter(opts *bind.CallOpts) (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getAmplificationParameter")

	outstruct := new(struct {
		Value      *big.Int
		IsUpdating bool
		Precision  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsUpdating = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Precision = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _BalancerV2StablePool.Contract.GetAmplificationParameter(&_BalancerV2StablePool.CallOpts)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _BalancerV2StablePool.Contract.GetAmplificationParameter(&_BalancerV2StablePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetAuthorizer() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetAuthorizer(&_BalancerV2StablePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetAuthorizer(&_BalancerV2StablePool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetOwner() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetOwner(&_BalancerV2StablePool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetOwner() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetOwner(&_BalancerV2StablePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getPausedState")

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
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerV2StablePool.Contract.GetPausedState(&_BalancerV2StablePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerV2StablePool.Contract.GetPausedState(&_BalancerV2StablePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetPoolId() ([32]byte, error) {
	return _BalancerV2StablePool.Contract.GetPoolId(&_BalancerV2StablePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetPoolId() ([32]byte, error) {
	return _BalancerV2StablePool.Contract.GetPoolId(&_BalancerV2StablePool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetRate() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.GetRate(&_BalancerV2StablePool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetRate() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.GetRate(&_BalancerV2StablePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.GetSwapFeePercentage(&_BalancerV2StablePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.GetSwapFeePercentage(&_BalancerV2StablePool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) GetVault() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetVault(&_BalancerV2StablePool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) GetVault() (common.Address, error) {
	return _BalancerV2StablePool.Contract.GetVault(&_BalancerV2StablePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Name() (string, error) {
	return _BalancerV2StablePool.Contract.Name(&_BalancerV2StablePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) Name() (string, error) {
	return _BalancerV2StablePool.Contract.Name(&_BalancerV2StablePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.Nonces(&_BalancerV2StablePool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.Nonces(&_BalancerV2StablePool.CallOpts, owner)
}

// OnSwap is a free data retrieval call binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) OnSwap(opts *bind.CallOpts, swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "onSwap", swapRequest, balances, indexIn, indexOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OnSwap is a free data retrieval call binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) OnSwap(swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.OnSwap(&_BalancerV2StablePool.CallOpts, swapRequest, balances, indexIn, indexOut)
}

// OnSwap is a free data retrieval call binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) OnSwap(swapRequest IPoolSwapStructsSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.OnSwap(&_BalancerV2StablePool.CallOpts, swapRequest, balances, indexIn, indexOut)
}

// OnSwap0 is a free data retrieval call binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) OnSwap0(opts *bind.CallOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "onSwap0", request, balanceTokenIn, balanceTokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OnSwap0 is a free data retrieval call binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) OnSwap0(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.OnSwap0(&_BalancerV2StablePool.CallOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap0 is a free data retrieval call binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) OnSwap0(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*big.Int, error) {
	return _BalancerV2StablePool.Contract.OnSwap0(&_BalancerV2StablePool.CallOpts, request, balanceTokenIn, balanceTokenOut)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Symbol() (string, error) {
	return _BalancerV2StablePool.Contract.Symbol(&_BalancerV2StablePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) Symbol() (string, error) {
	return _BalancerV2StablePool.Contract.Symbol(&_BalancerV2StablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerV2StablePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) TotalSupply() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.TotalSupply(&_BalancerV2StablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerV2StablePool *BalancerV2StablePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerV2StablePool.Contract.TotalSupply(&_BalancerV2StablePool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Approve(&_BalancerV2StablePool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Approve(&_BalancerV2StablePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.DecreaseAllowance(&_BalancerV2StablePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.DecreaseAllowance(&_BalancerV2StablePool.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.IncreaseAllowance(&_BalancerV2StablePool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.IncreaseAllowance(&_BalancerV2StablePool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.OnExitPool(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.OnExitPool(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.OnJoinPool(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.OnJoinPool(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Permit(&_BalancerV2StablePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Permit(&_BalancerV2StablePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) QueryExit(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.QueryExit(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.QueryExit(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) QueryJoin(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.QueryJoin(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.QueryJoin(&_BalancerV2StablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) SetAssetManagerPoolConfig(opts *bind.TransactOpts, token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "setAssetManagerPoolConfig", token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetAssetManagerPoolConfig(&_BalancerV2StablePool.TransactOpts, token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetAssetManagerPoolConfig(&_BalancerV2StablePool.TransactOpts, token, poolConfig)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetPaused(&_BalancerV2StablePool.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetPaused(&_BalancerV2StablePool.TransactOpts, paused)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetSwapFeePercentage(&_BalancerV2StablePool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.SetSwapFeePercentage(&_BalancerV2StablePool.TransactOpts, swapFeePercentage)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) StartAmplificationParameterUpdate(opts *bind.TransactOpts, rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "startAmplificationParameterUpdate", rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.StartAmplificationParameterUpdate(&_BalancerV2StablePool.TransactOpts, rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.StartAmplificationParameterUpdate(&_BalancerV2StablePool.TransactOpts, rawEndValue, endTime)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) StopAmplificationParameterUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "stopAmplificationParameterUpdate")
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerV2StablePool *BalancerV2StablePoolSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.StopAmplificationParameterUpdate(&_BalancerV2StablePool.TransactOpts)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.StopAmplificationParameterUpdate(&_BalancerV2StablePool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Transfer(&_BalancerV2StablePool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.Transfer(&_BalancerV2StablePool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.TransferFrom(&_BalancerV2StablePool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerV2StablePool *BalancerV2StablePoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerV2StablePool.Contract.TransferFrom(&_BalancerV2StablePool.TransactOpts, sender, recipient, amount)
}

// BalancerV2StablePoolAmpUpdateStartedIterator is returned from FilterAmpUpdateStarted and is used to iterate over the raw logs and unpacked data for AmpUpdateStarted events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolAmpUpdateStartedIterator struct {
	Event *BalancerV2StablePoolAmpUpdateStarted // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolAmpUpdateStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolAmpUpdateStarted)
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
		it.Event = new(BalancerV2StablePoolAmpUpdateStarted)
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
func (it *BalancerV2StablePoolAmpUpdateStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolAmpUpdateStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolAmpUpdateStarted represents a AmpUpdateStarted event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolAmpUpdateStarted struct {
	StartValue *big.Int
	EndValue   *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAmpUpdateStarted is a free log retrieval operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterAmpUpdateStarted(opts *bind.FilterOpts) (*BalancerV2StablePoolAmpUpdateStartedIterator, error) {

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "AmpUpdateStarted")
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolAmpUpdateStartedIterator{contract: _BalancerV2StablePool.contract, event: "AmpUpdateStarted", logs: logs, sub: sub}, nil
}

// WatchAmpUpdateStarted is a free log subscription operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchAmpUpdateStarted(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolAmpUpdateStarted) (event.Subscription, error) {

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "AmpUpdateStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolAmpUpdateStarted)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "AmpUpdateStarted", log); err != nil {
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

// ParseAmpUpdateStarted is a log parse operation binding the contract event 0x1835882ee7a34ac194f717a35e09bb1d24c82a3b9d854ab6c9749525b714cdf2.
//
// Solidity: event AmpUpdateStarted(uint256 startValue, uint256 endValue, uint256 startTime, uint256 endTime)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParseAmpUpdateStarted(log types.Log) (*BalancerV2StablePoolAmpUpdateStarted, error) {
	event := new(BalancerV2StablePoolAmpUpdateStarted)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "AmpUpdateStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2StablePoolAmpUpdateStoppedIterator is returned from FilterAmpUpdateStopped and is used to iterate over the raw logs and unpacked data for AmpUpdateStopped events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolAmpUpdateStoppedIterator struct {
	Event *BalancerV2StablePoolAmpUpdateStopped // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolAmpUpdateStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolAmpUpdateStopped)
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
		it.Event = new(BalancerV2StablePoolAmpUpdateStopped)
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
func (it *BalancerV2StablePoolAmpUpdateStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolAmpUpdateStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolAmpUpdateStopped represents a AmpUpdateStopped event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolAmpUpdateStopped struct {
	CurrentValue *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAmpUpdateStopped is a free log retrieval operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterAmpUpdateStopped(opts *bind.FilterOpts) (*BalancerV2StablePoolAmpUpdateStoppedIterator, error) {

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "AmpUpdateStopped")
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolAmpUpdateStoppedIterator{contract: _BalancerV2StablePool.contract, event: "AmpUpdateStopped", logs: logs, sub: sub}, nil
}

// WatchAmpUpdateStopped is a free log subscription operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchAmpUpdateStopped(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolAmpUpdateStopped) (event.Subscription, error) {

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "AmpUpdateStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolAmpUpdateStopped)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "AmpUpdateStopped", log); err != nil {
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

// ParseAmpUpdateStopped is a log parse operation binding the contract event 0xa0d01593e47e69d07e0ccd87bece09411e07dd1ed40ca8f2e7af2976542a0233.
//
// Solidity: event AmpUpdateStopped(uint256 currentValue)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParseAmpUpdateStopped(log types.Log) (*BalancerV2StablePoolAmpUpdateStopped, error) {
	event := new(BalancerV2StablePoolAmpUpdateStopped)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "AmpUpdateStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2StablePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolApprovalIterator struct {
	Event *BalancerV2StablePoolApproval // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolApproval)
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
		it.Event = new(BalancerV2StablePoolApproval)
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
func (it *BalancerV2StablePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolApproval represents a Approval event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BalancerV2StablePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolApprovalIterator{contract: _BalancerV2StablePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolApproval)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParseApproval(log types.Log) (*BalancerV2StablePoolApproval, error) {
	event := new(BalancerV2StablePoolApproval)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2StablePoolPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolPausedStateChangedIterator struct {
	Event *BalancerV2StablePoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolPausedStateChanged)
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
		it.Event = new(BalancerV2StablePoolPausedStateChanged)
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
func (it *BalancerV2StablePoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolPausedStateChanged represents a PausedStateChanged event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BalancerV2StablePoolPausedStateChangedIterator, error) {

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolPausedStateChangedIterator{contract: _BalancerV2StablePool.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolPausedStateChanged)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParsePausedStateChanged(log types.Log) (*BalancerV2StablePoolPausedStateChanged, error) {
	event := new(BalancerV2StablePoolPausedStateChanged)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2StablePoolSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolSwapFeePercentageChangedIterator struct {
	Event *BalancerV2StablePoolSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolSwapFeePercentageChanged)
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
		it.Event = new(BalancerV2StablePoolSwapFeePercentageChanged)
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
func (it *BalancerV2StablePoolSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BalancerV2StablePoolSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolSwapFeePercentageChangedIterator{contract: _BalancerV2StablePool.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolSwapFeePercentageChanged)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BalancerV2StablePoolSwapFeePercentageChanged, error) {
	event := new(BalancerV2StablePoolSwapFeePercentageChanged)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerV2StablePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolTransferIterator struct {
	Event *BalancerV2StablePoolTransfer // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolTransfer)
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
		it.Event = new(BalancerV2StablePoolTransfer)
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
func (it *BalancerV2StablePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolTransfer represents a Transfer event raised by the BalancerV2StablePool contract.
type BalancerV2StablePoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BalancerV2StablePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerV2StablePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolTransferIterator{contract: _BalancerV2StablePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerV2StablePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolTransfer)
				if err := _BalancerV2StablePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BalancerV2StablePool *BalancerV2StablePoolFilterer) ParseTransfer(log types.Log) (*BalancerV2StablePoolTransfer, error) {
	event := new(BalancerV2StablePoolTransfer)
	if err := _BalancerV2StablePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
