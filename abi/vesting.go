// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// VestingMetaData contains all meta data concerning the Vesting contract.
var VestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_durationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"VestingAlreadyTerminated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VestingNotTerminated\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VestingTerminated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"releaseRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"terminateVesting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"vestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"vestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingTerminated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VestingABI is the input ABI used to generate the binding from.
// Deprecated: Use VestingMetaData.ABI instead.
var VestingABI = VestingMetaData.ABI

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VestingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Vesting *VestingCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Vesting *VestingSession) Beneficiary() (common.Address, error) {
	return _Vesting.Contract.Beneficiary(&_Vesting.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Vesting *VestingCallerSession) Beneficiary() (common.Address, error) {
	return _Vesting.Contract.Beneficiary(&_Vesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Vesting *VestingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Vesting *VestingSession) Duration() (*big.Int, error) {
	return _Vesting.Contract.Duration(&_Vesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Vesting *VestingCallerSession) Duration() (*big.Int, error) {
	return _Vesting.Contract.Duration(&_Vesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingSession) Owner() (common.Address, error) {
	return _Vesting.Contract.Owner(&_Vesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingCallerSession) Owner() (common.Address, error) {
	return _Vesting.Contract.Owner(&_Vesting.CallOpts)
}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Vesting *VestingCaller) Released(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "released")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Vesting *VestingSession) Released() (*big.Int, error) {
	return _Vesting.Contract.Released(&_Vesting.CallOpts)
}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Vesting *VestingCallerSession) Released() (*big.Int, error) {
	return _Vesting.Contract.Released(&_Vesting.CallOpts)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) view returns(uint256)
func (_Vesting *VestingCaller) Released0(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "released0", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) view returns(uint256)
func (_Vesting *VestingSession) Released0(token common.Address) (*big.Int, error) {
	return _Vesting.Contract.Released0(&_Vesting.CallOpts, token)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) view returns(uint256)
func (_Vesting *VestingCallerSession) Released0(token common.Address) (*big.Int, error) {
	return _Vesting.Contract.Released0(&_Vesting.CallOpts, token)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Vesting *VestingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Vesting *VestingSession) Start() (*big.Int, error) {
	return _Vesting.Contract.Start(&_Vesting.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Vesting *VestingCallerSession) Start() (*big.Int, error) {
	return _Vesting.Contract.Start(&_Vesting.CallOpts)
}

// VestedAmount is a free data retrieval call binding the contract method 0x0a17b06b.
//
// Solidity: function vestedAmount(uint64 timestamp) view returns(uint256)
func (_Vesting *VestingCaller) VestedAmount(opts *bind.CallOpts, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "vestedAmount", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedAmount is a free data retrieval call binding the contract method 0x0a17b06b.
//
// Solidity: function vestedAmount(uint64 timestamp) view returns(uint256)
func (_Vesting *VestingSession) VestedAmount(timestamp uint64) (*big.Int, error) {
	return _Vesting.Contract.VestedAmount(&_Vesting.CallOpts, timestamp)
}

// VestedAmount is a free data retrieval call binding the contract method 0x0a17b06b.
//
// Solidity: function vestedAmount(uint64 timestamp) view returns(uint256)
func (_Vesting *VestingCallerSession) VestedAmount(timestamp uint64) (*big.Int, error) {
	return _Vesting.Contract.VestedAmount(&_Vesting.CallOpts, timestamp)
}

// VestedAmount0 is a free data retrieval call binding the contract method 0x810ec23b.
//
// Solidity: function vestedAmount(address token, uint64 timestamp) view returns(uint256)
func (_Vesting *VestingCaller) VestedAmount0(opts *bind.CallOpts, token common.Address, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "vestedAmount0", token, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedAmount0 is a free data retrieval call binding the contract method 0x810ec23b.
//
// Solidity: function vestedAmount(address token, uint64 timestamp) view returns(uint256)
func (_Vesting *VestingSession) VestedAmount0(token common.Address, timestamp uint64) (*big.Int, error) {
	return _Vesting.Contract.VestedAmount0(&_Vesting.CallOpts, token, timestamp)
}

// VestedAmount0 is a free data retrieval call binding the contract method 0x810ec23b.
//
// Solidity: function vestedAmount(address token, uint64 timestamp) view returns(uint256)
func (_Vesting *VestingCallerSession) VestedAmount0(token common.Address, timestamp uint64) (*big.Int, error) {
	return _Vesting.Contract.VestedAmount0(&_Vesting.CallOpts, token, timestamp)
}

// VestingTerminated is a free data retrieval call binding the contract method 0x8732e3a0.
//
// Solidity: function vestingTerminated() view returns(bool)
func (_Vesting *VestingCaller) VestingTerminated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "vestingTerminated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VestingTerminated is a free data retrieval call binding the contract method 0x8732e3a0.
//
// Solidity: function vestingTerminated() view returns(bool)
func (_Vesting *VestingSession) VestingTerminated() (bool, error) {
	return _Vesting.Contract.VestingTerminated(&_Vesting.CallOpts)
}

// VestingTerminated is a free data retrieval call binding the contract method 0x8732e3a0.
//
// Solidity: function vestingTerminated() view returns(bool)
func (_Vesting *VestingCallerSession) VestingTerminated() (bool, error) {
	return _Vesting.Contract.VestingTerminated(&_Vesting.CallOpts)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_Vesting *VestingTransactor) Release(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "release", token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_Vesting *VestingSession) Release(token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Release(&_Vesting.TransactOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_Vesting *VestingTransactorSession) Release(token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Release(&_Vesting.TransactOpts, token)
}

// Release0 is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingTransactor) Release0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "release0")
}

// Release0 is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingSession) Release0() (*types.Transaction, error) {
	return _Vesting.Contract.Release0(&_Vesting.TransactOpts)
}

// Release0 is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingTransactorSession) Release0() (*types.Transaction, error) {
	return _Vesting.Contract.Release0(&_Vesting.TransactOpts)
}

// ReleaseRemaining is a paid mutator transaction binding the contract method 0xeb491a88.
//
// Solidity: function releaseRemaining(address _token) returns()
func (_Vesting *VestingTransactor) ReleaseRemaining(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "releaseRemaining", _token)
}

// ReleaseRemaining is a paid mutator transaction binding the contract method 0xeb491a88.
//
// Solidity: function releaseRemaining(address _token) returns()
func (_Vesting *VestingSession) ReleaseRemaining(_token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.ReleaseRemaining(&_Vesting.TransactOpts, _token)
}

// ReleaseRemaining is a paid mutator transaction binding the contract method 0xeb491a88.
//
// Solidity: function releaseRemaining(address _token) returns()
func (_Vesting *VestingTransactorSession) ReleaseRemaining(_token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.ReleaseRemaining(&_Vesting.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vesting *VestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vesting *VestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vesting.Contract.RenounceOwnership(&_Vesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vesting *VestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vesting.Contract.RenounceOwnership(&_Vesting.TransactOpts)
}

// TerminateVesting is a paid mutator transaction binding the contract method 0x40c79fa7.
//
// Solidity: function terminateVesting(address[] _tokens) returns()
func (_Vesting *VestingTransactor) TerminateVesting(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "terminateVesting", _tokens)
}

// TerminateVesting is a paid mutator transaction binding the contract method 0x40c79fa7.
//
// Solidity: function terminateVesting(address[] _tokens) returns()
func (_Vesting *VestingSession) TerminateVesting(_tokens []common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.TerminateVesting(&_Vesting.TransactOpts, _tokens)
}

// TerminateVesting is a paid mutator transaction binding the contract method 0x40c79fa7.
//
// Solidity: function terminateVesting(address[] _tokens) returns()
func (_Vesting *VestingTransactorSession) TerminateVesting(_tokens []common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.TerminateVesting(&_Vesting.TransactOpts, _tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vesting *VestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vesting *VestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.TransferOwnership(&_Vesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vesting *VestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.TransferOwnership(&_Vesting.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vesting *VestingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vesting *VestingSession) Receive() (*types.Transaction, error) {
	return _Vesting.Contract.Receive(&_Vesting.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vesting *VestingTransactorSession) Receive() (*types.Transaction, error) {
	return _Vesting.Contract.Receive(&_Vesting.TransactOpts)
}

// VestingERC20ReleasedIterator is returned from FilterERC20Released and is used to iterate over the raw logs and unpacked data for ERC20Released events raised by the Vesting contract.
type VestingERC20ReleasedIterator struct {
	Event *VestingERC20Released // Event containing the contract specifics and raw log

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
func (it *VestingERC20ReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingERC20Released)
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
		it.Event = new(VestingERC20Released)
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
func (it *VestingERC20ReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingERC20ReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingERC20Released represents a ERC20Released event raised by the Vesting contract.
type VestingERC20Released struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20Released is a free log retrieval operation binding the contract event 0xc0e523490dd523c33b1878c9eb14ff46991e3f5b2cd33710918618f2a39cba1b.
//
// Solidity: event ERC20Released(address indexed token, uint256 amount)
func (_Vesting *VestingFilterer) FilterERC20Released(opts *bind.FilterOpts, token []common.Address) (*VestingERC20ReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "ERC20Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VestingERC20ReleasedIterator{contract: _Vesting.contract, event: "ERC20Released", logs: logs, sub: sub}, nil
}

// WatchERC20Released is a free log subscription operation binding the contract event 0xc0e523490dd523c33b1878c9eb14ff46991e3f5b2cd33710918618f2a39cba1b.
//
// Solidity: event ERC20Released(address indexed token, uint256 amount)
func (_Vesting *VestingFilterer) WatchERC20Released(opts *bind.WatchOpts, sink chan<- *VestingERC20Released, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "ERC20Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingERC20Released)
				if err := _Vesting.contract.UnpackLog(event, "ERC20Released", log); err != nil {
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

// ParseERC20Released is a log parse operation binding the contract event 0xc0e523490dd523c33b1878c9eb14ff46991e3f5b2cd33710918618f2a39cba1b.
//
// Solidity: event ERC20Released(address indexed token, uint256 amount)
func (_Vesting *VestingFilterer) ParseERC20Released(log types.Log) (*VestingERC20Released, error) {
	event := new(VestingERC20Released)
	if err := _Vesting.contract.UnpackLog(event, "ERC20Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingEtherReleasedIterator is returned from FilterEtherReleased and is used to iterate over the raw logs and unpacked data for EtherReleased events raised by the Vesting contract.
type VestingEtherReleasedIterator struct {
	Event *VestingEtherReleased // Event containing the contract specifics and raw log

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
func (it *VestingEtherReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingEtherReleased)
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
		it.Event = new(VestingEtherReleased)
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
func (it *VestingEtherReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingEtherReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingEtherReleased represents a EtherReleased event raised by the Vesting contract.
type VestingEtherReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherReleased is a free log retrieval operation binding the contract event 0xda9d4e5f101b8b9b1c5b76d0c5a9f7923571acfc02376aa076b75a8c080c956b.
//
// Solidity: event EtherReleased(uint256 amount)
func (_Vesting *VestingFilterer) FilterEtherReleased(opts *bind.FilterOpts) (*VestingEtherReleasedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "EtherReleased")
	if err != nil {
		return nil, err
	}
	return &VestingEtherReleasedIterator{contract: _Vesting.contract, event: "EtherReleased", logs: logs, sub: sub}, nil
}

// WatchEtherReleased is a free log subscription operation binding the contract event 0xda9d4e5f101b8b9b1c5b76d0c5a9f7923571acfc02376aa076b75a8c080c956b.
//
// Solidity: event EtherReleased(uint256 amount)
func (_Vesting *VestingFilterer) WatchEtherReleased(opts *bind.WatchOpts, sink chan<- *VestingEtherReleased) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "EtherReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingEtherReleased)
				if err := _Vesting.contract.UnpackLog(event, "EtherReleased", log); err != nil {
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

// ParseEtherReleased is a log parse operation binding the contract event 0xda9d4e5f101b8b9b1c5b76d0c5a9f7923571acfc02376aa076b75a8c080c956b.
//
// Solidity: event EtherReleased(uint256 amount)
func (_Vesting *VestingFilterer) ParseEtherReleased(log types.Log) (*VestingEtherReleased, error) {
	event := new(VestingEtherReleased)
	if err := _Vesting.contract.UnpackLog(event, "EtherReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vesting contract.
type VestingOwnershipTransferredIterator struct {
	Event *VestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingOwnershipTransferred)
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
		it.Event = new(VestingOwnershipTransferred)
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
func (it *VestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingOwnershipTransferred represents a OwnershipTransferred event raised by the Vesting contract.
type VestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vesting *VestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VestingOwnershipTransferredIterator{contract: _Vesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vesting *VestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingOwnershipTransferred)
				if err := _Vesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vesting *VestingFilterer) ParseOwnershipTransferred(log types.Log) (*VestingOwnershipTransferred, error) {
	event := new(VestingOwnershipTransferred)
	if err := _Vesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingVestingTerminatedIterator is returned from FilterVestingTerminated and is used to iterate over the raw logs and unpacked data for VestingTerminated events raised by the Vesting contract.
type VestingVestingTerminatedIterator struct {
	Event *VestingVestingTerminated // Event containing the contract specifics and raw log

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
func (it *VestingVestingTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingVestingTerminated)
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
		it.Event = new(VestingVestingTerminated)
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
func (it *VestingVestingTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingVestingTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingVestingTerminated represents a VestingTerminated event raised by the Vesting contract.
type VestingVestingTerminated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVestingTerminated is a free log retrieval operation binding the contract event 0x22ee45b273ba60d97bc6a3eb07766395877c427e0db75dae5e70615d6dd497ac.
//
// Solidity: event VestingTerminated()
func (_Vesting *VestingFilterer) FilterVestingTerminated(opts *bind.FilterOpts) (*VestingVestingTerminatedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "VestingTerminated")
	if err != nil {
		return nil, err
	}
	return &VestingVestingTerminatedIterator{contract: _Vesting.contract, event: "VestingTerminated", logs: logs, sub: sub}, nil
}

// WatchVestingTerminated is a free log subscription operation binding the contract event 0x22ee45b273ba60d97bc6a3eb07766395877c427e0db75dae5e70615d6dd497ac.
//
// Solidity: event VestingTerminated()
func (_Vesting *VestingFilterer) WatchVestingTerminated(opts *bind.WatchOpts, sink chan<- *VestingVestingTerminated) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "VestingTerminated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingVestingTerminated)
				if err := _Vesting.contract.UnpackLog(event, "VestingTerminated", log); err != nil {
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

// ParseVestingTerminated is a log parse operation binding the contract event 0x22ee45b273ba60d97bc6a3eb07766395877c427e0db75dae5e70615d6dd497ac.
//
// Solidity: event VestingTerminated()
func (_Vesting *VestingFilterer) ParseVestingTerminated(log types.Log) (*VestingVestingTerminated, error) {
	event := new(VestingVestingTerminated)
	if err := _Vesting.contract.UnpackLog(event, "VestingTerminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
