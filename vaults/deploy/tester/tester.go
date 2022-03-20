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
const ApiABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"amount0\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amount1\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"LP\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"callPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p1\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p4\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p5\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p6\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"LP\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"setPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b506106de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806381d01ed31161007157806381d01ed3146102545780639906d1751461025c578063c2a2747b14610264578063c3b9f21e1461026c578063e17b8b7f14610274578063f400fde41461027c576100a9565b80633203f9c8146100ae578063337055261461013c578063633ea0d0146101605780636e219667146101df5780637398ab18146101f9575b600080fd5b6100f1600480360360808110156100c457600080fd5b506001600160a01b0381358116916020810135909116906040810135600290810b9160600135900b610284565b604080516001600160801b03988916815260208101979097528681019590955292861660608601529085166080850152841660a084015290921660c082015290519081900360e00190f35b61014461048a565b604080516001600160801b039092168252519081900360200190f35b6101a36004803603608081101561017657600080fd5b506001600160a01b0381358116916020810135909116906040810135600290810b9160600135900b6104a0565b604080516001600160801b0396871681526020810195909552848101939093529084166060840152909216608082015290519081900360a00190f35b6101e76105b6565b60408051918252519081900360200190f35b6102016105bc565b604080516001600160801b03998a16815260208101989098528781019690965293871660608701529186166080860152851660a0850152841660c084015290921660e08201529051908190036101000190f35b6101e76105fa565b610144610600565b610144610616565b610144610625565b610144610634565b610144610643565b60008060008060008060008061029b8b8b8b610652565b90508b6001600160a01b0316634f1eb3d88c8c8c6000806040518663ffffffff1660e01b815260040180866001600160a01b031681526020018560020b81526020018460020b8152602001838152602001828152602001955050505050506040805180830381600087803b15801561031257600080fd5b505af1158015610326573d6000803e3d6000fd5b505050506040513d604081101561033c57600080fd5b508051602090910151600580546001600160801b039283166001600160801b031990911617905560048054928216600160801b02929091169190911781556040805163514ea4bf60e01b8152918201839052516001600160a01b038e169163514ea4bf9160248083019260a0929190829003018186803b1580156103bf57600080fd5b505afa1580156103d3573d6000803e3d6000fd5b505050506040513d60a08110156103e957600080fd5b508051602082015160408301516060840151608090940151600380546001600160801b03968716928716600160801b908102918816919091176001600160801b0319908116939093179182905560028490556001859055600080549688169690931695909517918290556004546005549287169e50939c50919a50818516995090839004841697509190048216945016915050949950949992975094509450565b600354600160801b90046001600160801b031681565b6000806000806000806104b4898989610652565b9050896001600160a01b031663514ea4bf826040518263ffffffff1660e01b81526004018082815260200191505060a06040518083038186803b1580156104fa57600080fd5b505afa15801561050e573d6000803e3d6000fd5b505050506040513d60a081101561052457600080fd5b508051602082015160408301516060840151608090940151600380546001600160801b03196001600160801b03918216600160801b948316850217811697821697909717918290556002849055600185905560008054881696821696909617958690556004805460069816979097179096559385169f929e50909c508284169b50909104909116975095505050505050565b60025481565b6000546001546002546003546004546005546001600160801b03958616969495939484841694600160801b9485900481169481851694048116921690565b60015481565b600454600160801b90046001600160801b031681565b6000546001600160801b031681565b6003546001600160801b031681565b6004546001600160801b031681565b6005546001600160801b031681565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a909301905281519101209056fea2646970667358221220808913d02c9402763ebc73b73085ed4b90c7de53187a129136329af2eb6eb30264736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend)
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

// Amount0 is a free data retrieval call binding the contract method 0x9906d175.
//
// Solidity: function amount0() view returns(uint128)
func (_Api *ApiCaller) Amount0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "amount0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount0 is a free data retrieval call binding the contract method 0x9906d175.
//
// Solidity: function amount0() view returns(uint128)
func (_Api *ApiSession) Amount0() (*big.Int, error) {
	return _Api.Contract.Amount0(&_Api.CallOpts)
}

// Amount0 is a free data retrieval call binding the contract method 0x9906d175.
//
// Solidity: function amount0() view returns(uint128)
func (_Api *ApiCallerSession) Amount0() (*big.Int, error) {
	return _Api.Contract.Amount0(&_Api.CallOpts)
}

// Amount1 is a free data retrieval call binding the contract method 0xf400fde4.
//
// Solidity: function amount1() view returns(uint128)
func (_Api *ApiCaller) Amount1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "amount1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount1 is a free data retrieval call binding the contract method 0xf400fde4.
//
// Solidity: function amount1() view returns(uint128)
func (_Api *ApiSession) Amount1() (*big.Int, error) {
	return _Api.Contract.Amount1(&_Api.CallOpts)
}

// Amount1 is a free data retrieval call binding the contract method 0xf400fde4.
//
// Solidity: function amount1() view returns(uint128)
func (_Api *ApiCallerSession) Amount1() (*big.Int, error) {
	return _Api.Contract.Amount1(&_Api.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0x7398ab18.
//
// Solidity: function getPosition() view returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_Api *ApiCaller) GetPosition(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPosition")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetPosition is a free data retrieval call binding the contract method 0x7398ab18.
//
// Solidity: function getPosition() view returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_Api *ApiSession) GetPosition() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetPosition(&_Api.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0x7398ab18.
//
// Solidity: function getPosition() view returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_Api *ApiCallerSession) GetPosition() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetPosition(&_Api.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint128)
func (_Api *ApiCaller) P1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint128)
func (_Api *ApiSession) P1() (*big.Int, error) {
	return _Api.Contract.P1(&_Api.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0xc2a2747b.
//
// Solidity: function p1() view returns(uint128)
func (_Api *ApiCallerSession) P1() (*big.Int, error) {
	return _Api.Contract.P1(&_Api.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_Api *ApiCaller) P2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_Api *ApiSession) P2() (*big.Int, error) {
	return _Api.Contract.P2(&_Api.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x81d01ed3.
//
// Solidity: function p2() view returns(uint256)
func (_Api *ApiCallerSession) P2() (*big.Int, error) {
	return _Api.Contract.P2(&_Api.CallOpts)
}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_Api *ApiCaller) P3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_Api *ApiSession) P3() (*big.Int, error) {
	return _Api.Contract.P3(&_Api.CallOpts)
}

// P3 is a free data retrieval call binding the contract method 0x6e219667.
//
// Solidity: function p3() view returns(uint256)
func (_Api *ApiCallerSession) P3() (*big.Int, error) {
	return _Api.Contract.P3(&_Api.CallOpts)
}

// P4 is a free data retrieval call binding the contract method 0xc3b9f21e.
//
// Solidity: function p4() view returns(uint128)
func (_Api *ApiCaller) P4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P4 is a free data retrieval call binding the contract method 0xc3b9f21e.
//
// Solidity: function p4() view returns(uint128)
func (_Api *ApiSession) P4() (*big.Int, error) {
	return _Api.Contract.P4(&_Api.CallOpts)
}

// P4 is a free data retrieval call binding the contract method 0xc3b9f21e.
//
// Solidity: function p4() view returns(uint128)
func (_Api *ApiCallerSession) P4() (*big.Int, error) {
	return _Api.Contract.P4(&_Api.CallOpts)
}

// P5 is a free data retrieval call binding the contract method 0x33705526.
//
// Solidity: function p5() view returns(uint128)
func (_Api *ApiCaller) P5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P5 is a free data retrieval call binding the contract method 0x33705526.
//
// Solidity: function p5() view returns(uint128)
func (_Api *ApiSession) P5() (*big.Int, error) {
	return _Api.Contract.P5(&_Api.CallOpts)
}

// P5 is a free data retrieval call binding the contract method 0x33705526.
//
// Solidity: function p5() view returns(uint128)
func (_Api *ApiCallerSession) P5() (*big.Int, error) {
	return _Api.Contract.P5(&_Api.CallOpts)
}

// P6 is a free data retrieval call binding the contract method 0xe17b8b7f.
//
// Solidity: function p6() view returns(uint128)
func (_Api *ApiCaller) P6(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "p6")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P6 is a free data retrieval call binding the contract method 0xe17b8b7f.
//
// Solidity: function p6() view returns(uint128)
func (_Api *ApiSession) P6() (*big.Int, error) {
	return _Api.Contract.P6(&_Api.CallOpts)
}

// P6 is a free data retrieval call binding the contract method 0xe17b8b7f.
//
// Solidity: function p6() view returns(uint128)
func (_Api *ApiCallerSession) P6() (*big.Int, error) {
	return _Api.Contract.P6(&_Api.CallOpts)
}

// CallPosition is a paid mutator transaction binding the contract method 0x3203f9c8.
//
// Solidity: function callPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128)
func (_Api *ApiTransactor) CallPosition(opts *bind.TransactOpts, pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "callPosition", pool, LP, tickLower, tickUpper)
}

// CallPosition is a paid mutator transaction binding the contract method 0x3203f9c8.
//
// Solidity: function callPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128)
func (_Api *ApiSession) CallPosition(pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CallPosition(&_Api.TransactOpts, pool, LP, tickLower, tickUpper)
}

// CallPosition is a paid mutator transaction binding the contract method 0x3203f9c8.
//
// Solidity: function callPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128, uint128, uint128)
func (_Api *ApiTransactorSession) CallPosition(pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CallPosition(&_Api.TransactOpts, pool, LP, tickLower, tickUpper)
}

// SetPosition is a paid mutator transaction binding the contract method 0x633ea0d0.
//
// Solidity: function setPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128)
func (_Api *ApiTransactor) SetPosition(opts *bind.TransactOpts, pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPosition", pool, LP, tickLower, tickUpper)
}

// SetPosition is a paid mutator transaction binding the contract method 0x633ea0d0.
//
// Solidity: function setPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128)
func (_Api *ApiSession) SetPosition(pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPosition(&_Api.TransactOpts, pool, LP, tickLower, tickUpper)
}

// SetPosition is a paid mutator transaction binding the contract method 0x633ea0d0.
//
// Solidity: function setPosition(address pool, address LP, int24 tickLower, int24 tickUpper) returns(uint128, uint256, uint256, uint128, uint128)
func (_Api *ApiTransactorSession) SetPosition(pool common.Address, LP common.Address, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPosition(&_Api.TransactOpts, pool, LP, tickLower, tickUpper)
}
