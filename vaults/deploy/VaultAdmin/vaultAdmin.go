// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiABI is the input ABI used to generate the binding from.
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_admins\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"authAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"delAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b506040516105393803806105398339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050919091016040525050600080546001600160a01b031916331781559150505b81518110156101685760018282815181106100ec57fe5b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b0390921691909117905583519083019160029185908590811061013c57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020556001016100d5565b50506103c0806101796000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806362d9185514610046578063704b6c021461006e578063fa9b082114610094575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100ce565b005b61006c6004803603602081101561008457600080fd5b50356001600160a01b03166102ab565b6100ba600480360360208110156100aa57600080fd5b50356001600160a01b031661035c565b604080519115158252519081900360200190f35b6000546001600160a01b03163314610115576040805162461bcd60e51b815260206004820152600560248201526437bbb732b960d91b604482015290519081900360640190fd5b6001600160a01b038116600090815260026020526040902054610171576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015290519081900360640190fd5b6001600160a01b03811660009081526002602052604090205460018054600019909201918290811061019f57fe5b600091825260209091200180546001600160a01b03191690556001805460001981019190829081106101cd57fe5b600091825260209091200154600180546001600160a01b0390921691849081106101f357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600180548061022c57fe5b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905580600101600260006001858154811061026957fe5b60009182526020808320909101546001600160a01b03908116845283820194909452604092830182209490945595909116855260029091528320929092555050565b6000546001600160a01b031633146102f2576040805162461bcd60e51b815260206004820152600560248201526437bbb732b960d91b604482015290519081900360640190fd5b6102fb8161036d565b610359576001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0384169081179091559054600091825260026020526040909120555b50565b60006103678261036d565b92915050565b6001600160a01b031660009081526002602052604090205415159056fea2646970667358221220354d73301892aa9f9a4c3339eb8e12d37ef95c4784ede7ebecc8c3c0cb422b6464736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _admins []common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _admins)
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

// AuthAdmin is a free data retrieval call binding the contract method 0xfa9b0821.
//
// Solidity: function authAdmin(address _admin) view returns(bool)
func (_Api *ApiCaller) AuthAdmin(opts *bind.CallOpts, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "authAdmin", _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthAdmin is a free data retrieval call binding the contract method 0xfa9b0821.
//
// Solidity: function authAdmin(address _admin) view returns(bool)
func (_Api *ApiSession) AuthAdmin(_admin common.Address) (bool, error) {
	return _Api.Contract.AuthAdmin(&_Api.CallOpts, _admin)
}

// AuthAdmin is a free data retrieval call binding the contract method 0xfa9b0821.
//
// Solidity: function authAdmin(address _admin) view returns(bool)
func (_Api *ApiCallerSession) AuthAdmin(_admin common.Address) (bool, error) {
	return _Api.Contract.AuthAdmin(&_Api.CallOpts, _admin)
}

// DelAdmin is a paid mutator transaction binding the contract method 0x62d91855.
//
// Solidity: function delAdmin(address _addr) returns()
func (_Api *ApiTransactor) DelAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "delAdmin", _addr)
}

// DelAdmin is a paid mutator transaction binding the contract method 0x62d91855.
//
// Solidity: function delAdmin(address _addr) returns()
func (_Api *ApiSession) DelAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Api.Contract.DelAdmin(&_Api.TransactOpts, _addr)
}

// DelAdmin is a paid mutator transaction binding the contract method 0x62d91855.
//
// Solidity: function delAdmin(address _addr) returns()
func (_Api *ApiTransactorSession) DelAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Api.Contract.DelAdmin(&_Api.TransactOpts, _addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _addr) returns()
func (_Api *ApiTransactor) SetAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setAdmin", _addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _addr) returns()
func (_Api *ApiSession) SetAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAdmin(&_Api.TransactOpts, _addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _addr) returns()
func (_Api *ApiTransactorSession) SetAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAdmin(&_Api.TransactOpts, _addr)
}
