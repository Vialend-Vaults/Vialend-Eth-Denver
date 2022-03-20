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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"SWAP_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_poolFee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"swapExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_poolFee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"name\":\"swapExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610ab5380380610ab583398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610a086100ad60003960008181608701528181610249015281816102c901528181610351015281816103ec015261046c0152610a086000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80639ffccba51461005c578063c31c9c0714610082578063c6005893146100c1578063ccdc1e89146100dc578063dfd00055146100ef575b600080fd5b61006f61006a366004610820565b610102565b6040519081526020015b60405180910390f35b6100a97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610079565b6100a97368b3465833fb72a70ecdf485e0e4c7bd8665fc4581565b61006f6100ea36600461086b565b610235565b61006f6100fd366004610820565b610395565b60007368b3465833fb72a70ecdf485e0e4c7bd8665fc45610125863330866104ee565b610144867368b3465833fb72a70ecdf485e0e4c7bd8665fc45856105f8565b6040805160e0810182526001600160a01b0388811682528781166020830190815262ffffff8881168486019081523360608601908152608086018a8152600060a0880181815260c0890191825298516304e45aaf60e01b815288518816600482015295518716602487015292519093166044850152518416606484015290516084830152935160a48201529251811660c48401529091908316906304e45aaf9060e4016020604051808303816000875af1158015610206573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022a91906108c0565b979650505050505050565b6000610243863330856104ee565b61026e867f0000000000000000000000000000000000000000000000000000000000000000846105f8565b60408051610100810182526001600160a01b038089168252878116602083015262ffffff87168284015233606083015242608083015260a0820186905260c08201859052600060e08301529151631b67c43360e31b815290917f0000000000000000000000000000000000000000000000000000000000000000169063db3e2198906102fe9084906004016108d9565b6020604051808303816000875af115801561031d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034191906108c0565b91508282101561038b57610377877f000000000000000000000000000000000000000000000000000000000000000060006105f8565b61038b87336103868587610949565b6106f8565b5095945050505050565b6000336103da5760405162461bcd60e51b815260206004820152600e60248201526d06d73672e73656e6465723d3078360941b60448201526064015b60405180910390fd5b6103e6853330856104ee565b610411857f0000000000000000000000000000000000000000000000000000000000000000846105f8565b60408051610100810182526001600160a01b038088168252868116602083015262ffffff86168284015233606083015242608083015260a08201859052600060c0830181905260e0830152915163414bf38960e01b815290917f0000000000000000000000000000000000000000000000000000000000000000169063414bf389906104a19084906004016108d9565b6020604051808303816000875af11580156104c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e491906108c0565b9695505050505050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b1790529151600092839290881691610552919061096e565b6000604051808303816000865af19150503d806000811461058f576040519150601f19603f3d011682016040523d82523d6000602084013e610594565b606091505b50915091508180156105be5750805115806105be5750808060200190518101906105be91906109a9565b6105f05760405162461bcd60e51b815260206004820152600360248201526229aa2360e91b60448201526064016103d1565b505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663095ea7b360e01b1790529151600092839290871691610654919061096e565b6000604051808303816000865af19150503d8060008114610691576040519150601f19603f3d011682016040523d82523d6000602084013e610696565b606091505b50915091508180156106c05750805115806106c05750808060200190518101906106c091906109a9565b6106f15760405162461bcd60e51b8152602060048201526002602482015261534160f01b60448201526064016103d1565b5050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151600092839290871691610754919061096e565b6000604051808303816000865af19150503d8060008114610791576040519150601f19603f3d011682016040523d82523d6000602084013e610796565b606091505b50915091508180156107c05750805115806107c05750808060200190518101906107c091906109a9565b6106f15760405162461bcd60e51b815260206004820152600260248201526114d560f21b60448201526064016103d1565b80356001600160a01b038116811461080857600080fd5b919050565b803562ffffff8116811461080857600080fd5b6000806000806080858703121561083657600080fd5b61083f856107f1565b935061084d602086016107f1565b925061085b6040860161080d565b9396929550929360600135925050565b600080600080600060a0868803121561088357600080fd5b61088c866107f1565b945061089a602087016107f1565b93506108a86040870161080d565b94979396509394606081013594506080013592915050565b6000602082840312156108d257600080fd5b5051919050565b6101008101610943828480516001600160a01b03908116835260208083015182169084015260408083015162ffffff16908401526060808301518216908401526080808301519084015260a0828101519084015260c0808301519084015260e09182015116910152565b92915050565b60008282101561096957634e487b7160e01b600052601160045260246000fd5b500390565b6000825160005b8181101561098f5760208186018101518583015201610975565b8181111561099e576000828501525b509190910192915050565b6000602082840312156109bb57600080fd5b815180151581146109cb57600080fd5b939250505056fea264697066735822122030fa431592be4a122fb74dd8cc14e9a5e33f70dd6190e65925587b2ab20838c964736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _swapRouter common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _swapRouter)
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

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_Api *ApiCaller) SWAPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "SWAP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_Api *ApiSession) SWAPROUTER() (common.Address, error) {
	return _Api.Contract.SWAPROUTER(&_Api.CallOpts)
}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_Api *ApiCallerSession) SWAPROUTER() (common.Address, error) {
	return _Api.Contract.SWAPROUTER(&_Api.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Api *ApiCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Api *ApiSession) SwapRouter() (common.Address, error) {
	return _Api.Contract.SwapRouter(&_Api.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Api *ApiCallerSession) SwapRouter() (common.Address, error) {
	return _Api.Contract.SwapRouter(&_Api.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x9ffccba5.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, uint24 _fee, uint256 _amountIn) returns(uint256)
func (_Api *ApiTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _fee *big.Int, _amountIn *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _fee, _amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0x9ffccba5.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, uint24 _fee, uint256 _amountIn) returns(uint256)
func (_Api *ApiSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _fee *big.Int, _amountIn *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, _tokenIn, _tokenOut, _fee, _amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0x9ffccba5.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, uint24 _fee, uint256 _amountIn) returns(uint256)
func (_Api *ApiTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _fee *big.Int, _amountIn *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Swap(&_Api.TransactOpts, _tokenIn, _tokenOut, _fee, _amountIn)
}

// SwapExactInputSingle is a paid mutator transaction binding the contract method 0xdfd00055.
//
// Solidity: function swapExactInputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountIn) returns(uint256 amountOut)
func (_Api *ApiTransactor) SwapExactInputSingle(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapExactInputSingle", _tokenIn, _tokenOut, _poolFee, amountIn)
}

// SwapExactInputSingle is a paid mutator transaction binding the contract method 0xdfd00055.
//
// Solidity: function swapExactInputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountIn) returns(uint256 amountOut)
func (_Api *ApiSession) SwapExactInputSingle(_tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExactInputSingle(&_Api.TransactOpts, _tokenIn, _tokenOut, _poolFee, amountIn)
}

// SwapExactInputSingle is a paid mutator transaction binding the contract method 0xdfd00055.
//
// Solidity: function swapExactInputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountIn) returns(uint256 amountOut)
func (_Api *ApiTransactorSession) SwapExactInputSingle(_tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExactInputSingle(&_Api.TransactOpts, _tokenIn, _tokenOut, _poolFee, amountIn)
}

// SwapExactOutputSingle is a paid mutator transaction binding the contract method 0xccdc1e89.
//
// Solidity: function swapExactOutputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountOut, uint256 amountInMaximum) returns(uint256 amountIn)
func (_Api *ApiTransactor) SwapExactOutputSingle(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountOut *big.Int, amountInMaximum *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "swapExactOutputSingle", _tokenIn, _tokenOut, _poolFee, amountOut, amountInMaximum)
}

// SwapExactOutputSingle is a paid mutator transaction binding the contract method 0xccdc1e89.
//
// Solidity: function swapExactOutputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountOut, uint256 amountInMaximum) returns(uint256 amountIn)
func (_Api *ApiSession) SwapExactOutputSingle(_tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountOut *big.Int, amountInMaximum *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExactOutputSingle(&_Api.TransactOpts, _tokenIn, _tokenOut, _poolFee, amountOut, amountInMaximum)
}

// SwapExactOutputSingle is a paid mutator transaction binding the contract method 0xccdc1e89.
//
// Solidity: function swapExactOutputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountOut, uint256 amountInMaximum) returns(uint256 amountIn)
func (_Api *ApiTransactorSession) SwapExactOutputSingle(_tokenIn common.Address, _tokenOut common.Address, _poolFee *big.Int, amountOut *big.Int, amountInMaximum *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SwapExactOutputSingle(&_Api.TransactOpts, _tokenIn, _tokenOut, _poolFee, amountOut, amountInMaximum)
}
