// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DelegatorPoolVestingSchedule is an auto generated low-level Go binding around an user-defined struct.
type DelegatorPoolVestingSchedule struct {
	TotalAmount     *big.Int
	StartTimestamp  uint64
	DurationSeconds uint64
}

// DelegatorPoolABI is the input ABI used to generate the binding from.
const DelegatorPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allowanceToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_dTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_feeCurve\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsPool\",\"type\":\"address\"}],\"name\":\"AddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AllowanceStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AllowanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"RedirectApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"RedirectApprovalRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsPool\",\"type\":\"address\"}],\"name\":\"RemoveToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"RewardsRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawRewards\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsPool\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"approveRedirect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"}],\"name\":\"currentRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_percentageBorrowed\",\"type\":\"uint256\"}],\"name\":\"currentRateAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"distributeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"distributeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCurve\",\"outputs\":[{\"internalType\":\"contractIFeeCurve\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getVestingSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"durationSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatorPool.VestingSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolRouter\",\"outputs\":[{\"internalType\":\"contractIPoolRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redirectApprovals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectRewardsFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redirectedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeRedirectApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardRedirects\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCurve\",\"type\":\"address\"}],\"name\":\"setFeeCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolRouter\",\"type\":\"address\"}],\"name\":\"setPoolRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenPools\",\"outputs\":[{\"internalType\":\"contractIRewardsPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawableRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DelegatorPool is an auto generated Go binding around an Ethereum contract.
type DelegatorPool struct {
	DelegatorPoolCaller     // Read-only binding to the contract
	DelegatorPoolTransactor // Write-only binding to the contract
	DelegatorPoolFilterer   // Log filterer for contract events
}

// DelegatorPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatorPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatorPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatorPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatorPoolSession struct {
	Contract     *DelegatorPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatorPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatorPoolCallerSession struct {
	Contract *DelegatorPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DelegatorPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatorPoolTransactorSession struct {
	Contract     *DelegatorPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DelegatorPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatorPoolRaw struct {
	Contract *DelegatorPool // Generic contract binding to access the raw methods on
}

// DelegatorPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatorPoolCallerRaw struct {
	Contract *DelegatorPoolCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatorPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatorPoolTransactorRaw struct {
	Contract *DelegatorPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegatorPool creates a new instance of DelegatorPool, bound to a specific deployed contract.
func NewDelegatorPool(address common.Address, backend bind.ContractBackend) (*DelegatorPool, error) {
	contract, err := bindDelegatorPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegatorPool{DelegatorPoolCaller: DelegatorPoolCaller{contract: contract}, DelegatorPoolTransactor: DelegatorPoolTransactor{contract: contract}, DelegatorPoolFilterer: DelegatorPoolFilterer{contract: contract}}, nil
}

// NewDelegatorPoolCaller creates a new read-only instance of DelegatorPool, bound to a specific deployed contract.
func NewDelegatorPoolCaller(address common.Address, caller bind.ContractCaller) (*DelegatorPoolCaller, error) {
	contract, err := bindDelegatorPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolCaller{contract: contract}, nil
}

// NewDelegatorPoolTransactor creates a new write-only instance of DelegatorPool, bound to a specific deployed contract.
func NewDelegatorPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatorPoolTransactor, error) {
	contract, err := bindDelegatorPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolTransactor{contract: contract}, nil
}

// NewDelegatorPoolFilterer creates a new log filterer instance of DelegatorPool, bound to a specific deployed contract.
func NewDelegatorPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatorPoolFilterer, error) {
	contract, err := bindDelegatorPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolFilterer{contract: contract}, nil
}

// bindDelegatorPool binds a generic wrapper to an already deployed contract.
func bindDelegatorPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegatorPool *DelegatorPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegatorPool.Contract.DelegatorPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegatorPool *DelegatorPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DelegatorPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegatorPool *DelegatorPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DelegatorPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegatorPool *DelegatorPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegatorPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegatorPool *DelegatorPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatorPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegatorPool *DelegatorPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegatorPool.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.Allowance(&_DelegatorPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.Allowance(&_DelegatorPool.CallOpts, owner, spender)
}

// AllowanceToken is a free data retrieval call binding the contract method 0xd5d67d51.
//
// Solidity: function allowanceToken() view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) AllowanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "allowanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowanceToken is a free data retrieval call binding the contract method 0xd5d67d51.
//
// Solidity: function allowanceToken() view returns(address)
func (_DelegatorPool *DelegatorPoolSession) AllowanceToken() (common.Address, error) {
	return _DelegatorPool.Contract.AllowanceToken(&_DelegatorPool.CallOpts)
}

// AllowanceToken is a free data retrieval call binding the contract method 0xd5d67d51.
//
// Solidity: function allowanceToken() view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) AllowanceToken() (common.Address, error) {
	return _DelegatorPool.Contract.AllowanceToken(&_DelegatorPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.BalanceOf(&_DelegatorPool.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.BalanceOf(&_DelegatorPool.CallOpts, _account)
}

// CurrentRate is a free data retrieval call binding the contract method 0x04e4c864.
//
// Solidity: function currentRate(address _token, uint16 _index) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) CurrentRate(opts *bind.CallOpts, _token common.Address, _index uint16) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "currentRate", _token, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRate is a free data retrieval call binding the contract method 0x04e4c864.
//
// Solidity: function currentRate(address _token, uint16 _index) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) CurrentRate(_token common.Address, _index uint16) (*big.Int, error) {
	return _DelegatorPool.Contract.CurrentRate(&_DelegatorPool.CallOpts, _token, _index)
}

// CurrentRate is a free data retrieval call binding the contract method 0x04e4c864.
//
// Solidity: function currentRate(address _token, uint16 _index) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) CurrentRate(_token common.Address, _index uint16) (*big.Int, error) {
	return _DelegatorPool.Contract.CurrentRate(&_DelegatorPool.CallOpts, _token, _index)
}

// CurrentRateAt is a free data retrieval call binding the contract method 0x6b5f8321.
//
// Solidity: function currentRateAt(uint256 _percentageBorrowed) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) CurrentRateAt(opts *bind.CallOpts, _percentageBorrowed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "currentRateAt", _percentageBorrowed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRateAt is a free data retrieval call binding the contract method 0x6b5f8321.
//
// Solidity: function currentRateAt(uint256 _percentageBorrowed) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) CurrentRateAt(_percentageBorrowed *big.Int) (*big.Int, error) {
	return _DelegatorPool.Contract.CurrentRateAt(&_DelegatorPool.CallOpts, _percentageBorrowed)
}

// CurrentRateAt is a free data retrieval call binding the contract method 0x6b5f8321.
//
// Solidity: function currentRateAt(uint256 _percentageBorrowed) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) CurrentRateAt(_percentageBorrowed *big.Int) (*big.Int, error) {
	return _DelegatorPool.Contract.CurrentRateAt(&_DelegatorPool.CallOpts, _percentageBorrowed)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelegatorPool *DelegatorPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelegatorPool *DelegatorPoolSession) Decimals() (uint8, error) {
	return _DelegatorPool.Contract.Decimals(&_DelegatorPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelegatorPool *DelegatorPoolCallerSession) Decimals() (uint8, error) {
	return _DelegatorPool.Contract.Decimals(&_DelegatorPool.CallOpts)
}

// FeeCurve is a free data retrieval call binding the contract method 0x856695ea.
//
// Solidity: function feeCurve() view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) FeeCurve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "feeCurve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCurve is a free data retrieval call binding the contract method 0x856695ea.
//
// Solidity: function feeCurve() view returns(address)
func (_DelegatorPool *DelegatorPoolSession) FeeCurve() (common.Address, error) {
	return _DelegatorPool.Contract.FeeCurve(&_DelegatorPool.CallOpts)
}

// FeeCurve is a free data retrieval call binding the contract method 0x856695ea.
//
// Solidity: function feeCurve() view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) FeeCurve() (common.Address, error) {
	return _DelegatorPool.Contract.FeeCurve(&_DelegatorPool.CallOpts)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _account) view returns((uint256,uint64,uint64))
func (_DelegatorPool *DelegatorPoolCaller) GetVestingSchedule(opts *bind.CallOpts, _account common.Address) (DelegatorPoolVestingSchedule, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "getVestingSchedule", _account)

	if err != nil {
		return *new(DelegatorPoolVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(DelegatorPoolVestingSchedule)).(*DelegatorPoolVestingSchedule)

	return out0, err

}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _account) view returns((uint256,uint64,uint64))
func (_DelegatorPool *DelegatorPoolSession) GetVestingSchedule(_account common.Address) (DelegatorPoolVestingSchedule, error) {
	return _DelegatorPool.Contract.GetVestingSchedule(&_DelegatorPool.CallOpts, _account)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _account) view returns((uint256,uint64,uint64))
func (_DelegatorPool *DelegatorPoolCallerSession) GetVestingSchedule(_account common.Address) (DelegatorPoolVestingSchedule, error) {
	return _DelegatorPool.Contract.GetVestingSchedule(&_DelegatorPool.CallOpts, _account)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_DelegatorPool *DelegatorPoolCaller) IsTokenSupported(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "isTokenSupported", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_DelegatorPool *DelegatorPoolSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _DelegatorPool.Contract.IsTokenSupported(&_DelegatorPool.CallOpts, _token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_DelegatorPool *DelegatorPoolCallerSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _DelegatorPool.Contract.IsTokenSupported(&_DelegatorPool.CallOpts, _token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelegatorPool *DelegatorPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelegatorPool *DelegatorPoolSession) Name() (string, error) {
	return _DelegatorPool.Contract.Name(&_DelegatorPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelegatorPool *DelegatorPoolCallerSession) Name() (string, error) {
	return _DelegatorPool.Contract.Name(&_DelegatorPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegatorPool *DelegatorPoolSession) Owner() (common.Address, error) {
	return _DelegatorPool.Contract.Owner(&_DelegatorPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) Owner() (common.Address, error) {
	return _DelegatorPool.Contract.Owner(&_DelegatorPool.CallOpts)
}

// PoolRouter is a free data retrieval call binding the contract method 0x5452b9ba.
//
// Solidity: function poolRouter() view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) PoolRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "poolRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolRouter is a free data retrieval call binding the contract method 0x5452b9ba.
//
// Solidity: function poolRouter() view returns(address)
func (_DelegatorPool *DelegatorPoolSession) PoolRouter() (common.Address, error) {
	return _DelegatorPool.Contract.PoolRouter(&_DelegatorPool.CallOpts)
}

// PoolRouter is a free data retrieval call binding the contract method 0x5452b9ba.
//
// Solidity: function poolRouter() view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) PoolRouter() (common.Address, error) {
	return _DelegatorPool.Contract.PoolRouter(&_DelegatorPool.CallOpts)
}

// RedirectApprovals is a free data retrieval call binding the contract method 0x611fae02.
//
// Solidity: function redirectApprovals(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) RedirectApprovals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "redirectApprovals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedirectApprovals is a free data retrieval call binding the contract method 0x611fae02.
//
// Solidity: function redirectApprovals(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolSession) RedirectApprovals(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RedirectApprovals(&_DelegatorPool.CallOpts, arg0)
}

// RedirectApprovals is a free data retrieval call binding the contract method 0x611fae02.
//
// Solidity: function redirectApprovals(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) RedirectApprovals(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RedirectApprovals(&_DelegatorPool.CallOpts, arg0)
}

// RedirectedStakes is a free data retrieval call binding the contract method 0x628a0f55.
//
// Solidity: function redirectedStakes(address ) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) RedirectedStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "redirectedStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedirectedStakes is a free data retrieval call binding the contract method 0x628a0f55.
//
// Solidity: function redirectedStakes(address ) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) RedirectedStakes(arg0 common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.RedirectedStakes(&_DelegatorPool.CallOpts, arg0)
}

// RedirectedStakes is a free data retrieval call binding the contract method 0x628a0f55.
//
// Solidity: function redirectedStakes(address ) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) RedirectedStakes(arg0 common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.RedirectedStakes(&_DelegatorPool.CallOpts, arg0)
}

// RewardRedirects is a free data retrieval call binding the contract method 0x5b45fac4.
//
// Solidity: function rewardRedirects(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) RewardRedirects(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "rewardRedirects", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRedirects is a free data retrieval call binding the contract method 0x5b45fac4.
//
// Solidity: function rewardRedirects(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolSession) RewardRedirects(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RewardRedirects(&_DelegatorPool.CallOpts, arg0)
}

// RewardRedirects is a free data retrieval call binding the contract method 0x5b45fac4.
//
// Solidity: function rewardRedirects(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) RewardRedirects(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RewardRedirects(&_DelegatorPool.CallOpts, arg0)
}

// RewardsAddress is a free data retrieval call binding the contract method 0x952db31f.
//
// Solidity: function rewardsAddress(address _account) view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) RewardsAddress(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "rewardsAddress", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsAddress is a free data retrieval call binding the contract method 0x952db31f.
//
// Solidity: function rewardsAddress(address _account) view returns(address)
func (_DelegatorPool *DelegatorPoolSession) RewardsAddress(_account common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RewardsAddress(&_DelegatorPool.CallOpts, _account)
}

// RewardsAddress is a free data retrieval call binding the contract method 0x952db31f.
//
// Solidity: function rewardsAddress(address _account) view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) RewardsAddress(_account common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.RewardsAddress(&_DelegatorPool.CallOpts, _account)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) Staked(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "staked", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) Staked(_account common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.Staked(&_DelegatorPool.CallOpts, _account)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) Staked(_account common.Address) (*big.Int, error) {
	return _DelegatorPool.Contract.Staked(&_DelegatorPool.CallOpts, _account)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_DelegatorPool *DelegatorPoolCaller) SupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "supportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_DelegatorPool *DelegatorPoolSession) SupportedTokens() ([]common.Address, error) {
	return _DelegatorPool.Contract.SupportedTokens(&_DelegatorPool.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_DelegatorPool *DelegatorPoolCallerSession) SupportedTokens() ([]common.Address, error) {
	return _DelegatorPool.Contract.SupportedTokens(&_DelegatorPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelegatorPool *DelegatorPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelegatorPool *DelegatorPoolSession) Symbol() (string, error) {
	return _DelegatorPool.Contract.Symbol(&_DelegatorPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelegatorPool *DelegatorPoolCallerSession) Symbol() (string, error) {
	return _DelegatorPool.Contract.Symbol(&_DelegatorPool.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x0e3bc974.
//
// Solidity: function tokenBalances() view returns(address[], uint256[])
func (_DelegatorPool *DelegatorPoolCaller) TokenBalances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "tokenBalances")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x0e3bc974.
//
// Solidity: function tokenBalances() view returns(address[], uint256[])
func (_DelegatorPool *DelegatorPoolSession) TokenBalances() ([]common.Address, []*big.Int, error) {
	return _DelegatorPool.Contract.TokenBalances(&_DelegatorPool.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x0e3bc974.
//
// Solidity: function tokenBalances() view returns(address[], uint256[])
func (_DelegatorPool *DelegatorPoolCallerSession) TokenBalances() ([]common.Address, []*big.Int, error) {
	return _DelegatorPool.Contract.TokenBalances(&_DelegatorPool.CallOpts)
}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCaller) TokenPools(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "tokenPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolSession) TokenPools(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.TokenPools(&_DelegatorPool.CallOpts, arg0)
}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_DelegatorPool *DelegatorPoolCallerSession) TokenPools(arg0 common.Address) (common.Address, error) {
	return _DelegatorPool.Contract.TokenPools(&_DelegatorPool.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) TotalStaked() (*big.Int, error) {
	return _DelegatorPool.Contract.TotalStaked(&_DelegatorPool.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) TotalStaked() (*big.Int, error) {
	return _DelegatorPool.Contract.TotalStaked(&_DelegatorPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelegatorPool *DelegatorPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelegatorPool *DelegatorPoolSession) TotalSupply() (*big.Int, error) {
	return _DelegatorPool.Contract.TotalSupply(&_DelegatorPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelegatorPool *DelegatorPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _DelegatorPool.Contract.TotalSupply(&_DelegatorPool.CallOpts)
}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_DelegatorPool *DelegatorPoolCaller) WithdrawableRewards(opts *bind.CallOpts, _account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegatorPool.contract.Call(opts, &out, "withdrawableRewards", _account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_DelegatorPool *DelegatorPoolSession) WithdrawableRewards(_account common.Address) ([]*big.Int, error) {
	return _DelegatorPool.Contract.WithdrawableRewards(&_DelegatorPool.CallOpts, _account)
}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_DelegatorPool *DelegatorPoolCallerSession) WithdrawableRewards(_account common.Address) ([]*big.Int, error) {
	return _DelegatorPool.Contract.WithdrawableRewards(&_DelegatorPool.CallOpts, _account)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_DelegatorPool *DelegatorPoolTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "addToken", _token, _rewardsPool)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_DelegatorPool *DelegatorPoolSession) AddToken(_token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.AddToken(&_DelegatorPool.TransactOpts, _token, _rewardsPool)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) AddToken(_token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.AddToken(&_DelegatorPool.TransactOpts, _token, _rewardsPool)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.Approve(&_DelegatorPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.Approve(&_DelegatorPool.TransactOpts, spender, amount)
}

// ApproveRedirect is a paid mutator transaction binding the contract method 0x5b0c9391.
//
// Solidity: function approveRedirect(address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactor) ApproveRedirect(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "approveRedirect", _to)
}

// ApproveRedirect is a paid mutator transaction binding the contract method 0x5b0c9391.
//
// Solidity: function approveRedirect(address _to) returns()
func (_DelegatorPool *DelegatorPoolSession) ApproveRedirect(_to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.ApproveRedirect(&_DelegatorPool.TransactOpts, _to)
}

// ApproveRedirect is a paid mutator transaction binding the contract method 0x5b0c9391.
//
// Solidity: function approveRedirect(address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) ApproveRedirect(_to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.ApproveRedirect(&_DelegatorPool.TransactOpts, _to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DecreaseAllowance(&_DelegatorPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DecreaseAllowance(&_DelegatorPool.TransactOpts, spender, subtractedValue)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolTransactor) DistributeToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "distributeToken", _token)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolSession) DistributeToken(_token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DistributeToken(&_DelegatorPool.TransactOpts, _token)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) DistributeToken(_token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DistributeToken(&_DelegatorPool.TransactOpts, _token)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolTransactor) DistributeTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "distributeTokens", _tokens)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolSession) DistributeTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DistributeTokens(&_DelegatorPool.TransactOpts, _tokens)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) DistributeTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.DistributeTokens(&_DelegatorPool.TransactOpts, _tokens)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.IncreaseAllowance(&_DelegatorPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.IncreaseAllowance(&_DelegatorPool.TransactOpts, spender, addedValue)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_DelegatorPool *DelegatorPoolTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "onTokenTransfer", _sender, _value, _calldata)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_DelegatorPool *DelegatorPoolSession) OnTokenTransfer(_sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _DelegatorPool.Contract.OnTokenTransfer(&_DelegatorPool.TransactOpts, _sender, _value, _calldata)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) OnTokenTransfer(_sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _DelegatorPool.Contract.OnTokenTransfer(&_DelegatorPool.TransactOpts, _sender, _value, _calldata)
}

// RedirectRewards is a paid mutator transaction binding the contract method 0xc53b585e.
//
// Solidity: function redirectRewards(address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactor) RedirectRewards(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "redirectRewards", _to)
}

// RedirectRewards is a paid mutator transaction binding the contract method 0xc53b585e.
//
// Solidity: function redirectRewards(address _to) returns()
func (_DelegatorPool *DelegatorPoolSession) RedirectRewards(_to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RedirectRewards(&_DelegatorPool.TransactOpts, _to)
}

// RedirectRewards is a paid mutator transaction binding the contract method 0xc53b585e.
//
// Solidity: function redirectRewards(address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) RedirectRewards(_to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RedirectRewards(&_DelegatorPool.TransactOpts, _to)
}

// RedirectRewardsFrom is a paid mutator transaction binding the contract method 0xef735f07.
//
// Solidity: function redirectRewardsFrom(address _from, address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactor) RedirectRewardsFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "redirectRewardsFrom", _from, _to)
}

// RedirectRewardsFrom is a paid mutator transaction binding the contract method 0xef735f07.
//
// Solidity: function redirectRewardsFrom(address _from, address _to) returns()
func (_DelegatorPool *DelegatorPoolSession) RedirectRewardsFrom(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RedirectRewardsFrom(&_DelegatorPool.TransactOpts, _from, _to)
}

// RedirectRewardsFrom is a paid mutator transaction binding the contract method 0xef735f07.
//
// Solidity: function redirectRewardsFrom(address _from, address _to) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) RedirectRewardsFrom(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RedirectRewardsFrom(&_DelegatorPool.TransactOpts, _from, _to)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "removeToken", _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RemoveToken(&_DelegatorPool.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.RemoveToken(&_DelegatorPool.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegatorPool *DelegatorPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegatorPool *DelegatorPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegatorPool.Contract.RenounceOwnership(&_DelegatorPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegatorPool.Contract.RenounceOwnership(&_DelegatorPool.TransactOpts)
}

// RevokeRedirectApproval is a paid mutator transaction binding the contract method 0xd387fd6c.
//
// Solidity: function revokeRedirectApproval() returns()
func (_DelegatorPool *DelegatorPoolTransactor) RevokeRedirectApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "revokeRedirectApproval")
}

// RevokeRedirectApproval is a paid mutator transaction binding the contract method 0xd387fd6c.
//
// Solidity: function revokeRedirectApproval() returns()
func (_DelegatorPool *DelegatorPoolSession) RevokeRedirectApproval() (*types.Transaction, error) {
	return _DelegatorPool.Contract.RevokeRedirectApproval(&_DelegatorPool.TransactOpts)
}

// RevokeRedirectApproval is a paid mutator transaction binding the contract method 0xd387fd6c.
//
// Solidity: function revokeRedirectApproval() returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) RevokeRedirectApproval() (*types.Transaction, error) {
	return _DelegatorPool.Contract.RevokeRedirectApproval(&_DelegatorPool.TransactOpts)
}

// SetFeeCurve is a paid mutator transaction binding the contract method 0xc5164881.
//
// Solidity: function setFeeCurve(address _feeCurve) returns()
func (_DelegatorPool *DelegatorPoolTransactor) SetFeeCurve(opts *bind.TransactOpts, _feeCurve common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "setFeeCurve", _feeCurve)
}

// SetFeeCurve is a paid mutator transaction binding the contract method 0xc5164881.
//
// Solidity: function setFeeCurve(address _feeCurve) returns()
func (_DelegatorPool *DelegatorPoolSession) SetFeeCurve(_feeCurve common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.SetFeeCurve(&_DelegatorPool.TransactOpts, _feeCurve)
}

// SetFeeCurve is a paid mutator transaction binding the contract method 0xc5164881.
//
// Solidity: function setFeeCurve(address _feeCurve) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) SetFeeCurve(_feeCurve common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.SetFeeCurve(&_DelegatorPool.TransactOpts, _feeCurve)
}

// SetPoolRouter is a paid mutator transaction binding the contract method 0x0c8d9cc7.
//
// Solidity: function setPoolRouter(address _poolRouter) returns()
func (_DelegatorPool *DelegatorPoolTransactor) SetPoolRouter(opts *bind.TransactOpts, _poolRouter common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "setPoolRouter", _poolRouter)
}

// SetPoolRouter is a paid mutator transaction binding the contract method 0x0c8d9cc7.
//
// Solidity: function setPoolRouter(address _poolRouter) returns()
func (_DelegatorPool *DelegatorPoolSession) SetPoolRouter(_poolRouter common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.SetPoolRouter(&_DelegatorPool.TransactOpts, _poolRouter)
}

// SetPoolRouter is a paid mutator transaction binding the contract method 0x0c8d9cc7.
//
// Solidity: function setPoolRouter(address _poolRouter) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) SetPoolRouter(_poolRouter common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.SetPoolRouter(&_DelegatorPool.TransactOpts, _poolRouter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.Transfer(&_DelegatorPool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.Transfer(&_DelegatorPool.TransactOpts, to, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) TransferAndCall(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "transferAndCall", _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferAndCall(&_DelegatorPool.TransactOpts, _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferAndCall(&_DelegatorPool.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferFrom(&_DelegatorPool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DelegatorPool *DelegatorPoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferFrom(&_DelegatorPool.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegatorPool *DelegatorPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegatorPool *DelegatorPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferOwnership(&_DelegatorPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.TransferOwnership(&_DelegatorPool.TransactOpts, newOwner)
}

// WithdrawAllowance is a paid mutator transaction binding the contract method 0xba650f79.
//
// Solidity: function withdrawAllowance(uint256 _amount) returns()
func (_DelegatorPool *DelegatorPoolTransactor) WithdrawAllowance(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "withdrawAllowance", _amount)
}

// WithdrawAllowance is a paid mutator transaction binding the contract method 0xba650f79.
//
// Solidity: function withdrawAllowance(uint256 _amount) returns()
func (_DelegatorPool *DelegatorPoolSession) WithdrawAllowance(_amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.WithdrawAllowance(&_DelegatorPool.TransactOpts, _amount)
}

// WithdrawAllowance is a paid mutator transaction binding the contract method 0xba650f79.
//
// Solidity: function withdrawAllowance(uint256 _amount) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) WithdrawAllowance(_amount *big.Int) (*types.Transaction, error) {
	return _DelegatorPool.Contract.WithdrawAllowance(&_DelegatorPool.TransactOpts, _amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolTransactor) WithdrawRewards(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.contract.Transact(opts, "withdrawRewards", _tokens)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolSession) WithdrawRewards(_tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.WithdrawRewards(&_DelegatorPool.TransactOpts, _tokens)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_DelegatorPool *DelegatorPoolTransactorSession) WithdrawRewards(_tokens []common.Address) (*types.Transaction, error) {
	return _DelegatorPool.Contract.WithdrawRewards(&_DelegatorPool.TransactOpts, _tokens)
}

// DelegatorPoolAddTokenIterator is returned from FilterAddToken and is used to iterate over the raw logs and unpacked data for AddToken events raised by the DelegatorPool contract.
type DelegatorPoolAddTokenIterator struct {
	Event *DelegatorPoolAddToken // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolAddToken)
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
		it.Event = new(DelegatorPoolAddToken)
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
func (it *DelegatorPoolAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolAddToken represents a AddToken event raised by the DelegatorPool contract.
type DelegatorPoolAddToken struct {
	Token       common.Address
	RewardsPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddToken is a free log retrieval operation binding the contract event 0xdbf34b45b47a653cf4940cccbec765f72d4d63de3237306905bfc0ee28832362.
//
// Solidity: event AddToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) FilterAddToken(opts *bind.FilterOpts, token []common.Address) (*DelegatorPoolAddTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "AddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolAddTokenIterator{contract: _DelegatorPool.contract, event: "AddToken", logs: logs, sub: sub}, nil
}

// WatchAddToken is a free log subscription operation binding the contract event 0xdbf34b45b47a653cf4940cccbec765f72d4d63de3237306905bfc0ee28832362.
//
// Solidity: event AddToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) WatchAddToken(opts *bind.WatchOpts, sink chan<- *DelegatorPoolAddToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "AddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolAddToken)
				if err := _DelegatorPool.contract.UnpackLog(event, "AddToken", log); err != nil {
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

// ParseAddToken is a log parse operation binding the contract event 0xdbf34b45b47a653cf4940cccbec765f72d4d63de3237306905bfc0ee28832362.
//
// Solidity: event AddToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) ParseAddToken(log types.Log) (*DelegatorPoolAddToken, error) {
	event := new(DelegatorPoolAddToken)
	if err := _DelegatorPool.contract.UnpackLog(event, "AddToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolAllowanceStakedIterator is returned from FilterAllowanceStaked and is used to iterate over the raw logs and unpacked data for AllowanceStaked events raised by the DelegatorPool contract.
type DelegatorPoolAllowanceStakedIterator struct {
	Event *DelegatorPoolAllowanceStaked // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolAllowanceStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolAllowanceStaked)
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
		it.Event = new(DelegatorPoolAllowanceStaked)
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
func (it *DelegatorPoolAllowanceStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolAllowanceStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolAllowanceStaked represents a AllowanceStaked event raised by the DelegatorPool contract.
type DelegatorPoolAllowanceStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowanceStaked is a free log retrieval operation binding the contract event 0x60df8ac38837b0ae1b717d64eba4a631341a9365ed654aa2cb3ed7657d65acee.
//
// Solidity: event AllowanceStaked(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) FilterAllowanceStaked(opts *bind.FilterOpts, user []common.Address) (*DelegatorPoolAllowanceStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "AllowanceStaked", userRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolAllowanceStakedIterator{contract: _DelegatorPool.contract, event: "AllowanceStaked", logs: logs, sub: sub}, nil
}

// WatchAllowanceStaked is a free log subscription operation binding the contract event 0x60df8ac38837b0ae1b717d64eba4a631341a9365ed654aa2cb3ed7657d65acee.
//
// Solidity: event AllowanceStaked(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) WatchAllowanceStaked(opts *bind.WatchOpts, sink chan<- *DelegatorPoolAllowanceStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "AllowanceStaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolAllowanceStaked)
				if err := _DelegatorPool.contract.UnpackLog(event, "AllowanceStaked", log); err != nil {
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

// ParseAllowanceStaked is a log parse operation binding the contract event 0x60df8ac38837b0ae1b717d64eba4a631341a9365ed654aa2cb3ed7657d65acee.
//
// Solidity: event AllowanceStaked(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) ParseAllowanceStaked(log types.Log) (*DelegatorPoolAllowanceStaked, error) {
	event := new(DelegatorPoolAllowanceStaked)
	if err := _DelegatorPool.contract.UnpackLog(event, "AllowanceStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolAllowanceWithdrawnIterator is returned from FilterAllowanceWithdrawn and is used to iterate over the raw logs and unpacked data for AllowanceWithdrawn events raised by the DelegatorPool contract.
type DelegatorPoolAllowanceWithdrawnIterator struct {
	Event *DelegatorPoolAllowanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolAllowanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolAllowanceWithdrawn)
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
		it.Event = new(DelegatorPoolAllowanceWithdrawn)
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
func (it *DelegatorPoolAllowanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolAllowanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolAllowanceWithdrawn represents a AllowanceWithdrawn event raised by the DelegatorPool contract.
type DelegatorPoolAllowanceWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowanceWithdrawn is a free log retrieval operation binding the contract event 0xcb4e4a5b4e544d9c587ec3ac0eead830ce8a4cccfe48b8dc331d447c481be897.
//
// Solidity: event AllowanceWithdrawn(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) FilterAllowanceWithdrawn(opts *bind.FilterOpts, user []common.Address) (*DelegatorPoolAllowanceWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "AllowanceWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolAllowanceWithdrawnIterator{contract: _DelegatorPool.contract, event: "AllowanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchAllowanceWithdrawn is a free log subscription operation binding the contract event 0xcb4e4a5b4e544d9c587ec3ac0eead830ce8a4cccfe48b8dc331d447c481be897.
//
// Solidity: event AllowanceWithdrawn(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) WatchAllowanceWithdrawn(opts *bind.WatchOpts, sink chan<- *DelegatorPoolAllowanceWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "AllowanceWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolAllowanceWithdrawn)
				if err := _DelegatorPool.contract.UnpackLog(event, "AllowanceWithdrawn", log); err != nil {
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

// ParseAllowanceWithdrawn is a log parse operation binding the contract event 0xcb4e4a5b4e544d9c587ec3ac0eead830ce8a4cccfe48b8dc331d447c481be897.
//
// Solidity: event AllowanceWithdrawn(address indexed user, uint256 amount)
func (_DelegatorPool *DelegatorPoolFilterer) ParseAllowanceWithdrawn(log types.Log) (*DelegatorPoolAllowanceWithdrawn, error) {
	event := new(DelegatorPoolAllowanceWithdrawn)
	if err := _DelegatorPool.contract.UnpackLog(event, "AllowanceWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DelegatorPool contract.
type DelegatorPoolApprovalIterator struct {
	Event *DelegatorPoolApproval // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolApproval)
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
		it.Event = new(DelegatorPoolApproval)
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
func (it *DelegatorPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolApproval represents a Approval event raised by the DelegatorPool contract.
type DelegatorPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DelegatorPool *DelegatorPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DelegatorPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolApprovalIterator{contract: _DelegatorPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DelegatorPool *DelegatorPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DelegatorPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolApproval)
				if err := _DelegatorPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DelegatorPool *DelegatorPoolFilterer) ParseApproval(log types.Log) (*DelegatorPoolApproval, error) {
	event := new(DelegatorPoolApproval)
	if err := _DelegatorPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelegatorPool contract.
type DelegatorPoolOwnershipTransferredIterator struct {
	Event *DelegatorPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolOwnershipTransferred)
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
		it.Event = new(DelegatorPoolOwnershipTransferred)
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
func (it *DelegatorPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolOwnershipTransferred represents a OwnershipTransferred event raised by the DelegatorPool contract.
type DelegatorPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegatorPool *DelegatorPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelegatorPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolOwnershipTransferredIterator{contract: _DelegatorPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegatorPool *DelegatorPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelegatorPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolOwnershipTransferred)
				if err := _DelegatorPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegatorPool *DelegatorPoolFilterer) ParseOwnershipTransferred(log types.Log) (*DelegatorPoolOwnershipTransferred, error) {
	event := new(DelegatorPoolOwnershipTransferred)
	if err := _DelegatorPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolRedirectApprovalIterator is returned from FilterRedirectApproval and is used to iterate over the raw logs and unpacked data for RedirectApproval events raised by the DelegatorPool contract.
type DelegatorPoolRedirectApprovalIterator struct {
	Event *DelegatorPoolRedirectApproval // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolRedirectApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolRedirectApproval)
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
		it.Event = new(DelegatorPoolRedirectApproval)
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
func (it *DelegatorPoolRedirectApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolRedirectApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolRedirectApproval represents a RedirectApproval event raised by the DelegatorPool contract.
type DelegatorPoolRedirectApproval struct {
	Approver common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedirectApproval is a free log retrieval operation binding the contract event 0xbc2ae327210da6d1f0946518a140d99a46f3a238ea185f565875ebfe507eb996.
//
// Solidity: event RedirectApproval(address indexed approver, address indexed to)
func (_DelegatorPool *DelegatorPoolFilterer) FilterRedirectApproval(opts *bind.FilterOpts, approver []common.Address, to []common.Address) (*DelegatorPoolRedirectApprovalIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "RedirectApproval", approverRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolRedirectApprovalIterator{contract: _DelegatorPool.contract, event: "RedirectApproval", logs: logs, sub: sub}, nil
}

// WatchRedirectApproval is a free log subscription operation binding the contract event 0xbc2ae327210da6d1f0946518a140d99a46f3a238ea185f565875ebfe507eb996.
//
// Solidity: event RedirectApproval(address indexed approver, address indexed to)
func (_DelegatorPool *DelegatorPoolFilterer) WatchRedirectApproval(opts *bind.WatchOpts, sink chan<- *DelegatorPoolRedirectApproval, approver []common.Address, to []common.Address) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "RedirectApproval", approverRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolRedirectApproval)
				if err := _DelegatorPool.contract.UnpackLog(event, "RedirectApproval", log); err != nil {
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

// ParseRedirectApproval is a log parse operation binding the contract event 0xbc2ae327210da6d1f0946518a140d99a46f3a238ea185f565875ebfe507eb996.
//
// Solidity: event RedirectApproval(address indexed approver, address indexed to)
func (_DelegatorPool *DelegatorPoolFilterer) ParseRedirectApproval(log types.Log) (*DelegatorPoolRedirectApproval, error) {
	event := new(DelegatorPoolRedirectApproval)
	if err := _DelegatorPool.contract.UnpackLog(event, "RedirectApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolRedirectApprovalRevokedIterator is returned from FilterRedirectApprovalRevoked and is used to iterate over the raw logs and unpacked data for RedirectApprovalRevoked events raised by the DelegatorPool contract.
type DelegatorPoolRedirectApprovalRevokedIterator struct {
	Event *DelegatorPoolRedirectApprovalRevoked // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolRedirectApprovalRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolRedirectApprovalRevoked)
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
		it.Event = new(DelegatorPoolRedirectApprovalRevoked)
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
func (it *DelegatorPoolRedirectApprovalRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolRedirectApprovalRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolRedirectApprovalRevoked represents a RedirectApprovalRevoked event raised by the DelegatorPool contract.
type DelegatorPoolRedirectApprovalRevoked struct {
	Approver common.Address
	From     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedirectApprovalRevoked is a free log retrieval operation binding the contract event 0x9dca899729a60a64e03e19b833c4ef7e1795b7d9ed899092fd59d5df5a8229e4.
//
// Solidity: event RedirectApprovalRevoked(address indexed approver, address indexed from)
func (_DelegatorPool *DelegatorPoolFilterer) FilterRedirectApprovalRevoked(opts *bind.FilterOpts, approver []common.Address, from []common.Address) (*DelegatorPoolRedirectApprovalRevokedIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "RedirectApprovalRevoked", approverRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolRedirectApprovalRevokedIterator{contract: _DelegatorPool.contract, event: "RedirectApprovalRevoked", logs: logs, sub: sub}, nil
}

// WatchRedirectApprovalRevoked is a free log subscription operation binding the contract event 0x9dca899729a60a64e03e19b833c4ef7e1795b7d9ed899092fd59d5df5a8229e4.
//
// Solidity: event RedirectApprovalRevoked(address indexed approver, address indexed from)
func (_DelegatorPool *DelegatorPoolFilterer) WatchRedirectApprovalRevoked(opts *bind.WatchOpts, sink chan<- *DelegatorPoolRedirectApprovalRevoked, approver []common.Address, from []common.Address) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "RedirectApprovalRevoked", approverRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolRedirectApprovalRevoked)
				if err := _DelegatorPool.contract.UnpackLog(event, "RedirectApprovalRevoked", log); err != nil {
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

// ParseRedirectApprovalRevoked is a log parse operation binding the contract event 0x9dca899729a60a64e03e19b833c4ef7e1795b7d9ed899092fd59d5df5a8229e4.
//
// Solidity: event RedirectApprovalRevoked(address indexed approver, address indexed from)
func (_DelegatorPool *DelegatorPoolFilterer) ParseRedirectApprovalRevoked(log types.Log) (*DelegatorPoolRedirectApprovalRevoked, error) {
	event := new(DelegatorPoolRedirectApprovalRevoked)
	if err := _DelegatorPool.contract.UnpackLog(event, "RedirectApprovalRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolRemoveTokenIterator is returned from FilterRemoveToken and is used to iterate over the raw logs and unpacked data for RemoveToken events raised by the DelegatorPool contract.
type DelegatorPoolRemoveTokenIterator struct {
	Event *DelegatorPoolRemoveToken // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolRemoveTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolRemoveToken)
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
		it.Event = new(DelegatorPoolRemoveToken)
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
func (it *DelegatorPoolRemoveTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolRemoveTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolRemoveToken represents a RemoveToken event raised by the DelegatorPool contract.
type DelegatorPoolRemoveToken struct {
	Token       common.Address
	RewardsPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveToken is a free log retrieval operation binding the contract event 0x39dcd754ec63af4b82c4c569ff1b6b4e55a8038e6545844747e54f2f2d4e8e50.
//
// Solidity: event RemoveToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) FilterRemoveToken(opts *bind.FilterOpts, token []common.Address) (*DelegatorPoolRemoveTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "RemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolRemoveTokenIterator{contract: _DelegatorPool.contract, event: "RemoveToken", logs: logs, sub: sub}, nil
}

// WatchRemoveToken is a free log subscription operation binding the contract event 0x39dcd754ec63af4b82c4c569ff1b6b4e55a8038e6545844747e54f2f2d4e8e50.
//
// Solidity: event RemoveToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) WatchRemoveToken(opts *bind.WatchOpts, sink chan<- *DelegatorPoolRemoveToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "RemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolRemoveToken)
				if err := _DelegatorPool.contract.UnpackLog(event, "RemoveToken", log); err != nil {
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

// ParseRemoveToken is a log parse operation binding the contract event 0x39dcd754ec63af4b82c4c569ff1b6b4e55a8038e6545844747e54f2f2d4e8e50.
//
// Solidity: event RemoveToken(address indexed token, address rewardsPool)
func (_DelegatorPool *DelegatorPoolFilterer) ParseRemoveToken(log types.Log) (*DelegatorPoolRemoveToken, error) {
	event := new(DelegatorPoolRemoveToken)
	if err := _DelegatorPool.contract.UnpackLog(event, "RemoveToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolRewardsRedirectedIterator is returned from FilterRewardsRedirected and is used to iterate over the raw logs and unpacked data for RewardsRedirected events raised by the DelegatorPool contract.
type DelegatorPoolRewardsRedirectedIterator struct {
	Event *DelegatorPoolRewardsRedirected // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolRewardsRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolRewardsRedirected)
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
		it.Event = new(DelegatorPoolRewardsRedirected)
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
func (it *DelegatorPoolRewardsRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolRewardsRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolRewardsRedirected represents a RewardsRedirected event raised by the DelegatorPool contract.
type DelegatorPoolRewardsRedirected struct {
	From common.Address
	To   common.Address
	By   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRewardsRedirected is a free log retrieval operation binding the contract event 0x6a8c4b3eeb6ce85acfca51b03ce7b05fdefe1ff2b7c0a45e7002b08e0cea1eb6.
//
// Solidity: event RewardsRedirected(address indexed from, address indexed to, address indexed by)
func (_DelegatorPool *DelegatorPoolFilterer) FilterRewardsRedirected(opts *bind.FilterOpts, from []common.Address, to []common.Address, by []common.Address) (*DelegatorPoolRewardsRedirectedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "RewardsRedirected", fromRule, toRule, byRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolRewardsRedirectedIterator{contract: _DelegatorPool.contract, event: "RewardsRedirected", logs: logs, sub: sub}, nil
}

// WatchRewardsRedirected is a free log subscription operation binding the contract event 0x6a8c4b3eeb6ce85acfca51b03ce7b05fdefe1ff2b7c0a45e7002b08e0cea1eb6.
//
// Solidity: event RewardsRedirected(address indexed from, address indexed to, address indexed by)
func (_DelegatorPool *DelegatorPoolFilterer) WatchRewardsRedirected(opts *bind.WatchOpts, sink chan<- *DelegatorPoolRewardsRedirected, from []common.Address, to []common.Address, by []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "RewardsRedirected", fromRule, toRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolRewardsRedirected)
				if err := _DelegatorPool.contract.UnpackLog(event, "RewardsRedirected", log); err != nil {
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

// ParseRewardsRedirected is a log parse operation binding the contract event 0x6a8c4b3eeb6ce85acfca51b03ce7b05fdefe1ff2b7c0a45e7002b08e0cea1eb6.
//
// Solidity: event RewardsRedirected(address indexed from, address indexed to, address indexed by)
func (_DelegatorPool *DelegatorPoolFilterer) ParseRewardsRedirected(log types.Log) (*DelegatorPoolRewardsRedirected, error) {
	event := new(DelegatorPoolRewardsRedirected)
	if err := _DelegatorPool.contract.UnpackLog(event, "RewardsRedirected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DelegatorPool contract.
type DelegatorPoolTransferIterator struct {
	Event *DelegatorPoolTransfer // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolTransfer)
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
		it.Event = new(DelegatorPoolTransfer)
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
func (it *DelegatorPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolTransfer represents a Transfer event raised by the DelegatorPool contract.
type DelegatorPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DelegatorPool *DelegatorPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DelegatorPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolTransferIterator{contract: _DelegatorPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DelegatorPool *DelegatorPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DelegatorPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolTransfer)
				if err := _DelegatorPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DelegatorPool *DelegatorPoolFilterer) ParseTransfer(log types.Log) (*DelegatorPoolTransfer, error) {
	event := new(DelegatorPoolTransfer)
	if err := _DelegatorPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorPoolWithdrawRewardsIterator is returned from FilterWithdrawRewards and is used to iterate over the raw logs and unpacked data for WithdrawRewards events raised by the DelegatorPool contract.
type DelegatorPoolWithdrawRewardsIterator struct {
	Event *DelegatorPoolWithdrawRewards // Event containing the contract specifics and raw log

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
func (it *DelegatorPoolWithdrawRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorPoolWithdrawRewards)
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
		it.Event = new(DelegatorPoolWithdrawRewards)
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
func (it *DelegatorPoolWithdrawRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorPoolWithdrawRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorPoolWithdrawRewards represents a WithdrawRewards event raised by the DelegatorPool contract.
type DelegatorPoolWithdrawRewards struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRewards is a free log retrieval operation binding the contract event 0xeeab3fd62be4cb59cbdc42d5b0f676a3597ff387b9a85e62330cc17c2a3603db.
//
// Solidity: event WithdrawRewards(address indexed account)
func (_DelegatorPool *DelegatorPoolFilterer) FilterWithdrawRewards(opts *bind.FilterOpts, account []common.Address) (*DelegatorPoolWithdrawRewardsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegatorPool.contract.FilterLogs(opts, "WithdrawRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorPoolWithdrawRewardsIterator{contract: _DelegatorPool.contract, event: "WithdrawRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawRewards is a free log subscription operation binding the contract event 0xeeab3fd62be4cb59cbdc42d5b0f676a3597ff387b9a85e62330cc17c2a3603db.
//
// Solidity: event WithdrawRewards(address indexed account)
func (_DelegatorPool *DelegatorPoolFilterer) WatchWithdrawRewards(opts *bind.WatchOpts, sink chan<- *DelegatorPoolWithdrawRewards, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegatorPool.contract.WatchLogs(opts, "WithdrawRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorPoolWithdrawRewards)
				if err := _DelegatorPool.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
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

// ParseWithdrawRewards is a log parse operation binding the contract event 0xeeab3fd62be4cb59cbdc42d5b0f676a3597ff387b9a85e62330cc17c2a3603db.
//
// Solidity: event WithdrawRewards(address indexed account)
func (_DelegatorPool *DelegatorPoolFilterer) ParseWithdrawRewards(log types.Log) (*DelegatorPoolWithdrawRewards, error) {
	event := new(DelegatorPoolWithdrawRewards)
	if err := _DelegatorPool.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
