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

// BalancerV2StablePoolFactoryMetaData contains all meta data concerning the BalancerV2StablePoolFactory contract.
var BalancerV2StablePoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amplificationParameter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerV2StablePoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2StablePoolFactoryMetaData.ABI instead.
var BalancerV2StablePoolFactoryABI = BalancerV2StablePoolFactoryMetaData.ABI

// BalancerV2StablePoolFactory is an auto generated Go binding around an Ethereum contract.
type BalancerV2StablePoolFactory struct {
	BalancerV2StablePoolFactoryCaller     // Read-only binding to the contract
	BalancerV2StablePoolFactoryTransactor // Write-only binding to the contract
	BalancerV2StablePoolFactoryFilterer   // Log filterer for contract events
}

// BalancerV2StablePoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2StablePoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2StablePoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2StablePoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2StablePoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2StablePoolFactorySession struct {
	Contract     *BalancerV2StablePoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BalancerV2StablePoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2StablePoolFactoryCallerSession struct {
	Contract *BalancerV2StablePoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BalancerV2StablePoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2StablePoolFactoryTransactorSession struct {
	Contract     *BalancerV2StablePoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BalancerV2StablePoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2StablePoolFactoryRaw struct {
	Contract *BalancerV2StablePoolFactory // Generic contract binding to access the raw methods on
}

// BalancerV2StablePoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2StablePoolFactoryCallerRaw struct {
	Contract *BalancerV2StablePoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2StablePoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2StablePoolFactoryTransactorRaw struct {
	Contract *BalancerV2StablePoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2StablePoolFactory creates a new instance of BalancerV2StablePoolFactory, bound to a specific deployed contract.
func NewBalancerV2StablePoolFactory(address common.Address, backend bind.ContractBackend) (*BalancerV2StablePoolFactory, error) {
	contract, err := bindBalancerV2StablePoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFactory{BalancerV2StablePoolFactoryCaller: BalancerV2StablePoolFactoryCaller{contract: contract}, BalancerV2StablePoolFactoryTransactor: BalancerV2StablePoolFactoryTransactor{contract: contract}, BalancerV2StablePoolFactoryFilterer: BalancerV2StablePoolFactoryFilterer{contract: contract}}, nil
}

// NewBalancerV2StablePoolFactoryCaller creates a new read-only instance of BalancerV2StablePoolFactory, bound to a specific deployed contract.
func NewBalancerV2StablePoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2StablePoolFactoryCaller, error) {
	contract, err := bindBalancerV2StablePoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFactoryCaller{contract: contract}, nil
}

// NewBalancerV2StablePoolFactoryTransactor creates a new write-only instance of BalancerV2StablePoolFactory, bound to a specific deployed contract.
func NewBalancerV2StablePoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2StablePoolFactoryTransactor, error) {
	contract, err := bindBalancerV2StablePoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFactoryTransactor{contract: contract}, nil
}

// NewBalancerV2StablePoolFactoryFilterer creates a new log filterer instance of BalancerV2StablePoolFactory, bound to a specific deployed contract.
func NewBalancerV2StablePoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2StablePoolFactoryFilterer, error) {
	contract, err := bindBalancerV2StablePoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFactoryFilterer{contract: contract}, nil
}

// bindBalancerV2StablePoolFactory binds a generic wrapper to an already deployed contract.
func bindBalancerV2StablePoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerV2StablePoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2StablePoolFactory.Contract.BalancerV2StablePoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.BalancerV2StablePoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.BalancerV2StablePoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2StablePoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2StablePoolFactory.contract.Call(opts, &out, "getPauseConfiguration")

	outstruct := new(struct {
		PauseWindowDuration  *big.Int
		BufferPeriodDuration *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PauseWindowDuration = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BalancerV2StablePoolFactory.Contract.GetPauseConfiguration(&_BalancerV2StablePoolFactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BalancerV2StablePoolFactory.Contract.GetPauseConfiguration(&_BalancerV2StablePoolFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2StablePoolFactory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactorySession) GetVault() (common.Address, error) {
	return _BalancerV2StablePoolFactory.Contract.GetVault(&_BalancerV2StablePoolFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCallerSession) GetVault() (common.Address, error) {
	return _BalancerV2StablePoolFactory.Contract.GetVault(&_BalancerV2StablePoolFactory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerV2StablePoolFactory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BalancerV2StablePoolFactory.Contract.IsPoolFromFactory(&_BalancerV2StablePoolFactory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BalancerV2StablePoolFactory.Contract.IsPoolFromFactory(&_BalancerV2StablePoolFactory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0x7932c7f3.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256 amplificationParameter, uint256 swapFeePercentage, address owner) returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, amplificationParameter *big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.contract.Transact(opts, "create", name, symbol, tokens, amplificationParameter, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0x7932c7f3.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256 amplificationParameter, uint256 swapFeePercentage, address owner) returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactorySession) Create(name string, symbol string, tokens []common.Address, amplificationParameter *big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.Create(&_BalancerV2StablePoolFactory.TransactOpts, name, symbol, tokens, amplificationParameter, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0x7932c7f3.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256 amplificationParameter, uint256 swapFeePercentage, address owner) returns(address)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, amplificationParameter *big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2StablePoolFactory.Contract.Create(&_BalancerV2StablePoolFactory.TransactOpts, name, symbol, tokens, amplificationParameter, swapFeePercentage, owner)
}

// BalancerV2StablePoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the BalancerV2StablePoolFactory contract.
type BalancerV2StablePoolFactoryPoolCreatedIterator struct {
	Event *BalancerV2StablePoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *BalancerV2StablePoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2StablePoolFactoryPoolCreated)
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
		it.Event = new(BalancerV2StablePoolFactoryPoolCreated)
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
func (it *BalancerV2StablePoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2StablePoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2StablePoolFactoryPoolCreated represents a PoolCreated event raised by the BalancerV2StablePoolFactory contract.
type BalancerV2StablePoolFactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*BalancerV2StablePoolFactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerV2StablePoolFactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2StablePoolFactoryPoolCreatedIterator{contract: _BalancerV2StablePoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *BalancerV2StablePoolFactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerV2StablePoolFactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2StablePoolFactoryPoolCreated)
				if err := _BalancerV2StablePoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BalancerV2StablePoolFactory *BalancerV2StablePoolFactoryFilterer) ParsePoolCreated(log types.Log) (*BalancerV2StablePoolFactoryPoolCreated, error) {
	event := new(BalancerV2StablePoolFactoryPoolCreated)
	if err := _BalancerV2StablePoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
