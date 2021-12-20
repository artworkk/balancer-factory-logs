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

// BalancerV2Wp2TokensFactoryMetaData contains all meta data concerning the BalancerV2Wp2TokensFactory contract.
var BalancerV2Wp2TokensFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerV2Wp2TokensFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2Wp2TokensFactoryMetaData.ABI instead.
var BalancerV2Wp2TokensFactoryABI = BalancerV2Wp2TokensFactoryMetaData.ABI

// BalancerV2Wp2TokensFactory is an auto generated Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactory struct {
	BalancerV2Wp2TokensFactoryCaller     // Read-only binding to the contract
	BalancerV2Wp2TokensFactoryTransactor // Write-only binding to the contract
	BalancerV2Wp2TokensFactoryFilterer   // Log filterer for contract events
}

// BalancerV2Wp2TokensFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2Wp2TokensFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2Wp2TokensFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2Wp2TokensFactorySession struct {
	Contract     *BalancerV2Wp2TokensFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BalancerV2Wp2TokensFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2Wp2TokensFactoryCallerSession struct {
	Contract *BalancerV2Wp2TokensFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// BalancerV2Wp2TokensFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2Wp2TokensFactoryTransactorSession struct {
	Contract     *BalancerV2Wp2TokensFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// BalancerV2Wp2TokensFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactoryRaw struct {
	Contract *BalancerV2Wp2TokensFactory // Generic contract binding to access the raw methods on
}

// BalancerV2Wp2TokensFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactoryCallerRaw struct {
	Contract *BalancerV2Wp2TokensFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2Wp2TokensFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2Wp2TokensFactoryTransactorRaw struct {
	Contract *BalancerV2Wp2TokensFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2Wp2TokensFactory creates a new instance of BalancerV2Wp2TokensFactory, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensFactory(address common.Address, backend bind.ContractBackend) (*BalancerV2Wp2TokensFactory, error) {
	contract, err := bindBalancerV2Wp2TokensFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFactory{BalancerV2Wp2TokensFactoryCaller: BalancerV2Wp2TokensFactoryCaller{contract: contract}, BalancerV2Wp2TokensFactoryTransactor: BalancerV2Wp2TokensFactoryTransactor{contract: contract}, BalancerV2Wp2TokensFactoryFilterer: BalancerV2Wp2TokensFactoryFilterer{contract: contract}}, nil
}

// NewBalancerV2Wp2TokensFactoryCaller creates a new read-only instance of BalancerV2Wp2TokensFactory, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensFactoryCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2Wp2TokensFactoryCaller, error) {
	contract, err := bindBalancerV2Wp2TokensFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFactoryCaller{contract: contract}, nil
}

// NewBalancerV2Wp2TokensFactoryTransactor creates a new write-only instance of BalancerV2Wp2TokensFactory, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2Wp2TokensFactoryTransactor, error) {
	contract, err := bindBalancerV2Wp2TokensFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFactoryTransactor{contract: contract}, nil
}

// NewBalancerV2Wp2TokensFactoryFilterer creates a new log filterer instance of BalancerV2Wp2TokensFactory, bound to a specific deployed contract.
func NewBalancerV2Wp2TokensFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2Wp2TokensFactoryFilterer, error) {
	contract, err := bindBalancerV2Wp2TokensFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFactoryFilterer{contract: contract}, nil
}

// bindBalancerV2Wp2TokensFactory binds a generic wrapper to an already deployed contract.
func bindBalancerV2Wp2TokensFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerV2Wp2TokensFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Wp2TokensFactory.Contract.BalancerV2Wp2TokensFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.BalancerV2Wp2TokensFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.BalancerV2Wp2TokensFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2Wp2TokensFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _BalancerV2Wp2TokensFactory.contract.Call(opts, &out, "getPauseConfiguration")

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
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BalancerV2Wp2TokensFactory.Contract.GetPauseConfiguration(&_BalancerV2Wp2TokensFactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _BalancerV2Wp2TokensFactory.Contract.GetPauseConfiguration(&_BalancerV2Wp2TokensFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerV2Wp2TokensFactory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactorySession) GetVault() (common.Address, error) {
	return _BalancerV2Wp2TokensFactory.Contract.GetVault(&_BalancerV2Wp2TokensFactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCallerSession) GetVault() (common.Address, error) {
	return _BalancerV2Wp2TokensFactory.Contract.GetVault(&_BalancerV2Wp2TokensFactory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerV2Wp2TokensFactory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BalancerV2Wp2TokensFactory.Contract.IsPoolFromFactory(&_BalancerV2Wp2TokensFactory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _BalancerV2Wp2TokensFactory.Contract.IsPoolFromFactory(&_BalancerV2Wp2TokensFactory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.contract.Transact(opts, "create", name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactorySession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.Create(&_BalancerV2Wp2TokensFactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _BalancerV2Wp2TokensFactory.Contract.Create(&_BalancerV2Wp2TokensFactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// BalancerV2Wp2TokensFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the BalancerV2Wp2TokensFactory contract.
type BalancerV2Wp2TokensFactoryPoolCreatedIterator struct {
	Event *BalancerV2Wp2TokensFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *BalancerV2Wp2TokensFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2Wp2TokensFactoryPoolCreated)
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
		it.Event = new(BalancerV2Wp2TokensFactoryPoolCreated)
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
func (it *BalancerV2Wp2TokensFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2Wp2TokensFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2Wp2TokensFactoryPoolCreated represents a PoolCreated event raised by the BalancerV2Wp2TokensFactory contract.
type BalancerV2Wp2TokensFactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*BalancerV2Wp2TokensFactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerV2Wp2TokensFactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &BalancerV2Wp2TokensFactoryPoolCreatedIterator{contract: _BalancerV2Wp2TokensFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *BalancerV2Wp2TokensFactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BalancerV2Wp2TokensFactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2Wp2TokensFactoryPoolCreated)
				if err := _BalancerV2Wp2TokensFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
func (_BalancerV2Wp2TokensFactory *BalancerV2Wp2TokensFactoryFilterer) ParsePoolCreated(log types.Log) (*BalancerV2Wp2TokensFactoryPoolCreated, error) {
	event := new(BalancerV2Wp2TokensFactoryPoolCreated)
	if err := _BalancerV2Wp2TokensFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
