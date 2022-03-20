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

// AMetaData contains all meta data concerning the A contract.
var AMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintable\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimal0\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimal1\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_B\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimal\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"changebyA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTotalSupply\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"param0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setParam0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uintparam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f2c57866": "_B()",
		"a74ea63f": "addToken(address,uint8)",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"bd3c818f": "changebyA(string)",
		"a511eb2c": "checkTotalSupply()",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"c24d942f": "param0()",
		"e0a068e8": "setParam0(uint256)",
		"95d89b41": "symbol()",
		"4f64b2be": "tokens(uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"52d74828": "uintparam()",
	},
	Bin: "0x60a06040523480156200001157600080fd5b50604051620011f6380380620011f6833981810160405260a08110156200003757600080fd5b50805160208083015160408085015160608601516080909601518251808401845260058152642a37b5b2b760d91b81870190815284518086019095526003808652622a25a760e91b978601979097528151979895979396929491939192620000a2929091906200033b565b508051620000b89060049060208401906200033b565b50506005805460ff1916601217905550620000e533620000d7620001bc565b60ff16600a0a8702620001c5565b6040805180820182526001600160a01b03958616815260ff94851660208083019182526008805460018181018355600083815295517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3928301805496518c16600160a01b90810260ff60a01b19938f166001600160a01b0319998a16178416179091558851808a01909952998c168852978a16938701938452825490810183559190945293519390920180549251909616909402919095169490921693909317909216919091179055503360601b608052620003e7565b60055460ff1690565b6001600160a01b03821662000221576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200022f60008383620002d4565b6200024b81600254620002d960201b6200096d1790919060201c565b6002556001600160a01b038216600090815260208181526040909120546200027e9183906200096d620002d9821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000334576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003735760008555620003be565b82601f106200038e57805160ff1916838001178555620003be565b82800160010185558215620003be579182015b82811115620003be578251825591602001919060010190620003a1565b50620003cc929150620003d0565b5090565b5b80821115620003cc5760008155600101620003d1565b60805160601c610dec6200040a6000398061085b528061094b5250610dec6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806395d89b41116100ad578063bd3c818f11610071578063bd3c818f146103a9578063c24d942f1461044f578063dd62ed3e14610457578063e0a068e814610485578063f2c57866146104a257610121565b806395d89b41146102ed578063a457c2d7146102f5578063a511eb2c14610321578063a74ea63f1461034c578063a9059cbb1461037d57610121565b8063313ce567116100f4578063313ce5671461023357806339509351146102515780634f64b2be1461027d57806352d74828146102bf57806370a08231146102c757610121565b806306fdde0314610126578063095ea7b3146101a357806318160ddd146101e357806323b872dd146101fd575b600080fd5b61012e6104c6565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610168578181015183820152602001610150565b50505050905090810190601f1680156101955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101cf600480360360408110156101b957600080fd5b506001600160a01b03813516906020013561055c565b604080519115158252519081900360200190f35b6101eb610579565b60408051918252519081900360200190f35b6101cf6004803603606081101561021357600080fd5b506001600160a01b0381358116916020810135909116906040013561057f565b61023b610606565b6040805160ff9092168252519081900360200190f35b6101cf6004803603604081101561026757600080fd5b506001600160a01b03813516906020013561060f565b61029a6004803603602081101561029357600080fd5b503561065d565b604080516001600160a01b03909316835260ff90911660208301528051918290030190f35b6101eb610692565b6101eb600480360360208110156102dd57600080fd5b50356001600160a01b0316610698565b61012e6106b3565b6101cf6004803603604081101561030b57600080fd5b506001600160a01b038135169060200135610714565b61032961077c565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61037b6004803603604081101561036257600080fd5b5080356001600160a01b0316906020013560ff16610798565b005b6101cf6004803603604081101561039357600080fd5b506001600160a01b03813516906020013561081e565b61037b600480360360208110156103bf57600080fd5b8101906020810181356401000000008111156103da57600080fd5b8201836020820111156103ec57600080fd5b8035906020019184600183028401116401000000008311171561040e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610832945050505050565b6101eb610913565b6101eb6004803603604081101561046d57600080fd5b506001600160a01b0381358116916020013516610919565b61037b6004803603602081101561049b57600080fd5b5035610944565b6104aa610949565b604080516001600160a01b039092168252519081900360200190f35b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105525780601f1061052757610100808354040283529160200191610552565b820191906000526020600020905b81548152906001019060200180831161053557829003601f168201915b5050505050905090565b60006105706105696109ce565b84846109d2565b50600192915050565b60025490565b600061058c848484610abe565b6105fc846105986109ce565b6105f785604051806060016040528060288152602001610d21602891396001600160a01b038a166000908152600160205260408120906105d66109ce565b6001600160a01b031681526020810191909152604001600020549190610c19565b6109d2565b5060019392505050565b60055460ff1690565b600061057061061c6109ce565b846105f7856001600061062d6109ce565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061096d565b6008818154811061066d57600080fd5b6000918252602090912001546001600160a01b0381169150600160a01b900460ff1682565b60075481565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105525780601f1061052757610100808354040283529160200191610552565b60006105706107216109ce565b846105f785604051806060016040528060258152602001610d92602591396001600061074b6109ce565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610c19565b6000806000610789610579565b60078190553393509150509091565b604080518082019091526001600160a01b03928316815260ff918216602082019081526008805460018101825560009190915291517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180549151909316600160a01b0260ff60a01b19929094166001600160a01b03199091161716919091179055565b600061057061082b6109ce565b8484610abe565b60405163bd3c818f60e01b81526020600482018181528351602484015283516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363bd3c818f9386939283926044019185019080838360005b838110156108ac578181015183820152602001610894565b50505050905090810190601f1680156108d95780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156108f857600080fd5b505af115801561090c573d6000803e3d6000fd5b5050505050565b60065481565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600655565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000828201838110156109c7576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b038316610a175760405162461bcd60e51b8152600401808060200182810382526024815260200180610d6e6024913960400191505060405180910390fd5b6001600160a01b038216610a5c5760405162461bcd60e51b8152600401808060200182810382526022815260200180610cd96022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610b035760405162461bcd60e51b8152600401808060200182810382526025815260200180610d496025913960400191505060405180910390fd5b6001600160a01b038216610b485760405162461bcd60e51b8152600401808060200182810382526023815260200180610cb66023913960400191505060405180910390fd5b610b53838383610cb0565b610b9081604051806060016040528060268152602001610cfb602691396001600160a01b0386166000908152602081905260409020549190610c19565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610bbf908261096d565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610ca85760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c6d578181015183820152602001610c55565b50505050905090810190601f168015610c9a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122013b62baa1bc27e631ecc9378f60daba3963a399121116fffe0fab49c3e1e088e64736f6c63430007060033",
}

// AABI is the input ABI used to generate the binding from.
// Deprecated: Use AMetaData.ABI instead.
var AABI = AMetaData.ABI

// Deprecated: Use AMetaData.Sigs instead.
// AFuncSigs maps the 4-byte function signature to its string representation.
var AFuncSigs = AMetaData.Sigs

// ABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AMetaData.Bin instead.
var ABin = AMetaData.Bin

// DeployA deploys a new Ethereum contract, binding an instance of A to it.
func DeployA(auth *bind.TransactOpts, backend bind.ContractBackend, mintable *big.Int, _token0 common.Address, _decimal0 uint8, _token1 common.Address, _decimal1 uint8) (common.Address, *types.Transaction, *A, error) {
	parsed, err := AMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ABin), backend, mintable, _token0, _decimal0, _token1, _decimal1)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// A is an auto generated Go binding around an Ethereum contract.
type A struct {
	ACaller     // Read-only binding to the contract
	ATransactor // Write-only binding to the contract
	AFilterer   // Log filterer for contract events
}

// ACaller is an auto generated read-only Go binding around an Ethereum contract.
type ACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ASession struct {
	Contract     *A                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACallerSession struct {
	Contract *ACaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATransactorSession struct {
	Contract     *ATransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARaw is an auto generated low-level Go binding around an Ethereum contract.
type ARaw struct {
	Contract *A // Generic contract binding to access the raw methods on
}

// ACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACallerRaw struct {
	Contract *ACaller // Generic read-only contract binding to access the raw methods on
}

// ATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATransactorRaw struct {
	Contract *ATransactor // Generic write-only contract binding to access the raw methods on
}

// NewA creates a new instance of A, bound to a specific deployed contract.
func NewA(address common.Address, backend bind.ContractBackend) (*A, error) {
	contract, err := bindA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &A{ACaller: ACaller{contract: contract}, ATransactor: ATransactor{contract: contract}, AFilterer: AFilterer{contract: contract}}, nil
}

// NewACaller creates a new read-only instance of A, bound to a specific deployed contract.
func NewACaller(address common.Address, caller bind.ContractCaller) (*ACaller, error) {
	contract, err := bindA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACaller{contract: contract}, nil
}

// NewATransactor creates a new write-only instance of A, bound to a specific deployed contract.
func NewATransactor(address common.Address, transactor bind.ContractTransactor) (*ATransactor, error) {
	contract, err := bindA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATransactor{contract: contract}, nil
}

// NewAFilterer creates a new log filterer instance of A, bound to a specific deployed contract.
func NewAFilterer(address common.Address, filterer bind.ContractFilterer) (*AFilterer, error) {
	contract, err := bindA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AFilterer{contract: contract}, nil
}

// bindA binds a generic wrapper to an already deployed contract.
func bindA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _A.Contract.ACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.ATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_A *ACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _A.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_A *ATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_A *ATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _A.Contract.contract.Transact(opts, method, params...)
}

// B is a free data retrieval call binding the contract method 0xf2c57866.
//
// Solidity: function _B() view returns(address)
func (_A *ACaller) B(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "_B")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// B is a free data retrieval call binding the contract method 0xf2c57866.
//
// Solidity: function _B() view returns(address)
func (_A *ASession) B() (common.Address, error) {
	return _A.Contract.B(&_A.CallOpts)
}

// B is a free data retrieval call binding the contract method 0xf2c57866.
//
// Solidity: function _B() view returns(address)
func (_A *ACallerSession) B() (common.Address, error) {
	return _A.Contract.B(&_A.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_A *ACaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_A *ASession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _A.Contract.Allowance(&_A.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_A *ACallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _A.Contract.Allowance(&_A.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_A *ACaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_A *ASession) BalanceOf(account common.Address) (*big.Int, error) {
	return _A.Contract.BalanceOf(&_A.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_A *ACallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _A.Contract.BalanceOf(&_A.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_A *ACaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_A *ASession) Decimals() (uint8, error) {
	return _A.Contract.Decimals(&_A.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_A *ACallerSession) Decimals() (uint8, error) {
	return _A.Contract.Decimals(&_A.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_A *ACaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_A *ASession) Name() (string, error) {
	return _A.Contract.Name(&_A.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_A *ACallerSession) Name() (string, error) {
	return _A.Contract.Name(&_A.CallOpts)
}

// Param0 is a free data retrieval call binding the contract method 0xc24d942f.
//
// Solidity: function param0() view returns(uint256)
func (_A *ACaller) Param0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "param0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Param0 is a free data retrieval call binding the contract method 0xc24d942f.
//
// Solidity: function param0() view returns(uint256)
func (_A *ASession) Param0() (*big.Int, error) {
	return _A.Contract.Param0(&_A.CallOpts)
}

// Param0 is a free data retrieval call binding the contract method 0xc24d942f.
//
// Solidity: function param0() view returns(uint256)
func (_A *ACallerSession) Param0() (*big.Int, error) {
	return _A.Contract.Param0(&_A.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_A *ACaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_A *ASession) Symbol() (string, error) {
	return _A.Contract.Symbol(&_A.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_A *ACallerSession) Symbol() (string, error) {
	return _A.Contract.Symbol(&_A.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, uint8 decimal)
func (_A *ACaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr    common.Address
	Decimal uint8
}, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "tokens", arg0)

	outstruct := new(struct {
		Addr    common.Address
		Decimal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Decimal = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, uint8 decimal)
func (_A *ASession) Tokens(arg0 *big.Int) (struct {
	Addr    common.Address
	Decimal uint8
}, error) {
	return _A.Contract.Tokens(&_A.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, uint8 decimal)
func (_A *ACallerSession) Tokens(arg0 *big.Int) (struct {
	Addr    common.Address
	Decimal uint8
}, error) {
	return _A.Contract.Tokens(&_A.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_A *ACaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_A *ASession) TotalSupply() (*big.Int, error) {
	return _A.Contract.TotalSupply(&_A.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_A *ACallerSession) TotalSupply() (*big.Int, error) {
	return _A.Contract.TotalSupply(&_A.CallOpts)
}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_A *ACaller) Uintparam(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _A.contract.Call(opts, &out, "uintparam")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_A *ASession) Uintparam() (*big.Int, error) {
	return _A.Contract.Uintparam(&_A.CallOpts)
}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_A *ACallerSession) Uintparam() (*big.Int, error) {
	return _A.Contract.Uintparam(&_A.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 _decimal) returns()
func (_A *ATransactor) AddToken(opts *bind.TransactOpts, newToken common.Address, _decimal uint8) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "addToken", newToken, _decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 _decimal) returns()
func (_A *ASession) AddToken(newToken common.Address, _decimal uint8) (*types.Transaction, error) {
	return _A.Contract.AddToken(&_A.TransactOpts, newToken, _decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 _decimal) returns()
func (_A *ATransactorSession) AddToken(newToken common.Address, _decimal uint8) (*types.Transaction, error) {
	return _A.Contract.AddToken(&_A.TransactOpts, newToken, _decimal)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_A *ATransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_A *ASession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.Approve(&_A.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_A *ATransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.Approve(&_A.TransactOpts, spender, amount)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string str) returns()
func (_A *ATransactor) ChangebyA(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "changebyA", str)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string str) returns()
func (_A *ASession) ChangebyA(str string) (*types.Transaction, error) {
	return _A.Contract.ChangebyA(&_A.TransactOpts, str)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string str) returns()
func (_A *ATransactorSession) ChangebyA(str string) (*types.Transaction, error) {
	return _A.Contract.ChangebyA(&_A.TransactOpts, str)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_A *ATransactor) CheckTotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "checkTotalSupply")
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_A *ASession) CheckTotalSupply() (*types.Transaction, error) {
	return _A.Contract.CheckTotalSupply(&_A.TransactOpts)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_A *ATransactorSession) CheckTotalSupply() (*types.Transaction, error) {
	return _A.Contract.CheckTotalSupply(&_A.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_A *ATransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_A *ASession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _A.Contract.DecreaseAllowance(&_A.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_A *ATransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _A.Contract.DecreaseAllowance(&_A.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_A *ATransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_A *ASession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _A.Contract.IncreaseAllowance(&_A.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_A *ATransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _A.Contract.IncreaseAllowance(&_A.TransactOpts, spender, addedValue)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_A *ATransactor) SetParam0(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "setParam0", value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_A *ASession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _A.Contract.SetParam0(&_A.TransactOpts, value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_A *ATransactorSession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _A.Contract.SetParam0(&_A.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_A *ATransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_A *ASession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.Transfer(&_A.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_A *ATransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.Transfer(&_A.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_A *ATransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_A *ASession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.TransferFrom(&_A.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_A *ATransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _A.Contract.TransferFrom(&_A.TransactOpts, sender, recipient, amount)
}

// AApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the A contract.
type AApprovalIterator struct {
	Event *AApproval // Event containing the contract specifics and raw log

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
func (it *AApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AApproval)
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
		it.Event = new(AApproval)
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
func (it *AApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AApproval represents a Approval event raised by the A contract.
type AApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_A *AFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _A.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AApprovalIterator{contract: _A.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_A *AFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _A.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AApproval)
				if err := _A.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_A *AFilterer) ParseApproval(log types.Log) (*AApproval, error) {
	event := new(AApproval)
	if err := _A.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ATransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the A contract.
type ATransferIterator struct {
	Event *ATransfer // Event containing the contract specifics and raw log

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
func (it *ATransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATransfer)
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
		it.Event = new(ATransfer)
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
func (it *ATransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATransfer represents a Transfer event raised by the A contract.
type ATransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_A *AFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ATransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _A.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ATransferIterator{contract: _A.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_A *AFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ATransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _A.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATransfer)
				if err := _A.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_A *AFilterer) ParseTransfer(log types.Log) (*ATransfer, error) {
	event := new(ATransfer)
	if err := _A.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b1c6af2111e7d0f27e67aedb161b704638496dfbcbe3f8226c39062689d846f64736f6c63430007060033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AiMetaData contains all meta data concerning the Ai contract.
var AiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTotalSupply\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setParam0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a74ea63f": "addToken(address,uint8)",
		"a511eb2c": "checkTotalSupply()",
		"e0a068e8": "setParam0(uint256)",
	},
}

// AiABI is the input ABI used to generate the binding from.
// Deprecated: Use AiMetaData.ABI instead.
var AiABI = AiMetaData.ABI

// Deprecated: Use AiMetaData.Sigs instead.
// AiFuncSigs maps the 4-byte function signature to its string representation.
var AiFuncSigs = AiMetaData.Sigs

// Ai is an auto generated Go binding around an Ethereum contract.
type Ai struct {
	AiCaller     // Read-only binding to the contract
	AiTransactor // Write-only binding to the contract
	AiFilterer   // Log filterer for contract events
}

// AiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AiSession struct {
	Contract     *Ai               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AiCallerSession struct {
	Contract *AiCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AiTransactorSession struct {
	Contract     *AiTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AiRaw struct {
	Contract *Ai // Generic contract binding to access the raw methods on
}

// AiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AiCallerRaw struct {
	Contract *AiCaller // Generic read-only contract binding to access the raw methods on
}

// AiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AiTransactorRaw struct {
	Contract *AiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAi creates a new instance of Ai, bound to a specific deployed contract.
func NewAi(address common.Address, backend bind.ContractBackend) (*Ai, error) {
	contract, err := bindAi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ai{AiCaller: AiCaller{contract: contract}, AiTransactor: AiTransactor{contract: contract}, AiFilterer: AiFilterer{contract: contract}}, nil
}

// NewAiCaller creates a new read-only instance of Ai, bound to a specific deployed contract.
func NewAiCaller(address common.Address, caller bind.ContractCaller) (*AiCaller, error) {
	contract, err := bindAi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AiCaller{contract: contract}, nil
}

// NewAiTransactor creates a new write-only instance of Ai, bound to a specific deployed contract.
func NewAiTransactor(address common.Address, transactor bind.ContractTransactor) (*AiTransactor, error) {
	contract, err := bindAi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AiTransactor{contract: contract}, nil
}

// NewAiFilterer creates a new log filterer instance of Ai, bound to a specific deployed contract.
func NewAiFilterer(address common.Address, filterer bind.ContractFilterer) (*AiFilterer, error) {
	contract, err := bindAi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AiFilterer{contract: contract}, nil
}

// bindAi binds a generic wrapper to an already deployed contract.
func bindAi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ai *AiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ai.Contract.AiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ai *AiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ai.Contract.AiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ai *AiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ai.Contract.AiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ai *AiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ai *AiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ai *AiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ai.Contract.contract.Transact(opts, method, params...)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 decimal) returns()
func (_Ai *AiTransactor) AddToken(opts *bind.TransactOpts, newToken common.Address, decimal uint8) (*types.Transaction, error) {
	return _Ai.contract.Transact(opts, "addToken", newToken, decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 decimal) returns()
func (_Ai *AiSession) AddToken(newToken common.Address, decimal uint8) (*types.Transaction, error) {
	return _Ai.Contract.AddToken(&_Ai.TransactOpts, newToken, decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address newToken, uint8 decimal) returns()
func (_Ai *AiTransactorSession) AddToken(newToken common.Address, decimal uint8) (*types.Transaction, error) {
	return _Ai.Contract.AddToken(&_Ai.TransactOpts, newToken, decimal)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_Ai *AiTransactor) CheckTotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ai.contract.Transact(opts, "checkTotalSupply")
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_Ai *AiSession) CheckTotalSupply() (*types.Transaction, error) {
	return _Ai.Contract.CheckTotalSupply(&_Ai.TransactOpts)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_Ai *AiTransactorSession) CheckTotalSupply() (*types.Transaction, error) {
	return _Ai.Contract.CheckTotalSupply(&_Ai.TransactOpts)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_Ai *AiTransactor) SetParam0(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Ai.contract.Transact(opts, "setParam0", value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_Ai *AiSession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _Ai.Contract.SetParam0(&_Ai.TransactOpts, value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_Ai *AiTransactorSession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _Ai.Contract.SetParam0(&_Ai.TransactOpts, value)
}

// BMetaData contains all meta data concerning the B contract.
var BMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"internalType\":\"contractAi\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addrparam\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"changebyA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTotalSupply\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setParam0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sparam\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"takeFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uintparam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0dbe671f": "a()",
		"a74ea63f": "addToken(address,uint8)",
		"bdac11ae": "addrparam()",
		"bd3c818f": "changebyA(string)",
		"a511eb2c": "checkTotalSupply()",
		"e0a068e8": "setParam0(uint256)",
		"b7c9fae3": "sparam()",
		"4bbcb421": "takeFund(uint256)",
		"52d74828": "uintparam()",
	},
	Bin: "0x60a060405234801561001057600080fd5b5061001961002e565b60601b6001600160601b031916608052610124565b6040805144602080830191909152428284015273b4fbf271143f4fbf7b91a5ded31805e42b2208d66060830181905273d87ba7a50b2e7e660f678a895e4b72e7cb4ccd9c60808085018290528551808603909101815260a090940194859052835193909201929092206000936107d093929160129160089190869086908590879086906100ba90610117565b9485526001600160a01b03938416602086015260ff9283166040808701919091529190931660608501529116608083015251829181900360a001906000f590508015801561010c573d6000803e3d6000fd5b509695505050505050565b6111f68061079583390190565b60805160601c6106406101556000398061027c52806102a0528061032c52806103c052806104ed52506106406000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a74ea63f11610066578063a74ea63f14610120578063b7c9fae31461014f578063bd3c818f146101cc578063bdac11ae14610272578063e0a068e8146100bc57610093565b80630dbe671f146100985780634bbcb421146100bc57806352d74828146100db578063a511eb2c146100f5575b600080fd5b6100a061027a565b604080516001600160a01b039092168252519081900360200190f35b6100d9600480360360208110156100d257600080fd5b503561029e565b005b6100e361031f565b60408051918252519081900360200190f35b6100fd610325565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6100d96004803603604081101561013657600080fd5b5080356001600160a01b0316906020013560ff166103be565b610157610454565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610191578181015183820152602001610179565b50505050905090810190601f1680156101be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100d9600480360360208110156101e257600080fd5b8101906020810181356401000000008111156101fd57600080fd5b82018360208201111561020f57600080fd5b8035906020019184600183028401116401000000008311171561023157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104e2945050505050565b6100a061055a565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e0a068e8826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561030457600080fd5b505af1158015610318573d6000803e3d6000fd5b5050505050565b60015481565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561038357600080fd5b505afa158015610397573d6000803e3d6000fd5b505050506040513d60208110156103ad57600080fd5b505160018190553393509150509091565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a74ea63f83836040518363ffffffff1660e01b815260040180836001600160a01b031681526020018260ff16815260200192505050600060405180830381600087803b15801561043857600080fd5b505af115801561044c573d6000803e3d6000fd5b505050505050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104da5780601f106104af576101008083540402835291602001916104da565b820191906000526020600020905b8154815290600101906020018083116104bd57829003601f168201915b505050505081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461054357604080518082019091526005808252646e6f74204160d81b602090920191825261053e91600091610569565b610556565b8051610556906000906020840190610569565b5050565b6002546001600160a01b031681565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261059f57600085556105e5565b82601f106105b857805160ff19168380011785556105e5565b828001600101855582156105e5579182015b828111156105e55782518255916020019190600101906105ca565b506105f19291506105f5565b5090565b5b808211156105f157600081556001016105f656fea2646970667358221220b203960194202d64714130d3f45e8f0a3d8490d244d1a0631e7451cac5d6b94d64736f6c6343000706003360a06040523480156200001157600080fd5b50604051620011f6380380620011f6833981810160405260a08110156200003757600080fd5b50805160208083015160408085015160608601516080909601518251808401845260058152642a37b5b2b760d91b81870190815284518086019095526003808652622a25a760e91b978601979097528151979895979396929491939192620000a2929091906200033b565b508051620000b89060049060208401906200033b565b50506005805460ff1916601217905550620000e533620000d7620001bc565b60ff16600a0a8702620001c5565b6040805180820182526001600160a01b03958616815260ff94851660208083019182526008805460018181018355600083815295517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3928301805496518c16600160a01b90810260ff60a01b19938f166001600160a01b0319998a16178416179091558851808a01909952998c168852978a16938701938452825490810183559190945293519390920180549251909616909402919095169490921693909317909216919091179055503360601b608052620003e7565b60055460ff1690565b6001600160a01b03821662000221576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200022f60008383620002d4565b6200024b81600254620002d960201b6200096d1790919060201c565b6002556001600160a01b038216600090815260208181526040909120546200027e9183906200096d620002d9821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000334576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003735760008555620003be565b82601f106200038e57805160ff1916838001178555620003be565b82800160010185558215620003be579182015b82811115620003be578251825591602001919060010190620003a1565b50620003cc929150620003d0565b5090565b5b80821115620003cc5760008155600101620003d1565b60805160601c610dec6200040a6000398061085b528061094b5250610dec6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806395d89b41116100ad578063bd3c818f11610071578063bd3c818f146103a9578063c24d942f1461044f578063dd62ed3e14610457578063e0a068e814610485578063f2c57866146104a257610121565b806395d89b41146102ed578063a457c2d7146102f5578063a511eb2c14610321578063a74ea63f1461034c578063a9059cbb1461037d57610121565b8063313ce567116100f4578063313ce5671461023357806339509351146102515780634f64b2be1461027d57806352d74828146102bf57806370a08231146102c757610121565b806306fdde0314610126578063095ea7b3146101a357806318160ddd146101e357806323b872dd146101fd575b600080fd5b61012e6104c6565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610168578181015183820152602001610150565b50505050905090810190601f1680156101955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101cf600480360360408110156101b957600080fd5b506001600160a01b03813516906020013561055c565b604080519115158252519081900360200190f35b6101eb610579565b60408051918252519081900360200190f35b6101cf6004803603606081101561021357600080fd5b506001600160a01b0381358116916020810135909116906040013561057f565b61023b610606565b6040805160ff9092168252519081900360200190f35b6101cf6004803603604081101561026757600080fd5b506001600160a01b03813516906020013561060f565b61029a6004803603602081101561029357600080fd5b503561065d565b604080516001600160a01b03909316835260ff90911660208301528051918290030190f35b6101eb610692565b6101eb600480360360208110156102dd57600080fd5b50356001600160a01b0316610698565b61012e6106b3565b6101cf6004803603604081101561030b57600080fd5b506001600160a01b038135169060200135610714565b61032961077c565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61037b6004803603604081101561036257600080fd5b5080356001600160a01b0316906020013560ff16610798565b005b6101cf6004803603604081101561039357600080fd5b506001600160a01b03813516906020013561081e565b61037b600480360360208110156103bf57600080fd5b8101906020810181356401000000008111156103da57600080fd5b8201836020820111156103ec57600080fd5b8035906020019184600183028401116401000000008311171561040e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610832945050505050565b6101eb610913565b6101eb6004803603604081101561046d57600080fd5b506001600160a01b0381358116916020013516610919565b61037b6004803603602081101561049b57600080fd5b5035610944565b6104aa610949565b604080516001600160a01b039092168252519081900360200190f35b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105525780601f1061052757610100808354040283529160200191610552565b820191906000526020600020905b81548152906001019060200180831161053557829003601f168201915b5050505050905090565b60006105706105696109ce565b84846109d2565b50600192915050565b60025490565b600061058c848484610abe565b6105fc846105986109ce565b6105f785604051806060016040528060288152602001610d21602891396001600160a01b038a166000908152600160205260408120906105d66109ce565b6001600160a01b031681526020810191909152604001600020549190610c19565b6109d2565b5060019392505050565b60055460ff1690565b600061057061061c6109ce565b846105f7856001600061062d6109ce565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061096d565b6008818154811061066d57600080fd5b6000918252602090912001546001600160a01b0381169150600160a01b900460ff1682565b60075481565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105525780601f1061052757610100808354040283529160200191610552565b60006105706107216109ce565b846105f785604051806060016040528060258152602001610d92602591396001600061074b6109ce565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610c19565b6000806000610789610579565b60078190553393509150509091565b604080518082019091526001600160a01b03928316815260ff918216602082019081526008805460018101825560009190915291517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180549151909316600160a01b0260ff60a01b19929094166001600160a01b03199091161716919091179055565b600061057061082b6109ce565b8484610abe565b60405163bd3c818f60e01b81526020600482018181528351602484015283516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363bd3c818f9386939283926044019185019080838360005b838110156108ac578181015183820152602001610894565b50505050905090810190601f1680156108d95780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156108f857600080fd5b505af115801561090c573d6000803e3d6000fd5b5050505050565b60065481565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600655565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000828201838110156109c7576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b038316610a175760405162461bcd60e51b8152600401808060200182810382526024815260200180610d6e6024913960400191505060405180910390fd5b6001600160a01b038216610a5c5760405162461bcd60e51b8152600401808060200182810382526022815260200180610cd96022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610b035760405162461bcd60e51b8152600401808060200182810382526025815260200180610d496025913960400191505060405180910390fd5b6001600160a01b038216610b485760405162461bcd60e51b8152600401808060200182810382526023815260200180610cb66023913960400191505060405180910390fd5b610b53838383610cb0565b610b9081604051806060016040528060268152602001610cfb602691396001600160a01b0386166000908152602081905260409020549190610c19565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610bbf908261096d565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610ca85760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c6d578181015183820152602001610c55565b50505050905090810190601f168015610c9a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122013b62baa1bc27e631ecc9378f60daba3963a399121116fffe0fab49c3e1e088e64736f6c63430007060033",
}

// BABI is the input ABI used to generate the binding from.
// Deprecated: Use BMetaData.ABI instead.
var BABI = BMetaData.ABI

// Deprecated: Use BMetaData.Sigs instead.
// BFuncSigs maps the 4-byte function signature to its string representation.
var BFuncSigs = BMetaData.Sigs

// BBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BMetaData.Bin instead.
var BBin = BMetaData.Bin

// DeployB deploys a new Ethereum contract, binding an instance of B to it.
func DeployB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *B, error) {
	parsed, err := BMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &B{BCaller: BCaller{contract: contract}, BTransactor: BTransactor{contract: contract}, BFilterer: BFilterer{contract: contract}}, nil
}

// B is an auto generated Go binding around an Ethereum contract.
type B struct {
	BCaller     // Read-only binding to the contract
	BTransactor // Write-only binding to the contract
	BFilterer   // Log filterer for contract events
}

// BCaller is an auto generated read-only Go binding around an Ethereum contract.
type BCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSession struct {
	Contract     *B                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BCallerSession struct {
	Contract *BCaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTransactorSession struct {
	Contract     *BTransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BRaw is an auto generated low-level Go binding around an Ethereum contract.
type BRaw struct {
	Contract *B // Generic contract binding to access the raw methods on
}

// BCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BCallerRaw struct {
	Contract *BCaller // Generic read-only contract binding to access the raw methods on
}

// BTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTransactorRaw struct {
	Contract *BTransactor // Generic write-only contract binding to access the raw methods on
}

// NewB creates a new instance of B, bound to a specific deployed contract.
func NewB(address common.Address, backend bind.ContractBackend) (*B, error) {
	contract, err := bindB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &B{BCaller: BCaller{contract: contract}, BTransactor: BTransactor{contract: contract}, BFilterer: BFilterer{contract: contract}}, nil
}

// NewBCaller creates a new read-only instance of B, bound to a specific deployed contract.
func NewBCaller(address common.Address, caller bind.ContractCaller) (*BCaller, error) {
	contract, err := bindB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BCaller{contract: contract}, nil
}

// NewBTransactor creates a new write-only instance of B, bound to a specific deployed contract.
func NewBTransactor(address common.Address, transactor bind.ContractTransactor) (*BTransactor, error) {
	contract, err := bindB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTransactor{contract: contract}, nil
}

// NewBFilterer creates a new log filterer instance of B, bound to a specific deployed contract.
func NewBFilterer(address common.Address, filterer bind.ContractFilterer) (*BFilterer, error) {
	contract, err := bindB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BFilterer{contract: contract}, nil
}

// bindB binds a generic wrapper to an already deployed contract.
func bindB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B *BRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _B.Contract.BCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B *BRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B.Contract.BTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B *BRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B.Contract.BTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B *BCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _B.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B *BTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B *BTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(address)
func (_B *BCaller) A(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _B.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(address)
func (_B *BSession) A() (common.Address, error) {
	return _B.Contract.A(&_B.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(address)
func (_B *BCallerSession) A() (common.Address, error) {
	return _B.Contract.A(&_B.CallOpts)
}

// Addrparam is a free data retrieval call binding the contract method 0xbdac11ae.
//
// Solidity: function addrparam() view returns(address)
func (_B *BCaller) Addrparam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _B.contract.Call(opts, &out, "addrparam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addrparam is a free data retrieval call binding the contract method 0xbdac11ae.
//
// Solidity: function addrparam() view returns(address)
func (_B *BSession) Addrparam() (common.Address, error) {
	return _B.Contract.Addrparam(&_B.CallOpts)
}

// Addrparam is a free data retrieval call binding the contract method 0xbdac11ae.
//
// Solidity: function addrparam() view returns(address)
func (_B *BCallerSession) Addrparam() (common.Address, error) {
	return _B.Contract.Addrparam(&_B.CallOpts)
}

// Sparam is a free data retrieval call binding the contract method 0xb7c9fae3.
//
// Solidity: function sparam() view returns(string)
func (_B *BCaller) Sparam(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _B.contract.Call(opts, &out, "sparam")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Sparam is a free data retrieval call binding the contract method 0xb7c9fae3.
//
// Solidity: function sparam() view returns(string)
func (_B *BSession) Sparam() (string, error) {
	return _B.Contract.Sparam(&_B.CallOpts)
}

// Sparam is a free data retrieval call binding the contract method 0xb7c9fae3.
//
// Solidity: function sparam() view returns(string)
func (_B *BCallerSession) Sparam() (string, error) {
	return _B.Contract.Sparam(&_B.CallOpts)
}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_B *BCaller) Uintparam(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _B.contract.Call(opts, &out, "uintparam")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_B *BSession) Uintparam() (*big.Int, error) {
	return _B.Contract.Uintparam(&_B.CallOpts)
}

// Uintparam is a free data retrieval call binding the contract method 0x52d74828.
//
// Solidity: function uintparam() view returns(uint256)
func (_B *BCallerSession) Uintparam() (*big.Int, error) {
	return _B.Contract.Uintparam(&_B.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address _token, uint8 decimal) returns()
func (_B *BTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, decimal uint8) (*types.Transaction, error) {
	return _B.contract.Transact(opts, "addToken", _token, decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address _token, uint8 decimal) returns()
func (_B *BSession) AddToken(_token common.Address, decimal uint8) (*types.Transaction, error) {
	return _B.Contract.AddToken(&_B.TransactOpts, _token, decimal)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address _token, uint8 decimal) returns()
func (_B *BTransactorSession) AddToken(_token common.Address, decimal uint8) (*types.Transaction, error) {
	return _B.Contract.AddToken(&_B.TransactOpts, _token, decimal)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_B *BTransactor) ChangebyA(opts *bind.TransactOpts, value string) (*types.Transaction, error) {
	return _B.contract.Transact(opts, "changebyA", value)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_B *BSession) ChangebyA(value string) (*types.Transaction, error) {
	return _B.Contract.ChangebyA(&_B.TransactOpts, value)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_B *BTransactorSession) ChangebyA(value string) (*types.Transaction, error) {
	return _B.Contract.ChangebyA(&_B.TransactOpts, value)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_B *BTransactor) CheckTotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B.contract.Transact(opts, "checkTotalSupply")
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_B *BSession) CheckTotalSupply() (*types.Transaction, error) {
	return _B.Contract.CheckTotalSupply(&_B.TransactOpts)
}

// CheckTotalSupply is a paid mutator transaction binding the contract method 0xa511eb2c.
//
// Solidity: function checkTotalSupply() returns(address, uint256)
func (_B *BTransactorSession) CheckTotalSupply() (*types.Transaction, error) {
	return _B.Contract.CheckTotalSupply(&_B.TransactOpts)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_B *BTransactor) SetParam0(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _B.contract.Transact(opts, "setParam0", value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_B *BSession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _B.Contract.SetParam0(&_B.TransactOpts, value)
}

// SetParam0 is a paid mutator transaction binding the contract method 0xe0a068e8.
//
// Solidity: function setParam0(uint256 value) returns()
func (_B *BTransactorSession) SetParam0(value *big.Int) (*types.Transaction, error) {
	return _B.Contract.SetParam0(&_B.TransactOpts, value)
}

// TakeFund is a paid mutator transaction binding the contract method 0x4bbcb421.
//
// Solidity: function takeFund(uint256 value) returns()
func (_B *BTransactor) TakeFund(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _B.contract.Transact(opts, "takeFund", value)
}

// TakeFund is a paid mutator transaction binding the contract method 0x4bbcb421.
//
// Solidity: function takeFund(uint256 value) returns()
func (_B *BSession) TakeFund(value *big.Int) (*types.Transaction, error) {
	return _B.Contract.TakeFund(&_B.TransactOpts, value)
}

// TakeFund is a paid mutator transaction binding the contract method 0x4bbcb421.
//
// Solidity: function takeFund(uint256 value) returns()
func (_B *BTransactorSession) TakeFund(value *big.Int) (*types.Transaction, error) {
	return _B.Contract.TakeFund(&_B.TransactOpts, value)
}

// BiMetaData contains all meta data concerning the Bi contract.
var BiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"changebyA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bd3c818f": "changebyA(string)",
	},
}

// BiABI is the input ABI used to generate the binding from.
// Deprecated: Use BiMetaData.ABI instead.
var BiABI = BiMetaData.ABI

// Deprecated: Use BiMetaData.Sigs instead.
// BiFuncSigs maps the 4-byte function signature to its string representation.
var BiFuncSigs = BiMetaData.Sigs

// Bi is an auto generated Go binding around an Ethereum contract.
type Bi struct {
	BiCaller     // Read-only binding to the contract
	BiTransactor // Write-only binding to the contract
	BiFilterer   // Log filterer for contract events
}

// BiCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiSession struct {
	Contract     *Bi               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiCallerSession struct {
	Contract *BiCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiTransactorSession struct {
	Contract     *BiTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiRaw struct {
	Contract *Bi // Generic contract binding to access the raw methods on
}

// BiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiCallerRaw struct {
	Contract *BiCaller // Generic read-only contract binding to access the raw methods on
}

// BiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiTransactorRaw struct {
	Contract *BiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBi creates a new instance of Bi, bound to a specific deployed contract.
func NewBi(address common.Address, backend bind.ContractBackend) (*Bi, error) {
	contract, err := bindBi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bi{BiCaller: BiCaller{contract: contract}, BiTransactor: BiTransactor{contract: contract}, BiFilterer: BiFilterer{contract: contract}}, nil
}

// NewBiCaller creates a new read-only instance of Bi, bound to a specific deployed contract.
func NewBiCaller(address common.Address, caller bind.ContractCaller) (*BiCaller, error) {
	contract, err := bindBi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiCaller{contract: contract}, nil
}

// NewBiTransactor creates a new write-only instance of Bi, bound to a specific deployed contract.
func NewBiTransactor(address common.Address, transactor bind.ContractTransactor) (*BiTransactor, error) {
	contract, err := bindBi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiTransactor{contract: contract}, nil
}

// NewBiFilterer creates a new log filterer instance of Bi, bound to a specific deployed contract.
func NewBiFilterer(address common.Address, filterer bind.ContractFilterer) (*BiFilterer, error) {
	contract, err := bindBi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiFilterer{contract: contract}, nil
}

// bindBi binds a generic wrapper to an already deployed contract.
func bindBi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bi *BiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bi.Contract.BiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bi *BiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bi.Contract.BiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bi *BiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bi.Contract.BiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bi *BiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bi *BiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bi *BiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bi.Contract.contract.Transact(opts, method, params...)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_Bi *BiTransactor) ChangebyA(opts *bind.TransactOpts, value string) (*types.Transaction, error) {
	return _Bi.contract.Transact(opts, "changebyA", value)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_Bi *BiSession) ChangebyA(value string) (*types.Transaction, error) {
	return _Bi.Contract.ChangebyA(&_Bi.TransactOpts, value)
}

// ChangebyA is a paid mutator transaction binding the contract method 0xbd3c818f.
//
// Solidity: function changebyA(string value) returns()
func (_Bi *BiTransactorSession) ChangebyA(value string) (*types.Transaction, error) {
	return _Bi.Contract.ChangebyA(&_Bi.TransactOpts, value)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DeployerMetaData contains all meta data concerning the Deployer contract.
var DeployerMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220380daadee47d07cdf421957983ce48a35d6d6efb3358c92c772ad8672312513d64736f6c63430007060033",
}

// DeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployerMetaData.ABI instead.
var DeployerABI = DeployerMetaData.ABI

// DeployerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeployerMetaData.Bin instead.
var DeployerBin = DeployerMetaData.Bin

// DeployDeployer deploys a new Ethereum contract, binding an instance of Deployer to it.
func DeployDeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Deployer, error) {
	parsed, err := DeployerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// Deployer is an auto generated Go binding around an Ethereum contract.
type Deployer struct {
	DeployerCaller     // Read-only binding to the contract
	DeployerTransactor // Write-only binding to the contract
	DeployerFilterer   // Log filterer for contract events
}

// DeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployerSession struct {
	Contract     *Deployer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployerCallerSession struct {
	Contract *DeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployerTransactorSession struct {
	Contract     *DeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployerRaw struct {
	Contract *Deployer // Generic contract binding to access the raw methods on
}

// DeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployerCallerRaw struct {
	Contract *DeployerCaller // Generic read-only contract binding to access the raw methods on
}

// DeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployerTransactorRaw struct {
	Contract *DeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployer creates a new instance of Deployer, bound to a specific deployed contract.
func NewDeployer(address common.Address, backend bind.ContractBackend) (*Deployer, error) {
	contract, err := bindDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// NewDeployerCaller creates a new read-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerCaller(address common.Address, caller bind.ContractCaller) (*DeployerCaller, error) {
	contract, err := bindDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerCaller{contract: contract}, nil
}

// NewDeployerTransactor creates a new write-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployerTransactor, error) {
	contract, err := bindDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerTransactor{contract: contract}, nil
}

// NewDeployerFilterer creates a new log filterer instance of Deployer, bound to a specific deployed contract.
func NewDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployerFilterer, error) {
	contract, err := bindDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployerFilterer{contract: contract}, nil
}

// bindDeployer binds a generic wrapper to an already deployed contract.
func bindDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.DeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transact(opts, method, params...)
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000c8e38038062000c8e833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200028c565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000218576000855562000263565b82601f106200023357805160ff191683800117855562000263565b8280016001018555821562000263579182015b828111156200026357825182559160200191906001019062000246565b506200027192915062000275565b5090565b5b8082111562000271576000815560010162000276565b6109f2806200029c6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063395093511161007157806339509351146101d957806370a082311461020557806395d89b411461022b578063a457c2d714610233578063a9059cbb1461025f578063dd62ed3e1461028b576100a9565b806306fdde03146100ae578063095ea7b31461012b57806318160ddd1461016b57806323b872dd14610185578063313ce567146101bb575b600080fd5b6100b66102b9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101576004803603604081101561014157600080fd5b506001600160a01b03813516906020013561034f565b604080519115158252519081900360200190f35b61017361036c565b60408051918252519081900360200190f35b6101576004803603606081101561019b57600080fd5b506001600160a01b03813581169160208101359091169060400135610372565b6101c36103f9565b6040805160ff9092168252519081900360200190f35b610157600480360360408110156101ef57600080fd5b506001600160a01b038135169060200135610402565b6101736004803603602081101561021b57600080fd5b50356001600160a01b0316610450565b6100b661046b565b6101576004803603604081101561024957600080fd5b506001600160a01b0381351690602001356104cc565b6101576004803603604081101561027557600080fd5b506001600160a01b038135169060200135610534565b610173600480360360408110156102a157600080fd5b506001600160a01b0381358116916020013516610548565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b820191906000526020600020905b81548152906001019060200180831161032857829003601f168201915b5050505050905090565b600061036361035c610573565b8484610577565b50600192915050565b60025490565b600061037f848484610663565b6103ef8461038b610573565b6103ea85604051806060016040528060288152602001610927602891396001600160a01b038a166000908152600160205260408120906103c9610573565b6001600160a01b0316815260208101919091526040016000205491906107be565b610577565b5060019392505050565b60055460ff1690565b600061036361040f610573565b846103ea8560016000610420610573565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610855565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b60006103636104d9610573565b846103ea856040518060600160405280602581526020016109986025913960016000610503610573565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906107be565b6000610363610541610573565b8484610663565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166105bc5760405162461bcd60e51b81526004018080602001828103825260248152602001806109746024913960400191505060405180910390fd5b6001600160a01b0382166106015760405162461bcd60e51b81526004018080602001828103825260228152602001806108df6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106a85760405162461bcd60e51b815260040180806020018281038252602581526020018061094f6025913960400191505060405180910390fd5b6001600160a01b0382166106ed5760405162461bcd60e51b81526004018080602001828103825260238152602001806108bc6023913960400191505060405180910390fd5b6106f88383836108b6565b61073581604051806060016040528060268152602001610901602691396001600160a01b03861660009081526020819052604090205491906107be565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107649082610855565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561084d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108125781810151838201526020016107fa565b50505050905090810190601f16801561083f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156108af576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122060792cbda3ba967fd9bbbb595d2735b600fddb89527b64097739631c835b4bc864736f6c63430007060033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// Deprecated: Use ERC20MetaData.Sigs instead.
// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = ERC20MetaData.Sigs

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"06fdde03": "name()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203f021cabf32240e70d7f3ec029e52149c1644b4a9e05e81be5157082f05e2dd164736f6c63430007060033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122065d5f9eece32ac298ee05490a79672f866925434c6003565bd3284de8f8541c564736f6c63430007060033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
