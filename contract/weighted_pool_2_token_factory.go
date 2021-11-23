// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Balancerv2wp2tokenfactoryMetaData contains all meta data concerning the Balancerv2wp2tokenfactory contract.
var Balancerv2wp2tokenfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolFromFactory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Balancerv2wp2tokenfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Balancerv2wp2tokenfactoryMetaData.ABI instead.
var Balancerv2wp2tokenfactoryABI = Balancerv2wp2tokenfactoryMetaData.ABI

// Balancerv2wp2tokenfactory is an auto generated Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactory struct {
	Balancerv2wp2tokenfactoryCaller     // Read-only binding to the contract
	Balancerv2wp2tokenfactoryTransactor // Write-only binding to the contract
	Balancerv2wp2tokenfactoryFilterer   // Log filterer for contract events
}

// Balancerv2wp2tokenfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2wp2tokenfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2wp2tokenfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Balancerv2wp2tokenfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2wp2tokenfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Balancerv2wp2tokenfactorySession struct {
	Contract     *Balancerv2wp2tokenfactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Balancerv2wp2tokenfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Balancerv2wp2tokenfactoryCallerSession struct {
	Contract *Balancerv2wp2tokenfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Balancerv2wp2tokenfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Balancerv2wp2tokenfactoryTransactorSession struct {
	Contract     *Balancerv2wp2tokenfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Balancerv2wp2tokenfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactoryRaw struct {
	Contract *Balancerv2wp2tokenfactory // Generic contract binding to access the raw methods on
}

// Balancerv2wp2tokenfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactoryCallerRaw struct {
	Contract *Balancerv2wp2tokenfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// Balancerv2wp2tokenfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Balancerv2wp2tokenfactoryTransactorRaw struct {
	Contract *Balancerv2wp2tokenfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerv2wp2tokenfactory creates a new instance of Balancerv2wp2tokenfactory, bound to a specific deployed contract.
func NewBalancerv2wp2tokenfactory(address common.Address, backend bind.ContractBackend) (*Balancerv2wp2tokenfactory, error) {
	contract, err := bindBalancerv2wp2tokenfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerv2wp2tokenfactory{Balancerv2wp2tokenfactoryCaller: Balancerv2wp2tokenfactoryCaller{contract: contract}, Balancerv2wp2tokenfactoryTransactor: Balancerv2wp2tokenfactoryTransactor{contract: contract}, Balancerv2wp2tokenfactoryFilterer: Balancerv2wp2tokenfactoryFilterer{contract: contract}}, nil
}

// NewBalancerv2wp2tokenfactoryCaller creates a new read-only instance of Balancerv2wp2tokenfactory, bound to a specific deployed contract.
func NewBalancerv2wp2tokenfactoryCaller(address common.Address, caller bind.ContractCaller) (*Balancerv2wp2tokenfactoryCaller, error) {
	contract, err := bindBalancerv2wp2tokenfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2wp2tokenfactoryCaller{contract: contract}, nil
}

// NewBalancerv2wp2tokenfactoryTransactor creates a new write-only instance of Balancerv2wp2tokenfactory, bound to a specific deployed contract.
func NewBalancerv2wp2tokenfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Balancerv2wp2tokenfactoryTransactor, error) {
	contract, err := bindBalancerv2wp2tokenfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2wp2tokenfactoryTransactor{contract: contract}, nil
}

// NewBalancerv2wp2tokenfactoryFilterer creates a new log filterer instance of Balancerv2wp2tokenfactory, bound to a specific deployed contract.
func NewBalancerv2wp2tokenfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Balancerv2wp2tokenfactoryFilterer, error) {
	contract, err := bindBalancerv2wp2tokenfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Balancerv2wp2tokenfactoryFilterer{contract: contract}, nil
}

// bindBalancerv2wp2tokenfactory binds a generic wrapper to an already deployed contract.
func bindBalancerv2wp2tokenfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Balancerv2wp2tokenfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2wp2tokenfactory.Contract.Balancerv2wp2tokenfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.Balancerv2wp2tokenfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.Balancerv2wp2tokenfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2wp2tokenfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.contract.Transact(opts, method, params...)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCaller) GetPauseConfiguration(opts *bind.CallOpts) (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	var out []interface{}
	err := _Balancerv2wp2tokenfactory.contract.Call(opts, &out, "getPauseConfiguration")

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
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactorySession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _Balancerv2wp2tokenfactory.Contract.GetPauseConfiguration(&_Balancerv2wp2tokenfactory.CallOpts)
}

// GetPauseConfiguration is a free data retrieval call binding the contract method 0x2da47c40.
//
// Solidity: function getPauseConfiguration() view returns(uint256 pauseWindowDuration, uint256 bufferPeriodDuration)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCallerSession) GetPauseConfiguration() (struct {
	PauseWindowDuration  *big.Int
	BufferPeriodDuration *big.Int
}, error) {
	return _Balancerv2wp2tokenfactory.Contract.GetPauseConfiguration(&_Balancerv2wp2tokenfactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2wp2tokenfactory.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactorySession) GetVault() (common.Address, error) {
	return _Balancerv2wp2tokenfactory.Contract.GetVault(&_Balancerv2wp2tokenfactory.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCallerSession) GetVault() (common.Address, error) {
	return _Balancerv2wp2tokenfactory.Contract.GetVault(&_Balancerv2wp2tokenfactory.CallOpts)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCaller) IsPoolFromFactory(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Balancerv2wp2tokenfactory.contract.Call(opts, &out, "isPoolFromFactory", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactorySession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _Balancerv2wp2tokenfactory.Contract.IsPoolFromFactory(&_Balancerv2wp2tokenfactory.CallOpts, pool)
}

// IsPoolFromFactory is a free data retrieval call binding the contract method 0x6634b753.
//
// Solidity: function isPoolFromFactory(address pool) view returns(bool)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryCallerSession) IsPoolFromFactory(pool common.Address) (bool, error) {
	return _Balancerv2wp2tokenfactory.Contract.IsPoolFromFactory(&_Balancerv2wp2tokenfactory.CallOpts, pool)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.contract.Transact(opts, "create", name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactorySession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.Create(&_Balancerv2wp2tokenfactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// Create is a paid mutator transaction binding the contract method 0x1596019b.
//
// Solidity: function create(string name, string symbol, address[] tokens, uint256[] weights, uint256 swapFeePercentage, bool oracleEnabled, address owner) returns(address)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryTransactorSession) Create(name string, symbol string, tokens []common.Address, weights []*big.Int, swapFeePercentage *big.Int, oracleEnabled bool, owner common.Address) (*types.Transaction, error) {
	return _Balancerv2wp2tokenfactory.Contract.Create(&_Balancerv2wp2tokenfactory.TransactOpts, name, symbol, tokens, weights, swapFeePercentage, oracleEnabled, owner)
}

// Balancerv2wp2tokenfactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Balancerv2wp2tokenfactory contract.
type Balancerv2wp2tokenfactoryPoolCreatedIterator struct {
	Event *Balancerv2wp2tokenfactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *Balancerv2wp2tokenfactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2wp2tokenfactoryPoolCreated)
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
		it.Event = new(Balancerv2wp2tokenfactoryPoolCreated)
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
func (it *Balancerv2wp2tokenfactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2wp2tokenfactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2wp2tokenfactoryPoolCreated represents a PoolCreated event raised by the Balancerv2wp2tokenfactory contract.
type Balancerv2wp2tokenfactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*Balancerv2wp2tokenfactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerv2wp2tokenfactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2wp2tokenfactoryPoolCreatedIterator{contract: _Balancerv2wp2tokenfactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address indexed pool)
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *Balancerv2wp2tokenfactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Balancerv2wp2tokenfactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2wp2tokenfactoryPoolCreated)
				if err := _Balancerv2wp2tokenfactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
func (_Balancerv2wp2tokenfactory *Balancerv2wp2tokenfactoryFilterer) ParsePoolCreated(log types.Log) (*Balancerv2wp2tokenfactoryPoolCreated, error) {
	event := new(Balancerv2wp2tokenfactoryPoolCreated)
	if err := _Balancerv2wp2tokenfactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
