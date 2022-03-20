package include

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	factory "viaroot/deploy/uniswap/v3/deploy/UniswapV3Factory"
	pool "viaroot/deploy/uniswap/v3/deploy/UniswapV3Pool"

	token "viaroot/deploy/Tokens/erc20/deploy/Token"

	swapCallee "viaroot/deploy/TestUniswapV3Callee"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	/*

		factory "../../uniswap/v3/deploy/UniswapV3Factory"
		pool "../../uniswap/v3/deploy/UniswapV3Pool"
		token "../../uniswap/v3/deploy/token"
	*/)

func CreatePool(do int) common.Address {
	if do <= 0 {
		return common.HexToAddress("0x0")
	}

	myPrintln("----------------------------------------------")
	myPrintln(".......................Create Pool. ..................")
	myPrintln("----------------------------------------------")

	factoryAddress := common.HexToAddress(Network.Factory)

	instance, err := factory.NewApi(factoryAddress, EthClient)
	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(Network.TokenA)
	tokenB := common.HexToAddress(Network.TokenB)
	fee := big.NewInt(Network.FeeTier)
	NonceGen()
	tx, err := instance.CreatePool(Auth,
		tokenA,
		tokenB,
		fee,
	)

	if err != nil {
		log.Fatal("createPool ", err)
	}

	Readstring("createpool done, wait for pending ... next getPool... ")

	//get the transaction hash
	_ = tx // reserve
	//myPrintln("createpool tx sent: %s", tx.Hash().Hex())

	poolAddress, err := instance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
	if err != nil {
		log.Fatal("getpool ", err)
	}

	Network.Pool = poolAddress.String()
	myPrintln("poolAddress:", poolAddress)

	AddSettingString("pool address:", poolAddress.String())
	AddSettingString("pool fee tier:", fee.String())

	Readstring("createpool done, wait for pending ... next... ")

	return poolAddress

}

func checkBalance(amt *big.Int, _token string, _owner string) bool {
	symbol, bal := GetBalance(_token, _owner)
	if amt.Cmp(bal) > 0 {
		fmt.Println("*Warning* There is not enough fund for ", symbol, ".  balance in wallet:", bal, ", fund required: ", amt)
		return false
	}

	return true

}

// use exact0
func Swap0(accountId int, swapAmount *big.Int, zeroForOne bool, _pool string) {

	myPrintln("----------------------------------------------")
	myPrintln(".......................swap0. ..................")
	myPrintln("----------------------------------------------")

	poolAddress := common.HexToAddress(_pool)
	calleeInstance, err := swapCallee.NewApi(common.HexToAddress(Network.Callee), EthClient)

	if err != nil {
		log.Fatal(err)
	}

	myPrintln("pool address:", poolAddress)
	myPrintln("callee address:", Network.Callee)

	poolInstance, err := pool.NewApi(poolAddress, EthClient)
	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})
	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	// myPrintln("TokenA address:", TokenA.String())
	// myPrintln("TokenB address:", TokenB.String())

	accountAddrss := GetAddress(accountId).String()
	myPrintln("acccount address:", accountAddrss)
	fundsEnough := true
	if zeroForOne {
		fundsEnough = checkBalance(swapAmount, TokenA.String(), accountAddrss)
	} else {
		fundsEnough = checkBalance(swapAmount, TokenB.String(), accountAddrss)
	}

	if !fundsEnough {
		return
	}

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ChangeAccount(accountId)
	myPrintln("from address:", FromAddress)

	ApproveToken(TokenA, maxToken0, Network.Callee)
	ApproveToken(TokenB, maxToken1, Network.Callee)

	recipient := FromAddress

	NonceGen()

	MIN_SQRT_RATIO := big.NewInt(4295128739)
	MAX_SQRT_RATIO, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)
	sqrtP0for1 := new(big.Int).Add(MIN_SQRT_RATIO, big.NewInt(1))
	sqrtP1for0 := new(big.Int).Sub(MAX_SQRT_RATIO, big.NewInt(1))

	amountIn := swapAmount.Mul(swapAmount, PowX(int64(1), int(Token[0].Decimals)))

	if zeroForOne {

		tx, err := calleeInstance.SwapExact0For1(Auth,
			poolAddress,
			amountIn,
			recipient,
			sqrtP0for1)

		myPrintln(">> SwapExact0For1 ", swapAmount)
		myPrintln("zeroForOne =", zeroForOne, swapAmount, sqrtP0for1, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		TxConfirm(tx.Hash())

	} else {

		tx, err := calleeInstance.Swap1ForExact0(Auth,
			poolAddress,
			amountIn,
			recipient,
			sqrtP1for0)

		myPrintln(">> Swap1ForExact0 ", swapAmount)
		myPrintln("zeroForOne =", zeroForOne, swapAmount, sqrtP1for0, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		TxConfirm(tx.Hash())
	}

	PrintPrice()

	ChangeAccount(Account)
}

//use exact1
func Swap1(accountId int, swapAmount *big.Int, zeroForOne bool, _pool string) {

	myPrintln("----------------------------------------------")
	myPrintln(".......................swap1. ..................")
	myPrintln("----------------------------------------------")

	poolAddress := common.HexToAddress(_pool)
	calleeInstance, err := swapCallee.NewApi(common.HexToAddress(Network.Callee), EthClient)

	if err != nil {
		log.Fatal(err)
	}

	myPrintln("pool address:", poolAddress)
	myPrintln("callee address:", Network.Callee)

	poolInstance, err := pool.NewApi(poolAddress, EthClient)
	if err != nil {
		log.Fatal("poolInstance err:", err)
	}

	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})
	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	accountAddrss := GetAddress(accountId).String()
	myPrintln("acccount address:", accountAddrss)
	fundsEnough := true
	if zeroForOne {
		fundsEnough = checkBalance(swapAmount, TokenA.String(), accountAddrss)
	} else {
		fundsEnough = checkBalance(swapAmount, TokenB.String(), accountAddrss)
	}

	if !fundsEnough {
		return
	}

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ChangeAccount(accountId)
	myPrintln("from address:", FromAddress)

	ApproveToken(TokenA, maxToken0, Network.Callee)
	ApproveToken(TokenB, maxToken1, Network.Callee)

	recipient := FromAddress

	NonceGen()

	MIN_SQRT_RATIO := big.NewInt(4295128739)
	MAX_SQRT_RATIO, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)

	sqrtP0for1 := new(big.Int).Add(MIN_SQRT_RATIO, big.NewInt(1))
	sqrtP1for0 := new(big.Int).Sub(MAX_SQRT_RATIO, big.NewInt(1))

	amountIn := swapAmount.Mul(swapAmount, PowX(int64(1), int(Token[1].Decimals)))

	if zeroForOne {

		tx, err := calleeInstance.Swap0ForExact1(Auth,
			poolAddress,
			amountIn,
			recipient,
			sqrtP0for1)

		myPrintln(">> swap0ForExact1 ", swapAmount)
		myPrintln("zeroForOne =", zeroForOne, swapAmount, sqrtP0for1, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		TxConfirm(tx.Hash())

	} else {

		tx, err := calleeInstance.SwapExact1For0(Auth,
			poolAddress,
			amountIn,
			recipient,
			sqrtP1for0)

		myPrintln(">> swapExact1For0 ", swapAmount)
		myPrintln("zeroForOne =", zeroForOne, swapAmount, sqrtP1for0, slot0.SqrtPriceX96, slot0.Tick)

		if err != nil {
			panic(err)
		}

		TxConfirm(tx.Hash())
	}

	PrintPrice()

	ChangeAccount(Account)
}

func InitialPool(do int) {

	if do <= 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Initialize Pool ..............")
	myPrintln("----------------------------------------------")

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	myPrintln("pool address:", common.HexToAddress(Network.Pool))

	poolInstance, err := pool.NewApi(common.HexToAddress(Network.Pool), EthClient)

	if err != nil {
		log.Fatal(err)
	}

	token0, err := poolInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	_, _, symbol, _, _ := GetTokenInstance(token0.String())
	//
	myPrintln("token0 is ", symbol)

	//# Set ETH/USDC price to 2000 . tokenA eth

	var price float64

	if token0 == common.HexToAddress(Network.TokenA) {
		price = 3000.0
	} else if token0 == common.HexToAddress(Network.TokenB) {
		price = (1.0 / 3000.0) //  1 eth = 2000usdc
	} else {
		log.Fatal("wrong token address ", token0, Network.TokenA, Network.TokenB)
	}

	myPrintln("price:", price)

	sqrtPriceX96 := getSqrtPriceX96(price)

	bigIntSqrtPX96 := Float64ToBigInt(sqrtPriceX96)

	myPrintln("bigInt . sqrtPriceX96:", bigIntSqrtPX96)

	//os.Exit(3)

	NonceGen()

	tx, err := poolInstance.Initialize(Auth, bigIntSqrtPX96)
	if err != nil {
		log.Fatal("initialize err:", err)
	}
	myPrintln("initial tx sent:", tx.Hash().Hex())

	Readstring("initial pool done, waiting for pending... next ... ")
	/*
		tx, err = poolInstance.IncreaseObservationCardinalityNext(Auth, 100)

		myPrintln("IncreaseObservationCardinalityNext sent:", tx.Hash().Hex())

		Readstring(" waiting for pending... next ... ")
	*/
	PoolInfo(Network.Pool)

}

// use current Network
func FindPool() common.Address {

	myPrintln("----------------------------------------------")
	myPrintln(".......................get Pool by presetting tokens and fee tier ..................")
	myPrintln("----------------------------------------------")

	tokenA := Network.TokenA
	tokenB := Network.TokenB
	fee := Network.FeeTier

	poolAddress := GetPool(tokenA, tokenB, fee)

	return poolAddress

}

// find pool by given parameters
func GetPool(token0 string, token1 string, feetier int64) common.Address {

	myPrintln("----------------------------------------------")
	myPrintln(".......................Get Pool ..................")
	myPrintln("----------------------------------------------")

	factoryAddress := common.HexToAddress(Network.Factory)

	factoryInstance, err := factory.NewApi(factoryAddress, EthClient)

	if err != nil {
		log.Fatal("factory.NewApi ", err)
	}

	tokenA := common.HexToAddress(token0)
	tokenB := common.HexToAddress(token1)
	fee := big.NewInt(feetier)

	poolAddress, err := factoryInstance.GetPool(&bind.CallOpts{}, tokenA, tokenB, fee)
	if err != nil {
		log.Fatal("getpool ", err)
	}

	Network.Pool = poolAddress.String()

	poolInstance := GetPoolInstance1(poolAddress.Hex())

	pooltoken0, _ := poolInstance.Token0(&bind.CallOpts{})
	pooltoken1, _ := poolInstance.Token1(&bind.CallOpts{})

	_, _, token0symbol, _, _ := GetTokenInstance(pooltoken0.Hex())
	_, _, token1symbol, _, _ := GetTokenInstance(pooltoken1.Hex())

	myPrintln("pool token0 - ", token0symbol, "  ", pooltoken0.Hex())
	myPrintln("pool token1 - ", token1symbol, "  ", pooltoken1.Hex())
	myPrintln("fee tier:", fee)

	myPrintln("pool address :", poolAddress)

	return poolAddress

}

func PoolInfo(_pool string) {

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	myPrintln("----------------------------------------------")
	myPrintln(".............. Pool Info ....... ")
	myPrintln("----------------------------------------------")

	myPrintln("pool address:", common.HexToAddress(_pool))

	poolInstance, err := pool.NewApi(common.HexToAddress(_pool), EthClient)

	if err != nil {
		log.Fatal("NEWAPI ERROR:", err)
	}

	token0, _ := poolInstance.Token0(&bind.CallOpts{})
	token1, _ := poolInstance.Token1(&bind.CallOpts{})
	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})

	myPrintln("Token0 address:", token0)
	myPrintln("Token1 address:", token1)
	myPrintln("Fee tier:", feeRate)

	_, name, symbol, decimals0, _ := GetTokenInstance(token0.String())
	myPrintln("Token0 info:", name, symbol, decimals0)

	_, name, symbol, decimals1, _ := GetTokenInstance(token1.String())
	myPrintln("Token1 info:", name, symbol, decimals0)

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	myPrintln("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	myPrintln("slot0.Tick:", slot0.Tick)

	priceBigInt, priceReadable := getPrice2(slot0.SqrtPriceX96, slot0.Tick, int(decimals0), int(decimals1))
	myPrintln("Price : ", priceBigInt, priceReadable)

	liquidity, err := poolInstance.Liquidity(&bind.CallOpts{})
	if err != nil {
		log.Fatal("liquidity : ", err)
	}
	//myPrintln("liquidity in pool:", Pricef(liquidity, int(1e18)))
	myPrintln("liquidity in pool:", liquidity)

	// liquidity0, symbol0 := TokenInfo(common.HexToAddress(Network.TokenA), common.HexToAddress(Network.Pool))
	// liquidity1, symbol1 := TokenInfo(common.HexToAddress(Network.TokenB), common.HexToAddress(Network.Pool))

	// myPrintln(symbol0, " in pool:", liquidity0)
	// myPrintln(symbol1, " in pool:", liquidity1)

}

func MintPool(liquidity int64, amount0 int64, amount1 int64) {

	myPrintln("----------------------------------------------")
	myPrintln(".........MintPool initil Pool from admin account.........  ")
	myPrintln("----------------------------------------------")

	//# Add some liquidity
	max_tick := int64(887272/60) * 60 // // 60 * 60
	min_tick := max_tick * -1
	_ = min_tick

	Auth = GetSignature(Networkid, 0)

	token0Instance, err := token.NewApi(common.HexToAddress(Network.TokenA), EthClient)
	if err != nil {
		log.Fatal("token0Instance,", err)
	}
	var maxToken0, _ = new(big.Int).SetString("900000000000000000000000000000", 10)

	NonceGen()
	_, err = token0Instance.Approve(Auth, common.HexToAddress(Network.Callee), maxToken0)

	if err != nil {
		log.Fatal("token0 approve callee err, ", err)
	}

	token1Instance, err := token.NewApi(common.HexToAddress(Network.TokenB), EthClient)
	if err != nil {
		log.Fatal("token1Instance,", err)
	}
	var maxToken1, _ = new(big.Int).SetString("900000000000000000000000000000", 10)

	NonceGen()
	_, err = token1Instance.Approve(Auth, common.HexToAddress(Network.Callee), maxToken1)
	if err != nil {
		log.Fatal("token1 approve callee err ", err)
	}

	Readstring("Approves sent.....  wait for pending..next .. ")

	balance0, _ := token0Instance.BalanceOf(&bind.CallOpts{}, FromAddress)
	balance1, _ := token1Instance.BalanceOf(&bind.CallOpts{}, FromAddress)

	myPrintln("balance0: ", balance0)
	myPrintln("balance1: ", balance1)
	//os.Exit(1)

	calleeInstance := GetCalleeInstance()

	myPrintln("max_tick, min_tick:", max_tick, min_tick)

	NonceGen()

	//uint128 amount

	var liquidityAmt *big.Int

	liquidityAmt = X1E18(liquidity)

	myPrintln("liquidityAmt:", liquidityAmt)

	_, err = calleeInstance.Mint(Auth,
		common.HexToAddress(Network.Pool),
		FromAddress,
		big.NewInt(min_tick),
		big.NewInt(max_tick),
		liquidityAmt,
	)

	if err != nil {
		log.Fatal("callee mint err ", err)
	}

	///require governance. redo auth
	Auth = GetSignature(Networkid, Account)

	PoolInfo(Network.Pool)
}

func TokenInfo(tokenAddress string, owner string) (*big.Int, string) {

	tokenInstance, err := token.NewApi(common.HexToAddress(tokenAddress), EthClient)

	if err != nil {
		log.Fatal("tokenInfo NewApi err ", err)
	}

	balance, _ := tokenInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(owner))
	symbol, _ := tokenInstance.Symbol(&bind.CallOpts{})

	return balance, symbol

}
func getSqrtPriceX96(price float64) float64 {
	return math.Floor(math.Sqrt(price) * (1 << 96))
}

func getBigFromFloat64(v float64) *big.Int {

	s := fmt.Sprintf("%.0f", v)
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		log.Fatal("getBigFromFloat64 Err:", ok)
	}
	return b

}

func PrintPrice() float64 {

	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	_, pf := getPrice(slot0.SqrtPriceX96, slot0.Tick)

	tn := time.Now().Format("15:04:05")
	fmt.Println("Price now:", pf, "  ", tn)

	return pf
}

// func Observe(secondsAgo int) {
// 	poolInstance := GetPoolInstance()

// 	var secondsAgos uint32[2]
//     secondsAgos[0] = 20000000   //
//     secondsAgos[1] = 0; // to (now)

// 	tick, err := poolInstance.Observe(&bind.CallOpts{}, uint32(secondsAgo))

// 	fmt.Fprintln("observed tick %d , %d seconds ago:", tick, secondsAgo)

// }
