// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ind\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ind\",\"type\":\"uint256\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"setPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556102b4806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063700629041461005c57806370db6902146100985780638c64ea4a146100ad578063b93f9b0a146100ee578063fb75142714610117575b600080fd5b61008561006a366004610219565b6001600160a01b031660009081526001602052604090205490565b6040519081526020015b60405180910390f35b6100ab6100a636600461023b565b61012a565b005b6100d66100bb366004610265565b6002602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161008f565b6100d66100fc366004610265565b6000908152600260205260409020546001600160a01b031690565b6100ab61012536600461023b565b61019f565b6000546001600160a01b031633146101715760405162461bcd60e51b815260206004820152600560248201526437bbb732b960d91b60448201526064015b60405180910390fd5b600090815260026020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146101e15760405162461bcd60e51b815260206004820152600560248201526437bbb732b960d91b6044820152606401610168565b6001600160a01b03909116600090815260016020526040902055565b80356001600160a01b038116811461021457600080fd5b919050565b60006020828403121561022b57600080fd5b610234826101fd565b9392505050565b6000806040838503121561024e57600080fd5b610257836101fd565b946020939093013593505050565b60006020828403121561027757600080fd5b503591905056fea2646970667358221220f948c2211aa780c2d870af003e8f6548933cc6b9652d7a8d696c31b411ab4b9264736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 ind) view returns(address)
func (_Api *ApiCaller) GetAddress(opts *bind.CallOpts, ind *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAddress", ind)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 ind) view returns(address)
func (_Api *ApiSession) GetAddress(ind *big.Int) (common.Address, error) {
	return _Api.Contract.GetAddress(&_Api.CallOpts, ind)
}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 ind) view returns(address)
func (_Api *ApiCallerSession) GetAddress(ind *big.Int) (common.Address, error) {
	return _Api.Contract.GetAddress(&_Api.CallOpts, ind)
}

// GetPermit is a free data retrieval call binding the contract method 0x70062904.
//
// Solidity: function getPermit(address addr) view returns(uint256)
func (_Api *ApiCaller) GetPermit(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPermit", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPermit is a free data retrieval call binding the contract method 0x70062904.
//
// Solidity: function getPermit(address addr) view returns(uint256)
func (_Api *ApiSession) GetPermit(addr common.Address) (*big.Int, error) {
	return _Api.Contract.GetPermit(&_Api.CallOpts, addr)
}

// GetPermit is a free data retrieval call binding the contract method 0x70062904.
//
// Solidity: function getPermit(address addr) view returns(uint256)
func (_Api *ApiCallerSession) GetPermit(addr common.Address) (*big.Int, error) {
	return _Api.Contract.GetPermit(&_Api.CallOpts, addr)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Api *ApiCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "vaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Api *ApiSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// SetAddress is a paid mutator transaction binding the contract method 0x70db6902.
//
// Solidity: function setAddress(address newAddress, uint256 ind) returns()
func (_Api *ApiTransactor) SetAddress(opts *bind.TransactOpts, newAddress common.Address, ind *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setAddress", newAddress, ind)
}

// SetAddress is a paid mutator transaction binding the contract method 0x70db6902.
//
// Solidity: function setAddress(address newAddress, uint256 ind) returns()
func (_Api *ApiSession) SetAddress(newAddress common.Address, ind *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetAddress(&_Api.TransactOpts, newAddress, ind)
}

// SetAddress is a paid mutator transaction binding the contract method 0x70db6902.
//
// Solidity: function setAddress(address newAddress, uint256 ind) returns()
func (_Api *ApiTransactorSession) SetAddress(newAddress common.Address, ind *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetAddress(&_Api.TransactOpts, newAddress, ind)
}

// SetPermit is a paid mutator transaction binding the contract method 0xfb751427.
//
// Solidity: function setPermit(address addr, uint256 level) returns()
func (_Api *ApiTransactor) SetPermit(opts *bind.TransactOpts, addr common.Address, level *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPermit", addr, level)
}

// SetPermit is a paid mutator transaction binding the contract method 0xfb751427.
//
// Solidity: function setPermit(address addr, uint256 level) returns()
func (_Api *ApiSession) SetPermit(addr common.Address, level *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPermit(&_Api.TransactOpts, addr, level)
}

// SetPermit is a paid mutator transaction binding the contract method 0xfb751427.
//
// Solidity: function setPermit(address addr, uint256 level) returns()
func (_Api *ApiTransactorSession) SetPermit(addr common.Address, level *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPermit(&_Api.TransactOpts, addr, level)
}
