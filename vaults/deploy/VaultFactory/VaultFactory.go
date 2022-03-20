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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_this\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"CreatedVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stat\",\"type\":\"uint256\"}],\"name\":\"changeStat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[10]\",\"name\":\"contracts\",\"type\":\"address[10]\"},{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_individualCap\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_compPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_protocolFee\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"_feetier\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"_quoteAmount\",\"type\":\"uint128\"}],\"name\":\"createVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPair0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sORv\",\"type\":\"address\"}],\"name\":\"getStat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getStat2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTeam\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b7c380380610b7c83398101604081905261002f91610062565b60008054336001600160a01b031991821617909155600180549091166001600160a01b0392909216919091179055610092565b60006020828403121561007457600080fd5b81516001600160a01b038116811461008b57600080fd5b9392505050565b610adb806100a16000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637afd53c41161008c578063a87d942c11610066578063a87d942c1461024c578063aa67735414610254578063b76a245a14610267578063fe33b302146102be57600080fd5b80637afd53c4146101dd5780638bce6edd146102085780638c64ea4a1461021957600080fd5b806324f0c771116100c857806324f0c7711461013c5780635eebb98614610175578063666358c8146101885780636e9960c3146101cc57600080fd5b806306e7cf18146100ef578063095cf5c6146101045780631627905514610117575b600080fd5b6101026100fd366004610803565b6102e7565b005b610102610112366004610844565b61036e565b610129610125366004610844565b3b90565b6040519081526020015b60405180910390f35b61012961014a366004610868565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6101026101833660046108f2565b6103a7565b6101b4610196366004610844565b6001600160a01b039081166000908152600360205260409020541690565b6040516001600160a01b039091168152602001610133565b6000546001600160a01b03166101b4565b6101296101eb366004610868565b600260209081526000928352604080842090915290825290205481565b6001546001600160a01b03166101b4565b61022c6102273660046109df565b61064f565b604080516001600160a01b03938416815292909116602083015201610133565b600454610129565b610102610262366004610868565b610688565b610129610275366004610844565b6001600160a01b03908116600081815260036020908152604080832054909416808352600280835285842094845293825284832054938252848320908352905291909120541790565b6101b46102cc366004610844565b6003602052600090815260409020546001600160a01b031681565b6000546001600160a01b031633146102fe57600080fd5b600081116103425760405162461bcd60e51b815260206004820152600c60248201526b1a5b9d985b1a59081cdd185d60a21b60448201526064015b60405180910390fd5b6001600160a01b0392831660009081526002602090815260408083209490951682529290925291902055565b6000546001600160a01b0316331461038557600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6101208801516000805460c08b015160e08c0151604051632a3c85d760e01b81523060048201526001600160a01b039384166024820152918316604483015282166064820152608481018b905260a481018a905291921690632a3c85d79060c4016020604051808303816000875af1158015610427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044b9190610a0e565b90506001600160a01b0381166104885760405162461bcd60e51b8152602060048201526002602482015261076360f41b6044820152606401610339565b60006040518061016001604052808b6000600a81106104a9576104a96109f8565b602090810291909101516001600160a01b039081168352339183019190915260408d8101518216818401526060808f01518316908401526080808f015183169084015260a0808f015183169084015260c0808f015183169084015260e0808f01518316908401528582166101008085019190915260008054841661012086015230610140909501949094528e01519051630873208160e11b815293945091929116906310e64102906105699085908c908c908c908c908c90600401610a2b565b6020604051808303816000875af1158015610588573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ac9190610a0e565b90506001600160a01b0381166105e95760405162461bcd60e51b8152602060048201526002602482015261073360f41b6044820152606401610339565b6105f38184610700565b604080513081523360208201526001600160a01b03838116828401528516606082015290517f341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a79181900360800190a15050505050505050505050565b6004818154811061065f57600080fd5b6000918252602090912060029091020180546001909101546001600160a01b0391821692501682565b6000546001600160a01b0316331461069f57600080fd5b6001600160a01b03818116600090815260036020526040902054811690831614156106f25760405162461bcd60e51b815260206004820152600360248201526206475760ec1b6044820152606401610339565b6106fc8282610700565b5050565b6001600160a01b0391821660008181526002602081815260408084209587168085529582528084208390558484526003825280842080546001600160a01b03199081168817909155868552818520805482168717905581518083019092529481529081019485526004805460018101825593525191027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b810180549286169284169290921790915591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c9092018054929093169116179055565b6001600160a01b03811681146107f057600080fd5b50565b80356107fe816107db565b919050565b60008060006060848603121561081857600080fd5b8335610823816107db565b92506020840135610833816107db565b929592945050506040919091013590565b60006020828403121561085657600080fd5b8135610861816107db565b9392505050565b6000806040838503121561087b57600080fd5b8235610886816107db565b91506020830135610896816107db565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b803560ff811681146107fe57600080fd5b803562ffffff811681146107fe57600080fd5b80356001600160801b03811681146107fe57600080fd5b600080600080600080600080610220898b03121561090f57600080fd5b89601f8a011261091e57600080fd5b60405161014080820182811067ffffffffffffffff82111715610943576109436108a1565b6040528a01818c82111561095657600080fd5b8b5b8281101561097757610969816107f3565b825260209182019101610958565b509199503597505061016089013595506109946101808a016108b7565b94506109a36101a08a016108b7565b93506109b26101c08a016108b7565b92506109c16101e08a016108c8565b91506109d06102008a016108db565b90509295985092959890939650565b6000602082840312156109f157600080fd5b5035919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610a2057600080fd5b8151610861816107db565b6102008101818860005b600b811015610a5d5781516001600160a01b0316835260209283019290910190600101610a35565b50505060ff871661016083015260ff861661018083015260ff85166101a083015262ffffff84166101c08301526001600160801b0383166101e083015297965050505050505056fea2646970667358221220e7e47e58d4a128e764d08e75c907c040f44725f2fa90fccf12acd6285057b81264736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _team common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _team)
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

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Api *ApiCallerSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiSession) GetCount() (*big.Int, error) {
	return _Api.Contract.GetCount(&_Api.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Api *ApiCallerSession) GetCount() (*big.Int, error) {
	return _Api.Contract.GetCount(&_Api.CallOpts)
}

// GetPair0 is a free data retrieval call binding the contract method 0x666358c8.
//
// Solidity: function getPair0(address _addr) view returns(address)
func (_Api *ApiCaller) GetPair0(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPair0", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair0 is a free data retrieval call binding the contract method 0x666358c8.
//
// Solidity: function getPair0(address _addr) view returns(address)
func (_Api *ApiSession) GetPair0(_addr common.Address) (common.Address, error) {
	return _Api.Contract.GetPair0(&_Api.CallOpts, _addr)
}

// GetPair0 is a free data retrieval call binding the contract method 0x666358c8.
//
// Solidity: function getPair0(address _addr) view returns(address)
func (_Api *ApiCallerSession) GetPair0(_addr common.Address) (common.Address, error) {
	return _Api.Contract.GetPair0(&_Api.CallOpts, _addr)
}

// GetStat is a free data retrieval call binding the contract method 0xb76a245a.
//
// Solidity: function getStat(address sORv) view returns(uint256)
func (_Api *ApiCaller) GetStat(opts *bind.CallOpts, sORv common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStat", sORv)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStat is a free data retrieval call binding the contract method 0xb76a245a.
//
// Solidity: function getStat(address sORv) view returns(uint256)
func (_Api *ApiSession) GetStat(sORv common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat(&_Api.CallOpts, sORv)
}

// GetStat is a free data retrieval call binding the contract method 0xb76a245a.
//
// Solidity: function getStat(address sORv) view returns(uint256)
func (_Api *ApiCallerSession) GetStat(sORv common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat(&_Api.CallOpts, sORv)
}

// GetStat2 is a free data retrieval call binding the contract method 0x24f0c771.
//
// Solidity: function getStat2(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiCaller) GetStat2(opts *bind.CallOpts, _strategy common.Address, _vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStat2", _strategy, _vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStat2 is a free data retrieval call binding the contract method 0x24f0c771.
//
// Solidity: function getStat2(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiSession) GetStat2(_strategy common.Address, _vault common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat2(&_Api.CallOpts, _strategy, _vault)
}

// GetStat2 is a free data retrieval call binding the contract method 0x24f0c771.
//
// Solidity: function getStat2(address _strategy, address _vault) view returns(uint256)
func (_Api *ApiCallerSession) GetStat2(_strategy common.Address, _vault common.Address) (*big.Int, error) {
	return _Api.Contract.GetStat2(&_Api.CallOpts, _strategy, _vault)
}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiCaller) GetTeam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTeam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiSession) GetTeam() (common.Address, error) {
	return _Api.Contract.GetTeam(&_Api.CallOpts)
}

// GetTeam is a free data retrieval call binding the contract method 0x8bce6edd.
//
// Solidity: function getTeam() view returns(address)
func (_Api *ApiCallerSession) GetTeam() (common.Address, error) {
	return _Api.Contract.GetTeam(&_Api.CallOpts)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiCaller) IsContract(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isContract", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiSession) IsContract(addr common.Address) (*big.Int, error) {
	return _Api.Contract.IsContract(&_Api.CallOpts, addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(uint256)
func (_Api *ApiCallerSession) IsContract(addr common.Address) (*big.Int, error) {
	return _Api.Contract.IsContract(&_Api.CallOpts, addr)
}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_Api *ApiCaller) Pairs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_Api *ApiSession) Pairs(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.Pairs(&_Api.CallOpts, arg0)
}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_Api *ApiCallerSession) Pairs(arg0 common.Address) (common.Address, error) {
	return _Api.Contract.Pairs(&_Api.CallOpts, arg0)
}

// Stat is a free data retrieval call binding the contract method 0x7afd53c4.
//
// Solidity: function stat(address , address ) view returns(uint256)
func (_Api *ApiCaller) Stat(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "stat", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stat is a free data retrieval call binding the contract method 0x7afd53c4.
//
// Solidity: function stat(address , address ) view returns(uint256)
func (_Api *ApiSession) Stat(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Api.Contract.Stat(&_Api.CallOpts, arg0, arg1)
}

// Stat is a free data retrieval call binding the contract method 0x7afd53c4.
//
// Solidity: function stat(address , address ) view returns(uint256)
func (_Api *ApiCallerSession) Stat(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Api.Contract.Stat(&_Api.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Strategy common.Address
		Vault    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Vault = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiSession) Vaults(arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address strategy, address vault)
func (_Api *ApiCallerSession) Vaults(arg0 *big.Int) (struct {
	Strategy common.Address
	Vault    common.Address
}, error) {
	return _Api.Contract.Vaults(&_Api.CallOpts, arg0)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiTransactor) ChangeStat(opts *bind.TransactOpts, _strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "changeStat", _strategy, _vault, _stat)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiSession) ChangeStat(_strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ChangeStat(&_Api.TransactOpts, _strategy, _vault, _stat)
}

// ChangeStat is a paid mutator transaction binding the contract method 0x06e7cf18.
//
// Solidity: function changeStat(address _strategy, address _vault, uint256 _stat) returns()
func (_Api *ApiTransactorSession) ChangeStat(_strategy common.Address, _vault common.Address, _stat *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ChangeStat(&_Api.TransactOpts, _strategy, _vault, _stat)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiTransactor) CreateVault(opts *bind.TransactOpts, contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createVault", contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiSession) CreateVault(contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateVault(&_Api.TransactOpts, contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5eebb986.
//
// Solidity: function createVault(address[10] contracts, uint256 _vaultCap, uint256 _individualCap, uint8 _uniPortion, uint8 _compPortion, uint8 _protocolFee, uint24 _feetier, uint128 _quoteAmount) returns()
func (_Api *ApiTransactorSession) CreateVault(contracts [10]common.Address, _vaultCap *big.Int, _individualCap *big.Int, _uniPortion uint8, _compPortion uint8, _protocolFee uint8, _feetier *big.Int, _quoteAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateVault(&_Api.TransactOpts, contracts, _vaultCap, _individualCap, _uniPortion, _compPortion, _protocolFee, _feetier, _quoteAmount)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _strategy, address _vault) returns()
func (_Api *ApiTransactor) Register(opts *bind.TransactOpts, _strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "register", _strategy, _vault)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _strategy, address _vault) returns()
func (_Api *ApiSession) Register(_strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.Contract.Register(&_Api.TransactOpts, _strategy, _vault)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _strategy, address _vault) returns()
func (_Api *ApiTransactorSession) Register(_strategy common.Address, _vault common.Address) (*types.Transaction, error) {
	return _Api.Contract.Register(&_Api.TransactOpts, _strategy, _vault)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiTransactor) SetTeam(opts *bind.TransactOpts, _team common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setTeam", _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiSession) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetTeam(&_Api.TransactOpts, _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_Api *ApiTransactorSession) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetTeam(&_Api.TransactOpts, _team)
}

// ApiCreatedVaultIterator is returned from FilterCreatedVault and is used to iterate over the raw logs and unpacked data for CreatedVault events raised by the Api contract.
type ApiCreatedVaultIterator struct {
	Event *ApiCreatedVault // Event containing the contract specifics and raw log

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
func (it *ApiCreatedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCreatedVault)
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
		it.Event = new(ApiCreatedVault)
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
func (it *ApiCreatedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCreatedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCreatedVault represents a CreatedVault event raised by the Api contract.
type ApiCreatedVault struct {
	This     common.Address
	Sender   common.Address
	Strategy common.Address
	Vault    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatedVault is a free log retrieval operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) FilterCreatedVault(opts *bind.FilterOpts) (*ApiCreatedVaultIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CreatedVault")
	if err != nil {
		return nil, err
	}
	return &ApiCreatedVaultIterator{contract: _Api.contract, event: "CreatedVault", logs: logs, sub: sub}, nil
}

// WatchCreatedVault is a free log subscription operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) WatchCreatedVault(opts *bind.WatchOpts, sink chan<- *ApiCreatedVault) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CreatedVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCreatedVault)
				if err := _Api.contract.UnpackLog(event, "CreatedVault", log); err != nil {
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

// ParseCreatedVault is a log parse operation binding the contract event 0x341f822f1e2a166284695f5b30ad55eb8e9932e740b7eed3ac9a1871f24a19a7.
//
// Solidity: event CreatedVault(address _this, address sender, address strategy, address vault)
func (_Api *ApiFilterer) ParseCreatedVault(log types.Log) (*ApiCreatedVault, error) {
	event := new(ApiCreatedVault)
	if err := _Api.contract.UnpackLog(event, "CreatedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
