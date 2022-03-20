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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"WETHAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"redeemType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"}],\"name\":\"redeemCErc20Tokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"redeemType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_cEtherContract\",\"type\":\"address\"}],\"name\":\"redeemCEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405234801561001057600080fd5b506040516112153803806112158339818101604052602081101561003357600080fd5b5051600180546001600160a01b031916331790556001600160a01b0381166100a2576040805162461bcd60e51b815260206004820152601860248201527f5745544820697320746865207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546001600160a01b039092166001600160a01b0319909216919091179055611143806100d26000396000f3fe6080604052600436106100955760003560e01c8063c311d04911610059578063c311d049146101b4578063c6fb5014146101de578063d0ace48e146101de578063d46eb1191461021f578063de0e9a3e146102275761009c565b80633b7ba25f1461009e5780635d29dab4146100f35780636e947298146101345780638da5cb5b14610149578063a82ffd881461017a5761009c565b3661009c57005b005b3480156100aa57600080fd5b506100e1600480360360608110156100c157600080fd5b506001600160a01b03813581169160208101359091169060400135610251565b60408051918252519081900360200190f35b3480156100ff57600080fd5b5061009c6004803603606081101561011657600080fd5b508035906001600160a01b03602082013581169160400135166105a0565b34801561014057600080fd5b506100e16106e7565b34801561015557600080fd5b5061015e6106eb565b604080516001600160a01b039092168252519081900360200190f35b6101a06004803603602081101561019057600080fd5b50356001600160a01b03166106fa565b604080519115158252519081900360200190f35b3480156101c057600080fd5b5061009c600480360360208110156101d757600080fd5b50356108d3565b3480156101ea57600080fd5b506101a06004803603606081101561020157600080fd5b508035906020810135151590604001356001600160a01b03166109e6565b61009c610b30565b34801561023357600080fd5b5061009c6004803603602081101561024a57600080fd5b5035610c6b565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156102a057600080fd5b505afa1580156102b4573d6000803e3d6000fd5b505050506040513d60208110156102ca57600080fd5b505182111561030a576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561035157600080fd5b505af1158015610365573d6000803e3d6000fd5b505050506040513d602081101561037b57600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a200000000000606082015290519192506000805160206110c4833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561041257600080fd5b505af1158015610426573d6000803e3d6000fd5b505050506040513d602081101561043c57600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c6564207570290000000000000000606082015290519192506000805160206110c4833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156104ed57600080fd5b505af1158015610501573d6000803e3d6000fd5b505050506040513d602081101561051757600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b15801561056557600080fd5b505af1158015610579573d6000803e3d6000fd5b505050506040513d602081101561058f57600080fd5b5051955050505050505b9392505050565b6001546001600160a01b031633146105e95760405162461bcd60e51b815260040180806020018281038252602e815260200180611029602e913960400191505060405180910390fd5b60008311610627576040805162461bcd60e51b8152602060048201526006602482015265185b5bdd5b9d60d21b604482015290519081900360640190fd5b6001600160a01b03811630148015906106525750816001600160a01b0316816001600160a01b031614155b610688576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b61069c6001600160a01b0383168285610d0e565b6040805160208101859052818152600f818301526e2bb4ba34323930bb9022b93199181d60891b606082015290516000805160206110c48339815191529181900360800190a1505050565b4790565b6001546001600160a01b031681565b6000808290506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561073d57600080fd5b505af1158015610751573d6000803e3d6000fd5b505050506040513d602081101561076757600080fd5b5051604080516020810183905281815260239181018290529192506000805160206110c48339815191529183918190606082019061107b82396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156107ec57600080fd5b505af1158015610800573d6000803e3d6000fd5b505050506040513d602081101561081657600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c6564207570206279203165313829606082015290519192506000805160206110c4833981519152919081900360800190a1826001600160a01b0316631249c58b6203d090346040518363ffffffff1660e01b81526004016000604051808303818589803b1580156108ae57600080fd5b5088f11580156108c2573d6000803e3d6000fd5b5060019a9950505050505050505050565b6001546001600160a01b0316331461091c5760405162461bcd60e51b815260040180806020018281038252602e815260200180611029602e913960400191505060405180910390fd5b6109246106e7565b811115610961576040805162461bcd60e51b8152602060048201526006602482015265185b5bdd5b9d60d21b604482015290519081900360640190fd5b604051339082156108fc029083906000818181858888f1935050505015801561098e573d6000803e3d6000fd5b5060408051602081018390528181526017818301527f5769746864726177457468206d73672e73656e6465723a000000000000000000606082015290516000805160206110c48339815191529181900360800190a150565b6000818160018515151415610a6e57816001600160a01b031663db006a75876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610a3b57600080fd5b505af1158015610a4f573d6000803e3d6000fd5b505050506040513d6020811015610a6557600080fd5b50519050610ae3565b816001600160a01b031663852a12e3876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610ab457600080fd5b505af1158015610ac8573d6000803e3d6000fd5b505050506040513d6020811015610ade57600080fd5b505190505b6000805160206110c483398151915281604051808060200183815260200182810382526024815260200180611057602491396040019250505060405180910390a150600195945050505050565b348015610b9f5760008054906101000a90046001600160a01b03166001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610b8557600080fd5b505af1158015610b99573d6000803e3d6000fd5b50505050505b60008054604080516370a0823160e01b8152306004820152905184936001600160a01b03909316926370a0823192602480820193602093909283900390910190829087803b158015610bf057600080fd5b505af1158015610c04573d6000803e3d6000fd5b505050506040513d6020811015610c1a57600080fd5b50511015610c68576040805162461bcd60e51b8152602060048201526016602482015275115d1a195c995d5b481b9bdd0819195c1bdcda5d195960521b604482015290519081900360640190fd5b50565b338115610d0a576000805460408051632e1a7d4d60e01b81526004810186905290516001600160a01b0390921692632e1a7d4d9260248084019382900301818387803b158015610cba57600080fd5b505af1158015610cce573d6000803e3d6000fd5b50506040516001600160a01b03841692504780156108fc029250906000818181858888f19350505050158015610d08573d6000803e3d6000fd5b505b5050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610d089084906000610db0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610e0c9092919063ffffffff16565b805190915015610d0857808060200190516020811015610dcf57600080fd5b5051610d085760405162461bcd60e51b815260040180806020018281038252602a8152602001806110e4602a913960400191505060405180910390fd5b6060610e1b8484600085610e23565b949350505050565b606082471015610e645760405162461bcd60e51b815260040180806020018281038252602681526020018061109e6026913960400191505060405180910390fd5b610e6d85610f7e565b610ebe576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b60208310610efc5780518252601f199092019160209182019101610edd565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610f5e576040519150601f19603f3d011682016040523d82523d6000602084013e610f63565b606091505b5091509150610f73828286610f84565b979650505050505050565b3b151590565b60608315610f93575081610599565b825115610fa35782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610fed578181015183820152602001610fd5565b50505050905090810190601f16801561101a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe546869732063616e206f6e6c792062652063616c6c65642062792074686520636f6e7472616374206f776e65722149662074686973206973206e6f7420302c2074686572652077617320616e206572726f7245786368616e6765205261746520287363616c65642075702062792031653138293a20416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212204f2b15a22b7e532072fb587400f8cf970a16d89f077cd751486979c343cacdf264736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, WETHAddr common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, WETHAddr)
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

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiCaller) GetETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiSession) GetETHBalance() (*big.Int, error) {
	return _Api.Contract.GetETHBalance(&_Api.CallOpts)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Api *ApiCallerSession) GetETHBalance() (*big.Int, error) {
	return _Api.Contract.GetETHBalance(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiTransactor) RedeemCErc20Tokens(opts *bind.TransactOpts, amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCErc20Tokens", amount, redeemType, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiSession) RedeemCErc20Tokens(amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, redeemType, _cErc20Contract)
}

// RedeemCErc20Tokens is a paid mutator transaction binding the contract method 0xd0ace48e.
//
// Solidity: function redeemCErc20Tokens(uint256 amount, bool redeemType, address _cErc20Contract) returns(bool)
func (_Api *ApiTransactorSession) RedeemCErc20Tokens(amount *big.Int, redeemType bool, _cErc20Contract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCErc20Tokens(&_Api.TransactOpts, amount, redeemType, _cErc20Contract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiTransactor) RedeemCEth(opts *bind.TransactOpts, amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemCEth", amount, redeemType, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiSession) RedeemCEth(amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, redeemType, _cEtherContract)
}

// RedeemCEth is a paid mutator transaction binding the contract method 0xc6fb5014.
//
// Solidity: function redeemCEth(uint256 amount, bool redeemType, address _cEtherContract) returns(bool)
func (_Api *ApiTransactorSession) RedeemCEth(amount *big.Int, redeemType bool, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.RedeemCEth(&_Api.TransactOpts, amount, redeemType, _cEtherContract)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiTransactor) SupplyErc20ToCompound(opts *bind.TransactOpts, _erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "supplyErc20ToCompound", _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiSession) SupplyErc20ToCompound(_erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SupplyErc20ToCompound(&_Api.TransactOpts, _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyErc20ToCompound is a paid mutator transaction binding the contract method 0x3b7ba25f.
//
// Solidity: function supplyErc20ToCompound(address _erc20Contract, address _cErc20Contract, uint256 _numTokensToSupply) returns(uint256)
func (_Api *ApiTransactorSession) SupplyErc20ToCompound(_erc20Contract common.Address, _cErc20Contract common.Address, _numTokensToSupply *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SupplyErc20ToCompound(&_Api.TransactOpts, _erc20Contract, _cErc20Contract, _numTokensToSupply)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0xa82ffd88.
//
// Solidity: function supplyEthToCompound(address _cEtherContract) payable returns(bool)
func (_Api *ApiTransactor) SupplyEthToCompound(opts *bind.TransactOpts, _cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "supplyEthToCompound", _cEtherContract)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0xa82ffd88.
//
// Solidity: function supplyEthToCompound(address _cEtherContract) payable returns(bool)
func (_Api *ApiSession) SupplyEthToCompound(_cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0xa82ffd88.
//
// Solidity: function supplyEthToCompound(address _cEtherContract) payable returns(bool)
func (_Api *ApiTransactorSession) SupplyEthToCompound(_cEtherContract common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 Amount) returns()
func (_Api *ApiTransactor) Unwrap(opts *bind.TransactOpts, Amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "unwrap", Amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 Amount) returns()
func (_Api *ApiSession) Unwrap(Amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Unwrap(&_Api.TransactOpts, Amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 Amount) returns()
func (_Api *ApiTransactorSession) Unwrap(Amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Unwrap(&_Api.TransactOpts, Amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawERC20", amount, erc20, to)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiSession) WithdrawERC20(amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawERC20(&_Api.TransactOpts, amount, erc20, to)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x5d29dab4.
//
// Solidity: function withdrawERC20(uint256 amount, address erc20, address to) returns()
func (_Api *ApiTransactorSession) WithdrawERC20(amount *big.Int, erc20 common.Address, to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawERC20(&_Api.TransactOpts, amount, erc20, to)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiTransactor) WithdrawEth(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawEth", amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiSession) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.WithdrawEth(&_Api.TransactOpts, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns()
func (_Api *ApiTransactorSession) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.WithdrawEth(&_Api.TransactOpts, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_Api *ApiTransactor) Wrap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "wrap")
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_Api *ApiSession) Wrap() (*types.Transaction, error) {
	return _Api.Contract.Wrap(&_Api.TransactOpts)
}

// Wrap is a paid mutator transaction binding the contract method 0xd46eb119.
//
// Solidity: function wrap() payable returns()
func (_Api *ApiTransactorSession) Wrap() (*types.Transaction, error) {
	return _Api.Contract.Wrap(&_Api.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactorSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// ApiMyLogIterator is returned from FilterMyLog and is used to iterate over the raw logs and unpacked data for MyLog events raised by the Api contract.
type ApiMyLogIterator struct {
	Event *ApiMyLog // Event containing the contract specifics and raw log

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
func (it *ApiMyLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog)
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
		it.Event = new(ApiMyLog)
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
func (it *ApiMyLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog represents a MyLog event raised by the Api contract.
type ApiMyLog struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog is a free log retrieval operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) FilterMyLog(opts *bind.FilterOpts) (*ApiMyLogIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog")
	if err != nil {
		return nil, err
	}
	return &ApiMyLogIterator{contract: _Api.contract, event: "MyLog", logs: logs, sub: sub}, nil
}

// WatchMyLog is a free log subscription operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) WatchMyLog(opts *bind.WatchOpts, sink chan<- *ApiMyLog) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog)
				if err := _Api.contract.UnpackLog(event, "MyLog", log); err != nil {
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

// ParseMyLog is a log parse operation binding the contract event 0x8d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad.
//
// Solidity: event MyLog(string arg0, uint256 arg1)
func (_Api *ApiFilterer) ParseMyLog(log types.Log) (*ApiMyLog, error) {
	event := new(ApiMyLog)
	if err := _Api.contract.UnpackLog(event, "MyLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
