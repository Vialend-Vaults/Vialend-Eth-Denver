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
const ApiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cToken1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cEth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"_quoteAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_uniPortion\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uFees1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lFees1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ErrorLogging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"GeneralA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GeneralB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"subject\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"GeneralS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyLog4\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uniAmt0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uniAmt1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cmpAmt0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cmpAmt1\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"withdrawPCT\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CEther\",\"outputs\":[{\"internalType\":\"contractIcEther\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken0\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CToken1\",\"outputs\":[{\"internalType\":\"contractIcErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"U3Fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"U3Fees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LcFees1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AccruedProtocolFees1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alloc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cHigh\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cLow\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curComp1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToken1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"doRebalance\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"baseAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"}],\"name\":\"getQuoteAtTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getSSLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ratio\",\"type\":\"uint8\"}],\"name\":\"setUniPortionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newHigh\",\"type\":\"int24\"}],\"name\":\"strategy1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cErc20Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numTokensToSupply\",\"type\":\"uint256\"}],\"name\":\"supplyErc20ToCompound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cEtherContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctoken\",\"type\":\"address\"}],\"name\":\"supplyEthToCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniPortion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"percent\",\"type\":\"uint8\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"percent\",\"type\":\"uint8\"}],\"name\":\"withdrawPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraws\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x60c0604052600180546001600160801b031916670de0b6b3a76400001790553480156200002b57600080fd5b50604051620060023803806200600283398181016040526101808110156200005257600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b01516101208c01516101408d0151610160909d01518951808b018b52600e81526d05669614c656e6420546f6b656e360941b818e019081528b51808d01909c5260048c52630564c54360e41b9d8c019d909d52600160005580519d9e9b9d999c989b979a96999598949793969295939491939092916200010291601e91906200052a565b5080516200011890601f9060208401906200052a565b506012602060006101000a81548160ff021916908360ff160217905550505033600260006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600360006101000a8154816001600160a01b0302191690836001600160a01b031602179055508b600660006101000a8154816001600160a01b0302191690836001600160a01b031602179055508b6001600160a01b0316630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015620001e657600080fd5b505afa158015620001fb573d6000803e3d6000fd5b505050506040513d60208110156200021257600080fd5b505160601b6001600160601b0319166080526040805163d21220a760e01b815290516001600160a01b038e169163d21220a7916004808301926020929190829003018186803b1580156200026557600080fd5b505afa1580156200027a573d6000803e3d6000fd5b505050506040513d60208110156200029157600080fd5b505160601b6001600160601b03191660a05260158054600160401b600160e01b031916680100000000000000006001600160a01b038d81169190910291909117909155601680546001600160a01b03199081168c841617909155601780549091168a83161790558b1662000335576040805162461bcd60e51b815260206004808301919091526024820152630ae8aa8960e31b604482015290519081900360640190fd5b8a601860006101000a8154816001600160a01b0302191690836001600160a01b03160217905550866005819055508b6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200039d57600080fd5b505afa158015620003b2573d6000803e3d6000fd5b505050506040513d6020811015620003c957600080fd5b50516006805460029290920b62ffffff16600160a01b0262ffffff60a01b19909216919091179055606487106200042d576040805162461bcd60e51b815260206004820152600360248201526228232960e91b604482015290519081900360640190fd5b601486905560158054600180546001600160801b0387166001600160801b031990911617905560ff84166701000000000000000260ff60381b1963ffffffff881663ffffffff1960028b900b62ffffff81166401000000000262ffffff60201b199096169590951716171617909155600012620004d7576040805162461bcd60e51b815260206004820152600360248201526213551160ea1b604482015290519081900360640190fd5b60008463ffffffff161162000518576040805162461bcd60e51b8152602060048201526002602482015261151160f21b604482015290519081900360640190fd5b505050505050505050505050620005d6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620005625760008555620005ad565b82601f106200057d57805160ff1916838001178555620005ad565b82800160010185558215620005ad579182015b82811115620005ad57825182559160200191906001019062000590565b50620005bb929150620005bf565b5090565b5b80821115620005bb5760008155600101620005c0565b60805160601c60a05160601c61599e62000664600039806118f75280611991528061215a528061283e52806128f15280612bdb528061313252806136955280613fee5280614154528061473652508061109a52806118d65280611956528061211d52806128b75280612b9e528061309952806136745280613f4b52806141335280614696525061599e6000f3fe6080604052600436106103855760003560e01c806370a08231116101d1578063c433c80a11610102578063d3487997116100a0578063dd62ed3e1161006f578063dd62ed3e14610d54578063f2a40db814610d8f578063f39c38a014610db9578063fa461e3314610dce5761038c565b8063d348799714610c8e578063d368867f14610d15578063d8d7f96f14610d2a578063d951021514610d3f5761038c565b8063cc5a02cb116100dc578063cc5a02cb14610bf1578063ccac7d8b14610c2d578063cda206b014610c42578063d21220a714610c795761038c565b8063c433c80a14610b97578063c8a746a614610bc7578063cbcf94aa14610bdc5761038c565b8063a9059cbb1161016f578063ab033ea911610149578063ab033ea914610b25578063ad5c464814610b58578063b0e21e8a14610b6d578063bbc001c314610b825761038c565b8063a9059cbb14610a82578063a91ef6eb14610abb578063aabfd57214610af25761038c565b806385f2aef2116101ab57806385f2aef214610a0a5780638e843c6c14610a1f57806395d89b4114610a34578063a457c2d714610a495761038c565b806370a082311461097f57806373232c60146109b2578063787dce3d146109e05761038c565b806332380717116102b657806343c57a2711610254578063624a177a11610223578063624a177a1461088c5780636253c718146108b957806366d7505b146108f05780636ea056a9146109465761038c565b806343c57a27146107a657806353257e00146107f95780635aa6e6751461080e5780635bb6aa85146108235761038c565b80633b7ba25f116102905780633b7ba25f146106d45780633cbff3fe146107175780633f3e4c111461074457806343a0d0661461076e5761038c565b8063323807171461063e57806339509351146106535780633aaa36e61461068c5761038c565b806318160ddd1161032357806323b872dd116102fd57806323b872dd146105a65780632ab4d052146105e9578063313ce567146105fe57806331dc5b95146106295761038c565b806318160ddd1461053d578063188b7ffa14610564578063238efcbc146105915761038c565b8063095ea7b31161035f578063095ea7b3146104795780630d40886d146104c65780630dfe1681146104f75780631755ff21146105285761038c565b806306b0b6561461038e57806306fdde03146103bc578063095cf5c6146104465761038c565b3661038c57005b005b34801561039a57600080fd5b506103a3610e55565b6040805192835260208301919091528051918290030190f35b3480156103c857600080fd5b506103d1610f66565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561040b5781810151838201526020016103f3565b50505050905090810190601f1680156104385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561045257600080fd5b5061038c6004803603602081101561046957600080fd5b50356001600160a01b0316610ffd565b34801561048557600080fd5b506104b26004803603604081101561049c57600080fd5b506001600160a01b03813516906020013561106b565b604080519115158252519081900360200190f35b3480156104d257600080fd5b506104db611089565b604080516001600160801b039092168252519081900360200190f35b34801561050357600080fd5b5061050c611098565b604080516001600160a01b039092168252519081900360200190f35b34801561053457600080fd5b5061050c6110bc565b34801561054957600080fd5b506105526110cb565b60408051918252519081900360200190f35b34801561057057600080fd5b5061038c6004803603602081101561058757600080fd5b503560ff166110d1565b34801561059d57600080fd5b5061038c611194565b3480156105b257600080fd5b506104b2600480360360608110156105c957600080fd5b506001600160a01b038135811691602081013590911690604001356111fb565b3480156105f557600080fd5b50610552611283565b34801561060a57600080fd5b50610613611289565b6040805160ff9092168252519081900360200190f35b34801561063557600080fd5b50610613611292565b34801561064a57600080fd5b5061050c6112a2565b34801561065f57600080fd5b506104b26004803603604081101561067657600080fd5b506001600160a01b0381351690602001356112b1565b34801561069857600080fd5b506106a16112ff565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b3480156106e057600080fd5b50610552600480360360608110156106f757600080fd5b506001600160a01b03813581169160208101359091169060400135611314565b34801561072357600080fd5b5061038c6004803603602081101561073a57600080fd5b503560020b611661565b34801561075057600080fd5b5061038c6004803603602081101561076757600080fd5b5035611721565b34801561077a57600080fd5b5061038c6004803603606081101561079157600080fd5b508035906020810135906040013515156117c2565b3480156107b257600080fd5b50610552600480360360808110156107c957600080fd5b50803560020b906001600160801b03602082013516906001600160a01b0360408201358116916060013516611aa7565b34801561080557600080fd5b5061050c611b99565b34801561081a57600080fd5b5061050c611ba8565b34801561082f57600080fd5b506108566004803603602081101561084657600080fd5b50356001600160a01b0316611bb7565b604080519687526020870195909552858501939093526060850191909152608084015260ff1660a0830152519081900360c00190f35b34801561089857600080fd5b5061038c600480360360208110156108af57600080fd5b503560ff16611bef565b3480156108c557600080fd5b5061038c600480360360408110156108dc57600080fd5b508035600290810b9160200135900b611ca0565b3480156108fc57600080fd5b5061092f6004803603604081101561091357600080fd5b5080356001600160a01b0316906020013563ffffffff16611dd2565b6040805160029290920b8252519081900360200190f35b34801561095257600080fd5b5061038c6004803603604081101561096957600080fd5b506001600160a01b0381351690602001356120cf565b34801561098b57600080fd5b50610552600480360360208110156109a257600080fd5b50356001600160a01b03166121e0565b6104b2600480360360408110156109c857600080fd5b506001600160a01b03813581169160200135166121ff565b3480156109ec57600080fd5b5061038c60048036036020811015610a0357600080fd5b50356123df565b348015610a1657600080fd5b5061050c612474565b348015610a2b57600080fd5b5061092f612483565b348015610a4057600080fd5b506103d1612493565b348015610a5557600080fd5b506104b260048036036040811015610a6c57600080fd5b506001600160a01b0381351690602001356124f3565b348015610a8e57600080fd5b506104b260048036036040811015610aa557600080fd5b506001600160a01b03813516906020013561255b565b348015610ac757600080fd5b506103a360048036036040811015610ade57600080fd5b508035600290810b9160200135900b61256f565b348015610afe57600080fd5b5061055260048036036020811015610b1557600080fd5b50356001600160a01b0316612606565b348015610b3157600080fd5b5061038c60048036036020811015610b4857600080fd5b50356001600160a01b0316612618565b348015610b6457600080fd5b5061050c612686565b348015610b7957600080fd5b50610552612695565b348015610b8e57600080fd5b5061038c61269b565b348015610ba357600080fd5b5061038c60048036036020811015610bba57600080fd5b503563ffffffff166126f2565b348015610bd357600080fd5b506105526127a4565b348015610be857600080fd5b5061050c6127aa565b348015610bfd57600080fd5b5061038c60048036036040811015610c1457600080fd5b5080356001600160a01b0316906020013560ff166127c0565b348015610c3957600080fd5b5061055261281e565b348015610c4e57600080fd5b506104db60048036036040811015610c6557600080fd5b508035600290810b9160200135900b612824565b348015610c8557600080fd5b5061050c61283c565b348015610c9a57600080fd5b5061038c60048036036060811015610cb157600080fd5b813591602081013591810190606081016040820135600160201b811115610cd757600080fd5b820183602082011115610ce957600080fd5b803590602001918460018302840111600160201b83111715610d0a57600080fd5b509092509050612860565b348015610d2157600080fd5b5061092f61291e565b348015610d3657600080fd5b5061038c61292e565b348015610d4b57600080fd5b5061038c61297c565b348015610d6057600080fd5b5061055260048036036040811015610d7757600080fd5b506001600160a01b0381358116916020013516612adf565b348015610d9b57600080fd5b5061050c60048036036020811015610db257600080fd5b5035612b0a565b348015610dc557600080fd5b5061050c612b34565b348015610dda57600080fd5b5061038c60048036036060811015610df157600080fd5b813591602081013591810190606081016040820135600160201b811115610e1757600080fd5b820183602082011115610e2957600080fd5b803590602001918460018302840111600160201b83111715610e4a57600080fd5b509092509050612b43565b600080601560089054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610eba57600080fd5b505afa158015610ece573d6000803e3d6000fd5b505050506040513d6020811015610ee457600080fd5b5051601654604080516370a0823160e01b815230600482015290519294506001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610f3457600080fd5b505afa158015610f48573d6000803e3d6000fd5b505050506040513d6020811015610f5e57600080fd5b505191929050565b601e8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610ff25780601f10610fc757610100808354040283529160200191610ff2565b820191906000526020600020905b815481529060010190602001808311610fd557829003601f168201915b505050505090505b90565b6002546001600160a01b03163314611049576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600061107f611078612c02565b8484612c06565b5060015b92915050565b6001546001600160801b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6006546001600160a01b031690565b601d5490565b60026000541415611117576040805162461bcd60e51b815260206004820152601f6024820152600080516020615740833981519152604482015290519081900360640190fd5b6002600055606460ff8216111561115b576040805162461bcd60e51b815260206004820152600360248201526270706360e81b604482015290519081900360640190fd5b6000611166336121e0565b111561118c57336000908152600d60205260409020600501805460ff191660ff83161790555b506001600055565b6004546001600160a01b031633146111e7576040805162461bcd60e51b815260206004820152601160248201527070656e64696e67476f7665726e616e636560781b604482015290519081900360640190fd5b600280546001600160a01b03191633179055565b6000611208848484612cf2565b61127884611214612c02565b61127385604051806060016040528060288152602001615868602891396001600160a01b038a166000908152601c6020526040812090611252612c02565b6001600160a01b031681526020810191909152604001600020549190612e4f565b612c06565b5060015b9392505050565b60145481565b60205460ff1690565b601554600160381b900460ff1681565b6016546001600160a01b031681565b600061107f6112be612c02565b8461127385601c60006112cf612c02565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490612ee6565b600e54600f5460105460115460125460135486565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561136357600080fd5b505afa158015611377573d6000803e3d6000fd5b505050506040513d602081101561138d57600080fd5b50518211156113cd576040805162461bcd60e51b815260206004820152600760248201526662616c616e636560c81b604482015290519081900360640190fd5b600084905060008490506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561141457600080fd5b505af1158015611428573d6000803e3d6000fd5b505050506040513d602081101561143e57600080fd5b50516040805160208101839052818152601b818301527f45786368616e6765205261746520287363616c6564207570293a200000000000606082015290519192506000805160206158fa833981519152919081900360800190a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156114d557600080fd5b505af11580156114e9573d6000803e3d6000fd5b505050506040513d60208110156114ff57600080fd5b505160408051602081018390528181526018818301527f537570706c7920526174653a20287363616c6564207570290000000000000000606082015290519192506000805160206158fa833981519152919081900360800190a1836001600160a01b031663095ea7b388886040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156115b057600080fd5b505af11580156115c4573d6000803e3d6000fd5b505050506040513d60208110156115da57600080fd5b50506040805163140e25ad60e31b81526004810188905290516000916001600160a01b0386169163a0712d689160248082019260209290919082900301818787803b15801561162857600080fd5b505af115801561163c573d6000803e3d6000fd5b505050506040513d602081101561165257600080fd5b50519998505050505050505050565b6002546001600160a01b031633146116ad576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008160020b136116f8576040805162461bcd60e51b815260206004820152601060248201526f36b0bc2a3bb0b82232bb34b0ba34b7b760811b604482015290519081900360640190fd5b6015805460029290920b62ffffff16600160201b0266ffffff0000000019909216919091179055565b60026000541415611767576040805162461bcd60e51b815260206004820152601f6024820152600080516020615740833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b031633146117b8576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6014556001600055565b60026000541415611808576040805162461bcd60e51b815260206004820152601f6024820152600080516020615740833981519152604482015290519081900360640190fd5b60026000558215158061181b5750600082115b611855576040805162461bcd60e51b815260206004808301919091526024820152630616d74360e41b604482015290519081900360640190fd5b6003546001600160a01b031633141561189a576040805162461bcd60e51b8152602060048201526002602482015261746f60f01b604482015290519081900360640190fd5b6118a333612f40565b60065460155461191b916118c8916001600160a01b039091169063ffffffff16611dd2565b6001546001600160801b03167f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611aa7565b6021556000808061192c8686612fc5565b92509250925061193a613028565b5061194361305c565b811561197e5761197e6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330856132fe565b80156119b9576119b96001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330846132fe565b6119c33384613358565b6014546119ce6110cb565b1115611a07576040805162461bcd60e51b815260206004820152600360248201526204341560ec1b604482015290519081900360640190fd5b336000908152600d6020526040902080548301815560018101805483019055600281018054840190556003018054820190558315611a5757611a52600080611a4d6110cb565b61344a565b505050505b6040805184815260208101849052808201839052905133917f36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e919081900360600190a25050600160005550505050565b600080611ab38661370c565b90506001600160801b036001600160a01b03821611611b22576001600160a01b0380821680029084811690861610611b0257611afd600160c01b876001600160801b031683613a3d565b611b1a565b611b1a81876001600160801b0316600160c01b613a3d565b925050611b90565b6000611b3c6001600160a01b03831680600160401b613a3d565b9050836001600160a01b0316856001600160a01b031610611b7457611b6f600160801b876001600160801b031683613a3d565b611b8c565b611b8c81876001600160801b0316600160801b613a3d565b9250505b50949350505050565b6017546001600160a01b031681565b6002546001600160a01b031681565b600d60205260009081526040902080546001820154600283015460038401546004850154600590950154939492939192909160ff1686565b6002546001600160a01b03163314611c3b576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60648160ff161115611c7c576040805162461bcd60e51b8152602060048201526005602482015264726174696f60d81b604482015290519081900360640190fd5b6015805460ff909216600160381b0267ff0000000000000019909216919091179055565b60026000541415611ce6576040805162461bcd60e51b815260206004820152601f6024820152600080516020615740833981519152604482015290519081900360640190fd5b60026000819055546001600160a01b03163314611d37576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b611d3f613028565b15611d4c57611d4c61297c565b6000611d566110cb565b90508015611dc857600080600080611d6f87878761344a565b6040805185815260208101859052808201849052606081018390529051949850929650909450925033917f90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b9181900360800190a2505050505b5050600160005550565b600063ffffffff8216611e12576040805162461bcd60e51b815260206004820152600360248201526207842560ec1b604482015290519081900360640190fd5b6040805160028082526060820183526000926020830190803683370190505090508281600081518110611e4157fe5b602002602001019063ffffffff16908163ffffffff1681525050600081600181518110611e6a57fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b8152600481018281528351602483015283516000936001600160a01b0389169363883bdbfd938793909283926044019185820191028083838b5b83811015611edb578181015183820152602001611ec3565b505050509050019250505060006040518083038186803b158015611efe57600080fd5b505afa158015611f12573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015611f3b57600080fd5b8101908080516040519392919084600160201b821115611f5a57600080fd5b908301906020820185811115611f6f57600080fd5b82518660208202830111600160201b82111715611f8b57600080fd5b82525081516020918201928201910280838360005b83811015611fb8578181015183820152602001611fa0565b5050505090500160405260200180516040519392919084600160201b821115611fe057600080fd5b908301906020820185811115611ff557600080fd5b82518660208202830111600160201b8211171561201157600080fd5b82525081516020918201928201910280838360005b8381101561203e578181015183820152602001612026565b5050505090500160405250505050905060008160008151811061205d57fe5b60200260200101518260018151811061207257fe5b60200260200101510390508463ffffffff168160060b8161208f57fe5b05935060008160060b1280156120b957508463ffffffff168160060b816120b257fe5b0760060b15155b156120c657600019909301925b50505092915050565b6002546001600160a01b0316331461211b576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161415801561218f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614155b6121c8576040805162461bcd60e51b81526020600482015260056024820152643a37b5b2b760d91b604482015290519081900360640190fd5b6121dc6001600160a01b0383163383613aec565b5050565b6001600160a01b0381166000908152601b60205260409020545b919050565b60008083905060008390506000816001600160a01b031663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561224757600080fd5b505af115801561225b573d6000803e3d6000fd5b505050506040513d602081101561227157600080fd5b5051604080516020810183905281815260239181018290529192506000805160206158fa833981519152918391819060608201906157a482396040019250505060405180910390a16000826001600160a01b031663ae9d70b06040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156122f657600080fd5b505af115801561230a573d6000803e3d6000fd5b505050506040513d602081101561232057600080fd5b5051604080516020818101849052828252818301527f537570706c7920526174653a20287363616c6564207570206279203165313829606082015290519192506000805160206158fa833981519152919081900360800190a1836001600160a01b0316631249c58b6203d090476040518363ffffffff1660e01b81526004016000604051808303818589803b1580156123b857600080fd5b5088f11580156123cc573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b6002546001600160a01b0316331461242b576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b601481111561246f576040805162461bcd60e51b815260206004820152600b60248201526a70726f746f636f6c46656560a81b604482015290519081900360640190fd5b600555565b6003546001600160a01b031681565b600654600160d01b900460020b81565b601f8054604080516020600260001961010060018716150201909416939093048085018490048402820184019092528181526060939092909190830182828015610ff25780601f10610fc757610100808354040283529160200191610ff2565b600061107f612500612c02565b846112738560405180606001604052806025815260200161594460259139601c600061252a612c02565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190612e4f565b600061107f612568612c02565b8484612cf2565b60008060008060006125818787613b43565b9450945050509250612594878785613bfb565b60055491965094506000906125ab90606490613ca5565b90506125d56125ce60646125c86001600160801b03871685613d02565b90613d5b565b8790612ee6565b95506125f96125f260646125c86001600160801b03861685613d02565b8690612ee6565b9450505050509250929050565b600c6020526000908152604090205481565b6002546001600160a01b03163314612664576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6018546001600160a01b031681565b60055481565b6002546001600160a01b031633146126e7576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b6126ef613028565b50565b6002546001600160a01b0316331461273e576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b60008163ffffffff1611612788576040805162461bcd60e51b815260206004820152600c60248201526b3a3bb0b8223ab930ba34b7b760a11b604482015290519081900360640190fd5b6015805463ffffffff191663ffffffff92909216919091179055565b60195481565b601554600160401b90046001600160a01b031681565b60026000541415612806576040805162461bcd60e51b815260206004820152601f6024820152600080516020615740833981519152604482015290519081900360640190fd5b60026000556128158282613dc2565b50506001600055565b601a5481565b60006128308383613b43565b50929695505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6006546001600160a01b031633146128a4576040805162461bcd60e51b8152602060048201526002602482015261505360f01b604482015290519081900360640190fd5b83156128de576128de6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163386613aec565b8215612918576129186001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163385613aec565b50505050565b600654600160b81b900460020b81565b6002546001600160a01b0316331461297a576040805162461bcd60e51b815260206004820152600a602482015269676f7665726e616e636560b01b604482015290519081900360640190fd5b565b600b5460008061298a610e55565b60065491935091506000906129b490600160b81b8104600290810b91600160d01b9004900b613b43565b5050505090508260001480156129c8575081155b80156129db57506001600160801b038116155b612a12576040805162461bcd60e51b81526020600482015260036024820152620756e360ec1b604482015290519081900360640190fd5b60005b84811015612ad8576000600d6000600b8481548110612a3057fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190206005015460ff161115612ad057612ad0600b8281548110612a7157fe5b9060005260206000200160009054906101000a90046001600160a01b0316600d6000600b8581548110612aa057fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190206005015460ff16613dc2565b600101612a15565b5050505050565b6001600160a01b039182166000908152601c6020908152604080832093909416825291909152205490565b600b8181548110612b1a57600080fd5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031681565b6006546001600160a01b03163314612b88576040805162461bcd60e51b815260206004820152600360248201526228299960e91b604482015290519081900360640190fd5b6000841315612bc557612bc56001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163386613aec565b6000831315612918576129186001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163385613aec565b3390565b6001600160a01b038316612c4b5760405162461bcd60e51b81526004018080602001828103825260248152602001806158d66024913960400191505060405180910390fd5b6001600160a01b038216612c905760405162461bcd60e51b81526004018080602001828103825260228152602001806157826022913960400191505060405180910390fd5b6001600160a01b038084166000818152601c6020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316612d375760405162461bcd60e51b81526004018080602001828103825260258152602001806158b16025913960400191505060405180910390fd5b6001600160a01b038216612d7c5760405162461bcd60e51b815260040180806020018281038252602381526020018061571d6023913960400191505060405180910390fd5b612d87838383613b3e565b612dc4816040518060600160405280602681526020016157c7602691396001600160a01b0386166000908152601b60205260409020549190612e4f565b6001600160a01b038085166000908152601b60205260408082209390935590841681522054612df39082612ee6565b6001600160a01b038084166000818152601b602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115612ede5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612ea3578181015183820152602001612e8b565b50505050905090810190601f168015612ed05780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60008282018381101561127c576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b612f49816140ee565b6126ef57600b80546001810182557f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db90180546001600160a01b0319166001600160a01b03939093169283179055546000918252600c6020908152604080842092909255600d90529020436004820155600501805460ff19169055565b60215460009083908390613004576040805162461bcd60e51b81526020600482015260016024820152600560fc1b604482015290519081900360640190fd5b6021546001546001600160801b031682028161301c57fe5b04820192509250925092565b6000806130336110cb565b905080613044576000915050610ffa565b61304c61410b565b613054614182565b600191505090565b60006130666110cb565b905080613073575061297a565b600b54604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b1580156130e057600080fd5b505afa1580156130f4573d6000803e3d6000fd5b505050506040513d602081101561310a57600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a08231916024808301926020929190829003018186803b15801561317857600080fd5b505afa15801561318c573d6000803e3d6000fd5b505050506040513d60208110156131a257600080fd5b5051905060005b83811015612ad85760006131dd600b83815481106131c357fe5b6000918252602090912001546001600160a01b03166121e0565b905080156132f55760006131f5876125c88785613d02565b90506000613207886125c88786613d02565b905060006132158383612fc5565b505090508381146132f15761324b600b868154811061323057fe5b6000918252602090912001546001600160a01b031685614311565b613276600b868154811061325b57fe5b6000918252602090912001546001600160a01b031682613358565b82600d6000600b888154811061328857fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812060020191909155600b80548492600d929091899081106132c857fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020600301555b5050505b506001016131a9565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261291890859061440d565b6001600160a01b0382166133b3576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6133bf60008383613b3e565b601d546133cc9082612ee6565b601d556001600160a01b0382166000908152601b60205260409020546133f29082612ee6565b6001600160a01b0383166000818152601b602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000806000806000806000806000600660009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b1580156134a657600080fd5b505afa1580156134ba573d6000803e3d6000fd5b505050506040513d60e08110156134d057600080fd5b50602001519050896134f45760008060008098509850985098505050505050613703565b8b60020b600014801561350a57508a60020b6000145b156135bf57600654600160d01b9004600290810b900b15801561353b5750600654600160b81b9004600290810b900b155b156135585760008060008098509850985098505050505050613703565b6006546002600160b81b8204810b600160d01b8304820b03810b81900591600160a01b9004810b9081810b90848401900b8161359057fe5b600654919005919091029c50600160a01b9004600290810b9081810b90838503900b816135b957fe5b05029c50505b6135da8c8c83600660149054906101000a900460020b6144be565b601554600160381b900460ff161561364357601554613612906064906125c890600160381b900460ff1661360c614692565b90613d02565b601554909550613635906064906125c890600160381b900460ff1661360c614732565b93506136438c8c87876147a1565b6015546064600160381b90910460ff1610156136f757613661614692565b925061366b614732565b915060006136bb7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008686614921565b9050806136f5576040805162461bcd60e51b815260206004820152600360248201526204d43760ec1b604482015290519081900360640190fd5b505b50929650909450925090505b93509350935093565b60008060008360020b12613723578260020b61372b565b8260020b6000035b9050620d89e8811115613769576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661377d57600160801b61378f565b6ffffcb933bd6fad37aa2d162d1a5940015b70ffffffffffffffffffffffffffffffffff16905060028216156137c3576ffff97272373d413259a46990580e213a0260801c5b60048216156137e2576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615613801576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615613820576fffcb9843d60f6159c9db58835c9266440260801c5b602082161561383f576fff973b41fa98c081472e6896dfb254c00260801c5b604082161561385e576fff2ea16466c96a3843ec78b326b528610260801c5b608082161561387d576ffe5dee046a99a2a811c461f1969c30530260801c5b61010082161561389d576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156138bd576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156138dd576ff3392b0822b70005940c7a398e4b70f30260801c5b6108008216156138fd576fe7159475a2c29b7443b29c7fa6e889d90260801c5b61100082161561391d576fd097f3bdfd2022b8845ad8f792aa58250260801c5b61200082161561393d576fa9f746462d870fdf8a65dc1f90e061e50260801c5b61400082161561395d576f70d869a156d2a1b890bb3df62baf32f70260801c5b61800082161561397d576f31be135f97d08fd981231505542fcfa60260801c5b6201000082161561399e576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b620200008216156139be576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156139dd576d2216e584f5fa1ea926041bedfe980260801c5b620800008216156139fa576b048a170391f7dc42444e8fa20260801c5b60008460020b1315613a15578060001981613a1157fe5b0490505b600160201b810615613a28576001613a2b565b60005b60ff16602082901c0192505050919050565b6000808060001985870986860292508281109083900303905080613a735760008411613a6857600080fd5b50829004905061127c565b808411613a7f57600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613b3e90849061440d565b505050565b600080600080600080613b573089896149f6565b6006546040805163514ea4bf60e01b81526004810184905290519293506001600160a01b039091169163514ea4bf9160248082019260a092909190829003018186803b158015613ba657600080fd5b505afa158015613bba573d6000803e3d6000fd5b505050506040513d60a0811015613bd057600080fd5b508051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6000806000600660009054906101000a90046001600160a01b03166001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b158015613c4e57600080fd5b505afa158015613c62573d6000803e3d6000fd5b505050506040513d60e0811015613c7857600080fd5b50519050613c9881613c898861370c565b613c928861370c565b87614a4c565b9250925050935093915050565b600082821115613cfc576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082613d1157506000611083565b82820282848281613d1e57fe5b041461127c5760405162461bcd60e51b81526004018080602001828103825260218152602001806158476021913960400191505060405180910390fd5b6000808211613db1576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381613dba57fe5b049392505050565b6000613dcc6110cb565b90506000613de560646125c88560ff1661360c886121e0565b905060008111613e22576040805162461bcd60e51b815260206004820152600360248201526207368360ec1b604482015290519081900360640190fd5b60008211613e5d576040805162461bcd60e51b815260206004820152600360248201526207473360ec1b604482015290519081900360640190fd5b81811115613e98576040805162461bcd60e51b815260206004820152600360248201526273683160e81b604482015290519081900360640190fd5b6000613eaa836125c88461360c614692565b90506000613ebe846125c88561360c614732565b905081158015613ecc575080155b15613f2e576040805160208101859052818152601c818301527f77697468647261772073686172653e302062757420616d74733d3d3000000000606082015290516000805160206158fa8339815191529181900360800190a1505050506121dc565b613f388684614311565b8115613fdb57613f726001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168784613aec565b6001600160a01b0386166000908152600d60205260409020600201548210613fb5576001600160a01b0386166000908152600d6020526040902060020154613fb7565b815b6001600160a01b0387166000908152600d6020526040902060020180549190910390555b801561407e576140156001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168783613aec565b6001600160a01b0386166000908152600d60205260409020600301548110614058576001600160a01b0386166000908152600d602052604090206003015461405a565b805b6001600160a01b0387166000908152600d6020526040902060030180549190910390555b614087866121e0565b6140c35761409486614ae8565b506001600160a01b0386166000908152600d602052604081208181556001810182905560028101829055600301555b5050506001600160a01b0383166000908152600d60205260409020600501805460ff19169055505050565b6001600160a01b03166000908152600c6020526040902054151590565b600080614116610e55565b9150915060008211806141295750600081115b1561417a5761417a7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008484614c0c565b6121dc614f37565b600554600954600754600091829161419991612ee6565b600a546008546141a891612ee6565b600754600854600954600a546040805194855260208501939093528383019190915260608301525192945090925030917fb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c39181900360800190a2600754600e5461421191612ee6565b600e55600854600f5461422391612ee6565b600f5560095460105461423591612ee6565b601055600a5460115461424791612ee6565b6011556000600a819055600981905560088190556007819055808415612ad85783156142905761427c60646125c88688613d02565b60125490925061428c9083612ee6565b6012555b82156142b9576142a560646125c88588613d02565b6013549091506142b59082612ee6565b6013555b60008211806142c85750600081115b15612ad85760006142d98383612fc5565b50506003549091506142f4906001600160a01b031682613358565b600354614309906001600160a01b0316612f40565b505050505050565b6001600160a01b0382166143565760405162461bcd60e51b81526004018080602001828103825260218152602001806158906021913960400191505060405180910390fd5b61436282600083613b3e565b61439f81604051806060016040528060228152602001615760602291396001600160a01b0385166000908152601b60205260409020549190612e4f565b6001600160a01b0383166000908152601b6020526040902055601d546143c59082613ca5565b601d556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6000614462826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166150769092919063ffffffff16565b805190915015613b3e5780806020019051602081101561448157600080fd5b5051613b3e5760405162461bcd60e51b815260040180806020018281038252602a81526020018061591a602a913960400191505060405180910390fd5b8260020b8460020b126144fd576040805162461bcd60e51b8152602060048201526002602482015261563160f01b604482015290519081900360640190fd5b8160020b8460020b1261453c576040805162461bcd60e51b81526020600482015260026024820152612b1960f11b604482015290519081900360640190fd5b8160020b8360020b1361457b576040805162461bcd60e51b8152602060048201526002602482015261563360f01b604482015290519081900360640190fd5b620d89e719600285900b12156145bd576040805162461bcd60e51b8152602060048201526002602482015261158d60f21b604482015290519081900360640190fd5b620d89e8600284900b13156145fe576040805162461bcd60e51b8152602060048201526002602482015261563560f01b604482015290519081900360640190fd5b8060020b8460020b8161460d57fe5b0760020b15614648576040805162461bcd60e51b81526020600482015260026024820152612b1b60f11b604482015290519081900360640190fd5b8060020b8360020b8161465757fe5b0760020b15612918576040805162461bcd60e51b8152602060048201526002602482015261563760f01b604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561470157600080fd5b505afa158015614715573d6000803e3d6000fd5b505050506040513d602081101561472b57600080fd5b5051905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561470157600080fd5b60065460408051633850c7bd60e01b815290516000926001600160a01b031691633850c7bd9160048083019260e0929190829003018186803b1580156147e657600080fd5b505afa1580156147fa573d6000803e3d6000fd5b505050506040513d60e081101561481057600080fd5b505190506000614833826148238861370c565b61482c8861370c565b878761508d565b60065460408051633c8a7d8d60e01b815230600482015260028a810b602483015289900b60448201526001600160801b038416606482015260a06084820152600060a4820181905282519495506001600160a01b0390931693633c8a7d8d9360c480840194938390030190829087803b1580156148af57600080fd5b505af11580156148c3573d6000803e3d6000fd5b505050506040513d60408110156148d957600080fd5b505060068054600297880b62ffffff908116600160b81b0262ffffff60b81b199890990b16600160d01b0262ffffff60d01b1990911617959095169590951790935550505050565b6018546015546016546019859055601a8490556000926001600160a01b0390811692600160401b9004811691811690881683141561498a5761496286615151565b601754614978906001600160a01b0316826121ff565b50614984878287611314565b506149e8565b826001600160a01b0316876001600160a01b031614156149cf576149ad85615151565b6017546149c3906001600160a01b0316836121ff565b50614984888388611314565b6149da888388611314565b506149e6878287611314565b505b506001979650505050505050565b6040805160609490941b6bffffffffffffffffffffffff1916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a9093019052815191012090565b600080836001600160a01b0316856001600160a01b03161115614a6d579293925b846001600160a01b0316866001600160a01b031611614a9857614a918585856151b8565b9150614adf565b836001600160a01b0316866001600160a01b03161015614ad157614abd8685856151b8565b9150614aca858785615221565b9050614adf565b614adc858585615221565b90505b94509492505050565b6000614af3826140ee565b614aff575060006121fa565b6001600160a01b0382166000908152600c6020526040902054600b8054600019928301928101908110614b2e57fe5b600091825260209091200154600b80546001600160a01b039092169183908110614b5457fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600101600c6000600b8481548110614b9757fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600b805480614bc757fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0385168252600c905260408120555060019050919050565b601854601554601654604080516370a0823160e01b815230600482015290516001600160a01b0394851694600160401b90940484169392831692600092908a16916370a0823191602480820192602092909190829003018186803b158015614c7357600080fd5b505afa158015614c87573d6000803e3d6000fd5b505050506040513d6020811015614c9d57600080fd5b5051604080516370a0823160e01b815230600482015290519192506000916001600160a01b038a16916370a08231916024808301926020929190829003018186803b158015614ceb57600080fd5b505afa158015614cff573d6000803e3d6000fd5b505050506040513d6020811015614d1557600080fd5b505190506001600160a01b038981169086161415614d4e57614d378785615264565b614d3f6152e2565b614d498684615264565b614d98565b846001600160a01b0316886001600160a01b03161415614d8457614d728684615264565b614d7a6152e2565b614d498785615264565b614d8e8684615264565b614d988785615264565b6000614e1d838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015614deb57600080fd5b505afa158015614dff573d6000803e3d6000fd5b505050506040513d6020811015614e1557600080fd5b505190613ca5565b90506000614e72838b6001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015614deb57600080fd5b9050600082118015614e85575060195482115b15614e9c57601954614e98908390613ca5565b6009555b600081118015614ead5750601a5481115b15614ec457601a54614ec0908290613ca5565b600a555b600954600a54604080516020810193909352828101919091526060808352600e908301526d3932b6b7bb32b632b73234b7339d60911b6080830152517fb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d9181900360a00190a15050505050505050505050565b600654600160b81b9004600290810b900b158015614f635750600654600160d01b9004600290810b900b155b15614f6d5761297a565b600654600090614f9290600160b81b8104600290810b91600160d01b9004900b613b43565b50506006549293506000928392508291508190614fc590600160b81b8104600290810b91600160d01b9004900b87615338565b93509350935093506000856001600160801b03161115614ffa57614fe98285613ca5565b600755614ff68184613ca5565b6008555b60065461501c90600160b81b8104600290810b91600160d01b9004900b613b43565b5050604080516001600160801b038516602082015281815260349181018290529398506000805160206158fa8339815191529389935091508190606082019061581382396040019250505060405180910390a15050505050565b606061508584846000856154aa565b949350505050565b6000836001600160a01b0316856001600160a01b031611156150ad579293925b846001600160a01b0316866001600160a01b0316116150d8576150d18585856155fa565b9050615148565b836001600160a01b0316866001600160a01b0316101561513a5760006150ff8786866155fa565b9050600061510e87898661565d565b9050806001600160801b0316826001600160801b03161061512f5780615131565b815b92505050615148565b61514585858461565d565b90505b95945050505050565b80156126ef5760185460408051632e1a7d4d60e01b81526004810184905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b1580156151a457600080fd5b505af1158015612ad8573d6000803e3d6000fd5b6000826001600160a01b0316846001600160a01b031611156151d8579192915b836001600160a01b0316615211606060ff16846001600160801b0316901b8686036001600160a01b0316866001600160a01b0316613a3d565b8161521857fe5b04949350505050565b6000826001600160a01b0316846001600160a01b03161115615241579192915b615085826001600160801b03168585036001600160a01b0316600160601b613a3d565b6000816001600160a01b031663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156152ac57600080fd5b505af11580156152c0573d6000803e3d6000fd5b505050506040513d60208110156152d657600080fd5b505190508015613b3e57fe5b471561297a57601860009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b1580156151a457600080fd5b60008080806001600160801b038516156153ea576006546040805163a34123a760e01b815260028a810b600483015289900b60248201526001600160801b038816604482015281516001600160a01b039093169263a34123a7926064808401939192918290030181600087803b1580156153b157600080fd5b505af11580156153c5573d6000803e3d6000fd5b505050506040513d60408110156153db57600080fd5b50805160209091015190945092505b600654604080516309e3d67b60e31b815230600482015260028a810b602483015289900b60448201526001600160801b0360648201819052608482015281516001600160a01b0390931692634f1eb3d89260a4808401939192918290030181600087803b15801561545a57600080fd5b505af115801561546e573d6000803e3d6000fd5b505050506040513d604081101561548457600080fd5b50805160209091015194989397506001600160801b039081169650909316935090915050565b6060824710156154eb5760405162461bcd60e51b81526004018080602001828103825260268152602001806157ed6026913960400191505060405180910390fd5b6154f48561569a565b615545576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b602083106155835780518252601f199092019160209182019101615564565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146155e5576040519150601f19603f3d011682016040523d82523d6000602084013e6155ea565b606091505b5091509150611b8c8282866156a0565b6000826001600160a01b0316846001600160a01b0316111561561a579192915b600061563d856001600160a01b0316856001600160a01b0316600160601b613a3d565b905061514861565884838888036001600160a01b0316613a3d565b615706565b6000826001600160a01b0316846001600160a01b0316111561567d579192915b61508561565883600160601b8787036001600160a01b0316613a3d565b3b151590565b606083156156af57508161127c565b8251156156bf5782518084602001fd5b60405162461bcd60e51b8152602060048201818152845160248401528451859391928392604401919085019080838360008315612ea3578181015183820152602001612e8b565b806001600160801b03811681146121fa57600080fdfe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345786368616e6765205261746520287363616c65642075702062792031653138293a2045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c756e69206c6971756964697479203d2049662074686973206973206e6f7420302c2074686572652077617320616e206572726f72536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573738d1cced004452bd270777a8c670f9f7e7c4fdde56f2db331fe289d39dc2624ad5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220c9fa2d193a602fa097bee9ddf17ef88ffe5cd8a3c2dc935fca7740b8e164c1a864736f6c63430007060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _pool common.Address, _weth common.Address, _cToken0 common.Address, _cToken1 common.Address, _cEth common.Address, _protocolFee *big.Int, _maxTotalSupply *big.Int, _maxTwapDeviation *big.Int, _twapDuration uint32, _quoteAmount *big.Int, _uniPortion uint8, _team common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _pool, _weth, _cToken0, _cToken1, _cEth, _protocolFee, _maxTotalSupply, _maxTwapDeviation, _twapDuration, _quoteAmount, _uniPortion, _team)
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

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint8 withdrawPCT)
func (_Api *ApiCaller) Assetholder(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit0    *big.Int
	Deposit1    *big.Int
	Current0    *big.Int
	Current1    *big.Int
	Block       *big.Int
	WithdrawPCT uint8
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Assetholder", arg0)

	outstruct := new(struct {
		Deposit0    *big.Int
		Deposit1    *big.Int
		Current0    *big.Int
		Current1    *big.Int
		Block       *big.Int
		WithdrawPCT uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deposit1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Current0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Current1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.WithdrawPCT = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint8 withdrawPCT)
func (_Api *ApiSession) Assetholder(arg0 common.Address) (struct {
	Deposit0    *big.Int
	Deposit1    *big.Int
	Current0    *big.Int
	Current1    *big.Int
	Block       *big.Int
	WithdrawPCT uint8
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint8 withdrawPCT)
func (_Api *ApiCallerSession) Assetholder(arg0 common.Address) (struct {
	Deposit0    *big.Int
	Deposit1    *big.Int
	Current0    *big.Int
	Current1    *big.Int
	Block       *big.Int
	WithdrawPCT uint8
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiCaller) CEther(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CEther")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiSession) CEther() (common.Address, error) {
	return _Api.Contract.CEther(&_Api.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x53257e00.
//
// Solidity: function CEther() view returns(address)
func (_Api *ApiCallerSession) CEther() (common.Address, error) {
	return _Api.Contract.CEther(&_Api.CallOpts)
}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiCaller) CToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CToken0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiSession) CToken0() (common.Address, error) {
	return _Api.Contract.CToken0(&_Api.CallOpts)
}

// CToken0 is a free data retrieval call binding the contract method 0xcbcf94aa.
//
// Solidity: function CToken0() view returns(address)
func (_Api *ApiCallerSession) CToken0() (common.Address, error) {
	return _Api.Contract.CToken0(&_Api.CallOpts)
}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiCaller) CToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "CToken1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiSession) CToken1() (common.Address, error) {
	return _Api.Contract.CToken1(&_Api.CallOpts)
}

// CToken1 is a free data retrieval call binding the contract method 0x32380717.
//
// Solidity: function CToken1() view returns(address)
func (_Api *ApiCallerSession) CToken1() (common.Address, error) {
	return _Api.Contract.CToken1(&_Api.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiCaller) Fees(opts *bind.CallOpts) (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Fees")

	outstruct := new(struct {
		U3Fees0              *big.Int
		U3Fees1              *big.Int
		LcFees0              *big.Int
		LcFees1              *big.Int
		AccruedProtocolFees0 *big.Int
		AccruedProtocolFees1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.U3Fees0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.U3Fees1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LcFees0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LcFees1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccruedProtocolFees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccruedProtocolFees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiSession) Fees() (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	return _Api.Contract.Fees(&_Api.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x3aaa36e6.
//
// Solidity: function Fees() view returns(uint256 U3Fees0, uint256 U3Fees1, uint256 LcFees0, uint256 LcFees1, uint256 AccruedProtocolFees0, uint256 AccruedProtocolFees1)
func (_Api *ApiCallerSession) Fees() (struct {
	U3Fees0              *big.Int
	U3Fees1              *big.Int
	LcFees0              *big.Int
	LcFees1              *big.Int
	AccruedProtocolFees0 *big.Int
	AccruedProtocolFees1 *big.Int
}, error) {
	return _Api.Contract.Fees(&_Api.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiSession) WETH() (common.Address, error) {
	return _Api.Contract.WETH(&_Api.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Api *ApiCallerSession) WETH() (common.Address, error) {
	return _Api.Contract.WETH(&_Api.CallOpts)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCaller) AccId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCallerSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
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
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiCaller) CHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "cHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiSession) CHigh() (*big.Int, error) {
	return _Api.Contract.CHigh(&_Api.CallOpts)
}

// CHigh is a free data retrieval call binding the contract method 0x8e843c6c.
//
// Solidity: function cHigh() view returns(int24)
func (_Api *ApiCallerSession) CHigh() (*big.Int, error) {
	return _Api.Contract.CHigh(&_Api.CallOpts)
}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiCaller) CLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "cLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiSession) CLow() (*big.Int, error) {
	return _Api.Contract.CLow(&_Api.CallOpts)
}

// CLow is a free data retrieval call binding the contract method 0xd368867f.
//
// Solidity: function cLow() view returns(int24)
func (_Api *ApiCallerSession) CLow() (*big.Int, error) {
	return _Api.Contract.CLow(&_Api.CallOpts)
}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiCaller) CurComp0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curComp0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiSession) CurComp0() (*big.Int, error) {
	return _Api.Contract.CurComp0(&_Api.CallOpts)
}

// CurComp0 is a free data retrieval call binding the contract method 0xc8a746a6.
//
// Solidity: function curComp0() view returns(uint256)
func (_Api *ApiCallerSession) CurComp0() (*big.Int, error) {
	return _Api.Contract.CurComp0(&_Api.CallOpts)
}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiCaller) CurComp1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "curComp1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiSession) CurComp1() (*big.Int, error) {
	return _Api.Contract.CurComp1(&_Api.CallOpts)
}

// CurComp1 is a free data retrieval call binding the contract method 0xccac7d8b.
//
// Solidity: function curComp1() view returns(uint256)
func (_Api *ApiCallerSession) CurComp1() (*big.Int, error) {
	return _Api.Contract.CurComp1(&_Api.CallOpts)
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

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiCaller) GetCAmounts(opts *bind.CallOpts) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCAmounts")

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiSession) GetCAmounts() (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Api.Contract.GetCAmounts(&_Api.CallOpts)
}

// GetCAmounts is a free data retrieval call binding the contract method 0x06b0b656.
//
// Solidity: function getCAmounts() view returns(uint256 amountA, uint256 amountB)
func (_Api *ApiCallerSession) GetCAmounts() (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Api.Contract.GetCAmounts(&_Api.CallOpts)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiCaller) GetPositionAmounts(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPositionAmounts", tickLower, tickUpper)

	outstruct := new(struct {
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Api.Contract.GetPositionAmounts(&_Api.CallOpts, tickLower, tickUpper)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_Api *ApiCallerSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Api.Contract.GetPositionAmounts(&_Api.CallOpts, tickLower, tickUpper)
}

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiCaller) GetQuoteAtTick(opts *bind.CallOpts, tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getQuoteAtTick", tick, baseAmount, baseToken, quoteToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiSession) GetQuoteAtTick(tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	return _Api.Contract.GetQuoteAtTick(&_Api.CallOpts, tick, baseAmount, baseToken, quoteToken)
}

// GetQuoteAtTick is a free data retrieval call binding the contract method 0x43c57a27.
//
// Solidity: function getQuoteAtTick(int24 tick, uint128 baseAmount, address baseToken, address quoteToken) pure returns(uint256 quoteAmount)
func (_Api *ApiCallerSession) GetQuoteAtTick(tick *big.Int, baseAmount *big.Int, baseToken common.Address, quoteToken common.Address) (*big.Int, error) {
	return _Api.Contract.GetQuoteAtTick(&_Api.CallOpts, tick, baseAmount, baseToken, quoteToken)
}

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiCaller) GetSSLiquidity(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSSLiquidity", tickLower, tickUpper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiSession) GetSSLiquidity(tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _Api.Contract.GetSSLiquidity(&_Api.CallOpts, tickLower, tickUpper)
}

// GetSSLiquidity is a free data retrieval call binding the contract method 0xcda206b0.
//
// Solidity: function getSSLiquidity(int24 tickLower, int24 tickUpper) view returns(uint128 liquidity)
func (_Api *ApiCallerSession) GetSSLiquidity(tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _Api.Contract.GetSSLiquidity(&_Api.CallOpts, tickLower, tickUpper)
}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiCaller) GetTwap(opts *bind.CallOpts, pool common.Address, period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTwap", pool, period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiSession) GetTwap(pool common.Address, period uint32) (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts, pool, period)
}

// GetTwap is a free data retrieval call binding the contract method 0x66d7505b.
//
// Solidity: function getTwap(address pool, uint32 period) view returns(int24 tick)
func (_Api *ApiCallerSession) GetTwap(pool common.Address, period uint32) (*big.Int, error) {
	return _Api.Contract.GetTwap(&_Api.CallOpts, pool, period)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiSession) Governance() (common.Address, error) {
	return _Api.Contract.Governance(&_Api.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Api *ApiCallerSession) Governance() (common.Address, error) {
	return _Api.Contract.Governance(&_Api.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiSession) MaxTotalSupply() (*big.Int, error) {
	return _Api.Contract.MaxTotalSupply(&_Api.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Api *ApiCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Api.Contract.MaxTotalSupply(&_Api.CallOpts)
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

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiSession) PendingGovernance() (common.Address, error) {
	return _Api.Contract.PendingGovernance(&_Api.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Api *ApiCallerSession) PendingGovernance() (common.Address, error) {
	return _Api.Contract.PendingGovernance(&_Api.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiCaller) PoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "poolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiSession) PoolAddress() (common.Address, error) {
	return _Api.Contract.PoolAddress(&_Api.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Api *ApiCallerSession) PoolAddress() (common.Address, error) {
	return _Api.Contract.PoolAddress(&_Api.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiSession) ProtocolFee() (*big.Int, error) {
	return _Api.Contract.ProtocolFee(&_Api.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Api *ApiCallerSession) ProtocolFee() (*big.Int, error) {
	return _Api.Contract.ProtocolFee(&_Api.CallOpts)
}

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiCaller) QuoteAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "quoteAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiSession) QuoteAmount() (*big.Int, error) {
	return _Api.Contract.QuoteAmount(&_Api.CallOpts)
}

// QuoteAmount is a free data retrieval call binding the contract method 0x0d40886d.
//
// Solidity: function quoteAmount() view returns(uint128)
func (_Api *ApiCallerSession) QuoteAmount() (*big.Int, error) {
	return _Api.Contract.QuoteAmount(&_Api.CallOpts)
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

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiCaller) Team(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "team")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiSession) Team() (common.Address, error) {
	return _Api.Contract.Team(&_Api.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_Api *ApiCallerSession) Team() (common.Address, error) {
	return _Api.Contract.Team(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCallerSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCallerSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
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

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiCaller) UniPortion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "uniPortion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiSession) UniPortion() (uint8, error) {
	return _Api.Contract.UniPortion(&_Api.CallOpts)
}

// UniPortion is a free data retrieval call binding the contract method 0x31dc5b95.
//
// Solidity: function uniPortion() view returns(uint8)
func (_Api *ApiCallerSession) UniPortion() (uint8, error) {
	return _Api.Contract.UniPortion(&_Api.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiSession) AcceptGovernance() (*types.Transaction, error) {
	return _Api.Contract.AcceptGovernance(&_Api.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Api *ApiTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Api.Contract.AcceptGovernance(&_Api.TransactOpts)
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiTransactor) Alloc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "alloc")
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiSession) Alloc() (*types.Transaction, error) {
	return _Api.Contract.Alloc(&_Api.TransactOpts)
}

// Alloc is a paid mutator transaction binding the contract method 0xbbc001c3.
//
// Solidity: function alloc() returns()
func (_Api *ApiTransactorSession) Alloc() (*types.Transaction, error) {
	return _Api.Contract.Alloc(&_Api.TransactOpts)
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

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", amountToken0, amountToken1, doRebalance)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalance)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 amountToken0, uint256 amountToken1, bool doRebalance) returns()
func (_Api *ApiTransactorSession) Deposit(amountToken0 *big.Int, amountToken1 *big.Int, doRebalance bool) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountToken0, amountToken1, doRebalance)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactor) EmergencyBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "emergencyBurn")
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactorSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetGovernance(&_Api.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Api *ApiTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetGovernance(&_Api.TransactOpts, _governance)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMaxTotalSupply", newMax)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiSession) SetMaxTotalSupply(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTotalSupply(&_Api.TransactOpts, newMax)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 newMax) returns()
func (_Api *ApiTransactorSession) SetMaxTotalSupply(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTotalSupply(&_Api.TransactOpts, newMax)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiTransactor) SetMaxTwapDeviation(opts *bind.TransactOpts, _maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMaxTwapDeviation", _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTwapDeviation(&_Api.TransactOpts, _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_Api *ApiTransactorSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetMaxTwapDeviation(&_Api.TransactOpts, _maxTwapDeviation)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiTransactor) SetProtocolFee(opts *bind.TransactOpts, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setProtocolFee", _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetProtocolFee(&_Api.TransactOpts, _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Api *ApiTransactorSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetProtocolFee(&_Api.TransactOpts, _protocolFee)
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

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiTransactor) SetTwapDuration(opts *bind.TransactOpts, _twapDuration uint32) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setTwapDuration", _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _Api.Contract.SetTwapDuration(&_Api.TransactOpts, _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_Api *ApiTransactorSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _Api.Contract.SetTwapDuration(&_Api.TransactOpts, _twapDuration)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiTransactor) SetUniPortionRatio(opts *bind.TransactOpts, ratio uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setUniPortionRatio", ratio)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiSession) SetUniPortionRatio(ratio uint8) (*types.Transaction, error) {
	return _Api.Contract.SetUniPortionRatio(&_Api.TransactOpts, ratio)
}

// SetUniPortionRatio is a paid mutator transaction binding the contract method 0x624a177a.
//
// Solidity: function setUniPortionRatio(uint8 ratio) returns()
func (_Api *ApiTransactorSession) SetUniPortionRatio(ratio uint8) (*types.Transaction, error) {
	return _Api.Contract.SetUniPortionRatio(&_Api.TransactOpts, ratio)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiTransactor) Strategy1(opts *bind.TransactOpts, newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "strategy1", newLow, newHigh)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiSession) Strategy1(newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Strategy1(&_Api.TransactOpts, newLow, newHigh)
}

// Strategy1 is a paid mutator transaction binding the contract method 0x6253c718.
//
// Solidity: function strategy1(int24 newLow, int24 newHigh) returns()
func (_Api *ApiTransactorSession) Strategy1(newLow *big.Int, newHigh *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Strategy1(&_Api.TransactOpts, newLow, newHigh)
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

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiTransactor) SupplyEthToCompound(opts *bind.TransactOpts, _cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "supplyEthToCompound", _cEtherContract, _ctoken)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiSession) SupplyEthToCompound(_cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract, _ctoken)
}

// SupplyEthToCompound is a paid mutator transaction binding the contract method 0x73232c60.
//
// Solidity: function supplyEthToCompound(address _cEtherContract, address _ctoken) payable returns(bool)
func (_Api *ApiTransactorSession) SupplyEthToCompound(_cEtherContract common.Address, _ctoken common.Address) (*types.Transaction, error) {
	return _Api.Contract.SupplyEthToCompound(&_Api.TransactOpts, _cEtherContract, _ctoken)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiTransactor) Sweep(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sweep", token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiSession) Sweep(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_Api *ApiTransactorSession) Sweep(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3MintCallback(&_Api.TransactOpts, amount0, amount1, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3SwapCallback(&_Api.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Api *ApiTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.UniswapV3SwapCallback(&_Api.TransactOpts, amount0Delta, amount1Delta, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address to, uint8 percent) returns()
func (_Api *ApiTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, percent uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdraw", to, percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address to, uint8 percent) returns()
func (_Api *ApiSession) Withdraw(to common.Address, percent uint8) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, to, percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address to, uint8 percent) returns()
func (_Api *ApiTransactorSession) Withdraw(to common.Address, percent uint8) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, to, percent)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x188b7ffa.
//
// Solidity: function withdrawPending(uint8 percent) returns()
func (_Api *ApiTransactor) WithdrawPending(opts *bind.TransactOpts, percent uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawPending", percent)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x188b7ffa.
//
// Solidity: function withdrawPending(uint8 percent) returns()
func (_Api *ApiSession) WithdrawPending(percent uint8) (*types.Transaction, error) {
	return _Api.Contract.WithdrawPending(&_Api.TransactOpts, percent)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x188b7ffa.
//
// Solidity: function withdrawPending(uint8 percent) returns()
func (_Api *ApiTransactorSession) WithdrawPending(percent uint8) (*types.Transaction, error) {
	return _Api.Contract.WithdrawPending(&_Api.TransactOpts, percent)
}

// Withdraws is a paid mutator transaction binding the contract method 0xd9510215.
//
// Solidity: function withdraws() returns()
func (_Api *ApiTransactor) Withdraws(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdraws")
}

// Withdraws is a paid mutator transaction binding the contract method 0xd9510215.
//
// Solidity: function withdraws() returns()
func (_Api *ApiSession) Withdraws() (*types.Transaction, error) {
	return _Api.Contract.Withdraws(&_Api.TransactOpts)
}

// Withdraws is a paid mutator transaction binding the contract method 0xd9510215.
//
// Solidity: function withdraws() returns()
func (_Api *ApiTransactorSession) Withdraws() (*types.Transaction, error) {
	return _Api.Contract.Withdraws(&_Api.TransactOpts)
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
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the Api contract.
type ApiCollectFeesIterator struct {
	Event *ApiCollectFees // Event containing the contract specifics and raw log

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
func (it *ApiCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCollectFees)
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
		it.Event = new(ApiCollectFees)
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
func (it *ApiCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCollectFees represents a CollectFees event raised by the Api contract.
type ApiCollectFees struct {
	Maker  common.Address
	UFees0 *big.Int
	UFees1 *big.Int
	LFees0 *big.Int
	LFees1 *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) FilterCollectFees(opts *bind.FilterOpts, maker []common.Address) (*ApiCollectFeesIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "CollectFees", makerRule)
	if err != nil {
		return nil, err
	}
	return &ApiCollectFeesIterator{contract: _Api.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *ApiCollectFees, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "CollectFees", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCollectFees)
				if err := _Api.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0xb9a5bd9c6fc68ddbb2525a58437661292930f5a4adc18d5fa410ac494a15d0c3.
//
// Solidity: event CollectFees(address indexed maker, uint256 uFees0, uint256 uFees1, uint256 lFees0, uint256 lFees1)
func (_Api *ApiFilterer) ParseCollectFees(log types.Log) (*ApiCollectFees, error) {
	event := new(ApiCollectFees)
	if err := _Api.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Api contract.
type ApiDepositIterator struct {
	Event *ApiDeposit // Event containing the contract specifics and raw log

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
func (it *ApiDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiDeposit)
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
		it.Event = new(ApiDeposit)
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
func (it *ApiDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiDeposit represents a Deposit event raised by the Api contract.
type ApiDeposit struct {
	Sender  common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ApiDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiDepositIterator{contract: _Api.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ApiDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiDeposit)
				if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseDeposit(log types.Log) (*ApiDeposit, error) {
	event := new(ApiDeposit)
	if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiErrorLoggingIterator is returned from FilterErrorLogging and is used to iterate over the raw logs and unpacked data for ErrorLogging events raised by the Api contract.
type ApiErrorLoggingIterator struct {
	Event *ApiErrorLogging // Event containing the contract specifics and raw log

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
func (it *ApiErrorLoggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiErrorLogging)
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
		it.Event = new(ApiErrorLogging)
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
func (it *ApiErrorLoggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiErrorLoggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiErrorLogging represents a ErrorLogging event raised by the Api contract.
type ApiErrorLogging struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterErrorLogging is a free log retrieval operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) FilterErrorLogging(opts *bind.FilterOpts) (*ApiErrorLoggingIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "ErrorLogging")
	if err != nil {
		return nil, err
	}
	return &ApiErrorLoggingIterator{contract: _Api.contract, event: "ErrorLogging", logs: logs, sub: sub}, nil
}

// WatchErrorLogging is a free log subscription operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) WatchErrorLogging(opts *bind.WatchOpts, sink chan<- *ApiErrorLogging) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "ErrorLogging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiErrorLogging)
				if err := _Api.contract.UnpackLog(event, "ErrorLogging", log); err != nil {
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

// ParseErrorLogging is a log parse operation binding the contract event 0xdcbe72357183be3854887085e70fd914c4c1733a94106878f48e8ba069d8fabc.
//
// Solidity: event ErrorLogging(string reason)
func (_Api *ApiFilterer) ParseErrorLogging(log types.Log) (*ApiErrorLogging, error) {
	event := new(ApiErrorLogging)
	if err := _Api.contract.UnpackLog(event, "ErrorLogging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralAIterator is returned from FilterGeneralA and is used to iterate over the raw logs and unpacked data for GeneralA events raised by the Api contract.
type ApiGeneralAIterator struct {
	Event *ApiGeneralA // Event containing the contract specifics and raw log

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
func (it *ApiGeneralAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralA)
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
		it.Event = new(ApiGeneralA)
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
func (it *ApiGeneralAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralA represents a GeneralA event raised by the Api contract.
type ApiGeneralA struct {
	Subject string
	Value   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralA is a free log retrieval operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) FilterGeneralA(opts *bind.FilterOpts) (*ApiGeneralAIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralA")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralAIterator{contract: _Api.contract, event: "GeneralA", logs: logs, sub: sub}, nil
}

// WatchGeneralA is a free log subscription operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) WatchGeneralA(opts *bind.WatchOpts, sink chan<- *ApiGeneralA) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralA)
				if err := _Api.contract.UnpackLog(event, "GeneralA", log); err != nil {
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

// ParseGeneralA is a log parse operation binding the contract event 0x3ab002a7eebd5665644b0637a05a350d75a661095d978c8094688b8913cb9c39.
//
// Solidity: event GeneralA(string subject, address value)
func (_Api *ApiFilterer) ParseGeneralA(log types.Log) (*ApiGeneralA, error) {
	event := new(ApiGeneralA)
	if err := _Api.contract.UnpackLog(event, "GeneralA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralBIterator is returned from FilterGeneralB and is used to iterate over the raw logs and unpacked data for GeneralB events raised by the Api contract.
type ApiGeneralBIterator struct {
	Event *ApiGeneralB // Event containing the contract specifics and raw log

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
func (it *ApiGeneralBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralB)
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
		it.Event = new(ApiGeneralB)
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
func (it *ApiGeneralBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralB represents a GeneralB event raised by the Api contract.
type ApiGeneralB struct {
	Subject string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralB is a free log retrieval operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) FilterGeneralB(opts *bind.FilterOpts) (*ApiGeneralBIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralB")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralBIterator{contract: _Api.contract, event: "GeneralB", logs: logs, sub: sub}, nil
}

// WatchGeneralB is a free log subscription operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) WatchGeneralB(opts *bind.WatchOpts, sink chan<- *ApiGeneralB) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralB")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralB)
				if err := _Api.contract.UnpackLog(event, "GeneralB", log); err != nil {
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

// ParseGeneralB is a log parse operation binding the contract event 0xb6bc63617433f8cf440bcc140feaaee6ef7d5796bef15eddbb581271d94754b2.
//
// Solidity: event GeneralB(string subject, uint256 value)
func (_Api *ApiFilterer) ParseGeneralB(log types.Log) (*ApiGeneralB, error) {
	event := new(ApiGeneralB)
	if err := _Api.contract.UnpackLog(event, "GeneralB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGeneralSIterator is returned from FilterGeneralS and is used to iterate over the raw logs and unpacked data for GeneralS events raised by the Api contract.
type ApiGeneralSIterator struct {
	Event *ApiGeneralS // Event containing the contract specifics and raw log

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
func (it *ApiGeneralSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGeneralS)
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
		it.Event = new(ApiGeneralS)
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
func (it *ApiGeneralSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGeneralSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGeneralS represents a GeneralS event raised by the Api contract.
type ApiGeneralS struct {
	Subject string
	Value   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGeneralS is a free log retrieval operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) FilterGeneralS(opts *bind.FilterOpts) (*ApiGeneralSIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "GeneralS")
	if err != nil {
		return nil, err
	}
	return &ApiGeneralSIterator{contract: _Api.contract, event: "GeneralS", logs: logs, sub: sub}, nil
}

// WatchGeneralS is a free log subscription operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) WatchGeneralS(opts *bind.WatchOpts, sink chan<- *ApiGeneralS) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "GeneralS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGeneralS)
				if err := _Api.contract.UnpackLog(event, "GeneralS", log); err != nil {
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

// ParseGeneralS is a log parse operation binding the contract event 0x461a711dabbb89f674949025eaa8bc10b8f56e252fe8cdbd4d8eeb3069f018b8.
//
// Solidity: event GeneralS(string subject, string value)
func (_Api *ApiFilterer) ParseGeneralS(log types.Log) (*ApiGeneralS, error) {
	event := new(ApiGeneralS)
	if err := _Api.contract.UnpackLog(event, "GeneralS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ApiMyLog2Iterator is returned from FilterMyLog2 and is used to iterate over the raw logs and unpacked data for MyLog2 events raised by the Api contract.
type ApiMyLog2Iterator struct {
	Event *ApiMyLog2 // Event containing the contract specifics and raw log

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
func (it *ApiMyLog2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog2)
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
		it.Event = new(ApiMyLog2)
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
func (it *ApiMyLog2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLog2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog2 represents a MyLog2 event raised by the Api contract.
type ApiMyLog2 struct {
	Arg0 string
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog2 is a free log retrieval operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) FilterMyLog2(opts *bind.FilterOpts) (*ApiMyLog2Iterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog2")
	if err != nil {
		return nil, err
	}
	return &ApiMyLog2Iterator{contract: _Api.contract, event: "MyLog2", logs: logs, sub: sub}, nil
}

// WatchMyLog2 is a free log subscription operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) WatchMyLog2(opts *bind.WatchOpts, sink chan<- *ApiMyLog2) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog2)
				if err := _Api.contract.UnpackLog(event, "MyLog2", log); err != nil {
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

// ParseMyLog2 is a log parse operation binding the contract event 0xb626a5b8d37253edd986272eff7a2541a1538cad44624e413d51713f0b2a414d.
//
// Solidity: event MyLog2(string arg0, uint256 arg1, uint256 arg2)
func (_Api *ApiFilterer) ParseMyLog2(log types.Log) (*ApiMyLog2, error) {
	event := new(ApiMyLog2)
	if err := _Api.contract.UnpackLog(event, "MyLog2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMyLog4Iterator is returned from FilterMyLog4 and is used to iterate over the raw logs and unpacked data for MyLog4 events raised by the Api contract.
type ApiMyLog4Iterator struct {
	Event *ApiMyLog4 // Event containing the contract specifics and raw log

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
func (it *ApiMyLog4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMyLog4)
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
		it.Event = new(ApiMyLog4)
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
func (it *ApiMyLog4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMyLog4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMyLog4 represents a MyLog4 event raised by the Api contract.
type ApiMyLog4 struct {
	Arg0 string
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 *big.Int
	Arg4 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyLog4 is a free log retrieval operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) FilterMyLog4(opts *bind.FilterOpts) (*ApiMyLog4Iterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MyLog4")
	if err != nil {
		return nil, err
	}
	return &ApiMyLog4Iterator{contract: _Api.contract, event: "MyLog4", logs: logs, sub: sub}, nil
}

// WatchMyLog4 is a free log subscription operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) WatchMyLog4(opts *bind.WatchOpts, sink chan<- *ApiMyLog4) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MyLog4")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMyLog4)
				if err := _Api.contract.UnpackLog(event, "MyLog4", log); err != nil {
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

// ParseMyLog4 is a log parse operation binding the contract event 0xa20514a72e3ef3eedfb172365769e68cd0eb83df0ef135feb5394a31045f7528.
//
// Solidity: event MyLog4(string arg0, uint256 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_Api *ApiFilterer) ParseMyLog4(log types.Log) (*ApiMyLog4, error) {
	event := new(ApiMyLog4)
	if err := _Api.contract.UnpackLog(event, "MyLog4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the Api contract.
type ApiRebalanceIterator struct {
	Event *ApiRebalance // Event containing the contract specifics and raw log

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
func (it *ApiRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRebalance)
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
		it.Event = new(ApiRebalance)
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
func (it *ApiRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRebalance represents a Rebalance event raised by the Api contract.
type ApiRebalance struct {
	Sender  common.Address
	UniAmt0 *big.Int
	UniAmt1 *big.Int
	CmpAmt0 *big.Int
	CmpAmt1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sender, uint256 uniAmt0, uint256 uniAmt1, uint256 cmpAmt0, uint256 cmpAmt1)
func (_Api *ApiFilterer) FilterRebalance(opts *bind.FilterOpts, sender []common.Address) (*ApiRebalanceIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Rebalance", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiRebalanceIterator{contract: _Api.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sender, uint256 uniAmt0, uint256 uniAmt1, uint256 cmpAmt0, uint256 cmpAmt1)
func (_Api *ApiFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *ApiRebalance, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Rebalance", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRebalance)
				if err := _Api.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sender, uint256 uniAmt0, uint256 uniAmt1, uint256 cmpAmt0, uint256 cmpAmt1)
func (_Api *ApiFilterer) ParseRebalance(log types.Log) (*ApiRebalance, error) {
	event := new(ApiRebalance)
	if err := _Api.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the Api contract.
type ApiSnapshotIterator struct {
	Event *ApiSnapshot // Event containing the contract specifics and raw log

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
func (it *ApiSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiSnapshot)
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
		it.Event = new(ApiSnapshot)
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
func (it *ApiSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiSnapshot represents a Snapshot event raised by the Api contract.
type ApiSnapshot struct {
	Tick         *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) FilterSnapshot(opts *bind.FilterOpts) (*ApiSnapshotIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &ApiSnapshotIterator{contract: _Api.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *ApiSnapshot) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiSnapshot)
				if err := _Api.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_Api *ApiFilterer) ParseSnapshot(log types.Log) (*ApiSnapshot, error) {
	event := new(ApiSnapshot)
	if err := _Api.contract.UnpackLog(event, "Snapshot", log); err != nil {
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
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Api contract.
type ApiWithdrawIterator struct {
	Event *ApiWithdraw // Event containing the contract specifics and raw log

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
func (it *ApiWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWithdraw)
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
		it.Event = new(ApiWithdraw)
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
func (it *ApiWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWithdraw represents a Withdraw event raised by the Api contract.
type ApiWithdraw struct {
	Sender  common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*ApiWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawIterator{contract: _Api.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ApiWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWithdraw)
				if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseWithdraw(log types.Log) (*ApiWithdraw, error) {
	event := new(ApiWithdraw)
	if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
