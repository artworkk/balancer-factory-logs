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

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IVaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// IVaultExitPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultExitPoolRequest struct {
	Assets            []common.Address
	MinAmountsOut     []*big.Int
	UserData          []byte
	ToInternalBalance bool
}

// IVaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
type IVaultFundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// IVaultJoinPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultJoinPoolRequest struct {
	Assets              []common.Address
	MaxAmountsIn        []*big.Int
	UserData            []byte
	FromInternalBalance bool
}

// IVaultPoolBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultPoolBalanceOp struct {
	Kind   uint8
	PoolId [32]byte
	Token  common.Address
	Amount *big.Int
}

// IVaultSingleSwap is an auto generated low-level Go binding around an user-defined struct.
type IVaultSingleSwap struct {
	PoolId   [32]byte
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
	Amount   *big.Int
	UserData []byte
}

// IVaultUserBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultUserBalanceOp struct {
	Kind      uint8
	Asset     common.Address
	Amount    *big.Int
	Sender    common.Address
	Recipient common.Address
}

// Balancerv2thevaultMetaData contains all meta data concerning the Balancerv2thevault contract.
var Balancerv2thevaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExternalBalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"}],\"name\":\"InternalBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"deltas\",\"type\":\"int256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"protocolFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"PoolBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cashDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"managedDelta\",\"type\":\"int256\"}],\"name\":\"PoolBalanceManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"RelayerApprovalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"batchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"assetDeltas\",\"type\":\"int256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"deregisterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.ExitPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getInternalBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"managed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"hasApprovedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.JoinPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.PoolBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.PoolBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"managePoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.UserBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIVault.UserBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"manageUserBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"queryBatchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"registerPool\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"setAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setRelayerApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.SingleSwap\",\"name\":\"singleSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculated\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Balancerv2thevaultABI is the input ABI used to generate the binding from.
// Deprecated: Use Balancerv2thevaultMetaData.ABI instead.
var Balancerv2thevaultABI = Balancerv2thevaultMetaData.ABI

// Balancerv2thevault is an auto generated Go binding around an Ethereum contract.
type Balancerv2thevault struct {
	Balancerv2thevaultCaller     // Read-only binding to the contract
	Balancerv2thevaultTransactor // Write-only binding to the contract
	Balancerv2thevaultFilterer   // Log filterer for contract events
}

// Balancerv2thevaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type Balancerv2thevaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2thevaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Balancerv2thevaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2thevaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Balancerv2thevaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2thevaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Balancerv2thevaultSession struct {
	Contract     *Balancerv2thevault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Balancerv2thevaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Balancerv2thevaultCallerSession struct {
	Contract *Balancerv2thevaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Balancerv2thevaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Balancerv2thevaultTransactorSession struct {
	Contract     *Balancerv2thevaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Balancerv2thevaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type Balancerv2thevaultRaw struct {
	Contract *Balancerv2thevault // Generic contract binding to access the raw methods on
}

// Balancerv2thevaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Balancerv2thevaultCallerRaw struct {
	Contract *Balancerv2thevaultCaller // Generic read-only contract binding to access the raw methods on
}

// Balancerv2thevaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Balancerv2thevaultTransactorRaw struct {
	Contract *Balancerv2thevaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerv2thevault creates a new instance of Balancerv2thevault, bound to a specific deployed contract.
func NewBalancerv2thevault(address common.Address, backend bind.ContractBackend) (*Balancerv2thevault, error) {
	contract, err := bindBalancerv2thevault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevault{Balancerv2thevaultCaller: Balancerv2thevaultCaller{contract: contract}, Balancerv2thevaultTransactor: Balancerv2thevaultTransactor{contract: contract}, Balancerv2thevaultFilterer: Balancerv2thevaultFilterer{contract: contract}}, nil
}

// NewBalancerv2thevaultCaller creates a new read-only instance of Balancerv2thevault, bound to a specific deployed contract.
func NewBalancerv2thevaultCaller(address common.Address, caller bind.ContractCaller) (*Balancerv2thevaultCaller, error) {
	contract, err := bindBalancerv2thevault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultCaller{contract: contract}, nil
}

// NewBalancerv2thevaultTransactor creates a new write-only instance of Balancerv2thevault, bound to a specific deployed contract.
func NewBalancerv2thevaultTransactor(address common.Address, transactor bind.ContractTransactor) (*Balancerv2thevaultTransactor, error) {
	contract, err := bindBalancerv2thevault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultTransactor{contract: contract}, nil
}

// NewBalancerv2thevaultFilterer creates a new log filterer instance of Balancerv2thevault, bound to a specific deployed contract.
func NewBalancerv2thevaultFilterer(address common.Address, filterer bind.ContractFilterer) (*Balancerv2thevaultFilterer, error) {
	contract, err := bindBalancerv2thevault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultFilterer{contract: contract}, nil
}

// bindBalancerv2thevault binds a generic wrapper to an already deployed contract.
func bindBalancerv2thevault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Balancerv2thevaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2thevault *Balancerv2thevaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2thevault.Contract.Balancerv2thevaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2thevault *Balancerv2thevaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Balancerv2thevaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2thevault *Balancerv2thevaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Balancerv2thevaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2thevault *Balancerv2thevaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2thevault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2thevault *Balancerv2thevaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2thevault *Balancerv2thevaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultSession) WETH() (common.Address, error) {
	return _Balancerv2thevault.Contract.WETH(&_Balancerv2thevault.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) WETH() (common.Address, error) {
	return _Balancerv2thevault.Contract.WETH(&_Balancerv2thevault.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Balancerv2thevault.Contract.GetActionId(&_Balancerv2thevault.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Balancerv2thevault.Contract.GetActionId(&_Balancerv2thevault.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetAuthorizer() (common.Address, error) {
	return _Balancerv2thevault.Contract.GetAuthorizer(&_Balancerv2thevault.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetAuthorizer() (common.Address, error) {
	return _Balancerv2thevault.Contract.GetAuthorizer(&_Balancerv2thevault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetDomainSeparator() ([32]byte, error) {
	return _Balancerv2thevault.Contract.GetDomainSeparator(&_Balancerv2thevault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Balancerv2thevault.Contract.GetDomainSeparator(&_Balancerv2thevault.CallOpts)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetInternalBalance(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getInternalBalance", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balancerv2thevault.Contract.GetInternalBalance(&_Balancerv2thevault.CallOpts, user, tokens)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balancerv2thevault.Contract.GetInternalBalance(&_Balancerv2thevault.CallOpts, user, tokens)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetNextNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getNextNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Balancerv2thevault.Contract.GetNextNonce(&_Balancerv2thevault.CallOpts, user)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Balancerv2thevault.Contract.GetNextNonce(&_Balancerv2thevault.CallOpts, user)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getPausedState")

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
func (_Balancerv2thevault *Balancerv2thevaultSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Balancerv2thevault.Contract.GetPausedState(&_Balancerv2thevault.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Balancerv2thevault.Contract.GetPausedState(&_Balancerv2thevault.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, uint8, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getPool", poolId)

	if err != nil {
		return *new(common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Balancerv2thevault.Contract.GetPool(&_Balancerv2thevault.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Balancerv2thevault.Contract.GetPool(&_Balancerv2thevault.CallOpts, poolId)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetPoolTokenInfo(opts *bind.CallOpts, poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getPoolTokenInfo", poolId, token)

	outstruct := new(struct {
		Cash            *big.Int
		Managed         *big.Int
		LastChangeBlock *big.Int
		AssetManager    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Managed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AssetManager = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Balancerv2thevault.Contract.GetPoolTokenInfo(&_Balancerv2thevault.CallOpts, poolId, token)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Balancerv2thevault.Contract.GetPoolTokenInfo(&_Balancerv2thevault.CallOpts, poolId, token)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetPoolTokens(opts *bind.CallOpts, poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getPoolTokens", poolId)

	outstruct := new(struct {
		Tokens          []common.Address
		Balances        []*big.Int
		LastChangeBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Balances = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Balancerv2thevault.Contract.GetPoolTokens(&_Balancerv2thevault.CallOpts, poolId)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Balancerv2thevault.Contract.GetPoolTokens(&_Balancerv2thevault.CallOpts, poolId)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Balancerv2thevault.Contract.GetProtocolFeesCollector(&_Balancerv2thevault.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Balancerv2thevault.Contract.GetProtocolFeesCollector(&_Balancerv2thevault.CallOpts)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Balancerv2thevault *Balancerv2thevaultCaller) HasApprovedRelayer(opts *bind.CallOpts, user common.Address, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Balancerv2thevault.contract.Call(opts, &out, "hasApprovedRelayer", user, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Balancerv2thevault *Balancerv2thevaultSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Balancerv2thevault.Contract.HasApprovedRelayer(&_Balancerv2thevault.CallOpts, user, relayer)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Balancerv2thevault *Balancerv2thevaultCallerSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Balancerv2thevault.Contract.HasApprovedRelayer(&_Balancerv2thevault.CallOpts, user, relayer)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Balancerv2thevault *Balancerv2thevaultTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "batchSwap", kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Balancerv2thevault *Balancerv2thevaultSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.BatchSwap(&_Balancerv2thevault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.BatchSwap(&_Balancerv2thevault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) DeregisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "deregisterTokens", poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.DeregisterTokens(&_Balancerv2thevault.TransactOpts, poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.DeregisterTokens(&_Balancerv2thevault.TransactOpts, poolId, tokens)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) ExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "exitPool", poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ExitPool(&_Balancerv2thevault.TransactOpts, poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ExitPool(&_Balancerv2thevault.TransactOpts, poolId, sender, recipient, request)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) FlashLoan(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "flashLoan", recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.FlashLoan(&_Balancerv2thevault.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.FlashLoan(&_Balancerv2thevault.TransactOpts, recipient, tokens, amounts, userData)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) JoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "joinPool", poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.JoinPool(&_Balancerv2thevault.TransactOpts, poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.JoinPool(&_Balancerv2thevault.TransactOpts, poolId, sender, recipient, request)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) ManagePoolBalance(opts *bind.TransactOpts, ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "managePoolBalance", ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ManagePoolBalance(&_Balancerv2thevault.TransactOpts, ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ManagePoolBalance(&_Balancerv2thevault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) ManageUserBalance(opts *bind.TransactOpts, ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "manageUserBalance", ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ManageUserBalance(&_Balancerv2thevault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.ManageUserBalance(&_Balancerv2thevault.TransactOpts, ops)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Balancerv2thevault *Balancerv2thevaultTransactor) QueryBatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "queryBatchSwap", kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Balancerv2thevault *Balancerv2thevaultSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.QueryBatchSwap(&_Balancerv2thevault.TransactOpts, kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.QueryBatchSwap(&_Balancerv2thevault.TransactOpts, kind, swaps, assets, funds)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultTransactor) RegisterPool(opts *bind.TransactOpts, specialization uint8) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "registerPool", specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.RegisterPool(&_Balancerv2thevault.TransactOpts, specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.RegisterPool(&_Balancerv2thevault.TransactOpts, specialization)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) RegisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "registerTokens", poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.RegisterTokens(&_Balancerv2thevault.TransactOpts, poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.RegisterTokens(&_Balancerv2thevault.TransactOpts, poolId, tokens, assetManagers)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) SetAuthorizer(opts *bind.TransactOpts, newAuthorizer common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "setAuthorizer", newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetAuthorizer(&_Balancerv2thevault.TransactOpts, newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetAuthorizer(&_Balancerv2thevault.TransactOpts, newAuthorizer)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetPaused(&_Balancerv2thevault.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetPaused(&_Balancerv2thevault.TransactOpts, paused)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) SetRelayerApproval(opts *bind.TransactOpts, sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "setRelayerApproval", sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetRelayerApproval(&_Balancerv2thevault.TransactOpts, sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.SetRelayerApproval(&_Balancerv2thevault.TransactOpts, sender, relayer, approved)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Balancerv2thevault *Balancerv2thevaultTransactor) Swap(opts *bind.TransactOpts, singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.Transact(opts, "swap", singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Balancerv2thevault *Balancerv2thevaultSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Swap(&_Balancerv2thevault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Swap(&_Balancerv2thevault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2thevault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancerv2thevault *Balancerv2thevaultSession) Receive() (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Receive(&_Balancerv2thevault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancerv2thevault *Balancerv2thevaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Balancerv2thevault.Contract.Receive(&_Balancerv2thevault.TransactOpts)
}

// Balancerv2thevaultAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultAuthorizerChangedIterator struct {
	Event *Balancerv2thevaultAuthorizerChanged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultAuthorizerChanged)
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
		it.Event = new(Balancerv2thevaultAuthorizerChanged)
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
func (it *Balancerv2thevaultAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultAuthorizerChanged represents a AuthorizerChanged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*Balancerv2thevaultAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultAuthorizerChangedIterator{contract: _Balancerv2thevault.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultAuthorizerChanged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
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

// ParseAuthorizerChanged is a log parse operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseAuthorizerChanged(log types.Log) (*Balancerv2thevaultAuthorizerChanged, error) {
	event := new(Balancerv2thevaultAuthorizerChanged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultExternalBalanceTransferIterator is returned from FilterExternalBalanceTransfer and is used to iterate over the raw logs and unpacked data for ExternalBalanceTransfer events raised by the Balancerv2thevault contract.
type Balancerv2thevaultExternalBalanceTransferIterator struct {
	Event *Balancerv2thevaultExternalBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultExternalBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultExternalBalanceTransfer)
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
		it.Event = new(Balancerv2thevaultExternalBalanceTransfer)
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
func (it *Balancerv2thevaultExternalBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultExternalBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultExternalBalanceTransfer represents a ExternalBalanceTransfer event raised by the Balancerv2thevault contract.
type Balancerv2thevaultExternalBalanceTransfer struct {
	Token     common.Address
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExternalBalanceTransfer is a free log retrieval operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterExternalBalanceTransfer(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*Balancerv2thevaultExternalBalanceTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultExternalBalanceTransferIterator{contract: _Balancerv2thevault.contract, event: "ExternalBalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchExternalBalanceTransfer is a free log subscription operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchExternalBalanceTransfer(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultExternalBalanceTransfer, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultExternalBalanceTransfer)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
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

// ParseExternalBalanceTransfer is a log parse operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseExternalBalanceTransfer(log types.Log) (*Balancerv2thevaultExternalBalanceTransfer, error) {
	event := new(Balancerv2thevaultExternalBalanceTransfer)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Balancerv2thevault contract.
type Balancerv2thevaultFlashLoanIterator struct {
	Event *Balancerv2thevaultFlashLoan // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultFlashLoan)
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
		it.Event = new(Balancerv2thevaultFlashLoan)
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
func (it *Balancerv2thevaultFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultFlashLoan represents a FlashLoan event raised by the Balancerv2thevault contract.
type Balancerv2thevaultFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*Balancerv2thevaultFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultFlashLoanIterator{contract: _Balancerv2thevault.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultFlashLoan)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseFlashLoan(log types.Log) (*Balancerv2thevaultFlashLoan, error) {
	event := new(Balancerv2thevaultFlashLoan)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultInternalBalanceChangedIterator is returned from FilterInternalBalanceChanged and is used to iterate over the raw logs and unpacked data for InternalBalanceChanged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultInternalBalanceChangedIterator struct {
	Event *Balancerv2thevaultInternalBalanceChanged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultInternalBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultInternalBalanceChanged)
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
		it.Event = new(Balancerv2thevaultInternalBalanceChanged)
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
func (it *Balancerv2thevaultInternalBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultInternalBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultInternalBalanceChanged represents a InternalBalanceChanged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultInternalBalanceChanged struct {
	User  common.Address
	Token common.Address
	Delta *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInternalBalanceChanged is a free log retrieval operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterInternalBalanceChanged(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*Balancerv2thevaultInternalBalanceChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultInternalBalanceChangedIterator{contract: _Balancerv2thevault.contract, event: "InternalBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchInternalBalanceChanged is a free log subscription operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchInternalBalanceChanged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultInternalBalanceChanged, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultInternalBalanceChanged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
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

// ParseInternalBalanceChanged is a log parse operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseInternalBalanceChanged(log types.Log) (*Balancerv2thevaultInternalBalanceChanged, error) {
	event := new(Balancerv2thevaultInternalBalanceChanged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultPausedStateChangedIterator struct {
	Event *Balancerv2thevaultPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultPausedStateChanged)
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
		it.Event = new(Balancerv2thevaultPausedStateChanged)
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
func (it *Balancerv2thevaultPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultPausedStateChanged represents a PausedStateChanged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*Balancerv2thevaultPausedStateChangedIterator, error) {

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultPausedStateChangedIterator{contract: _Balancerv2thevault.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultPausedStateChanged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParsePausedStateChanged(log types.Log) (*Balancerv2thevaultPausedStateChanged, error) {
	event := new(Balancerv2thevaultPausedStateChanged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultPoolBalanceChangedIterator is returned from FilterPoolBalanceChanged and is used to iterate over the raw logs and unpacked data for PoolBalanceChanged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolBalanceChangedIterator struct {
	Event *Balancerv2thevaultPoolBalanceChanged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultPoolBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultPoolBalanceChanged)
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
		it.Event = new(Balancerv2thevaultPoolBalanceChanged)
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
func (it *Balancerv2thevaultPoolBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultPoolBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultPoolBalanceChanged represents a PoolBalanceChanged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolBalanceChanged struct {
	PoolId             [32]byte
	LiquidityProvider  common.Address
	Tokens             []common.Address
	Deltas             []*big.Int
	ProtocolFeeAmounts []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceChanged is a free log retrieval operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterPoolBalanceChanged(opts *bind.FilterOpts, poolId [][32]byte, liquidityProvider []common.Address) (*Balancerv2thevaultPoolBalanceChangedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultPoolBalanceChangedIterator{contract: _Balancerv2thevault.contract, event: "PoolBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceChanged is a free log subscription operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchPoolBalanceChanged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultPoolBalanceChanged, poolId [][32]byte, liquidityProvider []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultPoolBalanceChanged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
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

// ParsePoolBalanceChanged is a log parse operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParsePoolBalanceChanged(log types.Log) (*Balancerv2thevaultPoolBalanceChanged, error) {
	event := new(Balancerv2thevaultPoolBalanceChanged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultPoolBalanceManagedIterator is returned from FilterPoolBalanceManaged and is used to iterate over the raw logs and unpacked data for PoolBalanceManaged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolBalanceManagedIterator struct {
	Event *Balancerv2thevaultPoolBalanceManaged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultPoolBalanceManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultPoolBalanceManaged)
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
		it.Event = new(Balancerv2thevaultPoolBalanceManaged)
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
func (it *Balancerv2thevaultPoolBalanceManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultPoolBalanceManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultPoolBalanceManaged represents a PoolBalanceManaged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolBalanceManaged struct {
	PoolId       [32]byte
	AssetManager common.Address
	Token        common.Address
	CashDelta    *big.Int
	ManagedDelta *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceManaged is a free log retrieval operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterPoolBalanceManaged(opts *bind.FilterOpts, poolId [][32]byte, assetManager []common.Address, token []common.Address) (*Balancerv2thevaultPoolBalanceManagedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultPoolBalanceManagedIterator{contract: _Balancerv2thevault.contract, event: "PoolBalanceManaged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceManaged is a free log subscription operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchPoolBalanceManaged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultPoolBalanceManaged, poolId [][32]byte, assetManager []common.Address, token []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultPoolBalanceManaged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
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

// ParsePoolBalanceManaged is a log parse operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParsePoolBalanceManaged(log types.Log) (*Balancerv2thevaultPoolBalanceManaged, error) {
	event := new(Balancerv2thevaultPoolBalanceManaged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolRegisteredIterator struct {
	Event *Balancerv2thevaultPoolRegistered // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultPoolRegistered)
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
		it.Event = new(Balancerv2thevaultPoolRegistered)
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
func (it *Balancerv2thevaultPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultPoolRegistered represents a PoolRegistered event raised by the Balancerv2thevault contract.
type Balancerv2thevaultPoolRegistered struct {
	PoolId         [32]byte
	PoolAddress    common.Address
	Specialization uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterPoolRegistered(opts *bind.FilterOpts, poolId [][32]byte, poolAddress []common.Address) (*Balancerv2thevaultPoolRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultPoolRegisteredIterator{contract: _Balancerv2thevault.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultPoolRegistered, poolId [][32]byte, poolAddress []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultPoolRegistered)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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

// ParsePoolRegistered is a log parse operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParsePoolRegistered(log types.Log) (*Balancerv2thevaultPoolRegistered, error) {
	event := new(Balancerv2thevaultPoolRegistered)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultRelayerApprovalChangedIterator is returned from FilterRelayerApprovalChanged and is used to iterate over the raw logs and unpacked data for RelayerApprovalChanged events raised by the Balancerv2thevault contract.
type Balancerv2thevaultRelayerApprovalChangedIterator struct {
	Event *Balancerv2thevaultRelayerApprovalChanged // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultRelayerApprovalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultRelayerApprovalChanged)
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
		it.Event = new(Balancerv2thevaultRelayerApprovalChanged)
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
func (it *Balancerv2thevaultRelayerApprovalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultRelayerApprovalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultRelayerApprovalChanged represents a RelayerApprovalChanged event raised by the Balancerv2thevault contract.
type Balancerv2thevaultRelayerApprovalChanged struct {
	Relayer  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerApprovalChanged is a free log retrieval operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterRelayerApprovalChanged(opts *bind.FilterOpts, relayer []common.Address, sender []common.Address) (*Balancerv2thevaultRelayerApprovalChangedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultRelayerApprovalChangedIterator{contract: _Balancerv2thevault.contract, event: "RelayerApprovalChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerApprovalChanged is a free log subscription operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchRelayerApprovalChanged(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultRelayerApprovalChanged, relayer []common.Address, sender []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultRelayerApprovalChanged)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
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

// ParseRelayerApprovalChanged is a log parse operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseRelayerApprovalChanged(log types.Log) (*Balancerv2thevaultRelayerApprovalChanged, error) {
	event := new(Balancerv2thevaultRelayerApprovalChanged)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Balancerv2thevault contract.
type Balancerv2thevaultSwapIterator struct {
	Event *Balancerv2thevaultSwap // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultSwap)
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
		it.Event = new(Balancerv2thevaultSwap)
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
func (it *Balancerv2thevaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultSwap represents a Swap event raised by the Balancerv2thevault contract.
type Balancerv2thevaultSwap struct {
	PoolId    [32]byte
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterSwap(opts *bind.FilterOpts, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (*Balancerv2thevaultSwapIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultSwapIterator{contract: _Balancerv2thevault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultSwap, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultSwap)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseSwap(log types.Log) (*Balancerv2thevaultSwap, error) {
	event := new(Balancerv2thevaultSwap)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultTokensDeregisteredIterator is returned from FilterTokensDeregistered and is used to iterate over the raw logs and unpacked data for TokensDeregistered events raised by the Balancerv2thevault contract.
type Balancerv2thevaultTokensDeregisteredIterator struct {
	Event *Balancerv2thevaultTokensDeregistered // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultTokensDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultTokensDeregistered)
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
		it.Event = new(Balancerv2thevaultTokensDeregistered)
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
func (it *Balancerv2thevaultTokensDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultTokensDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultTokensDeregistered represents a TokensDeregistered event raised by the Balancerv2thevault contract.
type Balancerv2thevaultTokensDeregistered struct {
	PoolId [32]byte
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensDeregistered is a free log retrieval operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterTokensDeregistered(opts *bind.FilterOpts, poolId [][32]byte) (*Balancerv2thevaultTokensDeregisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultTokensDeregisteredIterator{contract: _Balancerv2thevault.contract, event: "TokensDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokensDeregistered is a free log subscription operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchTokensDeregistered(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultTokensDeregistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultTokensDeregistered)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
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

// ParseTokensDeregistered is a log parse operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseTokensDeregistered(log types.Log) (*Balancerv2thevaultTokensDeregistered, error) {
	event := new(Balancerv2thevaultTokensDeregistered)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Balancerv2thevaultTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the Balancerv2thevault contract.
type Balancerv2thevaultTokensRegisteredIterator struct {
	Event *Balancerv2thevaultTokensRegistered // Event containing the contract specifics and raw log

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
func (it *Balancerv2thevaultTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2thevaultTokensRegistered)
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
		it.Event = new(Balancerv2thevaultTokensRegistered)
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
func (it *Balancerv2thevaultTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2thevaultTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2thevaultTokensRegistered represents a TokensRegistered event raised by the Balancerv2thevault contract.
type Balancerv2thevaultTokensRegistered struct {
	PoolId        [32]byte
	Tokens        []common.Address
	AssetManagers []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) FilterTokensRegistered(opts *bind.FilterOpts, poolId [][32]byte) (*Balancerv2thevaultTokensRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.FilterLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2thevaultTokensRegisteredIterator{contract: _Balancerv2thevault.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *Balancerv2thevaultTokensRegistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2thevault.contract.WatchLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2thevaultTokensRegistered)
				if err := _Balancerv2thevault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
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

// ParseTokensRegistered is a log parse operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Balancerv2thevault *Balancerv2thevaultFilterer) ParseTokensRegistered(log types.Log) (*Balancerv2thevaultTokensRegistered, error) {
	event := new(Balancerv2thevaultTokensRegistered)
	if err := _Balancerv2thevault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
