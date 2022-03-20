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
const ApiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b506103c2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635a9b0b891461003b5780638262963b1461005a575b600080fd5b61004361006f565b60405161005192919061024a565b60405180910390f35b61006d610068366004610282565b61010b565b005b606060008060015481805461008390610337565b80601f01602080910402602001604051908101604052809291908181526020018280546100af90610337565b80156100fc5780601f106100d1576101008083540402835291602001916100fc565b820191906000526020600020905b8154815290600101906020018083116100df57829003601f168201915b50505050509150915091509091565b815161011e906000906020850190610164565b506001819055604051819033907f4e3358e5bdb38f499a83fbcd33d16343569f3b6fe53021022dede6364017b4a890610158908690610372565b60405180910390a35050565b82805461017090610337565b90600052602060002090601f01602090048101928261019257600085556101d8565b82601f106101ab57805160ff19168380011785556101d8565b828001600101855582156101d8579182015b828111156101d85782518255916020019190600101906101bd565b506101e49291506101e8565b5090565b5b808211156101e457600081556001016101e9565b6000815180845260005b8181101561022357602081850181015186830182015201610207565b81811115610235576000602083870101525b50601f01601f19169290920160200192915050565b60408152600061025d60408301856101fd565b90508260208301529392505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561029557600080fd5b823567ffffffffffffffff808211156102ad57600080fd5b818501915085601f8301126102c157600080fd5b8135818111156102d3576102d361026c565b604051601f8201601f19908116603f011681019083821181831017156102fb576102fb61026c565b8160405282815288602084870101111561031457600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b600181811c9082168061034b57607f821691505b6020821081141561036c57634e487b7160e01b600052602260045260246000fd5b50919050565b60208152600061038560208301846101fd565b939250505056fea2646970667358221220473cf7e97189b4c349bea953ff064a9bb2dfeebd28e775247790e2a8728b8aa264736f6c634300080a0033"

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

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string, uint256)
func (_Api *ApiCaller) GetInfo(opts *bind.CallOpts) (string, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getInfo")

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string, uint256)
func (_Api *ApiSession) GetInfo() (string, *big.Int, error) {
	return _Api.Contract.GetInfo(&_Api.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string, uint256)
func (_Api *ApiCallerSession) GetInfo() (string, *big.Int, error) {
	return _Api.Contract.GetInfo(&_Api.CallOpts)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(string _fName, uint256 _age) returns()
func (_Api *ApiTransactor) SetInfo(opts *bind.TransactOpts, _fName string, _age *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setInfo", _fName, _age)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(string _fName, uint256 _age) returns()
func (_Api *ApiSession) SetInfo(_fName string, _age *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetInfo(&_Api.TransactOpts, _fName, _age)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(string _fName, uint256 _age) returns()
func (_Api *ApiTransactorSession) SetInfo(_fName string, _age *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetInfo(&_Api.TransactOpts, _fName, _age)
}

// ApiLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Api contract.
type ApiLogIterator struct {
	Event *ApiLog // Event containing the contract specifics and raw log

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
func (it *ApiLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLog)
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
		it.Event = new(ApiLog)
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
func (it *ApiLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLog represents a Log event raised by the Api contract.
type ApiLog struct {
	Arg0 common.Address
	Age  *big.Int
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x4e3358e5bdb38f499a83fbcd33d16343569f3b6fe53021022dede6364017b4a8.
//
// Solidity: event Log(address indexed arg0, uint256 indexed age, string name)
func (_Api *ApiFilterer) FilterLog(opts *bind.FilterOpts, arg0 []common.Address, age []*big.Int) (*ApiLogIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var ageRule []interface{}
	for _, ageItem := range age {
		ageRule = append(ageRule, ageItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Log", arg0Rule, ageRule)
	if err != nil {
		return nil, err
	}
	return &ApiLogIterator{contract: _Api.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x4e3358e5bdb38f499a83fbcd33d16343569f3b6fe53021022dede6364017b4a8.
//
// Solidity: event Log(address indexed arg0, uint256 indexed age, string name)
func (_Api *ApiFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ApiLog, arg0 []common.Address, age []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var ageRule []interface{}
	for _, ageItem := range age {
		ageRule = append(ageRule, ageItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Log", arg0Rule, ageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLog)
				if err := _Api.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x4e3358e5bdb38f499a83fbcd33d16343569f3b6fe53021022dede6364017b4a8.
//
// Solidity: event Log(address indexed arg0, uint256 indexed age, string name)
func (_Api *ApiFilterer) ParseLog(log types.Log) (*ApiLog, error) {
	event := new(ApiLog)
	if err := _Api.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
