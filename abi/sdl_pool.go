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

// SDLPoolLock is an auto generated low-level Go binding around an user-defined struct.
type SDLPoolLock struct {
	Amount      *big.Int
	BoostAmount *big.Int
	StartTime   uint64
	Duration    uint64
	Expiry      uint64
}

// SdlPoolMetaData contains all meta data concerning the SdlPool contract.
var SdlPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HalfDurationNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLockId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLockingDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalDurationNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721Implementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnlockAlreadyInitiated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnlockNotInitiated\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsPool\",\"type\":\"address\"}],\"name\":\"AddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lockingDuration\",\"type\":\"uint64\"}],\"name\":\"CreateLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"InitiateUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsPool\",\"type\":\"address\"}],\"name\":\"RemoveToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lockingDuration\",\"type\":\"uint64\"}],\"name\":\"UpdateLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawRewards\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"__RewardsPoolController_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsPool\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostController\",\"outputs\":[{\"internalType\":\"contractIBoostController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatorPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"distributeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"distributeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"effectiveBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_lockingDuration\",\"type\":\"uint64\"}],\"name\":\"extendLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getLockIdsByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_lockIds\",\"type\":\"uint256[]\"}],\"name\":\"getLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boostAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"internalType\":\"structSDLPool.Lock[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_sdlToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_boostController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegatorPool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"initiateUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLockId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_lockingDuration\",\"type\":\"uint64\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdlToken\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_boostController\",\"type\":\"address\"}],\"name\":\"setBoostController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenPools\",\"outputs\":[{\"internalType\":\"contractIRewardsPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEffectiveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawableRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SdlPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SdlPoolMetaData.ABI instead.
var SdlPoolABI = SdlPoolMetaData.ABI

// SdlPool is an auto generated Go binding around an Ethereum contract.
type SdlPool struct {
	SdlPoolCaller     // Read-only binding to the contract
	SdlPoolTransactor // Write-only binding to the contract
	SdlPoolFilterer   // Log filterer for contract events
}

// SdlPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SdlPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdlPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SdlPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdlPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SdlPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdlPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SdlPoolSession struct {
	Contract     *SdlPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SdlPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SdlPoolCallerSession struct {
	Contract *SdlPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SdlPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SdlPoolTransactorSession struct {
	Contract     *SdlPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SdlPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SdlPoolRaw struct {
	Contract *SdlPool // Generic contract binding to access the raw methods on
}

// SdlPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SdlPoolCallerRaw struct {
	Contract *SdlPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SdlPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SdlPoolTransactorRaw struct {
	Contract *SdlPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSdlPool creates a new instance of SdlPool, bound to a specific deployed contract.
func NewSdlPool(address common.Address, backend bind.ContractBackend) (*SdlPool, error) {
	contract, err := bindSdlPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SdlPool{SdlPoolCaller: SdlPoolCaller{contract: contract}, SdlPoolTransactor: SdlPoolTransactor{contract: contract}, SdlPoolFilterer: SdlPoolFilterer{contract: contract}}, nil
}

// NewSdlPoolCaller creates a new read-only instance of SdlPool, bound to a specific deployed contract.
func NewSdlPoolCaller(address common.Address, caller bind.ContractCaller) (*SdlPoolCaller, error) {
	contract, err := bindSdlPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SdlPoolCaller{contract: contract}, nil
}

// NewSdlPoolTransactor creates a new write-only instance of SdlPool, bound to a specific deployed contract.
func NewSdlPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SdlPoolTransactor, error) {
	contract, err := bindSdlPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SdlPoolTransactor{contract: contract}, nil
}

// NewSdlPoolFilterer creates a new log filterer instance of SdlPool, bound to a specific deployed contract.
func NewSdlPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SdlPoolFilterer, error) {
	contract, err := bindSdlPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SdlPoolFilterer{contract: contract}, nil
}

// bindSdlPool binds a generic wrapper to an already deployed contract.
func bindSdlPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SdlPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdlPool *SdlPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdlPool.Contract.SdlPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdlPool *SdlPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdlPool.Contract.SdlPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdlPool *SdlPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdlPool.Contract.SdlPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdlPool *SdlPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdlPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdlPool *SdlPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdlPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdlPool *SdlPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdlPool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.BalanceOf(&_SdlPool.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.BalanceOf(&_SdlPool.CallOpts, _account)
}

// BoostController is a free data retrieval call binding the contract method 0x7c898895.
//
// Solidity: function boostController() view returns(address)
func (_SdlPool *SdlPoolCaller) BoostController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "boostController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostController is a free data retrieval call binding the contract method 0x7c898895.
//
// Solidity: function boostController() view returns(address)
func (_SdlPool *SdlPoolSession) BoostController() (common.Address, error) {
	return _SdlPool.Contract.BoostController(&_SdlPool.CallOpts)
}

// BoostController is a free data retrieval call binding the contract method 0x7c898895.
//
// Solidity: function boostController() view returns(address)
func (_SdlPool *SdlPoolCallerSession) BoostController() (common.Address, error) {
	return _SdlPool.Contract.BoostController(&_SdlPool.CallOpts)
}

// DelegatorPool is a free data retrieval call binding the contract method 0x5deb761f.
//
// Solidity: function delegatorPool() view returns(address)
func (_SdlPool *SdlPoolCaller) DelegatorPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "delegatorPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatorPool is a free data retrieval call binding the contract method 0x5deb761f.
//
// Solidity: function delegatorPool() view returns(address)
func (_SdlPool *SdlPoolSession) DelegatorPool() (common.Address, error) {
	return _SdlPool.Contract.DelegatorPool(&_SdlPool.CallOpts)
}

// DelegatorPool is a free data retrieval call binding the contract method 0x5deb761f.
//
// Solidity: function delegatorPool() view returns(address)
func (_SdlPool *SdlPoolCallerSession) DelegatorPool() (common.Address, error) {
	return _SdlPool.Contract.DelegatorPool(&_SdlPool.CallOpts)
}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCaller) EffectiveBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "effectiveBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolSession) EffectiveBalanceOf(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.EffectiveBalanceOf(&_SdlPool.CallOpts, _account)
}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) EffectiveBalanceOf(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.EffectiveBalanceOf(&_SdlPool.CallOpts, _account)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolCaller) GetApproved(opts *bind.CallOpts, _lockId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "getApproved", _lockId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolSession) GetApproved(_lockId *big.Int) (common.Address, error) {
	return _SdlPool.Contract.GetApproved(&_SdlPool.CallOpts, _lockId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolCallerSession) GetApproved(_lockId *big.Int) (common.Address, error) {
	return _SdlPool.Contract.GetApproved(&_SdlPool.CallOpts, _lockId)
}

// GetLockIdsByOwner is a free data retrieval call binding the contract method 0xdd0228f5.
//
// Solidity: function getLockIdsByOwner(address _owner) view returns(uint256[])
func (_SdlPool *SdlPoolCaller) GetLockIdsByOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "getLockIdsByOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLockIdsByOwner is a free data retrieval call binding the contract method 0xdd0228f5.
//
// Solidity: function getLockIdsByOwner(address _owner) view returns(uint256[])
func (_SdlPool *SdlPoolSession) GetLockIdsByOwner(_owner common.Address) ([]*big.Int, error) {
	return _SdlPool.Contract.GetLockIdsByOwner(&_SdlPool.CallOpts, _owner)
}

// GetLockIdsByOwner is a free data retrieval call binding the contract method 0xdd0228f5.
//
// Solidity: function getLockIdsByOwner(address _owner) view returns(uint256[])
func (_SdlPool *SdlPoolCallerSession) GetLockIdsByOwner(_owner common.Address) ([]*big.Int, error) {
	return _SdlPool.Contract.GetLockIdsByOwner(&_SdlPool.CallOpts, _owner)
}

// GetLocks is a free data retrieval call binding the contract method 0x4a374d4d.
//
// Solidity: function getLocks(uint256[] _lockIds) view returns((uint256,uint256,uint64,uint64,uint64)[])
func (_SdlPool *SdlPoolCaller) GetLocks(opts *bind.CallOpts, _lockIds []*big.Int) ([]SDLPoolLock, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "getLocks", _lockIds)

	if err != nil {
		return *new([]SDLPoolLock), err
	}

	out0 := *abi.ConvertType(out[0], new([]SDLPoolLock)).(*[]SDLPoolLock)

	return out0, err

}

// GetLocks is a free data retrieval call binding the contract method 0x4a374d4d.
//
// Solidity: function getLocks(uint256[] _lockIds) view returns((uint256,uint256,uint64,uint64,uint64)[])
func (_SdlPool *SdlPoolSession) GetLocks(_lockIds []*big.Int) ([]SDLPoolLock, error) {
	return _SdlPool.Contract.GetLocks(&_SdlPool.CallOpts, _lockIds)
}

// GetLocks is a free data retrieval call binding the contract method 0x4a374d4d.
//
// Solidity: function getLocks(uint256[] _lockIds) view returns((uint256,uint256,uint64,uint64,uint64)[])
func (_SdlPool *SdlPoolCallerSession) GetLocks(_lockIds []*big.Int) ([]SDLPoolLock, error) {
	return _SdlPool.Contract.GetLocks(&_SdlPool.CallOpts, _lockIds)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_SdlPool *SdlPoolCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_SdlPool *SdlPoolSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _SdlPool.Contract.IsApprovedForAll(&_SdlPool.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_SdlPool *SdlPoolCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _SdlPool.Contract.IsApprovedForAll(&_SdlPool.CallOpts, _owner, _operator)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_SdlPool *SdlPoolCaller) IsTokenSupported(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "isTokenSupported", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_SdlPool *SdlPoolSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _SdlPool.Contract.IsTokenSupported(&_SdlPool.CallOpts, _token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_SdlPool *SdlPoolCallerSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _SdlPool.Contract.IsTokenSupported(&_SdlPool.CallOpts, _token)
}

// LastLockId is a free data retrieval call binding the contract method 0xd0cb39c6.
//
// Solidity: function lastLockId() view returns(uint256)
func (_SdlPool *SdlPoolCaller) LastLockId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "lastLockId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLockId is a free data retrieval call binding the contract method 0xd0cb39c6.
//
// Solidity: function lastLockId() view returns(uint256)
func (_SdlPool *SdlPoolSession) LastLockId() (*big.Int, error) {
	return _SdlPool.Contract.LastLockId(&_SdlPool.CallOpts)
}

// LastLockId is a free data retrieval call binding the contract method 0xd0cb39c6.
//
// Solidity: function lastLockId() view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) LastLockId() (*big.Int, error) {
	return _SdlPool.Contract.LastLockId(&_SdlPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SdlPool *SdlPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SdlPool *SdlPoolSession) Name() (string, error) {
	return _SdlPool.Contract.Name(&_SdlPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SdlPool *SdlPoolCallerSession) Name() (string, error) {
	return _SdlPool.Contract.Name(&_SdlPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SdlPool *SdlPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SdlPool *SdlPoolSession) Owner() (common.Address, error) {
	return _SdlPool.Contract.Owner(&_SdlPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SdlPool *SdlPoolCallerSession) Owner() (common.Address, error) {
	return _SdlPool.Contract.Owner(&_SdlPool.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolCaller) OwnerOf(opts *bind.CallOpts, _lockId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "ownerOf", _lockId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolSession) OwnerOf(_lockId *big.Int) (common.Address, error) {
	return _SdlPool.Contract.OwnerOf(&_SdlPool.CallOpts, _lockId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _lockId) view returns(address)
func (_SdlPool *SdlPoolCallerSession) OwnerOf(_lockId *big.Int) (common.Address, error) {
	return _SdlPool.Contract.OwnerOf(&_SdlPool.CallOpts, _lockId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SdlPool *SdlPoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SdlPool *SdlPoolSession) ProxiableUUID() ([32]byte, error) {
	return _SdlPool.Contract.ProxiableUUID(&_SdlPool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SdlPool *SdlPoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SdlPool.Contract.ProxiableUUID(&_SdlPool.CallOpts)
}

// SdlToken is a free data retrieval call binding the contract method 0x9b4140cc.
//
// Solidity: function sdlToken() view returns(address)
func (_SdlPool *SdlPoolCaller) SdlToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "sdlToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SdlToken is a free data retrieval call binding the contract method 0x9b4140cc.
//
// Solidity: function sdlToken() view returns(address)
func (_SdlPool *SdlPoolSession) SdlToken() (common.Address, error) {
	return _SdlPool.Contract.SdlToken(&_SdlPool.CallOpts)
}

// SdlToken is a free data retrieval call binding the contract method 0x9b4140cc.
//
// Solidity: function sdlToken() view returns(address)
func (_SdlPool *SdlPoolCallerSession) SdlToken() (common.Address, error) {
	return _SdlPool.Contract.SdlToken(&_SdlPool.CallOpts)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCaller) Staked(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "staked", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_SdlPool *SdlPoolSession) Staked(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.Staked(&_SdlPool.CallOpts, _account)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address _account) view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) Staked(_account common.Address) (*big.Int, error) {
	return _SdlPool.Contract.Staked(&_SdlPool.CallOpts, _account)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_SdlPool *SdlPoolCaller) SupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "supportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_SdlPool *SdlPoolSession) SupportedTokens() ([]common.Address, error) {
	return _SdlPool.Contract.SupportedTokens(&_SdlPool.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xb002249d.
//
// Solidity: function supportedTokens() view returns(address[])
func (_SdlPool *SdlPoolCallerSession) SupportedTokens() ([]common.Address, error) {
	return _SdlPool.Contract.SupportedTokens(&_SdlPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SdlPool *SdlPoolCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SdlPool *SdlPoolSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SdlPool.Contract.SupportsInterface(&_SdlPool.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SdlPool *SdlPoolCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SdlPool.Contract.SupportsInterface(&_SdlPool.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SdlPool *SdlPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SdlPool *SdlPoolSession) Symbol() (string, error) {
	return _SdlPool.Contract.Symbol(&_SdlPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SdlPool *SdlPoolCallerSession) Symbol() (string, error) {
	return _SdlPool.Contract.Symbol(&_SdlPool.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x0e3bc974.
//
// Solidity: function tokenBalances() view returns(address[], uint256[])
func (_SdlPool *SdlPoolCaller) TokenBalances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "tokenBalances")

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
func (_SdlPool *SdlPoolSession) TokenBalances() ([]common.Address, []*big.Int, error) {
	return _SdlPool.Contract.TokenBalances(&_SdlPool.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x0e3bc974.
//
// Solidity: function tokenBalances() view returns(address[], uint256[])
func (_SdlPool *SdlPoolCallerSession) TokenBalances() ([]common.Address, []*big.Int, error) {
	return _SdlPool.Contract.TokenBalances(&_SdlPool.CallOpts)
}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_SdlPool *SdlPoolCaller) TokenPools(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "tokenPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_SdlPool *SdlPoolSession) TokenPools(arg0 common.Address) (common.Address, error) {
	return _SdlPool.Contract.TokenPools(&_SdlPool.CallOpts, arg0)
}

// TokenPools is a free data retrieval call binding the contract method 0xc3d2c3c1.
//
// Solidity: function tokenPools(address ) view returns(address)
func (_SdlPool *SdlPoolCallerSession) TokenPools(arg0 common.Address) (common.Address, error) {
	return _SdlPool.Contract.TokenPools(&_SdlPool.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_SdlPool *SdlPoolCaller) TokenURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "tokenURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_SdlPool *SdlPoolSession) TokenURI(arg0 *big.Int) (string, error) {
	return _SdlPool.Contract.TokenURI(&_SdlPool.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_SdlPool *SdlPoolCallerSession) TokenURI(arg0 *big.Int) (string, error) {
	return _SdlPool.Contract.TokenURI(&_SdlPool.CallOpts, arg0)
}

// TotalEffectiveBalance is a free data retrieval call binding the contract method 0x89676073.
//
// Solidity: function totalEffectiveBalance() view returns(uint256)
func (_SdlPool *SdlPoolCaller) TotalEffectiveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "totalEffectiveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEffectiveBalance is a free data retrieval call binding the contract method 0x89676073.
//
// Solidity: function totalEffectiveBalance() view returns(uint256)
func (_SdlPool *SdlPoolSession) TotalEffectiveBalance() (*big.Int, error) {
	return _SdlPool.Contract.TotalEffectiveBalance(&_SdlPool.CallOpts)
}

// TotalEffectiveBalance is a free data retrieval call binding the contract method 0x89676073.
//
// Solidity: function totalEffectiveBalance() view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) TotalEffectiveBalance() (*big.Int, error) {
	return _SdlPool.Contract.TotalEffectiveBalance(&_SdlPool.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_SdlPool *SdlPoolCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_SdlPool *SdlPoolSession) TotalStaked() (*big.Int, error) {
	return _SdlPool.Contract.TotalStaked(&_SdlPool.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_SdlPool *SdlPoolCallerSession) TotalStaked() (*big.Int, error) {
	return _SdlPool.Contract.TotalStaked(&_SdlPool.CallOpts)
}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_SdlPool *SdlPoolCaller) WithdrawableRewards(opts *bind.CallOpts, _account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SdlPool.contract.Call(opts, &out, "withdrawableRewards", _account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_SdlPool *SdlPoolSession) WithdrawableRewards(_account common.Address) ([]*big.Int, error) {
	return _SdlPool.Contract.WithdrawableRewards(&_SdlPool.CallOpts, _account)
}

// WithdrawableRewards is a free data retrieval call binding the contract method 0x0f14b4d6.
//
// Solidity: function withdrawableRewards(address _account) view returns(uint256[])
func (_SdlPool *SdlPoolCallerSession) WithdrawableRewards(_account common.Address) ([]*big.Int, error) {
	return _SdlPool.Contract.WithdrawableRewards(&_SdlPool.CallOpts, _account)
}

// RewardsPoolControllerInit is a paid mutator transaction binding the contract method 0x7c9097de.
//
// Solidity: function __RewardsPoolController_init() returns()
func (_SdlPool *SdlPoolTransactor) RewardsPoolControllerInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "__RewardsPoolController_init")
}

// RewardsPoolControllerInit is a paid mutator transaction binding the contract method 0x7c9097de.
//
// Solidity: function __RewardsPoolController_init() returns()
func (_SdlPool *SdlPoolSession) RewardsPoolControllerInit() (*types.Transaction, error) {
	return _SdlPool.Contract.RewardsPoolControllerInit(&_SdlPool.TransactOpts)
}

// RewardsPoolControllerInit is a paid mutator transaction binding the contract method 0x7c9097de.
//
// Solidity: function __RewardsPoolController_init() returns()
func (_SdlPool *SdlPoolTransactorSession) RewardsPoolControllerInit() (*types.Transaction, error) {
	return _SdlPool.Contract.RewardsPoolControllerInit(&_SdlPool.TransactOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_SdlPool *SdlPoolTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "addToken", _token, _rewardsPool)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_SdlPool *SdlPoolSession) AddToken(_token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.AddToken(&_SdlPool.TransactOpts, _token, _rewardsPool)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _token, address _rewardsPool) returns()
func (_SdlPool *SdlPoolTransactorSession) AddToken(_token common.Address, _rewardsPool common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.AddToken(&_SdlPool.TransactOpts, _token, _rewardsPool)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "approve", _to, _lockId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolSession) Approve(_to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.Approve(&_SdlPool.TransactOpts, _to, _lockId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactorSession) Approve(_to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.Approve(&_SdlPool.TransactOpts, _to, _lockId)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_SdlPool *SdlPoolTransactor) DistributeToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "distributeToken", _token)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_SdlPool *SdlPoolSession) DistributeToken(_token common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.DistributeToken(&_SdlPool.TransactOpts, _token)
}

// DistributeToken is a paid mutator transaction binding the contract method 0x86d74037.
//
// Solidity: function distributeToken(address _token) returns()
func (_SdlPool *SdlPoolTransactorSession) DistributeToken(_token common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.DistributeToken(&_SdlPool.TransactOpts, _token)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_SdlPool *SdlPoolTransactor) DistributeTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "distributeTokens", _tokens)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_SdlPool *SdlPoolSession) DistributeTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.DistributeTokens(&_SdlPool.TransactOpts, _tokens)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x88951352.
//
// Solidity: function distributeTokens(address[] _tokens) returns()
func (_SdlPool *SdlPoolTransactorSession) DistributeTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.DistributeTokens(&_SdlPool.TransactOpts, _tokens)
}

// ExtendLockDuration is a paid mutator transaction binding the contract method 0xe231db76.
//
// Solidity: function extendLockDuration(uint256 _lockId, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolTransactor) ExtendLockDuration(opts *bind.TransactOpts, _lockId *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "extendLockDuration", _lockId, _lockingDuration)
}

// ExtendLockDuration is a paid mutator transaction binding the contract method 0xe231db76.
//
// Solidity: function extendLockDuration(uint256 _lockId, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolSession) ExtendLockDuration(_lockId *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.Contract.ExtendLockDuration(&_SdlPool.TransactOpts, _lockId, _lockingDuration)
}

// ExtendLockDuration is a paid mutator transaction binding the contract method 0xe231db76.
//
// Solidity: function extendLockDuration(uint256 _lockId, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolTransactorSession) ExtendLockDuration(_lockId *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.Contract.ExtendLockDuration(&_SdlPool.TransactOpts, _lockId, _lockingDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb0ed6a0.
//
// Solidity: function initialize(string _name, string _symbol, address _sdlToken, address _boostController, address _delegatorPool) returns()
func (_SdlPool *SdlPoolTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _sdlToken common.Address, _boostController common.Address, _delegatorPool common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "initialize", _name, _symbol, _sdlToken, _boostController, _delegatorPool)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb0ed6a0.
//
// Solidity: function initialize(string _name, string _symbol, address _sdlToken, address _boostController, address _delegatorPool) returns()
func (_SdlPool *SdlPoolSession) Initialize(_name string, _symbol string, _sdlToken common.Address, _boostController common.Address, _delegatorPool common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.Initialize(&_SdlPool.TransactOpts, _name, _symbol, _sdlToken, _boostController, _delegatorPool)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb0ed6a0.
//
// Solidity: function initialize(string _name, string _symbol, address _sdlToken, address _boostController, address _delegatorPool) returns()
func (_SdlPool *SdlPoolTransactorSession) Initialize(_name string, _symbol string, _sdlToken common.Address, _boostController common.Address, _delegatorPool common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.Initialize(&_SdlPool.TransactOpts, _name, _symbol, _sdlToken, _boostController, _delegatorPool)
}

// InitiateUnlock is a paid mutator transaction binding the contract method 0x2dc075e1.
//
// Solidity: function initiateUnlock(uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactor) InitiateUnlock(opts *bind.TransactOpts, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "initiateUnlock", _lockId)
}

// InitiateUnlock is a paid mutator transaction binding the contract method 0x2dc075e1.
//
// Solidity: function initiateUnlock(uint256 _lockId) returns()
func (_SdlPool *SdlPoolSession) InitiateUnlock(_lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.InitiateUnlock(&_SdlPool.TransactOpts, _lockId)
}

// InitiateUnlock is a paid mutator transaction binding the contract method 0x2dc075e1.
//
// Solidity: function initiateUnlock(uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactorSession) InitiateUnlock(_lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.InitiateUnlock(&_SdlPool.TransactOpts, _lockId)
}

// Migrate is a paid mutator transaction binding the contract method 0x76a19e5b.
//
// Solidity: function migrate(address _sender, uint256 _amount, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolTransactor) Migrate(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "migrate", _sender, _amount, _lockingDuration)
}

// Migrate is a paid mutator transaction binding the contract method 0x76a19e5b.
//
// Solidity: function migrate(address _sender, uint256 _amount, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolSession) Migrate(_sender common.Address, _amount *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.Contract.Migrate(&_SdlPool.TransactOpts, _sender, _amount, _lockingDuration)
}

// Migrate is a paid mutator transaction binding the contract method 0x76a19e5b.
//
// Solidity: function migrate(address _sender, uint256 _amount, uint64 _lockingDuration) returns()
func (_SdlPool *SdlPoolTransactorSession) Migrate(_sender common.Address, _amount *big.Int, _lockingDuration uint64) (*types.Transaction, error) {
	return _SdlPool.Contract.Migrate(&_SdlPool.TransactOpts, _sender, _amount, _lockingDuration)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_SdlPool *SdlPoolTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "onTokenTransfer", _sender, _value, _calldata)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_SdlPool *SdlPoolSession) OnTokenTransfer(_sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.OnTokenTransfer(&_SdlPool.TransactOpts, _sender, _value, _calldata)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes _calldata) returns()
func (_SdlPool *SdlPoolTransactorSession) OnTokenTransfer(_sender common.Address, _value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.OnTokenTransfer(&_SdlPool.TransactOpts, _sender, _value, _calldata)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_SdlPool *SdlPoolTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "removeToken", _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_SdlPool *SdlPoolSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.RemoveToken(&_SdlPool.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns()
func (_SdlPool *SdlPoolTransactorSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.RemoveToken(&_SdlPool.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SdlPool *SdlPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SdlPool *SdlPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _SdlPool.Contract.RenounceOwnership(&_SdlPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SdlPool *SdlPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SdlPool.Contract.RenounceOwnership(&_SdlPool.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "safeTransferFrom", _from, _to, _lockId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolSession) SafeTransferFrom(_from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.SafeTransferFrom(&_SdlPool.TransactOpts, _from, _to, _lockId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.SafeTransferFrom(&_SdlPool.TransactOpts, _from, _to, _lockId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId, bytes _data) returns()
func (_SdlPool *SdlPoolTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _lockId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "safeTransferFrom0", _from, _to, _lockId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId, bytes _data) returns()
func (_SdlPool *SdlPoolSession) SafeTransferFrom0(_from common.Address, _to common.Address, _lockId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.SafeTransferFrom0(&_SdlPool.TransactOpts, _from, _to, _lockId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _lockId, bytes _data) returns()
func (_SdlPool *SdlPoolTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _lockId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.SafeTransferFrom0(&_SdlPool.TransactOpts, _from, _to, _lockId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_SdlPool *SdlPoolTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_SdlPool *SdlPoolSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _SdlPool.Contract.SetApprovalForAll(&_SdlPool.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_SdlPool *SdlPoolTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _SdlPool.Contract.SetApprovalForAll(&_SdlPool.TransactOpts, _operator, _approved)
}

// SetBoostController is a paid mutator transaction binding the contract method 0x1b23009d.
//
// Solidity: function setBoostController(address _boostController) returns()
func (_SdlPool *SdlPoolTransactor) SetBoostController(opts *bind.TransactOpts, _boostController common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "setBoostController", _boostController)
}

// SetBoostController is a paid mutator transaction binding the contract method 0x1b23009d.
//
// Solidity: function setBoostController(address _boostController) returns()
func (_SdlPool *SdlPoolSession) SetBoostController(_boostController common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.SetBoostController(&_SdlPool.TransactOpts, _boostController)
}

// SetBoostController is a paid mutator transaction binding the contract method 0x1b23009d.
//
// Solidity: function setBoostController(address _boostController) returns()
func (_SdlPool *SdlPoolTransactorSession) SetBoostController(_boostController common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.SetBoostController(&_SdlPool.TransactOpts, _boostController)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "transferFrom", _from, _to, _lockId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolSession) TransferFrom(_from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.TransferFrom(&_SdlPool.TransactOpts, _from, _to, _lockId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _lockId) returns()
func (_SdlPool *SdlPoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _lockId *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.TransferFrom(&_SdlPool.TransactOpts, _from, _to, _lockId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SdlPool *SdlPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SdlPool *SdlPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.TransferOwnership(&_SdlPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SdlPool *SdlPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.TransferOwnership(&_SdlPool.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SdlPool *SdlPoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SdlPool *SdlPoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.UpgradeTo(&_SdlPool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SdlPool *SdlPoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.UpgradeTo(&_SdlPool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SdlPool *SdlPoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SdlPool *SdlPoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.UpgradeToAndCall(&_SdlPool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SdlPool *SdlPoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SdlPool.Contract.UpgradeToAndCall(&_SdlPool.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockId, uint256 _amount) returns()
func (_SdlPool *SdlPoolTransactor) Withdraw(opts *bind.TransactOpts, _lockId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "withdraw", _lockId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockId, uint256 _amount) returns()
func (_SdlPool *SdlPoolSession) Withdraw(_lockId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.Withdraw(&_SdlPool.TransactOpts, _lockId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockId, uint256 _amount) returns()
func (_SdlPool *SdlPoolTransactorSession) Withdraw(_lockId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SdlPool.Contract.Withdraw(&_SdlPool.TransactOpts, _lockId, _amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_SdlPool *SdlPoolTransactor) WithdrawRewards(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.contract.Transact(opts, "withdrawRewards", _tokens)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_SdlPool *SdlPoolSession) WithdrawRewards(_tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.WithdrawRewards(&_SdlPool.TransactOpts, _tokens)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x62d76d06.
//
// Solidity: function withdrawRewards(address[] _tokens) returns()
func (_SdlPool *SdlPoolTransactorSession) WithdrawRewards(_tokens []common.Address) (*types.Transaction, error) {
	return _SdlPool.Contract.WithdrawRewards(&_SdlPool.TransactOpts, _tokens)
}

// SdlPoolAddTokenIterator is returned from FilterAddToken and is used to iterate over the raw logs and unpacked data for AddToken events raised by the SdlPool contract.
type SdlPoolAddTokenIterator struct {
	Event *SdlPoolAddToken // Event containing the contract specifics and raw log

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
func (it *SdlPoolAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolAddToken)
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
		it.Event = new(SdlPoolAddToken)
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
func (it *SdlPoolAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolAddToken represents a AddToken event raised by the SdlPool contract.
type SdlPoolAddToken struct {
	Token       common.Address
	RewardsPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddToken is a free log retrieval operation binding the contract event 0xdbf34b45b47a653cf4940cccbec765f72d4d63de3237306905bfc0ee28832362.
//
// Solidity: event AddToken(address indexed token, address rewardsPool)
func (_SdlPool *SdlPoolFilterer) FilterAddToken(opts *bind.FilterOpts, token []common.Address) (*SdlPoolAddTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "AddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolAddTokenIterator{contract: _SdlPool.contract, event: "AddToken", logs: logs, sub: sub}, nil
}

// WatchAddToken is a free log subscription operation binding the contract event 0xdbf34b45b47a653cf4940cccbec765f72d4d63de3237306905bfc0ee28832362.
//
// Solidity: event AddToken(address indexed token, address rewardsPool)
func (_SdlPool *SdlPoolFilterer) WatchAddToken(opts *bind.WatchOpts, sink chan<- *SdlPoolAddToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "AddToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolAddToken)
				if err := _SdlPool.contract.UnpackLog(event, "AddToken", log); err != nil {
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
func (_SdlPool *SdlPoolFilterer) ParseAddToken(log types.Log) (*SdlPoolAddToken, error) {
	event := new(SdlPoolAddToken)
	if err := _SdlPool.contract.UnpackLog(event, "AddToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SdlPool contract.
type SdlPoolAdminChangedIterator struct {
	Event *SdlPoolAdminChanged // Event containing the contract specifics and raw log

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
func (it *SdlPoolAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolAdminChanged)
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
		it.Event = new(SdlPoolAdminChanged)
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
func (it *SdlPoolAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolAdminChanged represents a AdminChanged event raised by the SdlPool contract.
type SdlPoolAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SdlPool *SdlPoolFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SdlPoolAdminChangedIterator, error) {

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SdlPoolAdminChangedIterator{contract: _SdlPool.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SdlPool *SdlPoolFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SdlPoolAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolAdminChanged)
				if err := _SdlPool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SdlPool *SdlPoolFilterer) ParseAdminChanged(log types.Log) (*SdlPoolAdminChanged, error) {
	event := new(SdlPoolAdminChanged)
	if err := _SdlPool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SdlPool contract.
type SdlPoolApprovalIterator struct {
	Event *SdlPoolApproval // Event containing the contract specifics and raw log

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
func (it *SdlPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolApproval)
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
		it.Event = new(SdlPoolApproval)
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
func (it *SdlPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolApproval represents a Approval event raised by the SdlPool contract.
type SdlPoolApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SdlPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolApprovalIterator{contract: _SdlPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SdlPoolApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolApproval)
				if err := _SdlPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) ParseApproval(log types.Log) (*SdlPoolApproval, error) {
	event := new(SdlPoolApproval)
	if err := _SdlPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SdlPool contract.
type SdlPoolApprovalForAllIterator struct {
	Event *SdlPoolApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SdlPoolApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolApprovalForAll)
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
		it.Event = new(SdlPoolApprovalForAll)
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
func (it *SdlPoolApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolApprovalForAll represents a ApprovalForAll event raised by the SdlPool contract.
type SdlPoolApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SdlPool *SdlPoolFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SdlPoolApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolApprovalForAllIterator{contract: _SdlPool.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SdlPool *SdlPoolFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SdlPoolApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolApprovalForAll)
				if err := _SdlPool.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SdlPool *SdlPoolFilterer) ParseApprovalForAll(log types.Log) (*SdlPoolApprovalForAll, error) {
	event := new(SdlPoolApprovalForAll)
	if err := _SdlPool.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SdlPool contract.
type SdlPoolBeaconUpgradedIterator struct {
	Event *SdlPoolBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SdlPoolBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolBeaconUpgraded)
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
		it.Event = new(SdlPoolBeaconUpgraded)
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
func (it *SdlPoolBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolBeaconUpgraded represents a BeaconUpgraded event raised by the SdlPool contract.
type SdlPoolBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SdlPool *SdlPoolFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SdlPoolBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolBeaconUpgradedIterator{contract: _SdlPool.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SdlPool *SdlPoolFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SdlPoolBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolBeaconUpgraded)
				if err := _SdlPool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SdlPool *SdlPoolFilterer) ParseBeaconUpgraded(log types.Log) (*SdlPoolBeaconUpgraded, error) {
	event := new(SdlPoolBeaconUpgraded)
	if err := _SdlPool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolCreateLockIterator is returned from FilterCreateLock and is used to iterate over the raw logs and unpacked data for CreateLock events raised by the SdlPool contract.
type SdlPoolCreateLockIterator struct {
	Event *SdlPoolCreateLock // Event containing the contract specifics and raw log

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
func (it *SdlPoolCreateLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolCreateLock)
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
		it.Event = new(SdlPoolCreateLock)
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
func (it *SdlPoolCreateLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolCreateLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolCreateLock represents a CreateLock event raised by the SdlPool contract.
type SdlPoolCreateLock struct {
	Owner           common.Address
	LockId          *big.Int
	Amount          *big.Int
	BoostAmount     *big.Int
	LockingDuration uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateLock is a free log retrieval operation binding the contract event 0x0bd0744e02d09668fbdaf7745e9ae454f27484fcf18308ebc65d8508f21d5912.
//
// Solidity: event CreateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) FilterCreateLock(opts *bind.FilterOpts, owner []common.Address, lockId []*big.Int) (*SdlPoolCreateLockIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "CreateLock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolCreateLockIterator{contract: _SdlPool.contract, event: "CreateLock", logs: logs, sub: sub}, nil
}

// WatchCreateLock is a free log subscription operation binding the contract event 0x0bd0744e02d09668fbdaf7745e9ae454f27484fcf18308ebc65d8508f21d5912.
//
// Solidity: event CreateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) WatchCreateLock(opts *bind.WatchOpts, sink chan<- *SdlPoolCreateLock, owner []common.Address, lockId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "CreateLock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolCreateLock)
				if err := _SdlPool.contract.UnpackLog(event, "CreateLock", log); err != nil {
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

// ParseCreateLock is a log parse operation binding the contract event 0x0bd0744e02d09668fbdaf7745e9ae454f27484fcf18308ebc65d8508f21d5912.
//
// Solidity: event CreateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) ParseCreateLock(log types.Log) (*SdlPoolCreateLock, error) {
	event := new(SdlPoolCreateLock)
	if err := _SdlPool.contract.UnpackLog(event, "CreateLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SdlPool contract.
type SdlPoolInitializedIterator struct {
	Event *SdlPoolInitialized // Event containing the contract specifics and raw log

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
func (it *SdlPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolInitialized)
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
		it.Event = new(SdlPoolInitialized)
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
func (it *SdlPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolInitialized represents a Initialized event raised by the SdlPool contract.
type SdlPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SdlPool *SdlPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*SdlPoolInitializedIterator, error) {

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SdlPoolInitializedIterator{contract: _SdlPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SdlPool *SdlPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SdlPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolInitialized)
				if err := _SdlPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SdlPool *SdlPoolFilterer) ParseInitialized(log types.Log) (*SdlPoolInitialized, error) {
	event := new(SdlPoolInitialized)
	if err := _SdlPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolInitiateUnlockIterator is returned from FilterInitiateUnlock and is used to iterate over the raw logs and unpacked data for InitiateUnlock events raised by the SdlPool contract.
type SdlPoolInitiateUnlockIterator struct {
	Event *SdlPoolInitiateUnlock // Event containing the contract specifics and raw log

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
func (it *SdlPoolInitiateUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolInitiateUnlock)
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
		it.Event = new(SdlPoolInitiateUnlock)
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
func (it *SdlPoolInitiateUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolInitiateUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolInitiateUnlock represents a InitiateUnlock event raised by the SdlPool contract.
type SdlPoolInitiateUnlock struct {
	Owner  common.Address
	LockId *big.Int
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitiateUnlock is a free log retrieval operation binding the contract event 0x07728a9afd9c9ac8253019aedb66f63530f4fdeb57357c7eb64c601b528e6457.
//
// Solidity: event InitiateUnlock(address indexed owner, uint256 indexed lockId, uint64 expiry)
func (_SdlPool *SdlPoolFilterer) FilterInitiateUnlock(opts *bind.FilterOpts, owner []common.Address, lockId []*big.Int) (*SdlPoolInitiateUnlockIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "InitiateUnlock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolInitiateUnlockIterator{contract: _SdlPool.contract, event: "InitiateUnlock", logs: logs, sub: sub}, nil
}

// WatchInitiateUnlock is a free log subscription operation binding the contract event 0x07728a9afd9c9ac8253019aedb66f63530f4fdeb57357c7eb64c601b528e6457.
//
// Solidity: event InitiateUnlock(address indexed owner, uint256 indexed lockId, uint64 expiry)
func (_SdlPool *SdlPoolFilterer) WatchInitiateUnlock(opts *bind.WatchOpts, sink chan<- *SdlPoolInitiateUnlock, owner []common.Address, lockId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "InitiateUnlock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolInitiateUnlock)
				if err := _SdlPool.contract.UnpackLog(event, "InitiateUnlock", log); err != nil {
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

// ParseInitiateUnlock is a log parse operation binding the contract event 0x07728a9afd9c9ac8253019aedb66f63530f4fdeb57357c7eb64c601b528e6457.
//
// Solidity: event InitiateUnlock(address indexed owner, uint256 indexed lockId, uint64 expiry)
func (_SdlPool *SdlPoolFilterer) ParseInitiateUnlock(log types.Log) (*SdlPoolInitiateUnlock, error) {
	event := new(SdlPoolInitiateUnlock)
	if err := _SdlPool.contract.UnpackLog(event, "InitiateUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SdlPool contract.
type SdlPoolOwnershipTransferredIterator struct {
	Event *SdlPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SdlPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolOwnershipTransferred)
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
		it.Event = new(SdlPoolOwnershipTransferred)
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
func (it *SdlPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolOwnershipTransferred represents a OwnershipTransferred event raised by the SdlPool contract.
type SdlPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SdlPool *SdlPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SdlPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolOwnershipTransferredIterator{contract: _SdlPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SdlPool *SdlPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SdlPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolOwnershipTransferred)
				if err := _SdlPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SdlPool *SdlPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SdlPoolOwnershipTransferred, error) {
	event := new(SdlPoolOwnershipTransferred)
	if err := _SdlPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolRemoveTokenIterator is returned from FilterRemoveToken and is used to iterate over the raw logs and unpacked data for RemoveToken events raised by the SdlPool contract.
type SdlPoolRemoveTokenIterator struct {
	Event *SdlPoolRemoveToken // Event containing the contract specifics and raw log

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
func (it *SdlPoolRemoveTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolRemoveToken)
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
		it.Event = new(SdlPoolRemoveToken)
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
func (it *SdlPoolRemoveTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolRemoveTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolRemoveToken represents a RemoveToken event raised by the SdlPool contract.
type SdlPoolRemoveToken struct {
	Token       common.Address
	RewardsPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveToken is a free log retrieval operation binding the contract event 0x39dcd754ec63af4b82c4c569ff1b6b4e55a8038e6545844747e54f2f2d4e8e50.
//
// Solidity: event RemoveToken(address indexed token, address rewardsPool)
func (_SdlPool *SdlPoolFilterer) FilterRemoveToken(opts *bind.FilterOpts, token []common.Address) (*SdlPoolRemoveTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "RemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolRemoveTokenIterator{contract: _SdlPool.contract, event: "RemoveToken", logs: logs, sub: sub}, nil
}

// WatchRemoveToken is a free log subscription operation binding the contract event 0x39dcd754ec63af4b82c4c569ff1b6b4e55a8038e6545844747e54f2f2d4e8e50.
//
// Solidity: event RemoveToken(address indexed token, address rewardsPool)
func (_SdlPool *SdlPoolFilterer) WatchRemoveToken(opts *bind.WatchOpts, sink chan<- *SdlPoolRemoveToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "RemoveToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolRemoveToken)
				if err := _SdlPool.contract.UnpackLog(event, "RemoveToken", log); err != nil {
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
func (_SdlPool *SdlPoolFilterer) ParseRemoveToken(log types.Log) (*SdlPoolRemoveToken, error) {
	event := new(SdlPoolRemoveToken)
	if err := _SdlPool.contract.UnpackLog(event, "RemoveToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SdlPool contract.
type SdlPoolTransferIterator struct {
	Event *SdlPoolTransfer // Event containing the contract specifics and raw log

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
func (it *SdlPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolTransfer)
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
		it.Event = new(SdlPoolTransfer)
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
func (it *SdlPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolTransfer represents a Transfer event raised by the SdlPool contract.
type SdlPoolTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SdlPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolTransferIterator{contract: _SdlPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SdlPoolTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolTransfer)
				if err := _SdlPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SdlPool *SdlPoolFilterer) ParseTransfer(log types.Log) (*SdlPoolTransfer, error) {
	event := new(SdlPoolTransfer)
	if err := _SdlPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolUpdateLockIterator is returned from FilterUpdateLock and is used to iterate over the raw logs and unpacked data for UpdateLock events raised by the SdlPool contract.
type SdlPoolUpdateLockIterator struct {
	Event *SdlPoolUpdateLock // Event containing the contract specifics and raw log

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
func (it *SdlPoolUpdateLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolUpdateLock)
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
		it.Event = new(SdlPoolUpdateLock)
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
func (it *SdlPoolUpdateLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolUpdateLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolUpdateLock represents a UpdateLock event raised by the SdlPool contract.
type SdlPoolUpdateLock struct {
	Owner           common.Address
	LockId          *big.Int
	Amount          *big.Int
	BoostAmount     *big.Int
	LockingDuration uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateLock is a free log retrieval operation binding the contract event 0x8dbb057b85ca29d98e176658d7a7843f454205d9d623e5f2e7c93919334b0995.
//
// Solidity: event UpdateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) FilterUpdateLock(opts *bind.FilterOpts, owner []common.Address, lockId []*big.Int) (*SdlPoolUpdateLockIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "UpdateLock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolUpdateLockIterator{contract: _SdlPool.contract, event: "UpdateLock", logs: logs, sub: sub}, nil
}

// WatchUpdateLock is a free log subscription operation binding the contract event 0x8dbb057b85ca29d98e176658d7a7843f454205d9d623e5f2e7c93919334b0995.
//
// Solidity: event UpdateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) WatchUpdateLock(opts *bind.WatchOpts, sink chan<- *SdlPoolUpdateLock, owner []common.Address, lockId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "UpdateLock", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolUpdateLock)
				if err := _SdlPool.contract.UnpackLog(event, "UpdateLock", log); err != nil {
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

// ParseUpdateLock is a log parse operation binding the contract event 0x8dbb057b85ca29d98e176658d7a7843f454205d9d623e5f2e7c93919334b0995.
//
// Solidity: event UpdateLock(address indexed owner, uint256 indexed lockId, uint256 amount, uint256 boostAmount, uint64 lockingDuration)
func (_SdlPool *SdlPoolFilterer) ParseUpdateLock(log types.Log) (*SdlPoolUpdateLock, error) {
	event := new(SdlPoolUpdateLock)
	if err := _SdlPool.contract.UnpackLog(event, "UpdateLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SdlPool contract.
type SdlPoolUpgradedIterator struct {
	Event *SdlPoolUpgraded // Event containing the contract specifics and raw log

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
func (it *SdlPoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolUpgraded)
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
		it.Event = new(SdlPoolUpgraded)
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
func (it *SdlPoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolUpgraded represents a Upgraded event raised by the SdlPool contract.
type SdlPoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SdlPool *SdlPoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SdlPoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolUpgradedIterator{contract: _SdlPool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SdlPool *SdlPoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SdlPoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolUpgraded)
				if err := _SdlPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SdlPool *SdlPoolFilterer) ParseUpgraded(log types.Log) (*SdlPoolUpgraded, error) {
	event := new(SdlPoolUpgraded)
	if err := _SdlPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SdlPool contract.
type SdlPoolWithdrawIterator struct {
	Event *SdlPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *SdlPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolWithdraw)
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
		it.Event = new(SdlPoolWithdraw)
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
func (it *SdlPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolWithdraw represents a Withdraw event raised by the SdlPool contract.
type SdlPoolWithdraw struct {
	Owner  common.Address
	LockId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed lockId, uint256 amount)
func (_SdlPool *SdlPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, owner []common.Address, lockId []*big.Int) (*SdlPoolWithdrawIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "Withdraw", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolWithdrawIterator{contract: _SdlPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed lockId, uint256 amount)
func (_SdlPool *SdlPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SdlPoolWithdraw, owner []common.Address, lockId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var lockIdRule []interface{}
	for _, lockIdItem := range lockId {
		lockIdRule = append(lockIdRule, lockIdItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "Withdraw", ownerRule, lockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolWithdraw)
				if err := _SdlPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed lockId, uint256 amount)
func (_SdlPool *SdlPoolFilterer) ParseWithdraw(log types.Log) (*SdlPoolWithdraw, error) {
	event := new(SdlPoolWithdraw)
	if err := _SdlPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdlPoolWithdrawRewardsIterator is returned from FilterWithdrawRewards and is used to iterate over the raw logs and unpacked data for WithdrawRewards events raised by the SdlPool contract.
type SdlPoolWithdrawRewardsIterator struct {
	Event *SdlPoolWithdrawRewards // Event containing the contract specifics and raw log

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
func (it *SdlPoolWithdrawRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdlPoolWithdrawRewards)
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
		it.Event = new(SdlPoolWithdrawRewards)
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
func (it *SdlPoolWithdrawRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdlPoolWithdrawRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdlPoolWithdrawRewards represents a WithdrawRewards event raised by the SdlPool contract.
type SdlPoolWithdrawRewards struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRewards is a free log retrieval operation binding the contract event 0xeeab3fd62be4cb59cbdc42d5b0f676a3597ff387b9a85e62330cc17c2a3603db.
//
// Solidity: event WithdrawRewards(address indexed account)
func (_SdlPool *SdlPoolFilterer) FilterWithdrawRewards(opts *bind.FilterOpts, account []common.Address) (*SdlPoolWithdrawRewardsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SdlPool.contract.FilterLogs(opts, "WithdrawRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return &SdlPoolWithdrawRewardsIterator{contract: _SdlPool.contract, event: "WithdrawRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawRewards is a free log subscription operation binding the contract event 0xeeab3fd62be4cb59cbdc42d5b0f676a3597ff387b9a85e62330cc17c2a3603db.
//
// Solidity: event WithdrawRewards(address indexed account)
func (_SdlPool *SdlPoolFilterer) WatchWithdrawRewards(opts *bind.WatchOpts, sink chan<- *SdlPoolWithdrawRewards, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SdlPool.contract.WatchLogs(opts, "WithdrawRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdlPoolWithdrawRewards)
				if err := _SdlPool.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
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
func (_SdlPool *SdlPoolFilterer) ParseWithdrawRewards(log types.Log) (*SdlPoolWithdrawRewards, error) {
	event := new(SdlPoolWithdrawRewards)
	if err := _SdlPool.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
