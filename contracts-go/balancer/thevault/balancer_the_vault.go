// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thevault

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

// ThevaultMetaData contains all meta data concerning the Thevault contract.
var ThevaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExternalBalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"}],\"name\":\"InternalBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"deltas\",\"type\":\"int256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"protocolFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"PoolBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cashDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"managedDelta\",\"type\":\"int256\"}],\"name\":\"PoolBalanceManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"RelayerApprovalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"batchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"assetDeltas\",\"type\":\"int256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"deregisterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.ExitPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getInternalBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"managed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"hasApprovedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.JoinPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.PoolBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.PoolBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"managePoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.UserBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIVault.UserBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"manageUserBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"queryBatchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"registerPool\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"setAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setRelayerApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.SingleSwap\",\"name\":\"singleSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculated\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162006ed638038062006ed6833981016040819052620000359162000253565b8382826040518060400160405280601181526020017010985b185b98d95c88158c8815985d5b1d607a1b81525080604051806040016040528060018152602001603160f81b815250306001600160a01b031660001b89806001600160a01b03166080816001600160a01b031660601b815250505030604051620000b89062000245565b620000c491906200029f565b604051809103906000f080158015620000e1573d6000803e3d6000fd5b5060601b6001600160601b03191660a052600160005560c052815160209283012060e052805191012061010052507f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61012052620001486276a70083111561019462000181565b6200015c62278d0082111561019562000181565b429091016101408190520161016052620001768162000196565b5050505050620002cc565b8162000192576200019281620001f2565b5050565b6040516001600160a01b038216907f94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef90600090a2600380546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b62461bcd60e51b6000908152602060045260076024526642414c23000030600a808404818106603090810160081b95839006959095019082900491820690940160101b939093010160c81b604452606490fd5b610be680620062f083390190565b6000806000806080858703121562000269578384fd5b84516200027681620002b3565b60208601519094506200028981620002b3565b6040860151606090960151949790965092505050565b6001600160a01b0391909116815260200190565b6001600160a01b0381168114620002c957600080fd5b50565b60805160601c60a05160601c60c05160e05161010051610120516101405161016051615fc06200033060003980611aed525080611ac952508061289f5250806128e15250806128c05250806110fd5250806113b15250806105285250615fc06000f3fe6080604052600436106101a55760003560e01c8063945bcec9116100e1578063e6c460921161008a578063f84d066e11610064578063f84d066e1461048a578063f94d4668146104aa578063fa6e671d146104d9578063fec90d72146104f9576101d3565b8063e6c4609214610427578063ed24911d14610447578063f6c009271461045c576101d3565b8063b05f8e48116100bb578063b05f8e48146103cf578063b95cac28146103ff578063d2946c2b14610412576101d3565b8063945bcec914610385578063aaabadc514610398578063ad5c4648146103ba576101d3565b806352bbbe291161014e5780637d3aeb96116101285780637d3aeb9614610305578063851c1bb3146103255780638bdb39131461034557806390193b7c14610365576101d3565b806352bbbe29146102b25780635c38449e146102c557806366a9c7d2146102e5576101d3565b80630f5a6efa1161017f5780630f5a6efa1461024157806316c38b3c1461026e5780631c0de0511461028e576101d3565b8063058a628f146101d857806309b2760f146101f85780630e8e3e841461022e576101d3565b366101d3576101d16101b5610526565b6001600160a01b0316336001600160a01b03161461020661054b565b005b600080fd5b3480156101e457600080fd5b506101d16101f3366004615157565b61055d565b34801561020457600080fd5b506102186102133660046156e6565b610581565b6040516102259190615d3e565b60405180910390f35b6101d161023c36600461531e565b610634565b34801561024d57600080fd5b5061026161025c3660046151f5565b610770565b6040516102259190615d08565b34801561027a57600080fd5b506101d161028936600461545c565b610806565b34801561029a57600080fd5b506102a361081f565b60405161022593929190615d26565b6102186102c036600461588f565b610848565b3480156102d157600080fd5b506101d16102e036600461565b565b6109e9565b3480156102f157600080fd5b506101d1610300366004615545565b610e06565b34801561031157600080fd5b506101d1610320366004615516565b610fa5565b34801561033157600080fd5b50610218610340366004615633565b6110f9565b34801561035157600080fd5b506101d16103603660046154ac565b61114b565b34801561037157600080fd5b50610218610380366004615157565b611161565b610261610393366004615786565b61117c565b3480156103a457600080fd5b506103ad6112b0565b6040516102259190615b63565b3480156103c657600080fd5b506103ad6112c4565b3480156103db57600080fd5b506103ef6103ea36600461560f565b6112d3565b6040516102259493929190615eb9565b6101d161040d3660046154ac565b611396565b34801561041e57600080fd5b506103ad6113af565b34801561043357600080fd5b506101d1610442366004615243565b6113d3565b34801561045357600080fd5b506102186114ef565b34801561046857600080fd5b5061047c610477366004615494565b6114f9565b604051610225929190615b9b565b34801561049657600080fd5b506102616104a5366004615702565b611523565b3480156104b657600080fd5b506104ca6104c5366004615494565b611620565b60405161022593929190615cd2565b3480156104e557600080fd5b506101d16104f43660046151ab565b611654565b34801561050557600080fd5b50610519610514366004615173565b6116e6565b6040516102259190615d1b565b7f00000000000000000000000000000000000000000000000000000000000000005b90565b8161055957610559816116fb565b5050565b610565611768565b61056d611781565b610576816117af565b61057e611822565b50565b600061058b611768565b610593611829565b60006105a2338460065461183e565b6000818152600560205260409020549091506105c49060ff16156101f461054b565b60008181526005602052604090819020805460ff1916600190811790915560068054909101905551339082907f3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e9061061d908790615e3a565b60405180910390a3905061062f611822565b919050565b61063c611768565b6000806000805b845181101561075b5760008060008060006106718a878151811061066357fe5b60200260200101518961187d565b9c50939850919650945092509050600185600381111561068d57fe5b14156106a45761069f848383866118f5565b61074a565b866106b6576106b1611829565b600196505b60008560038111156106c457fe5b14156106f5576106d684838386611918565b6106df84611938565b1561069f576106ee8984611945565b985061074a565b61070a61070185611938565b1561020761054b565b600061071585610548565b9050600286600381111561072557fe5b141561073c5761073781848487611957565b610748565b61074881848487611970565b505b505060019093019250610643915050565b50610765836119de565b50505061057e611822565b6060815167ffffffffffffffff8111801561078a57600080fd5b506040519080825280602002602001820160405280156107b4578160200160208202803683370190505b50905060005b82518110156107ff576107e0848483815181106107d357fe5b6020026020010151611a01565b8282815181106107ec57fe5b60209081029190910101526001016107ba565b5092915050565b61080e611768565b610816611781565b61057681611a2c565b600080600061082c611aaa565b159250610837611ac7565b9150610841611aeb565b9050909192565b6000610852611768565b61085a611829565b835161086581611b0f565b610874834211156101fc61054b565b61088760008760800151116101fe61054b565b60006108968760400151611b41565b905060006108a78860600151611b41565b90506108ca816001600160a01b0316836001600160a01b031614156101fd61054b565b6108d2614ce1565b885160808201526020890151819060018111156108eb57fe5b908160018111156108f857fe5b9052506001600160a01b03808416602083015282811660408084019190915260808b0151606084015260a08b01516101008401528951821660c08401528901511660e082015260008061094a83611b66565b9198509250905061098160008c60200151600181111561096657fe5b146109745789831115610979565b898210155b6101fb61054b565b6109998b60400151838c600001518d60200151611c5a565b6109b18b60600151828c604001518d60600151611d38565b6109d36109c18c60400151611938565b6109cc5760006109ce565b825b6119de565b5050505050506109e1611822565b949350505050565b6109f1611768565b6109f9611829565b610a0583518351611e12565b6060835167ffffffffffffffff81118015610a1f57600080fd5b50604051908082528060200260200182016040528015610a49578160200160208202803683370190505b5090506060845167ffffffffffffffff81118015610a6657600080fd5b50604051908082528060200260200182016040528015610a90578160200160208202803683370190505b5090506000805b8651811015610c09576000878281518110610aae57fe5b602002602001015190506000878381518110610ac657fe5b60200260200101519050610b11846001600160a01b0316836001600160a01b03161160006001600160a01b0316846001600160a01b031614610b09576066610b0c565b60685b61054b565b819350816001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610b409190615b63565b60206040518083038186803b158015610b5857600080fd5b505afa158015610b6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b909190615968565b858481518110610b9c57fe5b602002602001018181525050610bb181611e1f565b868481518110610bbd57fe5b602002602001018181525050610beb81868581518110610bd957fe5b6020026020010151101561021061054b565b610bff6001600160a01b0383168b83611ea6565b5050600101610a97565b506040517ff04f27070000000000000000000000000000000000000000000000000000000081526001600160a01b0388169063f04f270790610c55908990899088908a90600401615c85565b600060405180830381600087803b158015610c6f57600080fd5b505af1158015610c83573d6000803e3d6000fd5b5050505060005b8651811015610df4576000878281518110610ca157fe5b602002602001015190506000848381518110610cb957fe5b602002602001015190506000826001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610cf19190615b63565b60206040518083038186803b158015610d0957600080fd5b505afa158015610d1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d419190615968565b9050610d528282101561020361054b565b60008282039050610d7b888681518110610d6857fe5b602002602001015182101561025a61054b565b610d858482611f11565b836001600160a01b03168c6001600160a01b03167f0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f08c8881518110610dc657fe5b602002602001015184604051610ddd929190615e4d565b60405180910390a350505050806001019050610c8a565b50505050610e00611822565b50505050565b610e0e611768565b610e16611829565b82610e2081611f33565b610e2c83518351611e12565b60005b8351811015610eca576000848281518110610e4657fe5b60200260200101519050610e7260006001600160a01b0316826001600160a01b0316141561013561054b565b838281518110610e7e57fe5b6020908102919091018101516000888152600a835260408082206001600160a01b0395861683529093529190912080546001600160a01b03191692909116919091179055600101610e2f565b506000610ed685611f64565b90506002816002811115610ee657fe5b1415610f3457610efc845160021461020c61054b565b610f2f8585600081518110610f0d57fe5b602002602001015186600181518110610f2257fe5b6020026020010151611f7e565b610f5c565b6001816002811115610f4257fe5b1415610f5257610f2f858561202a565b610f5c8585612082565b847ff5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec48585604051610f8e929190615bed565b60405180910390a25050610fa0611822565b505050565b610fad611768565b610fb5611829565b81610fbf81611f33565b6000610fca84611f64565b90506002816002811115610fda57fe5b141561102857610ff0835160021461020c61054b565b611023848460008151811061100157fe5b60200260200101518560018151811061101657fe5b60200260200101516120d7565b611050565b600181600281111561103657fe5b1415611046576110238484612145565b61105084846121ff565b60005b83518110156110b657600a6000868152602001908152602001600020600085838151811061107d57fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b0319169055600101611053565b50837f7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610846040516110e79190615bda565b60405180910390a25050610559611822565b60007f00000000000000000000000000000000000000000000000000000000000000008260405160200161112e929190615ac2565b604051602081830303815290604052805190602001209050919050565b610e00600185858561115c86612262565b61226e565b6001600160a01b031660009081526002602052604090205490565b6060611186611768565b61118e611829565b835161119981611b0f565b6111a8834211156101fc61054b565b6111b486518551611e12565b6111c08787878b6123f4565b91506000805b87518110156112925760008882815181106111dd57fe5b6020026020010151905060008583815181106111f557fe5b6020026020010151905061122188848151811061120e57fe5b60200260200101518213156101fb61054b565b600081131561126157885160208a015182916112409185918491611c5a565b61124983611938565b1561125b576112588582611945565b94505b50611288565b600081121561128857600081600003905061128683828c604001518d60600151611d38565b505b50506001016111c6565b5061129c816119de565b50506112a6611822565b9695505050505050565b60035461010090046001600160a01b031690565b60006112ce610526565b905090565b600080600080856112e381612683565b6000806112ef89611f64565b905060028160028111156112ff57fe5b14156113165761130f89896126a1565b9150611341565b600181600281111561132457fe5b14156113345761130f898961271b565b61133e8989612789565b91505b61134a826127a1565b9650611355826127b4565b9550611360826127ca565b6000998a52600a60209081526040808c206001600160a01b039b8c168d5290915290992054969995989796909616955050505050565b61139e611829565b610e00600085858561115c86612262565b7f000000000000000000000000000000000000000000000000000000000000000090565b6113db611768565b6113e3611829565b6113eb614d31565b60005b82518110156114e55782818151811061140357fe5b6020026020010151915060008260200151905061141f81612683565b604083015161143961143183836127d0565b61020961054b565b6000828152600a602090815260408083206001600160a01b03858116855292529091205461146c911633146101f661054b565b835160608501516000806114828487878661282c565b91509150846001600160a01b0316336001600160a01b0316877f6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a85856040516114cc929190615e4d565b60405180910390a45050505050508060010190506113ee565b505061057e611822565b60006112ce61289b565b6000808261150681612683565b61150f84612938565b61151885611f64565b925092505b50915091565b60603330146115f6576000306001600160a01b0316600036604051611549929190615ada565b6000604051808303816000865af19150503d8060008114611586576040519150601f19603f3d011682016040523d82523d6000602084013e61158b565b606091505b50509050806000811461159a57fe5b60046000803e6000516001600160e01b0319167ffa61cc120000000000000000000000000000000000000000000000000000000081146115de573d6000803e3d6000fd5b50602060005260043d0380600460203e602081016000f35b6060611604858585896123f4565b9050602081510263fa61cc126020830352600482036024820181fd5b60608060008361162f81612683565b606061163a8661293e565b9095509050611648816129a0565b95979096509350505050565b61165c611768565b611664611829565b8261166e81611b0f565b6001600160a01b0384811660008181526004602090815260408083209488168084529490915290819020805460ff1916861515179055519091907f46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8906116d5908690615d1b565b60405180910390a350610fa0611822565b60006116f28383612a4f565b90505b92915050565b7f08c379a0000000000000000000000000000000000000000000000000000000006000908152602060045260076024526642414c23000030600a808404818106603090810160081b95839006959095019082900491820690940160101b939093010160c81b604452606490fd5b61177a6002600054141561019061054b565b6002600055565b60006117986000356001600160e01b0319166110f9565b905061057e6117a78233612a7d565b61019161054b565b6040516001600160a01b038216907f94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef90600090a2600380546001600160a01b03909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b6001600055565b61183c611834611aaa565b61019261054b565b565b600069ffffffffffffffffffff8216605084600281111561185b57fe5b901b17606085901b6bffffffffffffffffffffffff19161790505b9392505050565b600080600080600080600088606001519050336001600160a01b0316816001600160a01b0316146118cf57876118ba576118b5611781565b600197505b6118cf6118c78233612a4f565b6101f761054b565b885160208a015160408b01516080909b0151919b909a9992985090965090945092505050565b61190a8361190286611b41565b836000612b20565b50610e008482846000611d38565b61192b8261192586611b41565b83612b76565b610e008482856000611c5a565b6001600160a01b03161590565b60008282016116f2848210158361054b565b6119648385836000612b20565b50610e00828583612b76565b8015610e005761198b6001600160a01b038516848484612ba6565b826001600160a01b0316846001600160a01b03167f540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c84846040516119d0929190615bc1565b60405180910390a350505050565b6119ed8134101561020461054b565b348190038015610559576105593382612bc7565b6001600160a01b039182166000908152600b6020908152604080832093909416825291909152205490565b8015611a4c57611a47611a3d611ac7565b421061019361054b565b611a61565b611a61611a57611aeb565b42106101a961054b565b6003805460ff19168215151790556040517f9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be6490611a9f908390615d1b565b60405180910390a150565b6000611ab4611aeb565b4211806112ce57505060035460ff161590565b7f000000000000000000000000000000000000000000000000000000000000000090565b7f000000000000000000000000000000000000000000000000000000000000000090565b336001600160a01b0382161461057e57611b27611781565b611b318133612a4f565b61057e5761057e816101f7612c41565b6000611b4c82611938565b611b5e57611b5982610548565b6116f5565b6116f5610526565b600080600080611b798560800151612938565b90506000611b8a8660800151611f64565b90506002816002811115611b9a57fe5b1415611bb157611baa8683612c75565b9450611bdc565b6001816002811115611bbf57fe5b1415611bcf57611baa8683612d25565b611bd98683612db8565b94505b611bef8660000151876060015187612ff7565b809450819550505085604001516001600160a01b031686602001516001600160a01b031687608001517f2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b8787604051611c49929190615e4d565b60405180910390a450509193909250565b82611c6457610e00565b611c6d84611938565b15611cee57611c7f811561020261054b565b611c8e8347101561020461054b565b611c96610526565b6001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b158015611cd057600080fd5b505af1158015611ce4573d6000803e3d6000fd5b5050505050610e00565b6000611cf985610548565b90508115611d16576000611d108483876001612b20565b90940393505b8315611d3157611d316001600160a01b038216843087612ba6565b5050505050565b82611d4257610e00565b611d4b84611938565b15611ddb57611d5d811561020261054b565b611d65610526565b6001600160a01b0316632e1a7d4d846040518263ffffffff1660e01b8152600401611d909190615d3e565b600060405180830381600087803b158015611daa57600080fd5b505af1158015611dbe573d6000803e3d6000fd5b50611dd6925050506001600160a01b03831684612bc7565b610e00565b6000611de685610548565b90508115611dfe57611df9838286612b76565b611d31565b611d316001600160a01b0382168486611ea6565b610559818314606761054b565b600080611e2a6113af565b6001600160a01b031663d877845c6040518163ffffffff1660e01b815260040160206040518083038186803b158015611e6257600080fd5b505afa158015611e76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e9a9190615968565b90506118768382613025565b610fa08363a9059cbb60e01b8484604051602401611ec5929190615bc1565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001600160e01b031990931692909217909152613072565b801561055957610559611f226113af565b6001600160a01b0384169083611ea6565b611f3c81612683565b61057e611f4882612938565b6001600160a01b0316336001600160a01b0316146101f561054b565b600061ffff605083901c166116f5600382106101f461054b565b611f9f816001600160a01b0316836001600160a01b0316141561020a61054b565b611fbe816001600160a01b0316836001600160a01b031610606661054b565b60008381526009602052604090208054611ffb906001600160a01b0316158015611ff3575060018201546001600160a01b0316155b61020b61054b565b80546001600160a01b039384166001600160a01b03199182161782556001909101805492909316911617905550565b6000828152600860205260408120905b8251811015610e0057600061206b84838151811061205457fe5b60200260200101518461311290919063ffffffff16565b90506120798161020a61054b565b5060010161203a565b6000828152600160205260408120905b8251811015610e005760006120c08483815181106120ac57fe5b602090810291909101015184906000613175565b90506120ce8161020a61054b565b50600101612092565b60008060006120e7868686613222565b9250925092506121116120f9846132e9565b80156121095750612109836132e9565b61020d61054b565b600095865260096020526040862080546001600160a01b031990811682556001909101805490911690559490945550505050565b6000828152600860205260408120905b8251811015610e0057600083828151811061216c57fe5b602002602001015190506121b8612109600760008881526020019081526020016000206000846001600160a01b03166001600160a01b03168152602001908152602001600020546132e9565b60008581526007602090815260408083206001600160a01b038516845290915281208190556121e7848361330b565b90506121f58161020961054b565b5050600101612155565b6000828152600160205260408120905b8251811015610e0057600083828151811061222657fe5b60200260200101519050600061223c8483613412565b905061224a612109826132e9565b6122548483613421565b50505080600101905061220f565b61226a614d5a565b5090565b612276611768565b8361228081612683565b8361228a81611b0f565b61229e836000015151846020015151611e12565b60606122ad84600001516134c3565b905060606122bb8883613552565b905060608060606122d08c8c8c8c8c896135e3565b92509250925060006122e18c611f64565b905060028160028111156122f157fe5b1415612359576123548c8760008151811061230857fe5b60200260200101518660008151811061231d57fe5b60200260200101518960018151811061233257fe5b60200260200101518860018151811061234757fe5b60200260200101516137a8565b612382565b600181600281111561236757fe5b1415612378576123548c87866137e7565b6123828c85613854565b6000808e600181111561239157fe5b1490508b6001600160a01b03168d7fe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78896123cb888661389d565b876040516123db93929190615c4c565b60405180910390a3505050505050505050611d31611822565b6060835167ffffffffffffffff8111801561240e57600080fd5b50604051908082528060200260200182016040528015612438578160200160208202803683370190505b509050612443614d84565b61244b614ce1565b60008060005b89518110156126765789818151811061246657fe5b6020026020010151945060008951866020015110801561248a575089518660400151105b905061249781606461054b565b60006124b98b8860200151815181106124ac57fe5b6020026020010151611b41565b905060006124d08c8960400151815181106124ac57fe5b90506124f3816001600160a01b0316836001600160a01b031614156101fd61054b565b60608801516125435761250b600085116101fe61054b565b60006125188b8484613945565b6001600160a01b0316876001600160a01b031614905061253a816101ff61054b565b50606088018590525b87516080880152868a600181111561255757fe5b9081600181111561256457fe5b9052506001600160a01b0380831660208901528181166040808a01919091526060808b0151908a015260808a01516101008a01528c51821660c08a01528c01511660e08801526000806125b689611b66565b919850925090506125c88c8585613967565b97506125fc6125d683613981565b8c8c60200151815181106125e657fe5b60200260200101516139b190919063ffffffff16565b8b8b602001518151811061260c57fe5b60200260200101818152505061264a61262482613981565b8c8c604001518151811061263457fe5b60200260200101516139e590919063ffffffff16565b8b8b604001518151811061265a57fe5b6020026020010181815250505050505050806001019050612451565b5050505050949350505050565b60008181526005602052604090205461057e9060ff166101f461054b565b60008060008060006126b287613a19565b945094509450945050836001600160a01b0316866001600160a01b031614156126e157829450505050506116f5565b816001600160a01b0316866001600160a01b031614156127065793506116f592505050565b6127116102096116fb565b5050505092915050565b60008281526007602090815260408083206001600160a01b03851684529091528120548161274882613a8f565b80612766575060008581526008602052604090206127669085613aa1565b9050806127815761277685612683565b6127816102096116fb565b509392505050565b60008281526001602052604081206109e18184613412565b6dffffffffffffffffffffffffffff1690565b60701c6dffffffffffffffffffffffffffff1690565b60e01c90565b6000806127dc84611f64565b905060028160028111156127ec57fe5b1415612804576127fc8484613ac2565b9150506116f5565b600181600281111561281257fe5b1415612822576127fc8484613b13565b6127fc8484613b2b565b600080600061283a86611f64565b9050600087600281111561284a57fe5b14156128665761285c86828787613b43565b9250925050612892565b600187600281111561287457fe5b14156128865761285c86828787613bbe565b61285c86828787613c3a565b94509492505050565b60007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000612908613c9d565b3060405160200161291d959493929190615df0565b60405160208183030381529060405280519060200120905090565b60601c90565b606080600061294c84611f64565b9050600281600281111561295c57fe5b14156129755761296b84613ca1565b925092505061299b565b600181600281111561298357fe5b14156129925761296b84613dd6565b61296b84613efd565b915091565b60606000825167ffffffffffffffff811180156129bc57600080fd5b506040519080825280602002602001820160405280156129e6578160200160208202803683370190505b5091506000905060005b825181101561151d576000848281518110612a0757fe5b60200260200101519050612a1a81613ff9565b848381518110612a2657fe5b602002602001018181525050612a4483612a3f836127ca565b614014565b9250506001016129f0565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6003546040517f9be2a88400000000000000000000000000000000000000000000000000000000815260009161010090046001600160a01b031690639be2a88490612ad090869086903090600401615d47565b60206040518083038186803b158015612ae857600080fd5b505afa158015612afc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f29190615478565b600080612b2d8686611a01565b9050612b468380612b3e5750848210155b61020161054b565b612b50818561402b565b9150818103612b6c878783612b6487613981565b60000361403a565b5050949350505050565b6000612b828484611a01565b90506000612b908284611945565b9050611d31858583612ba187613981565b61403a565b610e00846323b872dd60e01b858585604051602401611ec593929190615b77565b612bd6814710156101a361054b565b6000826001600160a01b031682604051612bef90610548565b60006040518083038185875af1925050503d8060008114612c2c576040519150601f19603f3d011682016040523d82523d6000602084013e612c31565b606091505b50509050610fa0816101a461054b565b6001600160a01b0382166000908152600260205260409020805460018101909155610fa0612c6f8483614095565b8361054b565b600080600080612c92866080015187602001518860400151613222565b92509250925060008087604001516001600160a01b031688602001516001600160a01b03161015612cc7575083905082612ccd565b50829050835b612cd9888884846141bb565b60408b015160208c01519199509294509092506001600160a01b03918216911610612d0d57612d0881836142d1565b612d17565b612d1782826142d1565b909255509295945050505050565b600080612d3a8460800151856020015161271b565b90506000612d508560800151866040015161271b565b9050612d5e858584846141bb565b6080880180516000908152600760208181526040808420828e01516001600160a01b03908116865290835281852098909855935183529081528282209a830151909516815298909352919096209590955550929392505050565b60808201516000908152600160209081526040822090840151829182918290612de290839061430c565b90506000612dfd88604001518461430c90919063ffffffff16565b9050811580612e0a575080155b15612e2757612e1c8860800151612683565b612e276102096116fb565b60001991820191016000612e3a8461432b565b905060608167ffffffffffffffff81118015612e5557600080fd5b50604051908082528060200260200182016040528015612e7f578160200160208202803683370190505b50600060a08c018190529091505b82811015612eff576000612ea1878361432f565b9050612eac81613ff9565b838381518110612eb857fe5b602002602001018181525050612ed58c60a00151612a3f836127ca565b60a08d015281861415612eea57809850612ef6565b84821415612ef6578097505b50600101612e8d565b506040517f01ec954a0000000000000000000000000000000000000000000000000000000081526001600160a01b038a16906301ec954a90612f4b908d90859089908990600401615e5b565b602060405180830381600087803b158015612f6557600080fd5b505af1158015612f79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f9d9190615968565b9750600080612fb58c600001518d606001518c612ff7565b9092509050612fc48983614345565b9850612fd08882614376565b9750612fdd87878b61438c565b612fe887868a61438c565b50505050505050505092915050565b6000808085600181111561300757fe5b141561301757508290508161301d565b50819050825b935093915050565b600082820261304984158061304257508385838161303f57fe5b04145b600361054b565b806130585760009150506116f5565b670de0b6b3a76400006000198201046001019150506116f5565b60006060836001600160a01b03168360405161308e9190615aea565b6000604051808303816000865af19150503d80600081146130cb576040519150601f19603f3d011682016040523d82523d6000602084013e6130d0565b606091505b509150915060008214156130e8573d6000803e3d6000fd5b610e0081516000148061310a57508180602001905181019061310a9190615478565b6101a261054b565b600061311e8383613aa1565b61316d57508154600180820184556000848152602080822090930180546001600160a01b0319166001600160a01b038616908117909155855490825282860190935260409020919091556116f5565b5060006116f5565b6001600160a01b03821660009081526002840160205260408120548061320257505082546040805180820182526001600160a01b03858116808352602080840187815260008781526001808c018452878220965187546001600160a01b03191696169590951786559051948401949094559482018089559083526002880190945291902091909155611876565b600019016000908152600180860160205260408220018390559050611876565b600080600080600061323487876143a4565b91509150600061324483836143d5565b60008a81526009602090815260408083208484526002019091528120805460018201549197509293509061327783613a8f565b80613286575061328682613a8f565b806132a757506132968c87613ac2565b80156132a757506132a78c86613ac2565b9050806132c2576132b78c612683565b6132c26102096116fb565b6132cc8383614408565b98506132d8838361442d565b975050505050505093509350939050565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161590565b6001600160a01b03811660009081526001830160205260408120548015613408578354600019808301919081019060009087908390811061334857fe5b60009182526020909120015487546001600160a01b039091169150819088908590811061337157fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559183168152600189810190925260409020908401905586548790806133ba57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03881682526001898101909152604082209190915594506116f59350505050565b60009150506116f5565b60006116f28383610209614444565b6001600160a01b0381166000908152600283016020526040812054801561340857835460001990810160008181526001878101602090815260408084209587018452808420865481546001600160a01b03199081166001600160a01b0392831617835588860180549387019390935588548216875260028d018086528488209a909a5588541690975584905593895593871682529390925281205590506116f5565b606080825167ffffffffffffffff811180156134de57600080fd5b50604051908082528060200260200182016040528015613508578160200160208202803683370190505b50905060005b83518110156107ff576135268482815181106124ac57fe5b82828151811061353257fe5b6001600160a01b039092166020928302919091019091015260010161350e565b60608060606135608561293e565b9150915061357082518551611e12565b613580600083511161020f61054b565b60005b82518110156135da576135d285828151811061359b57fe5b60200260200101516001600160a01b03168483815181106135b857fe5b60200260200101516001600160a01b03161461020861054b565b600101613583565b50949350505050565b60608060608060006135f4866129a0565b9150915060006136038b612938565b905060008c600181111561361357fe5b146136b657806001600160a01b03166374f3b0098c8c8c8787613634614481565b8f604001516040518863ffffffff1660e01b815260040161365b9796959493929190615d66565b600060405180830381600087803b15801561367557600080fd5b505af1158015613689573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526136b19190810190615405565b61374f565b806001600160a01b031663d5c096c48c8c8c87876136d2614481565b8f604001516040518863ffffffff1660e01b81526004016136f99796959493929190615d66565b600060405180830381600087803b15801561371357600080fd5b505af1158015613727573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261374f9190810190615405565b80955081965050506137658751865186516144fb565b60008c600181111561377357fe5b1461378a576137858989898888614513565b613797565b6137978a8989888861465a565b955050505096509650969350505050565b60006137b485846143d5565b600087815260096020908152604080832084845260020190915290209091506137dd85846142d1565b9055505050505050565b60005b8251811015610e00578181815181106137ff57fe5b602002602001015160076000868152602001908152602001600020600085848151811061382857fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020556001016137ea565b6000828152600160205260408120905b8251811015610e00576138958184838151811061387d57fe5b60200260200101518461438c9092919063ffffffff16565b600101613864565b6060825167ffffffffffffffff811180156138b757600080fd5b506040519080825280602002602001820160405280156138e1578160200160208202803683370190505b50905060005b83518110156107ff57826139115783818151811061390157fe5b6020026020010151600003613926565b83818151811061391d57fe5b60200260200101515b82828151811061393257fe5b60209081029190910101526001016138e7565b60008084600181111561395457fe5b1461395f57816109e1565b509092915050565b60008084600181111561397657fe5b146107ff57826109e1565b600061226a7f800000000000000000000000000000000000000000000000000000000000000083106101a561054b565b60008282016116f28284128015906139c95750848212155b806139de57506000841280156139de57508482125b600061054b565b60008183036116f28284128015906139fd5750848213155b80613a125750600084128015613a1257508482135b600161054b565b6000818152600960205260408120805460018201546001600160a01b0391821692849290911690829081613a4d86856143d5565b6000818152600284016020526040902080546001820154919950919250613a748282614408565b9650613a80828261442d565b94505050505091939590929450565b6000613a9a826132e9565b1592915050565b6001600160a01b031660009081526001919091016020526040902054151590565b600082815260096020526040812080546001600160a01b0384811691161480613afa575060018101546001600160a01b038481169116145b80156109e1575050506001600160a01b03161515919050565b60008281526008602052604081206109e18184613aa1565b60008281526001602052604081206109e181846147d0565b6000806002856002811115613b5457fe5b1415613b6a57613b658685856147f1565b613b94565b6001856002811115613b7857fe5b1415613b8957613b658685856147ff565b613b9486858561480d565b8215613bae57613bae6001600160a01b0385163385611ea6565b5050600081900394909350915050565b6000806002856002811115613bcf57fe5b1415613be557613be086858561481b565b613c0f565b6001856002811115613bf357fe5b1415613c0457613be0868585614829565b613c0f868585614837565b8215613c2a57613c2a6001600160a01b038516333086612ba6565b5090946000869003945092505050565b6000806002856002811115613c4b57fe5b1415613c6357613c5c868585614845565b9050613c90565b6001856002811115613c7157fe5b1415613c8257613c5c868585614855565b613c8d868585614865565b90505b6000915094509492505050565b4690565b606080600080600080613cb387613a19565b92975090955093509150506001600160a01b0384161580613cdb57506001600160a01b038216155b15613d04575050604080516000808252602082019081528183019092529450925061299b915050565b60408051600280825260608201835290916020830190803683370190505095508386600081518110613d3257fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508186600181518110613d6057fe5b6001600160a01b03929092166020928302919091018201526040805160028082526060820183529092909190830190803683370190505094508285600081518110613da757fe5b6020026020010181815250508085600181518110613dc157fe5b60200260200101818152505050505050915091565b60008181526008602052604090206060908190613df28161432b565b67ffffffffffffffff81118015613e0857600080fd5b50604051908082528060200260200182016040528015613e32578160200160208202803683370190505b509250825167ffffffffffffffff81118015613e4d57600080fd5b50604051908082528060200260200182016040528015613e77578160200160208202803683370190505b50915060005b8351811015613ef6576000613e928383614875565b905080858381518110613ea157fe5b6001600160a01b03928316602091820292909201810191909152600088815260078252604080822093851682529290915220548451859084908110613ee257fe5b602090810291909101015250600101613e7d565b5050915091565b60008181526001602052604090206060908190613f198161432b565b67ffffffffffffffff81118015613f2f57600080fd5b50604051908082528060200260200182016040528015613f59578160200160208202803683370190505b509250825167ffffffffffffffff81118015613f7457600080fd5b50604051908082528060200260200182016040528015613f9e578160200160208202803683370190505b50915060005b8351811015613ef657613fb782826148a2565b858381518110613fc357fe5b60200260200101858481518110613fd657fe5b60209081029190910101919091526001600160a01b039091169052600101613fa4565b6000614004826127b4565b61400d836127a1565b0192915050565b60008183101561402457816116f2565b5090919050565b600081831061402457816116f2565b6001600160a01b038085166000818152600b602090815260408083209488168084529490915290819020859055517f18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42906119d0908590615d3e565b6000806140a06148c6565b9050428110156140b45760009150506116f5565b60006140be6148d2565b9050806140d0576000925050506116f5565b6000816140db6149e3565b80516020918201206040516140f7939233918a91899101615dc4565b604051602081830303815290604052805190602001209050600061411a82614a32565b90506000806000614129614a4e565b9250925092506000600185858585604051600081526020016040526040516141549493929190615e1c565b6020604051602081039080840390855afa158015614176573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906141ac57508a6001600160a01b0316816001600160a01b0316145b9b9a5050505050505050505050565b6000806000806141ca86613ff9565b905060006141d786613ff9565b90506141ee6141e5886127ca565b612a3f886127ca565b60a08a01526040517f9d2c110c0000000000000000000000000000000000000000000000000000000081526001600160a01b03891690639d2c110c9061423c908c9086908690600401615e94565b602060405180830381600087803b15801561425657600080fd5b505af115801561426a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061428e9190615968565b92506000806142a68b600001518c6060015187612ff7565b90925090506142b58983614345565b96506142c18882614376565b9550505050509450945094915050565b6000806142e96142e0856127ca565b612a3f856127ca565b90506109e16142f7856127a1565b614300856127a1565b8363ffffffff16614a75565b6001600160a01b03166000908152600291909101602052604090205490565b5490565b6000908152600191820160205260409020015490565b60008061435b83614355866127a1565b90611945565b90506000614368856127b4565b9050436112a6838383614a83565b60008061435b83614386866127a1565b90614abc565b60009182526001928301602052604090912090910155565b600080826001600160a01b0316846001600160a01b0316106143c75782846143ca565b83835b915091509250929050565b600082826040516020016143ea929190615b06565b60405160208183030381529060405280519060200120905092915050565b60006116f2614416846127a1565b61441f846127a1565b614428866127ca565b614a83565b60006116f261443b846127b4565b61441f846127b4565b6001600160a01b038216600090815260028401602052604081205461446b8115158461054b565b614478856001830361432f565b95945050505050565b600061448b6113af565b6001600160a01b03166355c676286040518163ffffffff1660e01b815260040160206040518083038186803b1580156144c357600080fd5b505afa1580156144d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ce9190615968565b610fa0828414801561450c57508183145b606761054b565b6060835167ffffffffffffffff8111801561452d57600080fd5b50604051908082528060200260200182016040528015614557578160200160208202803683370190505b50905060005b85515181101561465057600084828151811061457557fe5b602002602001015190506145a58760200151838151811061459257fe5b60200260200101518210156101f961054b565b6000876000015183815181106145b757fe5b602002602001015190506145d181838b8b60600151611d38565b60008584815181106145df57fe5b602002602001015190506145fb6145f583611b41565b82611f11565b61462a6146088483611945565b89868151811061461457fe5b602002602001015161437690919063ffffffff16565b85858151811061463657fe5b60200260200101818152505050505080600101905061455d565b5095945050505050565b60606000845167ffffffffffffffff8111801561467657600080fd5b506040519080825280602002602001820160405280156146a0578160200160208202803683370190505b50915060005b8651518110156147c65760008582815181106146be57fe5b602002602001015190506146ee886020015183815181106146db57fe5b60200260200101518211156101fa61054b565b60008860000151838151811061470057fe5b6020026020010151905061471a81838c8c60600151611c5a565b61472381611938565b15614735576147328483611945565b93505b600086848151811061474357fe5b602002602001015190506147596145f583611b41565b80831015614778576147738382038a868151811061461457fe5b6147a0565b6147a08184038a868151811061478a57fe5b602002602001015161434590919063ffffffff16565b8685815181106147ac57fe5b6020026020010181815250505050508060010190506146a6565b50614650816119de565b6001600160a01b031660009081526002919091016020526040902054151590565b610e008383614ad284614b0d565b610e008383614ad284614bb8565b610e008383614ad284614c13565b610e008383614c6284614b0d565b610e008383614c6284614bb8565b610e008383614c6284614c13565b60006109e18484614c8385614b0d565b60006109e18484614c8385614bb8565b60006109e18484614c8385614c13565b600082600001828154811061488657fe5b6000918252602090912001546001600160a01b03169392505050565b600090815260019182016020526040902080549101546001600160a01b0390911691565b60006112ce6000614c9d565b6000803560e01c8063b95cac28811461491a57638bdb39138114614942576352bbbe29811461496a5763945bcec981146149925763fa6e671d81146149ba57600092506149de565b7f3f7b71252bd19113ff48c19c6e004a9bcfcca320a0d74d58e85877cbd7dcae5892506149de565b7f8bbc57f66ea936902f50a71ce12b92c43f3c5340bb40c27c4e90ab84eeae335392506149de565b7fe192dcbc143b1e244ad73b813fd3c097b832ad260a157340b4e5e5beda067abe92506149de565b7f9bfc43a4d98313c6766986ffd7c916c7481566d9f224c6819af0a53388aced3a92506149de565b7fa3f865aa351e51cfeb40f5178d1564bb629fe9030b83caf6361d1baaf5b90b5a92505b505090565b60606000368080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505082519293505050608010156105485760803603815290565b6000614a3c61289b565b8260405160200161112e929190615b2d565b6000806000614a5d6020614c9d565b9250614a696040614c9d565b91506108416060614c9d565b60e01b60709190911b010190565b6000838301614ab1858210801590614aa957506e01000000000000000000000000000082105b61020e61054b565b614478858585614a75565b6000614acc83831115600161054b565b50900390565b600080614ae283614386866127a1565b90506000614af384614355876127b4565b90506000614b00866127ca565b90506112a6838383614a83565b6000806000806000614b1e89613a19565b9450509350935093506000836001600160a01b0316896001600160a01b03161415614b69576000614b5384898b63ffffffff16565b9050614b5f8185614ca7565b9093509050614b8b565b6000614b7983898b63ffffffff16565b9050614b858184614ca7565b90925090505b614b9583836142d1565b8555614ba18383614cc3565b600190950194909455509192505050949350505050565b600080614bc5868661271b565b90506000614bd782858763ffffffff16565b60008881526007602090815260408083206001600160a01b038b16845290915290208190559050614c088183614ca7565b979650505050505050565b600084815260016020526040812081614c2c8287613412565b90506000614c3e82868863ffffffff16565b9050614c4b838883613175565b50614c568183614ca7565b98975050505050505050565b600080614c7283614355866127a1565b90506000614af384614386876127b4565b600080614c8f846127a1565b905043614478828583614a83565b3601607f19013590565b6000614cb2826127b4565b614cbb846127b4565b039392505050565b60006116f2614cd1846127b4565b614cda846127b4565b6000614a75565b60408051610120810190915280600081526000602082018190526040820181905260608083018290526080830182905260a0830182905260c0830182905260e08301919091526101009091015290565b604080516080810190915280600081526000602082018190526040820181905260609091015290565b60405180608001604052806060815260200160608152602001606081526020016000151581525090565b6040518060a0016040528060008019168152602001600081526020016000815260200160008152602001606081525090565b80356116f581615f5a565b600082601f830112614dd1578081fd5b8135614de4614ddf82615f04565b615edd565b818152915060208083019084810181840286018201871015614e0557600080fd5b60005b84811015614e2d578135614e1b81615f5a565b84529282019290820190600101614e08565b505050505092915050565b600082601f830112614e48578081fd5b8135614e56614ddf82615f04565b818152915060208083019084810160005b84811015614e2d578135870160a080601f19838c03011215614e8857600080fd5b614e9181615edd565b85830135815260408084013587830152606080850135828401526080915081850135818401525082840135925067ffffffffffffffff831115614ed357600080fd5b614ee18c8885870101614fc0565b90820152865250509282019290820190600101614e67565b600082601f830112614f09578081fd5b8135614f17614ddf82615f04565b818152915060208083019084810181840286018201871015614f3857600080fd5b60005b84811015614e2d57813584529282019290820190600101614f3b565b600082601f830112614f67578081fd5b8151614f75614ddf82615f04565b818152915060208083019084810181840286018201871015614f9657600080fd5b60005b84811015614e2d57815184529282019290820190600101614f99565b80356116f581615f6f565b600082601f830112614fd0578081fd5b813567ffffffffffffffff811115614fe6578182fd5b614ff9601f8201601f1916602001615edd565b915080825283602082850101111561501057600080fd5b8060208401602084013760009082016020015292915050565b80356116f581615f7d565b8035600281106116f557600080fd5b8035600481106116f557600080fd5b600060808284031215615063578081fd5b61506d6080615edd565b9050813567ffffffffffffffff8082111561508757600080fd5b61509385838601614dc1565b835260208401359150808211156150a957600080fd5b6150b585838601614ef9565b602084015260408401359150808211156150ce57600080fd5b506150db84828501614fc0565b6040830152506150ee8360608401614fb5565b606082015292915050565b60006080828403121561510a578081fd5b6151146080615edd565b9050813561512181615f5a565b8152602082013561513181615f6f565b6020820152604082013561514481615f5a565b604082015260608201356150ee81615f6f565b600060208284031215615168578081fd5b81356116f281615f5a565b60008060408385031215615185578081fd5b823561519081615f5a565b915060208301356151a081615f5a565b809150509250929050565b6000806000606084860312156151bf578081fd5b83356151ca81615f5a565b925060208401356151da81615f5a565b915060408401356151ea81615f6f565b809150509250925092565b60008060408385031215615207578182fd5b823561521281615f5a565b9150602083013567ffffffffffffffff81111561522d578182fd5b61523985828601614dc1565b9150509250929050565b60006020808385031215615255578182fd5b823567ffffffffffffffff81111561526b578283fd5b8301601f8101851361527b578283fd5b8035615289614ddf82615f04565b818152838101908385016080808502860187018a10156152a7578788fd5b8795505b848610156153105780828b0312156152c1578788fd5b6152ca81615edd565b6152d48b84615029565b8152878301358882015260406152ec8c828601614db6565b908201526060838101359082015284526001959095019492860192908101906152ab565b509098975050505050505050565b60006020808385031215615330578182fd5b823567ffffffffffffffff811115615346578283fd5b8301601f81018513615356578283fd5b8035615364614ddf82615f04565b8181528381019083850160a0808502860187018a1015615382578788fd5b8795505b848610156153105780828b03121561539c578788fd5b6153a581615edd565b6153af8b84615043565b81526153bd8b898501614db6565b818901526040838101359082015260606153d98c828601614db6565b9082015260806153eb8c858301614db6565b908201528452600195909501949286019290810190615386565b60008060408385031215615417578182fd5b825167ffffffffffffffff8082111561542e578384fd5b61543a86838701614f57565b9350602085015191508082111561544f578283fd5b5061523985828601614f57565b60006020828403121561546d578081fd5b81356116f281615f6f565b600060208284031215615489578081fd5b81516116f281615f6f565b6000602082840312156154a5578081fd5b5035919050565b600080600080608085870312156154c1578182fd5b8435935060208501356154d381615f5a565b925060408501356154e381615f5a565b9150606085013567ffffffffffffffff8111156154fe578182fd5b61550a87828801615052565b91505092959194509250565b60008060408385031215615528578182fd5b82359150602083013567ffffffffffffffff81111561522d578182fd5b600080600060608486031215615559578081fd5b8335925060208085013567ffffffffffffffff80821115615578578384fd5b61558488838901614dc1565b94506040870135915080821115615599578384fd5b508501601f810187136155aa578283fd5b80356155b8614ddf82615f04565b81815283810190838501858402850186018b10156155d4578687fd5b8694505b838510156155ff5780356155eb81615f5a565b8352600194909401939185019185016155d8565b5080955050505050509250925092565b60008060408385031215615621578182fd5b8235915060208301356151a081615f5a565b600060208284031215615644578081fd5b81356001600160e01b0319811681146116f2578182fd5b60008060008060808587031215615670578182fd5b843561567b81615f5a565b9350602085013567ffffffffffffffff80821115615697578384fd5b6156a388838901614dc1565b945060408701359150808211156156b8578384fd5b6156c488838901614ef9565b935060608701359150808211156156d9578283fd5b5061550a87828801614fc0565b6000602082840312156156f7578081fd5b81356116f281615f7d565b60008060008060e08587031215615717578182fd5b6157218686615034565b9350602085013567ffffffffffffffff8082111561573d578384fd5b61574988838901614e38565b9450604087013591508082111561575e578384fd5b5061576b87828801614dc1565b92505061577b86606087016150f9565b905092959194509250565b600080600080600080610120878903121561579f578384fd5b6157a98888615034565b955060208088013567ffffffffffffffff808211156157c6578687fd5b6157d28b838c01614e38565b975060408a01359150808211156157e7578687fd5b6157f38b838c01614dc1565b96506158028b60608c016150f9565b955060e08a0135915080821115615817578485fd5b508801601f81018a13615828578384fd5b8035615836614ddf82615f04565b81815283810190838501858402850186018e1015615852578788fd5b8794505b83851015615874578035835260019490940193918501918501615856565b50809650505050505061010087013590509295509295509295565b60008060008060e085870312156158a4578182fd5b843567ffffffffffffffff808211156158bb578384fd5b9086019060c082890312156158ce578384fd5b6158d860c0615edd565b823581526158e98960208501615034565b602082015260408301356158fc81615f5a565b604082015261590e8960608501614db6565b60608201526080830135608082015260a08301358281111561592e578586fd5b61593a8a828601614fc0565b60a08301525080965050505061595386602087016150f9565b939693955050505060a08201359160c0013590565b600060208284031215615979578081fd5b5051919050565b6001600160a01b03169052565b6000815180845260208085019450808401835b838110156159c55781516001600160a01b0316875295820195908201906001016159a0565b509495945050505050565b6000815180845260208085019450808401835b838110156159c5578151875295820195908201906001016159e3565b60008151808452615a17816020860160208601615f24565b601f01601f19169290920160200192915050565b6000610120825160028110615a3c57fe5b808552506020830151615a526020860182615980565b506040830151615a656040860182615980565b50606083015160608501526080830151608085015260a083015160a085015260c0830151615a9660c0860182615980565b5060e0830151615aa960e0860182615980565b506101008084015182828701526112a6838701826159ff565b9182526001600160e01b031916602082015260240190565b6000828483379101908152919050565b60008251615afc818460208701615f24565b9190910192915050565b6bffffffffffffffffffffffff19606093841b811682529190921b16601482015260280190565b7f190100000000000000000000000000000000000000000000000000000000000081526002810192909252602282015260420190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b038316815260408101615bb483615f50565b8260208301529392505050565b6001600160a01b03929092168252602082015260400190565b6000602082526116f2602083018461598d565b600060408252615c00604083018561598d565b828103602084810191909152845180835285820192820190845b81811015615c3f5784516001600160a01b031683529383019391830191600101615c1a565b5090979650505050505050565b600060608252615c5f606083018661598d565b8281036020840152615c7181866159d0565b905082810360408401526112a681856159d0565b600060808252615c98608083018761598d565b8281036020840152615caa81876159d0565b90508281036040840152615cbe81866159d0565b90508281036060840152614c0881856159ff565b600060608252615ce5606083018661598d565b8281036020840152615cf781866159d0565b915050826040830152949350505050565b6000602082526116f260208301846159d0565b901515815260200190565b92151583526020830191909152604082015260600190565b90815260200190565b9283526001600160a01b03918216602084015216604082015260600190565b60008882526001600160a01b03808916602084015280881660408401525060e06060830152615d9860e08301876159d0565b8560808401528460a084015282810360c0840152615db681856159ff565b9a9950505050505050505050565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b60208101615e4783615f50565b91905290565b918252602082015260400190565b600060808252615e6e6080830187615a2b565b8281036020840152615e8081876159d0565b604084019590955250506060015292915050565b600060608252615ea76060830186615a2b565b60208301949094525060400152919050565b938452602084019290925260408301526001600160a01b0316606082015260800190565b60405181810167ffffffffffffffff81118282101715615efc57600080fd5b604052919050565b600067ffffffffffffffff821115615f1a578081fd5b5060209081020190565b60005b83811015615f3f578181015183820152602001615f27565b83811115610e005750506000910152565b6003811061057e57fe5b6001600160a01b038116811461057e57600080fd5b801515811461057e57600080fd5b6003811061057e57600080fdfea2646970667358221220201e4f926e390fed8dd5318c58846af735c2bebc61b80693ae936a5fe76dcf1464736f6c6343000701003360c060405234801561001057600080fd5b50604051610be6380380610be683398101604081905261002f9161004d565b30608052600160005560601b6001600160601b03191660a05261007b565b60006020828403121561005e578081fd5b81516001600160a01b0381168114610074578182fd5b9392505050565b60805160a05160601c610b406100a66000398061041352806105495250806102a75250610b406000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c8063851c1bb311610076578063d877845c1161005b578063d877845c14610129578063e42abf3514610131578063fbfa77cf14610151576100a3565b8063851c1bb314610101578063aaabadc514610114576100a3565b806338e9922e146100a857806355c67628146100bd5780636b6b9f69146100db5780636daefab6146100ee575b600080fd5b6100bb6100b636600461099c565b610159565b005b6100c56101b8565b6040516100d29190610aa6565b60405180910390f35b6100bb6100e936600461099c565b6101be565b6100bb6100fc3660046107d1565b610211565b6100c561010f366004610924565b6102a3565b61011c6102f5565b6040516100d29190610a35565b6100c5610304565b61014461013f366004610852565b61030a565b6040516100d29190610a62565b61011c610411565b610161610435565b6101786706f05b59d3b2000082111561025861047e565b60018190556040517fa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc906101ad908390610aa6565b60405180910390a150565b60015490565b6101c6610435565b6101dc662386f26fc1000082111561025961047e565b60028190556040517f5a0b7386237e7f07fa741efc64e59c9387d2cccafec760efed4d53387f20e19a906101ad908390610aa6565b610219610490565b610221610435565b61022b84836104a9565b60005b8481101561029357600086868381811061024457fe5b90506020020160208101906102599190610980565b9050600085858481811061026957fe5b6020029190910135915061028990506001600160a01b03831685836104b6565b505060010161022e565b5061029c61053e565b5050505050565b60007f0000000000000000000000000000000000000000000000000000000000000000826040516020016102d89291906109cc565b604051602081830303815290604052805190602001209050919050565b60006102ff610545565b905090565b60025490565b6060815167ffffffffffffffff8111801561032457600080fd5b5060405190808252806020026020018201604052801561034e578160200160208202803683370190505b50905060005b825181101561040b5782818151811061036957fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161039c9190610a35565b60206040518083038186803b1580156103b457600080fd5b505afa1580156103c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ec91906109b4565b8282815181106103f857fe5b6020908102919091010152600101610354565b50919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006104646000357fffffffff00000000000000000000000000000000000000000000000000000000166102a3565b905061047b61047382336105d8565b61019161047e565b50565b8161048c5761048c8161066a565b5050565b6104a26002600054141561019061047e565b6002600055565b61048c818314606761047e565b6105398363a9059cbb60e01b84846040516024016104d5929190610a49565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526106d7565b505050565b6001600055565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663aaabadc56040518163ffffffff1660e01b815260040160206040518083038186803b1580156105a057600080fd5b505afa1580156105b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ff9190610964565b60006105e2610545565b6001600160a01b0316639be2a8848484306040518463ffffffff1660e01b815260040161061193929190610aaf565b60206040518083038186803b15801561062957600080fd5b505afa15801561063d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066191906108fd565b90505b92915050565b7f08c379a0000000000000000000000000000000000000000000000000000000006000908152602060045260076024526642414c23000030600a808404818106603090810160081b95839006959095019082900491820690940160101b939093010160c81b604452606490fd5b60006060836001600160a01b0316836040516106f391906109fc565b6000604051808303816000865af19150503d8060008114610730576040519150601f19603f3d011682016040523d82523d6000602084013e610735565b606091505b5091509150600082141561074d573d6000803e3d6000fd5b61077781516000148061076f57508180602001905181019061076f91906108fd565b6101a261047e565b50505050565b60008083601f84011261078e578182fd5b50813567ffffffffffffffff8111156107a5578182fd5b60208301915083602080830285010111156107bf57600080fd5b9250929050565b803561066481610af5565b6000806000806000606086880312156107e8578081fd5b853567ffffffffffffffff808211156107ff578283fd5b61080b89838a0161077d565b90975095506020880135915080821115610823578283fd5b506108308882890161077d565b909450925050604086013561084481610af5565b809150509295509295909350565b60006020808385031215610864578182fd5b823567ffffffffffffffff8082111561087b578384fd5b818501915085601f83011261088e578384fd5b81358181111561089c578485fd5b83810291506108ac848301610ace565b8181528481019084860184860187018a10156108c6578788fd5b8795505b838610156108f0576108dc8a826107c6565b8352600195909501949186019186016108ca565b5098975050505050505050565b60006020828403121561090e578081fd5b8151801515811461091d578182fd5b9392505050565b600060208284031215610935578081fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461091d578182fd5b600060208284031215610975578081fd5b815161091d81610af5565b600060208284031215610991578081fd5b813561091d81610af5565b6000602082840312156109ad578081fd5b5035919050565b6000602082840312156109c5578081fd5b5051919050565b9182527fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60008251815b81811015610a1c5760208186018101518583015201610a02565b81811115610a2a5782828501525b509190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6020808252825182820181905260009190848201906040850190845b81811015610a9a57835183529284019291840191600101610a7e565b50909695505050505050565b90815260200190565b9283526001600160a01b03918216602084015216604082015260600190565b60405181810167ffffffffffffffff81118282101715610aed57600080fd5b604052919050565b6001600160a01b038116811461047b57600080fdfea2646970667358221220be72bdf8e7a3c38606c5f954fbe2d77798347aaa1cfb76fe77ec2f6c245d24bc64736f6c63430007010033000000000000000000000000a331d84ec860bf466b4cdccfb4ac09a1b43f3ae6000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000000076a7000000000000000000000000000000000000000000000000000000000000278d00",
}

// ThevaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ThevaultMetaData.ABI instead.
var ThevaultABI = ThevaultMetaData.ABI

// ThevaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ThevaultMetaData.Bin instead.
var ThevaultBin = ThevaultMetaData.Bin

// DeployThevault deploys a new Ethereum contract, binding an instance of Thevault to it.
func DeployThevault(auth *bind.TransactOpts, backend bind.ContractBackend, authorizer common.Address, weth common.Address, pauseWindowDuration *big.Int, bufferPeriodDuration *big.Int) (common.Address, *types.Transaction, *Thevault, error) {
	parsed, err := ThevaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ThevaultBin), backend, authorizer, weth, pauseWindowDuration, bufferPeriodDuration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Thevault{ThevaultCaller: ThevaultCaller{contract: contract}, ThevaultTransactor: ThevaultTransactor{contract: contract}, ThevaultFilterer: ThevaultFilterer{contract: contract}}, nil
}

// Thevault is an auto generated Go binding around an Ethereum contract.
type Thevault struct {
	ThevaultCaller     // Read-only binding to the contract
	ThevaultTransactor // Write-only binding to the contract
	ThevaultFilterer   // Log filterer for contract events
}

// ThevaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThevaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThevaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThevaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThevaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThevaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThevaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThevaultSession struct {
	Contract     *Thevault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThevaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThevaultCallerSession struct {
	Contract *ThevaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ThevaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThevaultTransactorSession struct {
	Contract     *ThevaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ThevaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThevaultRaw struct {
	Contract *Thevault // Generic contract binding to access the raw methods on
}

// ThevaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThevaultCallerRaw struct {
	Contract *ThevaultCaller // Generic read-only contract binding to access the raw methods on
}

// ThevaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThevaultTransactorRaw struct {
	Contract *ThevaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThevault creates a new instance of Thevault, bound to a specific deployed contract.
func NewThevault(address common.Address, backend bind.ContractBackend) (*Thevault, error) {
	contract, err := bindThevault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Thevault{ThevaultCaller: ThevaultCaller{contract: contract}, ThevaultTransactor: ThevaultTransactor{contract: contract}, ThevaultFilterer: ThevaultFilterer{contract: contract}}, nil
}

// NewThevaultCaller creates a new read-only instance of Thevault, bound to a specific deployed contract.
func NewThevaultCaller(address common.Address, caller bind.ContractCaller) (*ThevaultCaller, error) {
	contract, err := bindThevault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThevaultCaller{contract: contract}, nil
}

// NewThevaultTransactor creates a new write-only instance of Thevault, bound to a specific deployed contract.
func NewThevaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ThevaultTransactor, error) {
	contract, err := bindThevault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThevaultTransactor{contract: contract}, nil
}

// NewThevaultFilterer creates a new log filterer instance of Thevault, bound to a specific deployed contract.
func NewThevaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ThevaultFilterer, error) {
	contract, err := bindThevault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThevaultFilterer{contract: contract}, nil
}

// bindThevault binds a generic wrapper to an already deployed contract.
func bindThevault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThevaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thevault *ThevaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thevault.Contract.ThevaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thevault *ThevaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thevault.Contract.ThevaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thevault *ThevaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thevault.Contract.ThevaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thevault *ThevaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thevault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thevault *ThevaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thevault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thevault *ThevaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thevault.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Thevault *ThevaultCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Thevault *ThevaultSession) WETH() (common.Address, error) {
	return _Thevault.Contract.WETH(&_Thevault.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Thevault *ThevaultCallerSession) WETH() (common.Address, error) {
	return _Thevault.Contract.WETH(&_Thevault.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Thevault *ThevaultCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Thevault *ThevaultSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Thevault.Contract.GetActionId(&_Thevault.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Thevault *ThevaultCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Thevault.Contract.GetActionId(&_Thevault.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Thevault *ThevaultCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Thevault *ThevaultSession) GetAuthorizer() (common.Address, error) {
	return _Thevault.Contract.GetAuthorizer(&_Thevault.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Thevault *ThevaultCallerSession) GetAuthorizer() (common.Address, error) {
	return _Thevault.Contract.GetAuthorizer(&_Thevault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Thevault *ThevaultCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Thevault *ThevaultSession) GetDomainSeparator() ([32]byte, error) {
	return _Thevault.Contract.GetDomainSeparator(&_Thevault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Thevault *ThevaultCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Thevault.Contract.GetDomainSeparator(&_Thevault.CallOpts)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Thevault *ThevaultCaller) GetInternalBalance(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getInternalBalance", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Thevault *ThevaultSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Thevault.Contract.GetInternalBalance(&_Thevault.CallOpts, user, tokens)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Thevault *ThevaultCallerSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Thevault.Contract.GetInternalBalance(&_Thevault.CallOpts, user, tokens)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Thevault *ThevaultCaller) GetNextNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getNextNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Thevault *ThevaultSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Thevault.Contract.GetNextNonce(&_Thevault.CallOpts, user)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Thevault *ThevaultCallerSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Thevault.Contract.GetNextNonce(&_Thevault.CallOpts, user)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Thevault *ThevaultCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getPausedState")

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
func (_Thevault *ThevaultSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Thevault.Contract.GetPausedState(&_Thevault.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Thevault *ThevaultCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Thevault.Contract.GetPausedState(&_Thevault.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Thevault *ThevaultCaller) GetPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, uint8, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getPool", poolId)

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
func (_Thevault *ThevaultSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Thevault.Contract.GetPool(&_Thevault.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Thevault *ThevaultCallerSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Thevault.Contract.GetPool(&_Thevault.CallOpts, poolId)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Thevault *ThevaultCaller) GetPoolTokenInfo(opts *bind.CallOpts, poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getPoolTokenInfo", poolId, token)

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
func (_Thevault *ThevaultSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Thevault.Contract.GetPoolTokenInfo(&_Thevault.CallOpts, poolId, token)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Thevault *ThevaultCallerSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Thevault.Contract.GetPoolTokenInfo(&_Thevault.CallOpts, poolId, token)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Thevault *ThevaultCaller) GetPoolTokens(opts *bind.CallOpts, poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getPoolTokens", poolId)

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
func (_Thevault *ThevaultSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Thevault.Contract.GetPoolTokens(&_Thevault.CallOpts, poolId)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Thevault *ThevaultCallerSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Thevault.Contract.GetPoolTokens(&_Thevault.CallOpts, poolId)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Thevault *ThevaultCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Thevault *ThevaultSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Thevault.Contract.GetProtocolFeesCollector(&_Thevault.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Thevault *ThevaultCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Thevault.Contract.GetProtocolFeesCollector(&_Thevault.CallOpts)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Thevault *ThevaultCaller) HasApprovedRelayer(opts *bind.CallOpts, user common.Address, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Thevault.contract.Call(opts, &out, "hasApprovedRelayer", user, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Thevault *ThevaultSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Thevault.Contract.HasApprovedRelayer(&_Thevault.CallOpts, user, relayer)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Thevault *ThevaultCallerSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Thevault.Contract.HasApprovedRelayer(&_Thevault.CallOpts, user, relayer)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Thevault *ThevaultTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "batchSwap", kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Thevault *ThevaultSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.Contract.BatchSwap(&_Thevault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Thevault *ThevaultTransactorSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.Contract.BatchSwap(&_Thevault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Thevault *ThevaultTransactor) DeregisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "deregisterTokens", poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Thevault *ThevaultSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.DeregisterTokens(&_Thevault.TransactOpts, poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Thevault *ThevaultTransactorSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.DeregisterTokens(&_Thevault.TransactOpts, poolId, tokens)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Thevault *ThevaultTransactor) ExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "exitPool", poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Thevault *ThevaultSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Thevault.Contract.ExitPool(&_Thevault.TransactOpts, poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Thevault *ThevaultTransactorSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Thevault.Contract.ExitPool(&_Thevault.TransactOpts, poolId, sender, recipient, request)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Thevault *ThevaultTransactor) FlashLoan(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "flashLoan", recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Thevault *ThevaultSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Thevault.Contract.FlashLoan(&_Thevault.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Thevault *ThevaultTransactorSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Thevault.Contract.FlashLoan(&_Thevault.TransactOpts, recipient, tokens, amounts, userData)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Thevault *ThevaultTransactor) JoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "joinPool", poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Thevault *ThevaultSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Thevault.Contract.JoinPool(&_Thevault.TransactOpts, poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Thevault *ThevaultTransactorSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Thevault.Contract.JoinPool(&_Thevault.TransactOpts, poolId, sender, recipient, request)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Thevault *ThevaultTransactor) ManagePoolBalance(opts *bind.TransactOpts, ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "managePoolBalance", ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Thevault *ThevaultSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Thevault.Contract.ManagePoolBalance(&_Thevault.TransactOpts, ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Thevault *ThevaultTransactorSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Thevault.Contract.ManagePoolBalance(&_Thevault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Thevault *ThevaultTransactor) ManageUserBalance(opts *bind.TransactOpts, ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "manageUserBalance", ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Thevault *ThevaultSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Thevault.Contract.ManageUserBalance(&_Thevault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Thevault *ThevaultTransactorSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Thevault.Contract.ManageUserBalance(&_Thevault.TransactOpts, ops)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Thevault *ThevaultTransactor) QueryBatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "queryBatchSwap", kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Thevault *ThevaultSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Thevault.Contract.QueryBatchSwap(&_Thevault.TransactOpts, kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Thevault *ThevaultTransactorSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Thevault.Contract.QueryBatchSwap(&_Thevault.TransactOpts, kind, swaps, assets, funds)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Thevault *ThevaultTransactor) RegisterPool(opts *bind.TransactOpts, specialization uint8) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "registerPool", specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Thevault *ThevaultSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Thevault.Contract.RegisterPool(&_Thevault.TransactOpts, specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Thevault *ThevaultTransactorSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Thevault.Contract.RegisterPool(&_Thevault.TransactOpts, specialization)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Thevault *ThevaultTransactor) RegisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "registerTokens", poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Thevault *ThevaultSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.RegisterTokens(&_Thevault.TransactOpts, poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Thevault *ThevaultTransactorSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.RegisterTokens(&_Thevault.TransactOpts, poolId, tokens, assetManagers)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Thevault *ThevaultTransactor) SetAuthorizer(opts *bind.TransactOpts, newAuthorizer common.Address) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "setAuthorizer", newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Thevault *ThevaultSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.SetAuthorizer(&_Thevault.TransactOpts, newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Thevault *ThevaultTransactorSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Thevault.Contract.SetAuthorizer(&_Thevault.TransactOpts, newAuthorizer)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Thevault *ThevaultTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Thevault *ThevaultSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Thevault.Contract.SetPaused(&_Thevault.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Thevault *ThevaultTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Thevault.Contract.SetPaused(&_Thevault.TransactOpts, paused)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Thevault *ThevaultTransactor) SetRelayerApproval(opts *bind.TransactOpts, sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "setRelayerApproval", sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Thevault *ThevaultSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Thevault.Contract.SetRelayerApproval(&_Thevault.TransactOpts, sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Thevault *ThevaultTransactorSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Thevault.Contract.SetRelayerApproval(&_Thevault.TransactOpts, sender, relayer, approved)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Thevault *ThevaultTransactor) Swap(opts *bind.TransactOpts, singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.contract.Transact(opts, "swap", singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Thevault *ThevaultSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.Contract.Swap(&_Thevault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Thevault *ThevaultTransactorSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Thevault.Contract.Swap(&_Thevault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Thevault *ThevaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thevault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Thevault *ThevaultSession) Receive() (*types.Transaction, error) {
	return _Thevault.Contract.Receive(&_Thevault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Thevault *ThevaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Thevault.Contract.Receive(&_Thevault.TransactOpts)
}

// ThevaultAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Thevault contract.
type ThevaultAuthorizerChangedIterator struct {
	Event *ThevaultAuthorizerChanged // Event containing the contract specifics and raw log

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
func (it *ThevaultAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultAuthorizerChanged)
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
		it.Event = new(ThevaultAuthorizerChanged)
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
func (it *ThevaultAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultAuthorizerChanged represents a AuthorizerChanged event raised by the Thevault contract.
type ThevaultAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Thevault *ThevaultFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*ThevaultAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultAuthorizerChangedIterator{contract: _Thevault.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Thevault *ThevaultFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *ThevaultAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultAuthorizerChanged)
				if err := _Thevault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseAuthorizerChanged(log types.Log) (*ThevaultAuthorizerChanged, error) {
	event := new(ThevaultAuthorizerChanged)
	if err := _Thevault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultExternalBalanceTransferIterator is returned from FilterExternalBalanceTransfer and is used to iterate over the raw logs and unpacked data for ExternalBalanceTransfer events raised by the Thevault contract.
type ThevaultExternalBalanceTransferIterator struct {
	Event *ThevaultExternalBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *ThevaultExternalBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultExternalBalanceTransfer)
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
		it.Event = new(ThevaultExternalBalanceTransfer)
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
func (it *ThevaultExternalBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultExternalBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultExternalBalanceTransfer represents a ExternalBalanceTransfer event raised by the Thevault contract.
type ThevaultExternalBalanceTransfer struct {
	Token     common.Address
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExternalBalanceTransfer is a free log retrieval operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Thevault *ThevaultFilterer) FilterExternalBalanceTransfer(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*ThevaultExternalBalanceTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultExternalBalanceTransferIterator{contract: _Thevault.contract, event: "ExternalBalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchExternalBalanceTransfer is a free log subscription operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Thevault *ThevaultFilterer) WatchExternalBalanceTransfer(opts *bind.WatchOpts, sink chan<- *ThevaultExternalBalanceTransfer, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultExternalBalanceTransfer)
				if err := _Thevault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseExternalBalanceTransfer(log types.Log) (*ThevaultExternalBalanceTransfer, error) {
	event := new(ThevaultExternalBalanceTransfer)
	if err := _Thevault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Thevault contract.
type ThevaultFlashLoanIterator struct {
	Event *ThevaultFlashLoan // Event containing the contract specifics and raw log

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
func (it *ThevaultFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultFlashLoan)
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
		it.Event = new(ThevaultFlashLoan)
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
func (it *ThevaultFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultFlashLoan represents a FlashLoan event raised by the Thevault contract.
type ThevaultFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Thevault *ThevaultFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*ThevaultFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultFlashLoanIterator{contract: _Thevault.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Thevault *ThevaultFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ThevaultFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultFlashLoan)
				if err := _Thevault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseFlashLoan(log types.Log) (*ThevaultFlashLoan, error) {
	event := new(ThevaultFlashLoan)
	if err := _Thevault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultInternalBalanceChangedIterator is returned from FilterInternalBalanceChanged and is used to iterate over the raw logs and unpacked data for InternalBalanceChanged events raised by the Thevault contract.
type ThevaultInternalBalanceChangedIterator struct {
	Event *ThevaultInternalBalanceChanged // Event containing the contract specifics and raw log

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
func (it *ThevaultInternalBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultInternalBalanceChanged)
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
		it.Event = new(ThevaultInternalBalanceChanged)
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
func (it *ThevaultInternalBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultInternalBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultInternalBalanceChanged represents a InternalBalanceChanged event raised by the Thevault contract.
type ThevaultInternalBalanceChanged struct {
	User  common.Address
	Token common.Address
	Delta *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInternalBalanceChanged is a free log retrieval operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Thevault *ThevaultFilterer) FilterInternalBalanceChanged(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*ThevaultInternalBalanceChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultInternalBalanceChangedIterator{contract: _Thevault.contract, event: "InternalBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchInternalBalanceChanged is a free log subscription operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Thevault *ThevaultFilterer) WatchInternalBalanceChanged(opts *bind.WatchOpts, sink chan<- *ThevaultInternalBalanceChanged, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultInternalBalanceChanged)
				if err := _Thevault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseInternalBalanceChanged(log types.Log) (*ThevaultInternalBalanceChanged, error) {
	event := new(ThevaultInternalBalanceChanged)
	if err := _Thevault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the Thevault contract.
type ThevaultPausedStateChangedIterator struct {
	Event *ThevaultPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *ThevaultPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultPausedStateChanged)
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
		it.Event = new(ThevaultPausedStateChanged)
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
func (it *ThevaultPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultPausedStateChanged represents a PausedStateChanged event raised by the Thevault contract.
type ThevaultPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Thevault *ThevaultFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*ThevaultPausedStateChangedIterator, error) {

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &ThevaultPausedStateChangedIterator{contract: _Thevault.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Thevault *ThevaultFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *ThevaultPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultPausedStateChanged)
				if err := _Thevault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParsePausedStateChanged(log types.Log) (*ThevaultPausedStateChanged, error) {
	event := new(ThevaultPausedStateChanged)
	if err := _Thevault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultPoolBalanceChangedIterator is returned from FilterPoolBalanceChanged and is used to iterate over the raw logs and unpacked data for PoolBalanceChanged events raised by the Thevault contract.
type ThevaultPoolBalanceChangedIterator struct {
	Event *ThevaultPoolBalanceChanged // Event containing the contract specifics and raw log

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
func (it *ThevaultPoolBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultPoolBalanceChanged)
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
		it.Event = new(ThevaultPoolBalanceChanged)
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
func (it *ThevaultPoolBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultPoolBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultPoolBalanceChanged represents a PoolBalanceChanged event raised by the Thevault contract.
type ThevaultPoolBalanceChanged struct {
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
func (_Thevault *ThevaultFilterer) FilterPoolBalanceChanged(opts *bind.FilterOpts, poolId [][32]byte, liquidityProvider []common.Address) (*ThevaultPoolBalanceChangedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultPoolBalanceChangedIterator{contract: _Thevault.contract, event: "PoolBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceChanged is a free log subscription operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Thevault *ThevaultFilterer) WatchPoolBalanceChanged(opts *bind.WatchOpts, sink chan<- *ThevaultPoolBalanceChanged, poolId [][32]byte, liquidityProvider []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultPoolBalanceChanged)
				if err := _Thevault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParsePoolBalanceChanged(log types.Log) (*ThevaultPoolBalanceChanged, error) {
	event := new(ThevaultPoolBalanceChanged)
	if err := _Thevault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultPoolBalanceManagedIterator is returned from FilterPoolBalanceManaged and is used to iterate over the raw logs and unpacked data for PoolBalanceManaged events raised by the Thevault contract.
type ThevaultPoolBalanceManagedIterator struct {
	Event *ThevaultPoolBalanceManaged // Event containing the contract specifics and raw log

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
func (it *ThevaultPoolBalanceManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultPoolBalanceManaged)
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
		it.Event = new(ThevaultPoolBalanceManaged)
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
func (it *ThevaultPoolBalanceManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultPoolBalanceManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultPoolBalanceManaged represents a PoolBalanceManaged event raised by the Thevault contract.
type ThevaultPoolBalanceManaged struct {
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
func (_Thevault *ThevaultFilterer) FilterPoolBalanceManaged(opts *bind.FilterOpts, poolId [][32]byte, assetManager []common.Address, token []common.Address) (*ThevaultPoolBalanceManagedIterator, error) {

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

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultPoolBalanceManagedIterator{contract: _Thevault.contract, event: "PoolBalanceManaged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceManaged is a free log subscription operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Thevault *ThevaultFilterer) WatchPoolBalanceManaged(opts *bind.WatchOpts, sink chan<- *ThevaultPoolBalanceManaged, poolId [][32]byte, assetManager []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultPoolBalanceManaged)
				if err := _Thevault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParsePoolBalanceManaged(log types.Log) (*ThevaultPoolBalanceManaged, error) {
	event := new(ThevaultPoolBalanceManaged)
	if err := _Thevault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Thevault contract.
type ThevaultPoolRegisteredIterator struct {
	Event *ThevaultPoolRegistered // Event containing the contract specifics and raw log

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
func (it *ThevaultPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultPoolRegistered)
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
		it.Event = new(ThevaultPoolRegistered)
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
func (it *ThevaultPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultPoolRegistered represents a PoolRegistered event raised by the Thevault contract.
type ThevaultPoolRegistered struct {
	PoolId         [32]byte
	PoolAddress    common.Address
	Specialization uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Thevault *ThevaultFilterer) FilterPoolRegistered(opts *bind.FilterOpts, poolId [][32]byte, poolAddress []common.Address) (*ThevaultPoolRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultPoolRegisteredIterator{contract: _Thevault.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Thevault *ThevaultFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *ThevaultPoolRegistered, poolId [][32]byte, poolAddress []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultPoolRegistered)
				if err := _Thevault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParsePoolRegistered(log types.Log) (*ThevaultPoolRegistered, error) {
	event := new(ThevaultPoolRegistered)
	if err := _Thevault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultRelayerApprovalChangedIterator is returned from FilterRelayerApprovalChanged and is used to iterate over the raw logs and unpacked data for RelayerApprovalChanged events raised by the Thevault contract.
type ThevaultRelayerApprovalChangedIterator struct {
	Event *ThevaultRelayerApprovalChanged // Event containing the contract specifics and raw log

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
func (it *ThevaultRelayerApprovalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultRelayerApprovalChanged)
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
		it.Event = new(ThevaultRelayerApprovalChanged)
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
func (it *ThevaultRelayerApprovalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultRelayerApprovalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultRelayerApprovalChanged represents a RelayerApprovalChanged event raised by the Thevault contract.
type ThevaultRelayerApprovalChanged struct {
	Relayer  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerApprovalChanged is a free log retrieval operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Thevault *ThevaultFilterer) FilterRelayerApprovalChanged(opts *bind.FilterOpts, relayer []common.Address, sender []common.Address) (*ThevaultRelayerApprovalChangedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultRelayerApprovalChangedIterator{contract: _Thevault.contract, event: "RelayerApprovalChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerApprovalChanged is a free log subscription operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Thevault *ThevaultFilterer) WatchRelayerApprovalChanged(opts *bind.WatchOpts, sink chan<- *ThevaultRelayerApprovalChanged, relayer []common.Address, sender []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultRelayerApprovalChanged)
				if err := _Thevault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseRelayerApprovalChanged(log types.Log) (*ThevaultRelayerApprovalChanged, error) {
	event := new(ThevaultRelayerApprovalChanged)
	if err := _Thevault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Thevault contract.
type ThevaultSwapIterator struct {
	Event *ThevaultSwap // Event containing the contract specifics and raw log

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
func (it *ThevaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultSwap)
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
		it.Event = new(ThevaultSwap)
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
func (it *ThevaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultSwap represents a Swap event raised by the Thevault contract.
type ThevaultSwap struct {
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
func (_Thevault *ThevaultFilterer) FilterSwap(opts *bind.FilterOpts, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (*ThevaultSwapIterator, error) {

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

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultSwapIterator{contract: _Thevault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Thevault *ThevaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ThevaultSwap, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultSwap)
				if err := _Thevault.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseSwap(log types.Log) (*ThevaultSwap, error) {
	event := new(ThevaultSwap)
	if err := _Thevault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultTokensDeregisteredIterator is returned from FilterTokensDeregistered and is used to iterate over the raw logs and unpacked data for TokensDeregistered events raised by the Thevault contract.
type ThevaultTokensDeregisteredIterator struct {
	Event *ThevaultTokensDeregistered // Event containing the contract specifics and raw log

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
func (it *ThevaultTokensDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultTokensDeregistered)
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
		it.Event = new(ThevaultTokensDeregistered)
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
func (it *ThevaultTokensDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultTokensDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultTokensDeregistered represents a TokensDeregistered event raised by the Thevault contract.
type ThevaultTokensDeregistered struct {
	PoolId [32]byte
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensDeregistered is a free log retrieval operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Thevault *ThevaultFilterer) FilterTokensDeregistered(opts *bind.FilterOpts, poolId [][32]byte) (*ThevaultTokensDeregisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultTokensDeregisteredIterator{contract: _Thevault.contract, event: "TokensDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokensDeregistered is a free log subscription operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Thevault *ThevaultFilterer) WatchTokensDeregistered(opts *bind.WatchOpts, sink chan<- *ThevaultTokensDeregistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultTokensDeregistered)
				if err := _Thevault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseTokensDeregistered(log types.Log) (*ThevaultTokensDeregistered, error) {
	event := new(ThevaultTokensDeregistered)
	if err := _Thevault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThevaultTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the Thevault contract.
type ThevaultTokensRegisteredIterator struct {
	Event *ThevaultTokensRegistered // Event containing the contract specifics and raw log

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
func (it *ThevaultTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThevaultTokensRegistered)
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
		it.Event = new(ThevaultTokensRegistered)
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
func (it *ThevaultTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThevaultTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThevaultTokensRegistered represents a TokensRegistered event raised by the Thevault contract.
type ThevaultTokensRegistered struct {
	PoolId        [32]byte
	Tokens        []common.Address
	AssetManagers []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Thevault *ThevaultFilterer) FilterTokensRegistered(opts *bind.FilterOpts, poolId [][32]byte) (*ThevaultTokensRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Thevault.contract.FilterLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &ThevaultTokensRegisteredIterator{contract: _Thevault.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Thevault *ThevaultFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *ThevaultTokensRegistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Thevault.contract.WatchLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThevaultTokensRegistered)
				if err := _Thevault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
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
func (_Thevault *ThevaultFilterer) ParseTokensRegistered(log types.Log) (*ThevaultTokensRegistered, error) {
	event := new(ThevaultTokensRegistered)
	if err := _Thevault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
