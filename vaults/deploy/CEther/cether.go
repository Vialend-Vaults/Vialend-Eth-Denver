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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolSeizeShareMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"repayBorrow\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60806040523480156200001157600080fd5b50604051620059a3380380620059a3833981810160405260e08110156200003757600080fd5b8151602083015160408085015160608601805192519496939591949391820192846401000000008211156200006b57600080fd5b9083019060208201858111156200008157600080fd5b82516401000000008111828201881017156200009c57600080fd5b82525081516020918201929091019080838360005b83811015620000cb578181015183820152602001620000b1565b50505050905090810190601f168015620000f95780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011d57600080fd5b9083019060208201858111156200013357600080fd5b82516401000000008111828201881017156200014e57600080fd5b82525081516020918201929091019080838360005b838110156200017d57818101518382015260200162000163565b50505050905090810190601f168015620001ab5780820380516001836020036101000a031916815260200191505b506040908152602082015191015160038054610100600160a81b03191633610100021790559092509050620001e587878787878762000218565b600380546001600160a01b0390921661010002610100600160a81b03199092169190911790555062000853945050505050565b60035461010090046001600160a01b03163314620002685760405162461bcd60e51b81526004018080602001828103825260248152602001806200590a6024913960400191505060405180910390fd5b600954158015620002795750600a54155b620002b65760405162461bcd60e51b81526004018080602001828103825260238152602001806200592e6023913960400191505060405180910390fd5b600784905583620002f95760405162461bcd60e51b8152600401808060200182810382526030815260200180620059516030913960400191505060405180910390fd5b60006200030f876001600160e01b036200042e16565b9050801562000365576040805162461bcd60e51b815260206004820152601a60248201527f73657474696e6720636f6d7074726f6c6c6572206661696c6564000000000000604482015290519081900360640190fd5b620003786001600160e01b036200059616565b600955670de0b6b3a7640000600a556200039b866001600160e01b036200059b16565b90508015620003dc5760405162461bcd60e51b8152600401808060200182810382526022815260200180620059816022913960400191505060405180910390fd5b8351620003f1906001906020870190620007b1565b50825162000407906002906020860190620007b1565b50506003805460ff90921660ff199283161790556000805490911660011790555050505050565b60035460009061010090046001600160a01b031633146200046857620004606001603f6001600160e01b036200074116565b905062000591565b60055460408051623f1ee960e11b815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b158015620004ae57600080fd5b505afa158015620004c3573d6000803e3d6000fd5b505050506040513d6020811015620004da57600080fd5b50516200052e576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160005b9150505b919050565b435b90565b600354600090819061010090046001600160a01b03163314620005d857620005cf600160426001600160e01b036200074116565b91505062000591565b620005eb6001600160e01b036200059616565b600954146200060b57620005cf600a60416001600160e01b036200074116565b600660009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200065d57600080fd5b505afa15801562000672573d6000803e3d6000fd5b505050506040513d60208110156200068957600080fd5b5051620006dd576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a160006200058d565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa08360108111156200077157fe5b8360508111156200077e57fe5b604080519283526020830191909152600082820152519081900360600190a1826010811115620007aa57fe5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620007f457805160ff191683800117855562000824565b8280016001018555821562000824579182015b828111156200082457825182559160200191906001019062000807565b506200083292915062000836565b5090565b6200059891905b808211156200083257600081556001016200083d565b6150a780620008636000396000f3fe6080604052600436106102885760003560e01c806395d89b411161015a578063c5ebeaec116100c1578063f3fdb15a1161007a578063f3fdb15a14610a80578063f851a44014610a95578063f8f9da2814610aaa578063fca7820b14610abf578063fcb6414714610ae9578063fe9c44ae14610af157610288565b8063c5ebeaec14610983578063db006a75146109ad578063dd62ed3e146109d7578063e597461914610a12578063e9c714f214610a38578063f2b3abbd14610a4d57610288565b8063aae40a2a11610113578063aae40a2a1461085c578063ae9d70b01461088a578063b2a02ff11461089f578063b71d1a0c146108e2578063bd6d894d14610915578063c37f68e21461092a57610288565b806395d89b411461065257806395dd91931461066757806399d8c1b41461069a578063a6afed95146107f9578063a9059cbb1461080e578063aa5af0fd1461084757610288565b80633b1d21a2116101fe5780636752e702116101b75780636752e702146105a15780636c540baf146105b657806370a08231146105cb57806373acee98146105fe578063852a12e3146106135780638f840ddd1461063d57610288565b80633b1d21a2146104fd5780634576b5db1461051257806347bd3718146105455780634e4d9fea1461055a5780635fe3b56714610562578063601a0bf11461057757610288565b806318160ddd1161025057806318160ddd14610401578063182df0f51461041657806323b872dd1461042b578063267822471461046e578063313ce5671461049f5780633af9e669146104ca57610288565b806306fdde03146102c6578063095ea7b3146103505780631249c58b1461039d578063173b9904146103a757806317bfdfbc146103ce575b600061029334610b06565b5090506102c3816040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b815250610bae565b50005b3480156102d257600080fd5b506102db610dae565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103155781810151838201526020016102fd565b50505050905090810190601f1680156103425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035c57600080fd5b506103896004803603604081101561037357600080fd5b506001600160a01b038135169060200135610e3b565b604080519115158252519081900360200190f35b6103a5610ea8565b005b3480156103b357600080fd5b506103bc610ee6565b60408051918252519081900360200190f35b3480156103da57600080fd5b506103bc600480360360208110156103f157600080fd5b50356001600160a01b0316610eec565b34801561040d57600080fd5b506103bc610fac565b34801561042257600080fd5b506103bc610fb2565b34801561043757600080fd5b506103896004803603606081101561044e57600080fd5b506001600160a01b03813581169160208101359091169060400135611015565b34801561047a57600080fd5b50610483611087565b604080516001600160a01b039092168252519081900360200190f35b3480156104ab57600080fd5b506104b4611096565b6040805160ff9092168252519081900360200190f35b3480156104d657600080fd5b506103bc600480360360208110156104ed57600080fd5b50356001600160a01b031661109f565b34801561050957600080fd5b506103bc611157565b34801561051e57600080fd5b506103bc6004803603602081101561053557600080fd5b50356001600160a01b0316611166565b34801561055157600080fd5b506103bc6112bb565b6103a56112c1565b34801561056e57600080fd5b50610483611303565b34801561058357600080fd5b506103bc6004803603602081101561059a57600080fd5b5035611312565b3480156105ad57600080fd5b506103bc6113ad565b3480156105c257600080fd5b506103bc6113b8565b3480156105d757600080fd5b506103bc600480360360208110156105ee57600080fd5b50356001600160a01b03166113be565b34801561060a57600080fd5b506103bc6113d9565b34801561061f57600080fd5b506103bc6004803603602081101561063657600080fd5b503561148f565b34801561064957600080fd5b506103bc61149a565b34801561065e57600080fd5b506102db6114a0565b34801561067357600080fd5b506103bc6004803603602081101561068a57600080fd5b50356001600160a01b03166114f8565b3480156106a657600080fd5b506103a5600480360360c08110156106bd57600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156106f857600080fd5b82018360208201111561070a57600080fd5b8035906020019184600183028401116401000000008311171561072c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561077f57600080fd5b82018360208201111561079157600080fd5b803590602001918460018302840111640100000000831117156107b357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506115559050565b34801561080557600080fd5b506103bc61173c565b34801561081a57600080fd5b506103896004803603604081101561083157600080fd5b506001600160a01b038135169060200135611a94565b34801561085357600080fd5b506103bc611b05565b6103a56004803603604081101561087257600080fd5b506001600160a01b0381358116916020013516611b0b565b34801561089657600080fd5b506103bc611b58565b3480156108ab57600080fd5b506103bc600480360360608110156108c257600080fd5b506001600160a01b03813581169160208101359091169060400135611bf7565b3480156108ee57600080fd5b506103bc6004803603602081101561090557600080fd5b50356001600160a01b0316611c68565b34801561092157600080fd5b506103bc611cf4565b34801561093657600080fd5b5061095d6004803603602081101561094d57600080fd5b50356001600160a01b0316611db0565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561098f57600080fd5b506103bc600480360360208110156109a657600080fd5b5035611e45565b3480156109b957600080fd5b506103bc600480360360208110156109d057600080fd5b5035611e50565b3480156109e357600080fd5b506103bc600480360360408110156109fa57600080fd5b506001600160a01b0381358116916020013516611e5b565b6103a560048036036020811015610a2857600080fd5b50356001600160a01b0316611e86565b348015610a4457600080fd5b506103bc611ed4565b348015610a5957600080fd5b506103bc60048036036020811015610a7057600080fd5b50356001600160a01b0316611fd7565b348015610a8c57600080fd5b50610483612011565b348015610aa157600080fd5b50610483612020565b348015610ab657600080fd5b506103bc612034565b348015610acb57600080fd5b506103bc60048036036020811015610ae257600080fd5b5035612098565b6103bc612116565b348015610afd57600080fd5b50610389612121565b60008054819060ff16610b4d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610b5f61173c565b90508015610b8a57610b7d816010811115610b7657fe5b601e612126565b925060009150610b9a9050565b610b94338561218c565b92509250505b6000805460ff191660011790559092909150565b81610bb857610daa565b606081516005016040519080825280601f01601f191660200182016040528015610be9576020820181803883390190505b50905060005b8251811015610c3a57828181518110610c0457fe5b602001015160f81c60f81b828281518110610c1b57fe5b60200101906001600160f81b031916908160001a905350600101610bef565b8151600160fd1b90839083908110610c4e57fe5b60200101906001600160f81b031916908160001a905350602860f81b828260010181518110610c7957fe5b60200101906001600160f81b031916908160001a905350600a840460300160f81b828260020181518110610ca957fe5b60200101906001600160f81b031916908160001a905350600a840660300160f81b828260030181518110610cd957fe5b60200101906001600160f81b031916908160001a905350602960f81b828260040181518110610d0457fe5b60200101906001600160f81b031916908160001a905350818415610da65760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d6b578181015183820152602001610d53565b50505050905090810190601f168015610d985780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5050505b5050565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610e335780601f10610e0857610100808354040283529160200191610e33565b820191906000526020600020905b815481529060010190602001808311610e1657829003601f168201915b505050505081565b336000818152600f602090815260408083206001600160a01b03871680855290835281842086905581518681529151939493909284927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a360019150505b92915050565b6000610eb334610b06565b509050610ee3816040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b815250610bae565b50565b60085481565b6000805460ff16610f31576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610f4361173c565b14610f8e576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b610f97826114f8565b90505b6000805460ff19166001179055919050565b600d5481565b6000806000610fbf61255e565b90925090506000826003811115610fd257fe5b1461100e5760405162461bcd60e51b8152600401808060200182810382526035815260200180614fbe6035913960400191505060405180910390fd5b9150505b90565b6000805460ff1661105a576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556110703386868661260d565b1490506000805460ff191660011790559392505050565b6004546001600160a01b031681565b60035460ff1681565b60006110a9614c5f565b60405180602001604052806110bc611cf4565b90526001600160a01b0384166000908152600e60205260408120549192509081906110e890849061289b565b909250905060008260038111156110fb57fe5b1461114d576040805162461bcd60e51b815260206004820152601f60248201527f62616c616e636520636f756c64206e6f742062652063616c63756c6174656400604482015290519081900360640190fd5b925050505b919050565b60006111616128ee565b905090565b60035460009061010090046001600160a01b031633146111935761118c6001603f612126565b9050611152565b60055460408051623f1ee960e11b815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b1580156111d857600080fd5b505afa1580156111ec573d6000803e3d6000fd5b505050506040513d602081101561120257600080fd5b5051611255576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160005b9392505050565b600b5481565b60006112cc3461291a565b509050610ee381604051806040016040528060128152602001711c995c185e509bdc9c9bddc819985a5b195960721b815250610bae565b6005546001600160a01b031681565b6000805460ff16611357576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561136961173c565b9050801561138f5761138781601081111561138057fe5b6030612126565b915050610f9a565b6113988361299c565b9150506000805460ff19166001179055919050565b666379da05b6000081565b60095481565b6001600160a01b03166000908152600e602052604090205490565b6000805460ff1661141e576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561143061173c565b1461147b576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b50600b546000805460ff1916600117905590565b6000610ea282612acf565b600c5481565b6002805460408051602060018416156101000260001901909316849004601f81018490048402820184019092528181529291830182828015610e335780601f10610e0857610100808354040283529160200191610e33565b600080600061150684612b50565b9092509050600082600381111561151957fe5b146112b45760405162461bcd60e51b8152600401808060200182810382526037815260200180614ec96037913960400191505060405180910390fd5b60035461010090046001600160a01b031633146115a35760405162461bcd60e51b8152600401808060200182810382526024815260200180614e056024913960400191505060405180910390fd5b6009541580156115b35750600a54155b6115ee5760405162461bcd60e51b8152600401808060200182810382526023815260200180614e296023913960400191505060405180910390fd5b60078490558361162f5760405162461bcd60e51b8152600401808060200182810382526030815260200180614e4c6030913960400191505060405180910390fd5b600061163a87611166565b9050801561168f576040805162461bcd60e51b815260206004820152601a60248201527f73657474696e6720636f6d7074726f6c6c6572206661696c6564000000000000604482015290519081900360640190fd5b611697612c04565b600955670de0b6b3a7640000600a556116af86612c08565b905080156116ee5760405162461bcd60e51b8152600401808060200182810382526022815260200180614e7c6022913960400191505060405180910390fd5b8351611701906001906020870190614c72565b508251611715906002906020860190614c72565b50506003805460ff90921660ff199283161790556000805490911660011790555050505050565b600080611747612c04565b6009549091508082141561176057600092505050611012565b600061176a6128ee565b600b54600c54600a54600654604080516315f2405360e01b815260048101879052602481018690526044810185905290519596509394929391926000926001600160a01b03909216916315f24053916064808301926020929190829003018186803b1580156117d857600080fd5b505afa1580156117ec573d6000803e3d6000fd5b505050506040513d602081101561180257600080fd5b5051905065048c27395000811115611861576040805162461bcd60e51b815260206004820152601c60248201527f626f72726f772072617465206973206162737572646c79206869676800000000604482015290519081900360640190fd5b60008061186e8989612d7d565b9092509050600082600381111561188157fe5b146118d3576040805162461bcd60e51b815260206004820152601f60248201527f636f756c64206e6f742063616c63756c61746520626c6f636b2064656c746100604482015290519081900360640190fd5b6118db614c5f565b6000806000806118f960405180602001604052808a81525087612da0565b9097509450600087600381111561190c57fe5b1461193e576119296009600689600381111561192457fe5b612e08565b9e505050505050505050505050505050611012565b611948858c61289b565b9097509350600087600381111561195b57fe5b14611973576119296009600189600381111561192457fe5b61197d848c612e6e565b9097509250600087600381111561199057fe5b146119a8576119296009600489600381111561192457fe5b6119c36040518060200160405280600854815250858c612e94565b909750915060008760038111156119d657fe5b146119ee576119296009600589600381111561192457fe5b6119f9858a8b612e94565b90975090506000876003811115611a0c57fe5b14611a24576119296009600389600381111561192457fe5b60098e9055600a819055600b839055600c829055604080518d8152602081018690528082018390526060810185905290517f4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc049181900360800190a160009e50505050505050505050505050505090565b6000805460ff16611ad9576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611aef3333868661260d565b1490506000805460ff1916600117905592915050565b600a5481565b6000611b18833484612ef0565b509050611b5381604051806040016040528060168152602001751b1a5c5d5a59185d19509bdc9c9bddc819985a5b195960521b815250610bae565b505050565b6006546000906001600160a01b031663b8168816611b746128ee565b600b54600c546008546040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060206040518083038186803b158015611bc657600080fd5b505afa158015611bda573d6000803e3d6000fd5b505050506040513d6020811015611bf057600080fd5b5051905090565b6000805460ff16611c3c576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19169055611c5233858585613022565b90506000805460ff191660011790559392505050565b60035460009061010090046001600160a01b03163314611c8e5761118c60016045612126565b600480546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160006112b4565b6000805460ff16611d39576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611d4b61173c565b14611d96576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b611d9e610fb2565b90506000805460ff1916600117905590565b6001600160a01b0381166000908152600e6020526040812054819081908190818080611ddb89612b50565b935090506000816003811115611ded57fe5b14611e0b5760095b975060009650869550859450611e3e9350505050565b611e1361255e565b925090506000816003811115611e2557fe5b14611e31576009611df5565b5060009650919450925090505b9193509193565b6000610ea2826133fc565b6000610ea28261347b565b6001600160a01b039182166000908152600f6020908152604080832093909416825291909152205490565b6000611e9282346134f5565b509050610daa816040518060400160405280601881526020017f7265706179426f72726f77426568616c66206661696c65640000000000000000815250610bae565b6004546000906001600160a01b031633141580611eef575033155b15611f0757611f0060016000612126565b9050611012565b60038054600480546001600160a01b03818116610100818102610100600160a81b0319871617968790556001600160a01b031990931690935560408051948390048216808652929095041660208401528351909391927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600454604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160009250505090565b600080611fe261173c565b9050801561200857612000816010811115611ff957fe5b6040612126565b915050611152565b6112b483612c08565b6006546001600160a01b031681565b60035461010090046001600160a01b031681565b6006546000906001600160a01b03166315f240536120506128ee565b600b54600c546040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060206040518083038186803b158015611bc657600080fd5b6000805460ff166120dd576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556120ef61173c565b9050801561210d5761138781601081111561210657fe5b6046612126565b611398836135a0565b600061116134613648565b600181565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601081111561215557fe5b83605081111561216157fe5b604080519283526020830191909152600082820152519081900360600190a18260108111156112b457fe5b60055460408051634ef4c3e160e01b81523060048201526001600160a01b03858116602483015260448201859052915160009384938493911691634ef4c3e19160648082019260209290919082900301818787803b1580156121ed57600080fd5b505af1158015612201573d6000803e3d6000fd5b505050506040513d602081101561221757600080fd5b50519050801561223b5761222e6003601f83612e08565b9250600091506125579050565b612243612c04565b600954146122575761222e600a6022612126565b61225f614cf0565b61226761255e565b604083018190526020830182600381111561227e57fe5b600381111561228957fe5b90525060009050816020015160038111156122a057fe5b146122ca576122bc600960218360200151600381111561192457fe5b935060009250612557915050565b6122d486866136dc565b60c08201819052604080516020810182529083015181526122f59190613778565b606083018190526020830182600381111561230c57fe5b600381111561231757fe5b905250600090508160200151600381111561232e57fe5b14612380576040805162461bcd60e51b815260206004820181905260248201527f4d494e545f45584348414e47455f43414c43554c4154494f4e5f4641494c4544604482015290519081900360640190fd5b612390600d548260600151612e6e565b60808301819052602083018260038111156123a757fe5b60038111156123b257fe5b90525060009050816020015160038111156123c957fe5b146124055760405162461bcd60e51b8152600401808060200182810382526028815260200180614ff36028913960400191505060405180910390fd5b6001600160a01b0386166000908152600e6020526040902054606082015161242d9190612e6e565b60a083018190526020830182600381111561244457fe5b600381111561244f57fe5b905250600090508160200151600381111561246657fe5b146124a25760405162461bcd60e51b815260040180806020018281038252602b815260200180614e9e602b913960400191505060405180910390fd5b6080810151600d5560a08101516001600160a01b0387166000818152600e60209081526040918290209390935560c084015160608086015183519485529484019190915282820193909352517f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f929181900390910190a1606081015160408051918252516001600160a01b038816913091600080516020614f3a8339815191529181900360200190a360c00151600093509150505b9250929050565b600d5460009081908061257957505060075460009150612609565b60006125836128ee565b9050600061258f614c5f565b60006125a084600b54600c5461378f565b9350905060008160038111156125b257fe5b146125c7579550600094506126099350505050565b6125d183866137cd565b9250905060008160038111156125e357fe5b146125f8579550600094506126099350505050565b505160009550935061260992505050565b9091565b600554604080516317b9b84b60e31b81523060048201526001600160a01b03868116602483015285811660448301526064820185905291516000938493169163bdcdc25891608480830192602092919082900301818787803b15801561267257600080fd5b505af1158015612686573d6000803e3d6000fd5b505050506040513d602081101561269c57600080fd5b5051905080156126bb576126b36003604a83612e08565b915050612893565b836001600160a01b0316856001600160a01b031614156126e1576126b36002604b612126565b60006001600160a01b0387811690871614156127005750600019612728565b506001600160a01b038086166000908152600f60209081526040808320938a16835292905220545b6000806000806127388589612d7d565b9094509250600084600381111561274b57fe5b146127695761275c6009604b612126565b9650505050505050612893565b6001600160a01b038a166000908152600e602052604090205461278c9089612d7d565b9094509150600084600381111561279f57fe5b146127b05761275c6009604c612126565b6001600160a01b0389166000908152600e60205260409020546127d39089612e6e565b909450905060008460038111156127e657fe5b146127f75761275c6009604d612126565b6001600160a01b03808b166000908152600e6020526040808220859055918b16815220819055600019851461284f576001600160a01b03808b166000908152600f60209081526040808320938f168352929052208390555b886001600160a01b03168a6001600160a01b0316600080516020614f3a8339815191528a6040518082815260200191505060405180910390a3600096505050505050505b949350505050565b60008060006128a8614c5f565b6128b28686612da0565b909250905060008260038111156128c557fe5b146128d65750915060009050612557565b60006128e18261387d565b9350935050509250929050565b60008060006128fd4734612d7d565b9092509050600082600381111561291057fe5b1461100e57600080fd5b60008054819060ff16612961576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561297361173c565b9050801561299157610b7d81601081111561298a57fe5b6036612126565b610b9433338661388c565b600354600090819061010090046001600160a01b031633146129c45761200060016031612126565b6129cc612c04565b600954146129e057612000600a6033612126565b826129e96128ee565b10156129fb57612000600e6032612126565b600c54831115612a115761200060026034612126565b50600c5482810390811115612a575760405162461bcd60e51b815260040180806020018281038252602481526020018061504f6024913960400191505060405180910390fd5b600c819055600354612a779061010090046001600160a01b031684613bda565b600354604080516101009092046001600160a01b0316825260208201859052818101839052517f3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e916060908290030190a160006112b4565b6000805460ff16612b14576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612b2661173c565b90508015612b4457611387816010811115612b3d57fe5b6027612126565b61139833600085613c10565b6001600160a01b038116600090815260106020526040812080548291829182918291612b87575060009450849350612bff92505050565b612b978160000154600a546140d7565b90945092506000846003811115612baa57fe5b14612bbf575091935060009250612bff915050565b612bcd838260010154614116565b90945091506000846003811115612be057fe5b14612bf5575091935060009250612bff915050565b5060009450925050505b915091565b4390565b600354600090819061010090046001600160a01b03163314612c305761200060016042612126565b612c38612c04565b60095414612c4c57612000600a6041612126565b600660009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b158015612c9d57600080fd5b505afa158015612cb1573d6000803e3d6000fd5b505050506040513d6020811015612cc757600080fd5b5051612d1a576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a160006112b4565b600080838311612d94575060009050818303612557565b50600390506000612557565b6000612daa614c5f565b600080612dbb8660000151866140d7565b90925090506000826003811115612dce57fe5b14612ded57506040805160208101909152600081529092509050612557565b60408051602081019091529081526000969095509350505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0846010811115612e3757fe5b846050811115612e4357fe5b604080519283526020830191909152818101859052519081900360600190a183601081111561289357fe5b600080838301848110612e8657600092509050612557565b506002915060009050612557565b6000806000612ea1614c5f565b612eab8787612da0565b90925090506000826003811115612ebe57fe5b14612ecf5750915060009050612ee8565b612ee1612edb8261387d565b86612e6e565b9350935050505b935093915050565b60008054819060ff16612f37576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612f4961173c565b90508015612f7457612f67816010811115612f6057fe5b600f612126565b92506000915061300b9050565b836001600160a01b031663a6afed956040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612faf57600080fd5b505af1158015612fc3573d6000803e3d6000fd5b505050506040513d6020811015612fd957600080fd5b505190508015612ff957612f67816010811115612ff257fe5b6010612126565b61300533878787614141565b92509250505b6000805460ff191660011790559094909350915050565b6005546040805163d02f735160e01b81523060048201526001600160a01b038781166024830152868116604483015285811660648301526084820185905291516000938493169163d02f73519160a480830192602092919082900301818787803b15801561308f57600080fd5b505af11580156130a3573d6000803e3d6000fd5b505050506040513d60208110156130b957600080fd5b5051905080156130d0576126b36003601b83612e08565b846001600160a01b0316846001600160a01b031614156130f6576126b36006601c612126565b6130fe614d2e565b6001600160a01b0385166000908152600e60205260409020546131219085612d7d565b602083018190528282600381111561313557fe5b600381111561314057fe5b905250600090508151600381111561315457fe5b14613179576131706009601a8360000151600381111561192457fe5b92505050612893565b613198846040518060200160405280666379da05b60000815250614633565b608082018190526131aa90859061465b565b60608201526131b761255e565b60c08301819052828260038111156131cb57fe5b60038111156131d657fe5b90525060009050815160038111156131ea57fe5b1461323c576040805162461bcd60e51b815260206004820152601860248201527f65786368616e67652072617465206d617468206572726f720000000000000000604482015290519081900360640190fd5b61325c60405180602001604052808360c001518152508260800151614695565b60a08201819052600c5461326f916146b4565b60e0820152600d546080820151613286919061465b565b6101008201526001600160a01b0386166000908152600e602052604090205460608201516132b49190612e6e565b60408301819052828260038111156132c857fe5b60038111156132d357fe5b90525060009050815160038111156132e757fe5b1461330357613170600960198360000151600381111561192457fe5b60e0810151600c55610100810151600d556020808201516001600160a01b038088166000818152600e855260408082209490945583860151928b16808252908490209290925560608501518351908152925191939092600080516020614f3a833981519152929081900390910190a36080810151604080519182525130916001600160a01b03881691600080516020614f3a8339815191529181900360200190a360a081015160e082015160408051308152602081019390935282810191909152517fa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc59181900360600190a16000979650505050505050565b6000805460ff16613441576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561345361173c565b905080156134715761138781601081111561346a57fe5b6008612126565b61139833846146ea565b6000805460ff166134c0576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556134d261173c565b905080156134e957611387816010811115612b3d57fe5b61139833846000613c10565b60008054819060ff1661353c576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561354e61173c565b905080156135795761356c81601081111561356557fe5b6035612126565b92506000915061358a9050565b61358433868661388c565b92509250505b6000805460ff1916600117905590939092509050565b60035460009061010090046001600160a01b031633146135c65761118c60016047612126565b6135ce612c04565b600954146135e25761118c600a6048612126565b670de0b6b3a76400008211156135fe5761118c60026049612126565b6008805490839055604080518281526020810185905281517faaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460929181900390910190a160006112b4565b6000805460ff1661368d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561369f61173c565b905080156136bd576113878160108111156136b657fe5b604e612126565b6136c68361497e565b509150506000805460ff19166001179055919050565b6000336001600160a01b0384161461372d576040805162461bcd60e51b815260206004820152600f60248201526e0e6cadcc8cae440dad2e6dac2e8c6d608b1b604482015290519081900360640190fd5b813414613772576040805162461bcd60e51b815260206004820152600e60248201526d0ecc2d8eaca40dad2e6dac2e8c6d60931b604482015290519081900360640190fd5b50919050565b6000806000613785614c5f565b6128b28686614a66565b60008060008061379f8787612e6e565b909250905060008260038111156137b257fe5b146137c35750915060009050612ee8565b612ee18186612d7d565b60006137d7614c5f565b6000806137ec86670de0b6b3a76400006140d7565b909250905060008260038111156137ff57fe5b1461381e57506040805160208101909152600081529092509050612557565b60008061382b8388614116565b9092509050600082600381111561383e57fe5b1461386057506040805160208101909152600081529094509250612557915050565b604080516020810190915290815260009890975095505050505050565b51670de0b6b3a7640000900490565b60055460408051631200453160e11b81523060048201526001600160a01b0386811660248301528581166044830152606482018590529151600093849384939116916324008a629160848082019260209290919082900301818787803b1580156138f557600080fd5b505af1158015613909573d6000803e3d6000fd5b505050506040513d602081101561391f57600080fd5b505190508015613943576139366003603883612e08565b925060009150612ee89050565b61394b612c04565b6009541461395f57613936600a6039612126565b613967614d7b565b6001600160a01b038616600090815260106020526040902060010154606082015261399186612b50565b60808301819052602083018260038111156139a857fe5b60038111156139b357fe5b90525060009050816020015160038111156139ca57fe5b146139f4576139e6600960378360200151600381111561192457fe5b935060009250612ee8915050565b600019851415613a0d5760808101516040820152613a15565b604081018590525b613a238782604001516136dc565b60e082018190526080820151613a3891612d7d565b60a0830181905260208301826003811115613a4f57fe5b6003811115613a5a57fe5b9052506000905081602001516003811115613a7157fe5b14613aad5760405162461bcd60e51b815260040180806020018281038252603a815260200180614f00603a913960400191505060405180910390fd5b613abd600b548260e00151612d7d565b60c0830181905260208301826003811115613ad457fe5b6003811115613adf57fe5b9052506000905081602001516003811115613af657fe5b14613b325760405162461bcd60e51b8152600401808060200182810382526031815260200180614f5a6031913960400191505060405180910390fd5b60a080820180516001600160a01b03808a16600081815260106020908152604091829020948555600a5460019095019490945560c0870151600b81905560e088015195518251948f16855294840192909252828101949094526060820192909252608081019190915290517f1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1929181900390910190a160e00151600097909650945050505050565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015611b53573d6000803e3d6000fd5b6000821580613c1d575081155b613c585760405162461bcd60e51b815260040180806020018281038252603481526020018061501b6034913960400191505060405180910390fd5b613c60614cf0565b613c6861255e565b6040830181905260208301826003811115613c7f57fe5b6003811115613c8a57fe5b9052506000905081602001516003811115613ca157fe5b14613cc557613cbd6009602b8360200151600381111561192457fe5b9150506112b4565b8315613d46576060810184905260408051602081018252908201518152613cec908561289b565b6080830181905260208301826003811115613d0357fe5b6003811115613d0e57fe5b9052506000905081602001516003811115613d2557fe5b14613d4157613cbd600960298360200151600381111561192457fe5b613dbf565b613d628360405180602001604052808460400151815250613778565b6060830181905260208301826003811115613d7957fe5b6003811115613d8457fe5b9052506000905081602001516003811115613d9b57fe5b14613db757613cbd6009602a8360200151600381111561192457fe5b608081018390525b60055460608201516040805163eabe7d9160e01b81523060048201526001600160a01b03898116602483015260448201939093529051600093929092169163eabe7d919160648082019260209290919082900301818787803b158015613e2457600080fd5b505af1158015613e38573d6000803e3d6000fd5b505050506040513d6020811015613e4e57600080fd5b505190508015613e6e57613e656003602883612e08565b925050506112b4565b613e76612c04565b60095414613e8a57613e65600a602c612126565b613e9a600d548360600151612d7d565b60a0840181905260208401826003811115613eb157fe5b6003811115613ebc57fe5b9052506000905082602001516003811115613ed357fe5b14613eef57613e656009602e8460200151600381111561192457fe5b6001600160a01b0386166000908152600e60205260409020546060830151613f179190612d7d565b60c0840181905260208401826003811115613f2e57fe5b6003811115613f3957fe5b9052506000905082602001516003811115613f5057fe5b14613f6c57613e656009602d8460200151600381111561192457fe5b8160800151613f796128ee565b1015613f8b57613e65600e602f612126565b613f99868360800151613bda565b60a0820151600d5560c08201516001600160a01b0387166000818152600e6020908152604091829020939093556060850151815190815290513093600080516020614f3a833981519152928290030190a36080820151606080840151604080516001600160a01b038b168152602081019490945283810191909152517fe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a9299281900390910190a160055460808301516060840151604080516351dff98960e01b81523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916351dff98991608480830192600092919082900301818387803b1580156140ac57600080fd5b505af11580156140c0573d6000803e3d6000fd5b50600092506140cd915050565b9695505050505050565b600080836140ea57506000905080612557565b838302838582816140f757fe5b041461410b57506002915060009050612557565b600092509050612557565b6000808261412a5750600190506000612557565b600083858161413557fe5b04915091509250929050565b60055460408051632fe3f38f60e11b81523060048201526001600160a01b0384811660248301528781166044830152868116606483015260848201869052915160009384938493911691635fc7e71e9160a48082019260209290919082900301818787803b1580156141b257600080fd5b505af11580156141c6573d6000803e3d6000fd5b505050506040513d60208110156141dc57600080fd5b505190508015614200576141f36003601283612e08565b92506000915061462a9050565b614208612c04565b6009541461421c576141f3600a6016612126565b614224612c04565b846001600160a01b0316636c540baf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561425d57600080fd5b505afa158015614271573d6000803e3d6000fd5b505050506040513d602081101561428757600080fd5b50511461429a576141f3600a6011612126565b866001600160a01b0316866001600160a01b031614156142c0576141f360066017612126565b846142d1576141f360076015612126565b6000198514156142e7576141f360076014612126565b6000806142f589898961388c565b909250905081156143255761431682601081111561430f57fe5b6018612126565b94506000935061462a92505050565b6005546040805163c488847b60e01b81523060048201526001600160a01b038981166024830152604482018590528251600094859492169263c488847b926064808301939192829003018186803b15801561437f57600080fd5b505afa158015614393573d6000803e3d6000fd5b505050506040513d60408110156143a957600080fd5b508051602090910151909250905081156143f45760405162461bcd60e51b8152600401808060200182810382526033815260200180614f8b6033913960400191505060405180910390fd5b80886001600160a01b03166370a082318c6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561444b57600080fd5b505afa15801561445f573d6000803e3d6000fd5b505050506040513d602081101561447557600080fd5b505110156144ca576040805162461bcd60e51b815260206004820152601860248201527f4c49515549444154455f5345495a455f544f4f5f4d5543480000000000000000604482015290519081900360640190fd5b60006001600160a01b0389163014156144f0576144e9308d8d85613022565b905061457a565b6040805163b2a02ff160e01b81526001600160a01b038e811660048301528d81166024830152604482018590529151918b169163b2a02ff1916064808201926020929091908290030181600087803b15801561454b57600080fd5b505af115801561455f573d6000803e3d6000fd5b505050506040513d602081101561457557600080fd5b505190505b80156145c4576040805162461bcd60e51b81526020600482015260146024820152731d1bdad95b881cd95a5e9d5c994819985a5b195960621b604482015290519081900360640190fd5b604080516001600160a01b03808f168252808e1660208301528183018790528b1660608201526080810184905290517f298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb529181900360a00190a16000975092955050505050505b94509492505050565b6000670de0b6b3a764000061464c848460000151614ac5565b8161465357fe5b049392505050565b60006112b48383604051806040016040528060158152602001747375627472616374696f6e20756e646572666c6f7760581b815250614b07565b600061469f614c5f565b6146a98484614b61565b90506128938161387d565b60006112b48383604051806040016040528060118152602001706164646974696f6e206f766572666c6f7760781b815250614b8b565b6005546040805163368f515360e21b81523060048201526001600160a01b0385811660248301526044820185905291516000938493169163da3d454c91606480830192602092919082900301818787803b15801561474757600080fd5b505af115801561475b573d6000803e3d6000fd5b505050506040513d602081101561477157600080fd5b505190508015614790576147886003600e83612e08565b915050610ea2565b614798612c04565b600954146147ab57614788600a80612126565b826147b46128ee565b10156147c657614788600e6009612126565b6147ce614dc1565b6147d785612b50565b60208301819052828260038111156147eb57fe5b60038111156147f657fe5b905250600090508151600381111561480a57fe5b1461482f57614826600960078360000151600381111561192457fe5b92505050610ea2565b61483d816020015185612e6e565b604083018190528282600381111561485157fe5b600381111561485c57fe5b905250600090508151600381111561487057fe5b1461488c576148266009600c8360000151600381111561192457fe5b614898600b5485612e6e565b60608301819052828260038111156148ac57fe5b60038111156148b757fe5b90525060009050815160038111156148cb57fe5b146148e7576148266009600b8360000151600381111561192457fe5b6148f18585613bda565b604080820180516001600160a01b03881660008181526010602090815290859020928355600a54600190930192909255606080860151600b81905593518551928352928201899052818501929092529081019190915290517f13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab809181900360800190a1600095945050505050565b60008060008061498c612c04565b600954146149ab576149a0600a604f612126565b93509150612bff9050565b6149b533866136dc565b905080600c54019150600c54821015614a15576040805162461bcd60e51b815260206004820181905260248201527f61646420726573657276657320756e6578706563746564206f766572666c6f77604482015290519081900360640190fd5b600c829055604080513381526020810183905280820184905290517fa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc59181900360600190a160009350915050915091565b6000614a70614c5f565b600080614a85670de0b6b3a7640000876140d7565b90925090506000826003811115614a9857fe5b14614ab757506040805160208101909152600081529092509050612557565b6128e18186600001516137cd565b60006112b483836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f77000000000000000000815250614be9565b60008184841115614b595760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610d6b578181015183820152602001610d53565b505050900390565b614b69614c5f565b6040518060200160405280614b82856000015185614ac5565b90529392505050565b60008383018285821015614be05760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610d6b578181015183820152602001610d53565b50949350505050565b6000831580614bf6575082155b15614c03575060006112b4565b83830283858281614c1057fe5b04148390614be05760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610d6b578181015183820152602001610d53565b6040518060200160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614cb357805160ff1916838001178555614ce0565b82800160010185558215614ce0579182015b82811115614ce0578251825591602001919060010190614cc5565b50614cec929150614dea565b5090565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516101208101909152806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805161010081019091528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516080810190915280600081526020016000815260200160008152602001600081525090565b61101291905b80821115614cec5760008155600101614df056fe6f6e6c792061646d696e206d617920696e697469616c697a6520746865206d61726b65746d61726b6574206d6179206f6e6c7920626520696e697469616c697a6564206f6e6365696e697469616c2065786368616e67652072617465206d7573742062652067726561746572207468616e207a65726f2e73657474696e6720696e7465726573742072617465206d6f64656c206661696c65644d494e545f4e45575f4143434f554e545f42414c414e43455f43414c43554c4154494f4e5f4641494c4544626f72726f7742616c616e636553746f7265643a20626f72726f7742616c616e636553746f726564496e7465726e616c206661696c656452455041595f424f52524f575f4e45575f4143434f554e545f424f52524f575f42414c414e43455f43414c43554c4154494f4e5f4641494c4544ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef52455041595f424f52524f575f4e45575f544f54414c5f42414c414e43455f43414c43554c4154494f4e5f4641494c45444c49515549444154455f434f4d5054524f4c4c45525f43414c43554c4154455f414d4f554e545f5345495a455f4641494c454465786368616e67655261746553746f7265643a2065786368616e67655261746553746f726564496e7465726e616c206661696c65644d494e545f4e45575f544f54414c5f535550504c595f43414c43554c4154494f4e5f4641494c45446f6e65206f662072656465656d546f6b656e73496e206f722072656465656d416d6f756e74496e206d757374206265207a65726f72656475636520726573657276657320756e657870656374656420756e646572666c6f77a265627a7a7231582090908e0feb271b6d25392b7b700e39906630e4f91f111e647f3b57db4bd4cab164736f6c634300051000326f6e6c792061646d696e206d617920696e697469616c697a6520746865206d61726b65746d61726b6574206d6179206f6e6c7920626520696e697469616c697a6564206f6e6365696e697469616c2065786368616e67652072617465206d7573742062652067726561746572207468616e207a65726f2e73657474696e6720696e7465726573742072617465206d6f64656c206661696c6564"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8, admin_ common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_, admin_)
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

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Api *ApiCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Api *ApiSession) AccrualBlockNumber() (*big.Int, error) {
	return _Api.Contract.AccrualBlockNumber(&_Api.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Api *ApiCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Api.Contract.AccrualBlockNumber(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCallerSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Api *ApiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Api *ApiCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Api *ApiSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Api.Contract.BorrowBalanceStored(&_Api.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Api *ApiCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Api.Contract.BorrowBalanceStored(&_Api.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Api *ApiCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Api *ApiSession) BorrowIndex() (*big.Int, error) {
	return _Api.Contract.BorrowIndex(&_Api.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Api *ApiCallerSession) BorrowIndex() (*big.Int, error) {
	return _Api.Contract.BorrowIndex(&_Api.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Api *ApiCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Api *ApiSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Api.Contract.BorrowRatePerBlock(&_Api.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Api *ApiCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Api.Contract.BorrowRatePerBlock(&_Api.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Api *ApiCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Api *ApiSession) Comptroller() (common.Address, error) {
	return _Api.Contract.Comptroller(&_Api.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Api *ApiCallerSession) Comptroller() (common.Address, error) {
	return _Api.Contract.Comptroller(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCallerSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Api *ApiCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Api *ApiSession) ExchangeRateStored() (*big.Int, error) {
	return _Api.Contract.ExchangeRateStored(&_Api.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Api *ApiCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Api.Contract.ExchangeRateStored(&_Api.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetAccountSnapshot(&_Api.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Api *ApiCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Api.Contract.GetAccountSnapshot(&_Api.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Api *ApiCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Api *ApiSession) GetCash() (*big.Int, error) {
	return _Api.Contract.GetCash(&_Api.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Api *ApiCallerSession) GetCash() (*big.Int, error) {
	return _Api.Contract.GetCash(&_Api.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Api *ApiCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Api *ApiSession) InterestRateModel() (common.Address, error) {
	return _Api.Contract.InterestRateModel(&_Api.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Api *ApiCallerSession) InterestRateModel() (common.Address, error) {
	return _Api.Contract.InterestRateModel(&_Api.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Api *ApiCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Api *ApiSession) IsCToken() (bool, error) {
	return _Api.Contract.IsCToken(&_Api.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Api *ApiCallerSession) IsCToken() (bool, error) {
	return _Api.Contract.IsCToken(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Api *ApiCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Api *ApiSession) PendingAdmin() (common.Address, error) {
	return _Api.Contract.PendingAdmin(&_Api.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Api *ApiCallerSession) PendingAdmin() (common.Address, error) {
	return _Api.Contract.PendingAdmin(&_Api.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Api *ApiCaller) ProtocolSeizeShareMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "protocolSeizeShareMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Api *ApiSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _Api.Contract.ProtocolSeizeShareMantissa(&_Api.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_Api *ApiCallerSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _Api.Contract.ProtocolSeizeShareMantissa(&_Api.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Api *ApiCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Api *ApiSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Api.Contract.ReserveFactorMantissa(&_Api.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Api *ApiCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Api.Contract.ReserveFactorMantissa(&_Api.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Api *ApiCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Api *ApiSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Api.Contract.SupplyRatePerBlock(&_Api.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Api *ApiCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Api.Contract.SupplyRatePerBlock(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Api *ApiCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Api *ApiSession) TotalBorrows() (*big.Int, error) {
	return _Api.Contract.TotalBorrows(&_Api.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Api *ApiCallerSession) TotalBorrows() (*big.Int, error) {
	return _Api.Contract.TotalBorrows(&_Api.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Api *ApiCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Api *ApiSession) TotalReserves() (*big.Int, error) {
	return _Api.Contract.TotalReserves(&_Api.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Api *ApiCallerSession) TotalReserves() (*big.Int, error) {
	return _Api.Contract.TotalReserves(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Api *ApiTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Api *ApiSession) AcceptAdmin() (*types.Transaction, error) {
	return _Api.Contract.AcceptAdmin(&_Api.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Api *ApiTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Api.Contract.AcceptAdmin(&_Api.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0xfcb64147.
//
// Solidity: function _addReserves() payable returns(uint256)
func (_Api *ApiTransactor) AddReserves(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_addReserves")
}

// AddReserves is a paid mutator transaction binding the contract method 0xfcb64147.
//
// Solidity: function _addReserves() payable returns(uint256)
func (_Api *ApiSession) AddReserves() (*types.Transaction, error) {
	return _Api.Contract.AddReserves(&_Api.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0xfcb64147.
//
// Solidity: function _addReserves() payable returns(uint256)
func (_Api *ApiTransactorSession) AddReserves() (*types.Transaction, error) {
	return _Api.Contract.AddReserves(&_Api.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Api *ApiTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Api *ApiSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReduceReserves(&_Api.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Api *ApiTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReduceReserves(&_Api.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Api *ApiTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Api *ApiSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetComptroller(&_Api.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Api *ApiTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetComptroller(&_Api.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Api *ApiTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Api *ApiSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetInterestRateModel(&_Api.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Api *ApiTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetInterestRateModel(&_Api.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Api *ApiTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Api *ApiSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPendingAdmin(&_Api.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Api *ApiTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPendingAdmin(&_Api.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Api *ApiTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Api *ApiSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetReserveFactor(&_Api.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Api *ApiTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetReserveFactor(&_Api.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Api *ApiTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Api *ApiSession) AccrueInterest() (*types.Transaction, error) {
	return _Api.Contract.AccrueInterest(&_Api.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Api *ApiTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Api.Contract.AccrueInterest(&_Api.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Api *ApiTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Api *ApiSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Api.Contract.BalanceOfUnderlying(&_Api.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Api *ApiTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Api.Contract.BalanceOfUnderlying(&_Api.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Api *ApiTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Api *ApiSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Borrow(&_Api.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Api *ApiTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Borrow(&_Api.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Api *ApiTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Api *ApiSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Api.Contract.BorrowBalanceCurrent(&_Api.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Api *ApiTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Api.Contract.BorrowBalanceCurrent(&_Api.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Api *ApiTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Api *ApiSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Api.Contract.ExchangeRateCurrent(&_Api.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Api *ApiTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Api.Contract.ExchangeRateCurrent(&_Api.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Api *ApiTransactor) Initialize(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "initialize", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Api *ApiSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Api.Contract.Initialize(&_Api.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Api *ApiTransactorSession) Initialize(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Api.Contract.Initialize(&_Api.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address cTokenCollateral) payable returns()
func (_Api *ApiTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "liquidateBorrow", borrower, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address cTokenCollateral) payable returns()
func (_Api *ApiSession) LiquidateBorrow(borrower common.Address, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Api.Contract.LiquidateBorrow(&_Api.TransactOpts, borrower, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xaae40a2a.
//
// Solidity: function liquidateBorrow(address borrower, address cTokenCollateral) payable returns()
func (_Api *ApiTransactorSession) LiquidateBorrow(borrower common.Address, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Api.Contract.LiquidateBorrow(&_Api.TransactOpts, borrower, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Api *ApiTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Api *ApiSession) Mint() (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Api *ApiTransactorSession) Mint() (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Api *ApiTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Api *ApiSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Redeem(&_Api.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Api *ApiTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Redeem(&_Api.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Api *ApiTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Api *ApiSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.RedeemUnderlying(&_Api.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Api *ApiTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.RedeemUnderlying(&_Api.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Api *ApiTransactor) RepayBorrow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "repayBorrow")
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Api *ApiSession) RepayBorrow() (*types.Transaction, error) {
	return _Api.Contract.RepayBorrow(&_Api.TransactOpts)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x4e4d9fea.
//
// Solidity: function repayBorrow() payable returns()
func (_Api *ApiTransactorSession) RepayBorrow() (*types.Transaction, error) {
	return _Api.Contract.RepayBorrow(&_Api.TransactOpts)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Api *ApiTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "repayBorrowBehalf", borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Api *ApiSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _Api.Contract.RepayBorrowBehalf(&_Api.TransactOpts, borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xe5974619.
//
// Solidity: function repayBorrowBehalf(address borrower) payable returns()
func (_Api *ApiTransactorSession) RepayBorrowBehalf(borrower common.Address) (*types.Transaction, error) {
	return _Api.Contract.RepayBorrowBehalf(&_Api.TransactOpts, borrower)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Api *ApiTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Api *ApiSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Seize(&_Api.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Api *ApiTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Seize(&_Api.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Api *ApiTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Api *ApiSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Api.Contract.TotalBorrowsCurrent(&_Api.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Api *ApiTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Api.Contract.TotalBorrowsCurrent(&_Api.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Api *ApiSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Api *ApiSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, src, dst, amount)
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

// ApiAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Api contract.
type ApiAccrueInterestIterator struct {
	Event *ApiAccrueInterest // Event containing the contract specifics and raw log

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
func (it *ApiAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiAccrueInterest)
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
		it.Event = new(ApiAccrueInterest)
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
func (it *ApiAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiAccrueInterest represents a AccrueInterest event raised by the Api contract.
type ApiAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Api *ApiFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*ApiAccrueInterestIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &ApiAccrueInterestIterator{contract: _Api.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Api *ApiFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *ApiAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiAccrueInterest)
				if err := _Api.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Api *ApiFilterer) ParseAccrueInterest(log types.Log) (*ApiAccrueInterest, error) {
	event := new(ApiAccrueInterest)
	if err := _Api.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Api contract.
type ApiApprovalIterator struct {
	Event *ApiApproval // Event containing the contract specifics and raw log

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
func (it *ApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApproval)
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
		it.Event = new(ApiApproval)
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
func (it *ApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApproval represents a Approval event raised by the Api contract.
type ApiApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Api *ApiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ApiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalIterator{contract: _Api.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Api *ApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApproval)
				if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Api contract.
type ApiBorrowIterator struct {
	Event *ApiBorrow // Event containing the contract specifics and raw log

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
func (it *ApiBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiBorrow)
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
		it.Event = new(ApiBorrow)
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
func (it *ApiBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiBorrow represents a Borrow event raised by the Api contract.
type ApiBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) FilterBorrow(opts *bind.FilterOpts) (*ApiBorrowIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &ApiBorrowIterator{contract: _Api.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *ApiBorrow) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiBorrow)
				if err := _Api.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) ParseBorrow(log types.Log) (*ApiBorrow, error) {
	event := new(ApiBorrow)
	if err := _Api.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Api contract.
type ApiFailureIterator struct {
	Event *ApiFailure // Event containing the contract specifics and raw log

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
func (it *ApiFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiFailure)
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
		it.Event = new(ApiFailure)
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
func (it *ApiFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiFailure represents a Failure event raised by the Api contract.
type ApiFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Api *ApiFilterer) FilterFailure(opts *bind.FilterOpts) (*ApiFailureIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ApiFailureIterator{contract: _Api.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Api *ApiFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ApiFailure) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiFailure)
				if err := _Api.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Api *ApiFilterer) ParseFailure(log types.Log) (*ApiFailure, error) {
	event := new(ApiFailure)
	if err := _Api.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Api contract.
type ApiLiquidateBorrowIterator struct {
	Event *ApiLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *ApiLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLiquidateBorrow)
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
		it.Event = new(ApiLiquidateBorrow)
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
func (it *ApiLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLiquidateBorrow represents a LiquidateBorrow event raised by the Api contract.
type ApiLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Api *ApiFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*ApiLiquidateBorrowIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &ApiLiquidateBorrowIterator{contract: _Api.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Api *ApiFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *ApiLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLiquidateBorrow)
				if err := _Api.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Api *ApiFilterer) ParseLiquidateBorrow(log types.Log) (*ApiLiquidateBorrow, error) {
	event := new(ApiLiquidateBorrow)
	if err := _Api.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Api contract.
type ApiMintIterator struct {
	Event *ApiMint // Event containing the contract specifics and raw log

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
func (it *ApiMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMint)
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
		it.Event = new(ApiMint)
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
func (it *ApiMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMint represents a Mint event raised by the Api contract.
type ApiMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Api *ApiFilterer) FilterMint(opts *bind.FilterOpts) (*ApiMintIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &ApiMintIterator{contract: _Api.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Api *ApiFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ApiMint) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMint)
				if err := _Api.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Api *ApiFilterer) ParseMint(log types.Log) (*ApiMint, error) {
	event := new(ApiMint)
	if err := _Api.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Api contract.
type ApiNewAdminIterator struct {
	Event *ApiNewAdmin // Event containing the contract specifics and raw log

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
func (it *ApiNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiNewAdmin)
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
		it.Event = new(ApiNewAdmin)
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
func (it *ApiNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiNewAdmin represents a NewAdmin event raised by the Api contract.
type ApiNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Api *ApiFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*ApiNewAdminIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &ApiNewAdminIterator{contract: _Api.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Api *ApiFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *ApiNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiNewAdmin)
				if err := _Api.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Api *ApiFilterer) ParseNewAdmin(log types.Log) (*ApiNewAdmin, error) {
	event := new(ApiNewAdmin)
	if err := _Api.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Api contract.
type ApiNewComptrollerIterator struct {
	Event *ApiNewComptroller // Event containing the contract specifics and raw log

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
func (it *ApiNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiNewComptroller)
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
		it.Event = new(ApiNewComptroller)
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
func (it *ApiNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiNewComptroller represents a NewComptroller event raised by the Api contract.
type ApiNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Api *ApiFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*ApiNewComptrollerIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &ApiNewComptrollerIterator{contract: _Api.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Api *ApiFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *ApiNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiNewComptroller)
				if err := _Api.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Api *ApiFilterer) ParseNewComptroller(log types.Log) (*ApiNewComptroller, error) {
	event := new(ApiNewComptroller)
	if err := _Api.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Api contract.
type ApiNewMarketInterestRateModelIterator struct {
	Event *ApiNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *ApiNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiNewMarketInterestRateModel)
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
		it.Event = new(ApiNewMarketInterestRateModel)
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
func (it *ApiNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Api contract.
type ApiNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Api *ApiFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*ApiNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &ApiNewMarketInterestRateModelIterator{contract: _Api.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Api *ApiFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *ApiNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiNewMarketInterestRateModel)
				if err := _Api.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Api *ApiFilterer) ParseNewMarketInterestRateModel(log types.Log) (*ApiNewMarketInterestRateModel, error) {
	event := new(ApiNewMarketInterestRateModel)
	if err := _Api.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Api contract.
type ApiNewPendingAdminIterator struct {
	Event *ApiNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *ApiNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiNewPendingAdmin)
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
		it.Event = new(ApiNewPendingAdmin)
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
func (it *ApiNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiNewPendingAdmin represents a NewPendingAdmin event raised by the Api contract.
type ApiNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Api *ApiFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*ApiNewPendingAdminIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &ApiNewPendingAdminIterator{contract: _Api.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Api *ApiFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *ApiNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiNewPendingAdmin)
				if err := _Api.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Api *ApiFilterer) ParseNewPendingAdmin(log types.Log) (*ApiNewPendingAdmin, error) {
	event := new(ApiNewPendingAdmin)
	if err := _Api.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Api contract.
type ApiNewReserveFactorIterator struct {
	Event *ApiNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *ApiNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiNewReserveFactor)
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
		it.Event = new(ApiNewReserveFactor)
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
func (it *ApiNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiNewReserveFactor represents a NewReserveFactor event raised by the Api contract.
type ApiNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Api *ApiFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*ApiNewReserveFactorIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &ApiNewReserveFactorIterator{contract: _Api.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Api *ApiFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *ApiNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiNewReserveFactor)
				if err := _Api.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Api *ApiFilterer) ParseNewReserveFactor(log types.Log) (*ApiNewReserveFactor, error) {
	event := new(ApiNewReserveFactor)
	if err := _Api.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Api contract.
type ApiRedeemIterator struct {
	Event *ApiRedeem // Event containing the contract specifics and raw log

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
func (it *ApiRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRedeem)
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
		it.Event = new(ApiRedeem)
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
func (it *ApiRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRedeem represents a Redeem event raised by the Api contract.
type ApiRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Api *ApiFilterer) FilterRedeem(opts *bind.FilterOpts) (*ApiRedeemIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &ApiRedeemIterator{contract: _Api.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Api *ApiFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ApiRedeem) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRedeem)
				if err := _Api.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Api *ApiFilterer) ParseRedeem(log types.Log) (*ApiRedeem, error) {
	event := new(ApiRedeem)
	if err := _Api.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Api contract.
type ApiRepayBorrowIterator struct {
	Event *ApiRepayBorrow // Event containing the contract specifics and raw log

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
func (it *ApiRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRepayBorrow)
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
		it.Event = new(ApiRepayBorrow)
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
func (it *ApiRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRepayBorrow represents a RepayBorrow event raised by the Api contract.
type ApiRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*ApiRepayBorrowIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &ApiRepayBorrowIterator{contract: _Api.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *ApiRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRepayBorrow)
				if err := _Api.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Api *ApiFilterer) ParseRepayBorrow(log types.Log) (*ApiRepayBorrow, error) {
	event := new(ApiRepayBorrow)
	if err := _Api.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Api contract.
type ApiReservesAddedIterator struct {
	Event *ApiReservesAdded // Event containing the contract specifics and raw log

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
func (it *ApiReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiReservesAdded)
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
		it.Event = new(ApiReservesAdded)
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
func (it *ApiReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiReservesAdded represents a ReservesAdded event raised by the Api contract.
type ApiReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*ApiReservesAddedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &ApiReservesAddedIterator{contract: _Api.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *ApiReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiReservesAdded)
				if err := _Api.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) ParseReservesAdded(log types.Log) (*ApiReservesAdded, error) {
	event := new(ApiReservesAdded)
	if err := _Api.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Api contract.
type ApiReservesReducedIterator struct {
	Event *ApiReservesReduced // Event containing the contract specifics and raw log

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
func (it *ApiReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiReservesReduced)
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
		it.Event = new(ApiReservesReduced)
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
func (it *ApiReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiReservesReduced represents a ReservesReduced event raised by the Api contract.
type ApiReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*ApiReservesReducedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &ApiReservesReducedIterator{contract: _Api.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *ApiReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiReservesReduced)
				if err := _Api.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Api *ApiFilterer) ParseReservesReduced(log types.Log) (*ApiReservesReduced, error) {
	event := new(ApiReservesReduced)
	if err := _Api.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

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
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
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
		it.Event = new(ApiTransfer)
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
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
