// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_Factory *FactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getPauseConfiguration")

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
func (_Factory *FactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _Factory.Contract.GetPauseConfiguration(&_Factory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_Factory *FactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _Factory.Contract.GetPauseConfiguration(&_Factory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Factory *FactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Factory *FactorySession) GetVault() (common.Address, error) {
	return _Factory.Contract.GetVault(&_Factory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Factory *FactoryCallerSession) GetVault() (common.Address, error) {
	return _Factory.Contract.GetVault(&_Factory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Factory *FactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Factory *FactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _Factory.Contract.IsPoolFromFactory(&_Factory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Factory *FactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _Factory.Contract.IsPoolFromFactory(&_Factory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_Factory *FactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "create", name, symbol, tokens, weights, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_Factory *FactorySession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Create(&_Factory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_Factory *FactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Create(&_Factory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, owner)
}

// FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Factory contract.
type FactoryPoolCreatedIterator struct {
	Event *FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPoolCreated)
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
		it.Event = new(FactoryPoolCreated)
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
func (it *FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPoolCreated represents a PoolCreated event raised by the Factory contract.
type FactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_Factory *FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*FactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &FactoryPoolCreatedIterator{contract: _Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_Factory *FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *FactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPoolCreated)
				if err := _Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
func (_Factory *FactoryFilterer) ParsePoolCreated(log types.Log) (*FactoryPoolCreated, error) {
	event := new(FactoryPoolCreated)
	if err := _Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
