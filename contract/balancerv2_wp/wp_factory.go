// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weightedpool

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

// WpFactoryMetaData contains all meta data concerning the WpFactory contract.
var WpFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WpFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WpFactoryMetaData.ABI instead.
var WpFactoryABI = WpFactoryMetaData.ABI

// WpFactory is an auto generated Go binding around an Ethereum contract.
type WpFactory struct {
	WpFactoryCaller     // Read-only binding to the contract
	WpFactoryTransactor // Write-only binding to the contract
	WpFactoryFilterer   // Log filterer for contract events
}

// WpFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WpFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WpFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WpFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WpFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WpFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WpFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WpFactorySession struct {
	Contract     *WpFactory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WpFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WpFactoryCallerSession struct {
	Contract *WpFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WpFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WpFactoryTransactorSession struct {
	Contract     *WpFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WpFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WpFactoryRaw struct {
	Contract *WpFactory // Generic contract binding to access the raw methods on
}

// WpFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WpFactoryCallerRaw struct {
	Contract *WpFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WpFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WpFactoryTransactorRaw struct {
	Contract *WpFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWpFactory creates a new instance of WpFactory, bound to a specific deployed contract.
func NewWpFactory(address common.Address, backend bind.ContractBackend) (*WpFactory, error) {
	contract, err := bindWpFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WpFactory{WpFactoryCaller: WpFactoryCaller{contract: contract}, WpFactoryTransactor: WpFactoryTransactor{contract: contract}, WpFactoryFilterer: WpFactoryFilterer{contract: contract}}, nil
}

// NewWpFactoryCaller creates a new read-only instance of WpFactory, bound to a specific deployed contract.
func NewWpFactoryCaller(address common.Address, caller bind.ContractCaller) (*WpFactoryCaller, error) {
	contract, err := bindWpFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WpFactoryCaller{contract: contract}, nil
}

// NewWpFactoryTransactor creates a new write-only instance of WpFactory, bound to a specific deployed contract.
func NewWpFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WpFactoryTransactor, error) {
	contract, err := bindWpFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WpFactoryTransactor{contract: contract}, nil
}

// NewWpFactoryFilterer creates a new log filterer instance of WpFactory, bound to a specific deployed contract.
func NewWpFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WpFactoryFilterer, error) {
	contract, err := bindWpFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WpFactoryFilterer{contract: contract}, nil
}

// bindWpFactory binds a generic wrapper to an already deployed contract.
func bindWpFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WpFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WpFactory *WpFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WpFactory.Contract.WpFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WpFactory *WpFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WpFactory.Contract.WpFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WpFactory *WpFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WpFactory.Contract.WpFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WpFactory *WpFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WpFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WpFactory *WpFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WpFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WpFactory *WpFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WpFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_WpFactory *WpFactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _WpFactory.contract.Call(opts, &out, "getPauseConfiguration")

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
func (_WpFactory *WpFactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _WpFactory.Contract.GetPauseConfiguration(&_WpFactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_WpFactory *WpFactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _WpFactory.Contract.GetPauseConfiguration(&_WpFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_WpFactory *WpFactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WpFactory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_WpFactory *WpFactorySession) GetVault() (common.Address, error) {
	return _WpFactory.Contract.GetVault(&_WpFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_WpFactory *WpFactoryCallerSession) GetVault() (common.Address, error) {
	return _WpFactory.Contract.GetVault(&_WpFactory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_WpFactory *WpFactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _WpFactory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_WpFactory *WpFactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _WpFactory.Contract.IsPoolFromFactory(&_WpFactory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_WpFactory *WpFactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _WpFactory.Contract.IsPoolFromFactory(&_WpFactory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_WpFactory *WpFactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _WpFactory.contract.Transact(opts, "create", name, symbol, tokens, weights, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_WpFactory *WpFactorySession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _WpFactory.Contract.Create(&_WpFactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, owner)
}

// Create is a paid mutator transaction binding the contract method 0xfbce0393.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, address owner) returns(address)
func (_WpFactory *WpFactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, owner common.Address) (*types.Transaction, error) {
	return _WpFactory.Contract.Create(&_WpFactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, owner)
}

// WpFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the WpFactory contract.
type WpFactoryPoolCreatedIterator struct {
	Event *WpFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *WpFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WpFactoryPoolCreated)
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
		it.Event = new(WpFactoryPoolCreated)
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
func (it *WpFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WpFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WpFactoryPoolCreated represents a PoolCreated event raised by the WpFactory contract.
type WpFactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_WpFactory *WpFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*WpFactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WpFactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &WpFactoryPoolCreatedIterator{contract: _WpFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_WpFactory *WpFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *WpFactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WpFactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WpFactoryPoolCreated)
				if err := _WpFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
func (_WpFactory *WpFactoryFilterer) ParsePoolCreated(log types.Log) (*WpFactoryPoolCreated, error) {
	event := new(WpFactoryPoolCreated)
	if err := _WpFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
