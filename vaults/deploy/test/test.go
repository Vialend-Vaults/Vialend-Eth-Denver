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

// UniCompHelperStratInfo is an auto generated low-level Go binding around an user-defined struct.
type UniCompHelperStratInfo struct {
	VaultCap         *big.Int
	IndividualCap    *big.Int
	QuoteAmount      *big.Int
	Decimal0         uint8
	Decimal1         uint8
	UniPortion       uint8
	CompPortion      uint8
	CreatorFee       uint8
	AdminFee         uint8
	Feetier          *big.Int
	MaxTwapDeviation *big.Int
	TwapDuration     uint32
	Threshold        uint32
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_individualCap\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_quoteAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_decimal0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_decimal1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_compPortion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_creatorFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_adminFee\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"_feetier\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_threshold\",\"type\":\"uint32\"}],\"internalType\":\"structUniCompHelper.StratInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516103bb3803806103bb83398101604081905261002f9161027e565b8060008082015181600001556020820151816001015560408201518160020160006101000a8154816001600160801b0302191690836001600160801b0316021790555060608201518160020160106101000a81548160ff021916908360ff16021790555060808201518160020160116101000a81548160ff021916908360ff16021790555060a08201518160020160126101000a81548160ff021916908360ff16021790555060c08201518160020160136101000a81548160ff021916908360ff16021790555060e08201518160020160146101000a81548160ff021916908360ff1602179055506101008201518160020160156101000a81548160ff021916908360ff1602179055506101208201518160020160166101000a81548162ffffff021916908362ffffff1602179055506101408201518160020160196101000a81548162ffffff021916908360020b62ffffff16021790555061016082015181600201601c6101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160030160006101000a81548163ffffffff021916908363ffffffff1602179055509050505061036e565b6040516101a081016001600160401b038111828210171561021257634e487b7160e01b600052604160045260246000fd5b60405290565b80516001600160801b038116811461022f57600080fd5b919050565b805160ff8116811461022f57600080fd5b805162ffffff8116811461022f57600080fd5b8051600281900b811461022f57600080fd5b805163ffffffff8116811461022f57600080fd5b60006101a0828403121561029157600080fd5b6102996101e1565b82518152602083015160208201526102b360408401610218565b60408201526102c460608401610234565b60608201526102d560808401610234565b60808201526102e660a08401610234565b60a08201526102f760c08401610234565b60c082015261030860e08401610234565b60e082015261010061031b818501610234565b9082015261012061032d848201610245565b9082015261014061033f848201610258565b9082015261016061035184820161026a565b9082015261018061036384820161026a565b908201529392505050565b603f8061037c6000396000f3fe6080604052600080fdfea2646970667358221220f3abcfdb6b4809d42cce86dfb0d2533c467f28c7d8fb670afdfc7366565cad3c64736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _info UniCompHelperStratInfo) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _info)
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
