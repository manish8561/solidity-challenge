// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package biddingwar

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

// BiddingwarMetaData contains all meta data concerning the Biddingwar contract.
var BiddingwarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"bid\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"when\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"BidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"bid\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"when\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"WinEvent\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"bidCounter\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"when\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"declareResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameEndTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BiddingwarABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingwarMetaData.ABI instead.
var BiddingwarABI = BiddingwarMetaData.ABI

// Biddingwar is an auto generated Go binding around an Ethereum contract.
type Biddingwar struct {
	BiddingwarCaller     // Read-only binding to the contract
	BiddingwarTransactor // Write-only binding to the contract
	BiddingwarFilterer   // Log filterer for contract events
}

// BiddingwarCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingwarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingwarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingwarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingwarSession struct {
	Contract     *Biddingwar       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingwarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingwarCallerSession struct {
	Contract *BiddingwarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BiddingwarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingwarTransactorSession struct {
	Contract     *BiddingwarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BiddingwarRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingwarRaw struct {
	Contract *Biddingwar // Generic contract binding to access the raw methods on
}

// BiddingwarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingwarCallerRaw struct {
	Contract *BiddingwarCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingwarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingwarTransactorRaw struct {
	Contract *BiddingwarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiddingwar creates a new instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwar(address common.Address, backend bind.ContractBackend) (*Biddingwar, error) {
	contract, err := bindBiddingwar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Biddingwar{BiddingwarCaller: BiddingwarCaller{contract: contract}, BiddingwarTransactor: BiddingwarTransactor{contract: contract}, BiddingwarFilterer: BiddingwarFilterer{contract: contract}}, nil
}

// NewBiddingwarCaller creates a new read-only instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarCaller(address common.Address, caller bind.ContractCaller) (*BiddingwarCaller, error) {
	contract, err := bindBiddingwar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarCaller{contract: contract}, nil
}

// NewBiddingwarTransactor creates a new write-only instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingwarTransactor, error) {
	contract, err := bindBiddingwar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarTransactor{contract: contract}, nil
}

// NewBiddingwarFilterer creates a new log filterer instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingwarFilterer, error) {
	contract, err := bindBiddingwar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingwarFilterer{contract: contract}, nil
}

// bindBiddingwar binds a generic wrapper to an already deployed contract.
func bindBiddingwar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BiddingwarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwar *BiddingwarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwar.Contract.BiddingwarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwar *BiddingwarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.Contract.BiddingwarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwar *BiddingwarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwar.Contract.BiddingwarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwar *BiddingwarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwar *BiddingwarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwar *BiddingwarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwar.Contract.contract.Transact(opts, method, params...)
}

// BidCounter is a free data retrieval call binding the contract method 0x03f69354.
//
// Solidity: function bidCounter() view returns(uint128)
func (_Biddingwar *BiddingwarCaller) BidCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "bidCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCounter is a free data retrieval call binding the contract method 0x03f69354.
//
// Solidity: function bidCounter() view returns(uint128)
func (_Biddingwar *BiddingwarSession) BidCounter() (*big.Int, error) {
	return _Biddingwar.Contract.BidCounter(&_Biddingwar.CallOpts)
}

// BidCounter is a free data retrieval call binding the contract method 0x03f69354.
//
// Solidity: function bidCounter() view returns(uint128)
func (_Biddingwar *BiddingwarCallerSession) BidCounter() (*big.Int, error) {
	return _Biddingwar.Contract.BidCounter(&_Biddingwar.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, address user, uint64 when)
func (_Biddingwar *BiddingwarCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	When   uint64
}, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "bids", arg0)

	outstruct := new(struct {
		Amount *big.Int
		User   common.Address
		When   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.When = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, address user, uint64 when)
func (_Biddingwar *BiddingwarSession) Bids(arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	When   uint64
}, error) {
	return _Biddingwar.Contract.Bids(&_Biddingwar.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, address user, uint64 when)
func (_Biddingwar *BiddingwarCallerSession) Bids(arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	When   uint64
}, error) {
	return _Biddingwar.Contract.Bids(&_Biddingwar.CallOpts, arg0)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarSession) Commission() (*big.Int, error) {
	return _Biddingwar.Contract.Commission(&_Biddingwar.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) Commission() (*big.Int, error) {
	return _Biddingwar.Contract.Commission(&_Biddingwar.CallOpts)
}

// GameEndTime is a free data retrieval call binding the contract method 0x44d9bc5f.
//
// Solidity: function gameEndTime() view returns(uint64)
func (_Biddingwar *BiddingwarCaller) GameEndTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "gameEndTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GameEndTime is a free data retrieval call binding the contract method 0x44d9bc5f.
//
// Solidity: function gameEndTime() view returns(uint64)
func (_Biddingwar *BiddingwarSession) GameEndTime() (uint64, error) {
	return _Biddingwar.Contract.GameEndTime(&_Biddingwar.CallOpts)
}

// GameEndTime is a free data retrieval call binding the contract method 0x44d9bc5f.
//
// Solidity: function gameEndTime() view returns(uint64)
func (_Biddingwar *BiddingwarCallerSession) GameEndTime() (uint64, error) {
	return _Biddingwar.Contract.GameEndTime(&_Biddingwar.CallOpts)
}

// GameStatus is a free data retrieval call binding the contract method 0x722fed7e.
//
// Solidity: function gameStatus() view returns(bool)
func (_Biddingwar *BiddingwarCaller) GameStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "gameStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameStatus is a free data retrieval call binding the contract method 0x722fed7e.
//
// Solidity: function gameStatus() view returns(bool)
func (_Biddingwar *BiddingwarSession) GameStatus() (bool, error) {
	return _Biddingwar.Contract.GameStatus(&_Biddingwar.CallOpts)
}

// GameStatus is a free data retrieval call binding the contract method 0x722fed7e.
//
// Solidity: function gameStatus() view returns(bool)
func (_Biddingwar *BiddingwarCallerSession) GameStatus() (bool, error) {
	return _Biddingwar.Contract.GameStatus(&_Biddingwar.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwar *BiddingwarSession) HighestBid() (*big.Int, error) {
	return _Biddingwar.Contract.HighestBid(&_Biddingwar.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) HighestBid() (*big.Int, error) {
	return _Biddingwar.Contract.HighestBid(&_Biddingwar.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingwar *BiddingwarCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingwar *BiddingwarSession) HighestBidder() (common.Address, error) {
	return _Biddingwar.Contract.HighestBidder(&_Biddingwar.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingwar *BiddingwarCallerSession) HighestBidder() (common.Address, error) {
	return _Biddingwar.Contract.HighestBidder(&_Biddingwar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarSession) Owner() (common.Address, error) {
	return _Biddingwar.Contract.Owner(&_Biddingwar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarCallerSession) Owner() (common.Address, error) {
	return _Biddingwar.Contract.Owner(&_Biddingwar.CallOpts)
}

// DeclareResult is a paid mutator transaction binding the contract method 0x8574c296.
//
// Solidity: function declareResult() returns()
func (_Biddingwar *BiddingwarTransactor) DeclareResult(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "declareResult")
}

// DeclareResult is a paid mutator transaction binding the contract method 0x8574c296.
//
// Solidity: function declareResult() returns()
func (_Biddingwar *BiddingwarSession) DeclareResult() (*types.Transaction, error) {
	return _Biddingwar.Contract.DeclareResult(&_Biddingwar.TransactOpts)
}

// DeclareResult is a paid mutator transaction binding the contract method 0x8574c296.
//
// Solidity: function declareResult() returns()
func (_Biddingwar *BiddingwarTransactorSession) DeclareResult() (*types.Transaction, error) {
	return _Biddingwar.Contract.DeclareResult(&_Biddingwar.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_Biddingwar *BiddingwarTransactor) PlaceBid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "placeBid")
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_Biddingwar *BiddingwarSession) PlaceBid() (*types.Transaction, error) {
	return _Biddingwar.Contract.PlaceBid(&_Biddingwar.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() payable returns()
func (_Biddingwar *BiddingwarTransactorSession) PlaceBid() (*types.Transaction, error) {
	return _Biddingwar.Contract.PlaceBid(&_Biddingwar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Biddingwar.Contract.RenounceOwnership(&_Biddingwar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Biddingwar.Contract.RenounceOwnership(&_Biddingwar.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.TransferOwnership(&_Biddingwar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.TransferOwnership(&_Biddingwar.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Biddingwar *BiddingwarTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Biddingwar.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Biddingwar *BiddingwarSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Biddingwar.Contract.Fallback(&_Biddingwar.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Biddingwar *BiddingwarTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Biddingwar.Contract.Fallback(&_Biddingwar.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarSession) Receive() (*types.Transaction, error) {
	return _Biddingwar.Contract.Receive(&_Biddingwar.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarTransactorSession) Receive() (*types.Transaction, error) {
	return _Biddingwar.Contract.Receive(&_Biddingwar.TransactOpts)
}

// BiddingwarBidEventIterator is returned from FilterBidEvent and is used to iterate over the raw logs and unpacked data for BidEvent events raised by the Biddingwar contract.
type BiddingwarBidEventIterator struct {
	Event *BiddingwarBidEvent // Event containing the contract specifics and raw log

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
func (it *BiddingwarBidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarBidEvent)
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
		it.Event = new(BiddingwarBidEvent)
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
func (it *BiddingwarBidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarBidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarBidEvent represents a BidEvent event raised by the Biddingwar contract.
type BiddingwarBidEvent struct {
	Bid    *big.Int
	Amount *big.Int
	When   uint64
	User   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidEvent is a free log retrieval operation binding the contract event 0x042ced85334596236356fb01034b15f35a1abbc3c7c51ba1b8d3ad152bfeda79.
//
// Solidity: event BidEvent(uint128 indexed bid, uint256 amount, uint64 when, address indexed user)
func (_Biddingwar *BiddingwarFilterer) FilterBidEvent(opts *bind.FilterOpts, bid []*big.Int, user []common.Address) (*BiddingwarBidEventIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "BidEvent", bidRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BiddingwarBidEventIterator{contract: _Biddingwar.contract, event: "BidEvent", logs: logs, sub: sub}, nil
}

// WatchBidEvent is a free log subscription operation binding the contract event 0x042ced85334596236356fb01034b15f35a1abbc3c7c51ba1b8d3ad152bfeda79.
//
// Solidity: event BidEvent(uint128 indexed bid, uint256 amount, uint64 when, address indexed user)
func (_Biddingwar *BiddingwarFilterer) WatchBidEvent(opts *bind.WatchOpts, sink chan<- *BiddingwarBidEvent, bid []*big.Int, user []common.Address) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "BidEvent", bidRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarBidEvent)
				if err := _Biddingwar.contract.UnpackLog(event, "BidEvent", log); err != nil {
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

// ParseBidEvent is a log parse operation binding the contract event 0x042ced85334596236356fb01034b15f35a1abbc3c7c51ba1b8d3ad152bfeda79.
//
// Solidity: event BidEvent(uint128 indexed bid, uint256 amount, uint64 when, address indexed user)
func (_Biddingwar *BiddingwarFilterer) ParseBidEvent(log types.Log) (*BiddingwarBidEvent, error) {
	event := new(BiddingwarBidEvent)
	if err := _Biddingwar.contract.UnpackLog(event, "BidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Biddingwar contract.
type BiddingwarOwnershipTransferredIterator struct {
	Event *BiddingwarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BiddingwarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarOwnershipTransferred)
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
		it.Event = new(BiddingwarOwnershipTransferred)
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
func (it *BiddingwarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarOwnershipTransferred represents a OwnershipTransferred event raised by the Biddingwar contract.
type BiddingwarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Biddingwar *BiddingwarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BiddingwarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BiddingwarOwnershipTransferredIterator{contract: _Biddingwar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Biddingwar *BiddingwarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BiddingwarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarOwnershipTransferred)
				if err := _Biddingwar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Biddingwar *BiddingwarFilterer) ParseOwnershipTransferred(log types.Log) (*BiddingwarOwnershipTransferred, error) {
	event := new(BiddingwarOwnershipTransferred)
	if err := _Biddingwar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarWinEventIterator is returned from FilterWinEvent and is used to iterate over the raw logs and unpacked data for WinEvent events raised by the Biddingwar contract.
type BiddingwarWinEventIterator struct {
	Event *BiddingwarWinEvent // Event containing the contract specifics and raw log

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
func (it *BiddingwarWinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarWinEvent)
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
		it.Event = new(BiddingwarWinEvent)
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
func (it *BiddingwarWinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarWinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarWinEvent represents a WinEvent event raised by the Biddingwar contract.
type BiddingwarWinEvent struct {
	Bid        *big.Int
	HighestBid *big.Int
	WinAmount  *big.Int
	When       uint64
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWinEvent is a free log retrieval operation binding the contract event 0xe900abebb29da21b2ce54f9bf92ace3e1f761a0a9a918473570dd7166057bc1f.
//
// Solidity: event WinEvent(uint128 indexed bid, uint256 highestBid, uint256 winAmount, uint64 when, address indexed winner)
func (_Biddingwar *BiddingwarFilterer) FilterWinEvent(opts *bind.FilterOpts, bid []*big.Int, winner []common.Address) (*BiddingwarWinEventIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "WinEvent", bidRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &BiddingwarWinEventIterator{contract: _Biddingwar.contract, event: "WinEvent", logs: logs, sub: sub}, nil
}

// WatchWinEvent is a free log subscription operation binding the contract event 0xe900abebb29da21b2ce54f9bf92ace3e1f761a0a9a918473570dd7166057bc1f.
//
// Solidity: event WinEvent(uint128 indexed bid, uint256 highestBid, uint256 winAmount, uint64 when, address indexed winner)
func (_Biddingwar *BiddingwarFilterer) WatchWinEvent(opts *bind.WatchOpts, sink chan<- *BiddingwarWinEvent, bid []*big.Int, winner []common.Address) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "WinEvent", bidRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarWinEvent)
				if err := _Biddingwar.contract.UnpackLog(event, "WinEvent", log); err != nil {
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

// ParseWinEvent is a log parse operation binding the contract event 0xe900abebb29da21b2ce54f9bf92ace3e1f761a0a9a918473570dd7166057bc1f.
//
// Solidity: event WinEvent(uint128 indexed bid, uint256 highestBid, uint256 winAmount, uint64 when, address indexed winner)
func (_Biddingwar *BiddingwarFilterer) ParseWinEvent(log types.Log) (*BiddingwarWinEvent, error) {
	event := new(BiddingwarWinEvent)
	if err := _Biddingwar.contract.UnpackLog(event, "WinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
