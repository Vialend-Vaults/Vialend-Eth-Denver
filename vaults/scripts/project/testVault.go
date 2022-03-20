package include

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	token "viaroot/deploy/Tokens/erc20/deploy/Token"

	//vault "viaroot/deploy/FeeMaker"

	VaultStrategy2 "viaroot/deploy/VaultStrategy2"
	arb "viaroot/deploy/arb"
	admin "viaroot/deploy/vaultAdmin"
	bridge "viaroot/deploy/vaultBridge"
	vault "viaroot/deploy/vialendFeeMaker"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// global variables
var tickLower = new(big.Int)
var tickUpper = new(big.Int)

var qTickLower = new(big.Int) // for query only, monitor
var qTickUpper = new(big.Int) // for query only, monitor

var prevFees0 = new(big.Int)
var prevFees1 = new(big.Int)

func Deposit(_amount0 int, _amount1 int, acc int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Deposit.........  ")
	myPrintln("----------------------------------------------")

	ChangeAccount(acc)

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))
	myPrintln("TokenA:", Network.TokenA)
	myPrintln("TokenB:", Network.TokenB)
	myPrintln("fromAddress to deposit: ", FromAddress)

	// instance, err := vault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	_, _, instance := GetInstance3()

	tokenAInstance, err := token.NewApi(common.HexToAddress(Network.TokenA), EthClient)
	if err != nil {
		log.Fatal("tokenAInstance,", err)
	}

	bal0, err := tokenAInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	myPrintln("tokenA in Wallet ", bal0)

	if err != nil {
		log.Fatal("bal0 err ", err)
	}

	tokenBInstance, err := token.NewApi(common.HexToAddress(Network.TokenB), EthClient)
	if err != nil {
		log.Fatal("tokenBInstance,", err)
	}

	bal1, err := tokenBInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	myPrintln("tokenB in Wallet ", bal1)

	//  amount0 * 10^decimals
	amount0 := big.NewInt(int64(_amount0))
	amount1 := big.NewInt(int64(_amount1))

	if bal1.Cmp(amount1) < 0 {
		log.Fatal("there is not enough tokenB: required:", amount1, ", available:", bal1)
	}

	if bal0.Cmp(amount0) < 0 {
		log.Fatal("there is not enough tokenA: required: ", amount0, ", available:", bal0)
	}

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	myPrintln("token0 decimals:", Token[0].Decimals)
	myPrintln("token1 decimals:", Token[1].Decimals)

	ApproveToken(common.HexToAddress(Network.TokenA), maxToken0, Network.Vault)
	ApproveToken(common.HexToAddress(Network.TokenB), maxToken1, Network.Vault)

	// not multi 10^decimals
	amountToken0 := amount0
	amountToken1 := amount1

	myPrintln("amountToken0 to Vault:", amountToken0)
	myPrintln("amountToken1 to Vault:", amountToken1)

	NonceGen()
	tx, err := instance.Deposit(Auth,
		amountToken0,
		amountToken1,
	)

	if err != nil {
		log.Fatal("deposit err: ", err)
	}

	ChangeAccount(Account)

	//get the transaction hash
	myPrintln("deposit tx: %s", tx.Hash().Hex())

	//	time.Sleep(Network.PendingTime * time.Second)
	//Readstring("deposit sent...wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func WithdrawPending(pct int64, acc int64) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Withdraw pending with percentage.........  ")
	myPrintln("----------------------------------------------")

	ChangeAccount(int(acc))

	NonceGen()
	vaultInstance := GetVaultInstance()
	tx, err := vaultInstance.WithdrawPending(Auth, uint8(pct))

	if err != nil {
		log.Fatal("withdraw pending: ", err)
	}

	/// reset account back
	ChangeAccount(Account)

	//	Readstring("withdraw sent.... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func Withdraw(_percent int, acc int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Withdraw.........  ")
	myPrintln("----------------------------------------------")

	/// set new account Auth
	ChangeAccount(acc)

	recipient := FromAddress

	myPrintln("Withdraw to  ", recipient)

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	percent := uint8(_percent)

	// instance, err := vault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }
	_, _, instance := GetInstance3()

	///get account's available shares
	myshares, err := instance.BalanceOf(&bind.CallOpts{}, recipient)
	if err != nil {
		log.Fatal("balance of myshare ", err)
	}

	if myshares.Cmp(big.NewInt(0)) == 0 {
		myPrintln("share==0 ", myshares)
		return
	}

	NonceGen()
	tx, err := instance.Withdraw(Auth, percent)

	if err != nil {
		log.Fatal("withdraw: ", err)
	}

	/// reset account back
	ChangeAccount(Account)

	//get the transaction hash
	myPrintln("withdraw tx: ", tx.Hash().Hex())

	//	Readstring("withdraw sent.... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func GetSwapInfo(rangeRatio int64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick
	sqrtPriceX96 := slot0.SqrtPriceX96

	//myPrintln("tick: ", tick)
	//myPrintln("sqrtPriceX96: ", sqrtPriceX96)

	Totals := GetTVL()

	myPrintln("total locked token0: ", Totals.Total0)
	myPrintln("total locked token1: ", Totals.Total1)

	_, pf := getPrice(sqrtPriceX96, tick)

	min := pf * (1 - float64(rangeRatio)/100)
	max := pf * (1 + float64(rangeRatio)/100)
	x := BigIntToFloat64(Totals.Total0) / math.Pow10(int(Token[0].Decimals))
	y := BigIntToFloat64(Totals.Total1) / math.Pow10(int(Token[1].Decimals))

	xDecimals, yDecimals := Token[0].Decimals, Token[1].Decimals

	//myPrintln("pf, min, max, rangeRatio: ", pf, min, max, rangeRatio)

	//myPrintln("pf,min, max, rangeRatio: ", pf, min, max, rangeRatio)

	a, b := getTicks(pf, min, max, float64(xDecimals), float64(yDecimals))
	tickA := math.Round(a/60) * 60
	tickB := math.Round(b/60) * 60

	//myPrintln("tick a b:", tickA, tickB)

	tickLower = big.NewInt(int64(tickA)) //tickA) //big.NewInt(-1140) //
	tickUpper = big.NewInt(int64(tickB)) //tickB) //big.NewInt(840)   //

	//myPrintln("---abminmax:", pf-float64(rangeRatio)/100, pf, a, b, tickLower, tickUpper)
	//os.Exit(3)
	// myPrintln(pf, min, max, x, y)
	// myPrintln(priceFromsqrtP)
	// myPrintln(BigIntToFloat64(Total0))
	// myPrintln(BigIntToFloat64(Total1))

	//amt0, amt1, swapAmount, zeroForOne := getBestAmounts(pf, min, max, x, y)
	return getBestAmounts(pf, min, max, x, y)
}

func Sweep(tokenAddr string, amount *big.Int) {

	myTitle("sweep")

	vaultInstance := GetVaultInstance2(Network.Vault)

	tx, err := vaultInstance.Sweep(Auth, common.HexToAddress(tokenAddr), amount)

	if err != nil {
		log.Fatal("sweep err ", err)
	}

	TxConfirm(tx.Hash())

}

// func EmergencyCall() {

// 	myPrintln("----------------------------------------------")
// 	myPrintln(".........Emergency call setup emergency stat.........  ")
// 	myPrintln("----------------------------------------------")

// 	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

// 	_, _, vaultInstance := GetInstance3()

// 	// tickLower, err := vaultInstance.CLow(&bind.CallOpts{})
// 	// tickUpper, err := vaultInstance.CHigh(&bind.CallOpts{})
// 	// liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)

// 	// lendingAmount0 := checkCTokenBalance("CETH", Network.LendingContracts.CETH)
// 	// lendingAmount1 := checkCTokenBalance("CUSDC", Network.LendingContracts.CUSDC)
// 	NonceGen()
// 	tx, err := vaultInstance.EmergencyCall(Auth)

// 	if err != nil {
// 		log.Fatal("emergency tx err ", err)
// 	}

// 	myPrintln("emergency tx: ", tx.Hash().Hex())

// 	//Readstring("emergency sent sent.....  wait for pending..next .. white hacker to withdraw ")
// 	TxConfirm(tx.Hash())

// }

func EmergencyWithdraw(acc int) {
	myPrintln("----------------------------------------------")
	myPrintln(".........Emergency withdraw .........  ")
	myPrintln("----------------------------------------------")
	_, _, vaultInstance := GetInstance3()

	ChangeAccount(acc)

	tx, _ := vaultInstance.EmergencyWithdraw(Auth)

	ChangeAccount(Account)

	TxConfirm(tx.Hash())

}

/// param0: fullRangeSize,
/// param1: tickspacing,
/// param2: accId
func Strategy1(_range int64, acc int64) {

	myPrintln("----------------------------------------------")
	fmt.Println(".........Strategy1 .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	vaultInstance, err := vault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Fatal("vault.NewApi ", err)
	}

	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(_range), big.NewInt(2))
	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := Network.FeeTier / 50 // ie 3000/50= 60, 500/50 = 10, 10000/50 = 200
	myPrintln("tickspacing:", tickSpacing)

	if tickSpacing < 10 {
		log.Fatal("wrong tickSpacing = ", tickSpacing)
	}
	//tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	ChangeAccount(int(acc))

	myPrintln("range size:", _range)
	myPrintln("ticklower, TICK,  tickupper in...", tickLower, tick, tickUpper)
	myPrintln("in range? ", tick.Cmp(tickLower) > 0 && tick.Cmp(tickUpper) < 0)

	// set ticklower and tickupper
	//setRange(param)
	i := 0
	for {
		NonceGen()

		tx, err := vaultInstance.Strategy1(Auth,
			tickLower,
			tickUpper)

		if err != nil {
			fmt.Println("strateg1 tx err , trying one more time in 2 seconds..", err)
			time.Sleep(2 * time.Second)
		} else {
			//myPrintln("strategy1 tx: ", tx.Hash().Hex())
			TxConfirm(tx.Hash())
			break
		}

		if i > 10 {
			log.Fatal("strateg1 tx err ", err)
		}

	}

	///require governance. redo auth
	ChangeAccount(Account)

	fmt.Println()

	//Readstring("Rebalance by Strategy1 sent....... ")

}

func LendingInfo() {
	myPrintln("----------------------------------------------")
	myPrintln(".........Lending pool info .........  ")
	myPrintln("----------------------------------------------")

	myPrintln(".........Compound info .........  ")
	checkCTokenBalance(Network.Vault, "CUSDC", Network.LendingContracts.CUSDC)
	checkCTokenBalance(Network.Vault, "CETH", Network.LendingContracts.CETH)
	checkETHBalance()

	//	_, stratInstance, _ := GetInstance3()
	//	CAmounts, _ := stratInstance.GetCAmounts(&bind.CallOpts{})
	//	myPrintln("CToken0, Ctoken1:", CAmounts)

	exchangeRateStored, _ := GetCTokenInstance(Network.CTOKEN0).ExchangeRateStored(&bind.CallOpts{})

	//dem := int(18 + int(Token[0].Decimals) - 8)
	//	ctoken0Underlying := BigIntToFloat64(CAmounts.Amount0) * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored0:", exchangeRateStored)
	//	myPrintln("ctoken0Underlying", ctoken0Underlying)

	exchangeRateStored, _ = GetCTokenInstance(Network.CTOKEN1).ExchangeRateStored(&bind.CallOpts{})
	//dem = int(18 + int(Token[1].Decimals) - 8)
	//	ctoken1Underlying := BigIntToFloat64(CAmounts.Amount1) * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored1:", exchangeRateStored)
	//	myPrintln("ctoken1Underlying", ctoken1Underlying)

	myPrintln("counter check with GetCompAmounts() from contract:")
	//	GetCompAmounts()

	myPrintln("wbtc info")
	exchangeRateStored, _ = GetCTokenInstance(Network.LendingContracts.CWBTC).ExchangeRateStored(&bind.CallOpts{})
	cwbtcAmount := float64(141008860)
	cwbtcUnderlying := cwbtcAmount * (BigIntToFloat64(exchangeRateStored) / 1e18) /// 1 * BigIntToFloat64(PowX(10, dem)))
	myPrintln("exchangeRateStored0:", exchangeRateStored)
	myPrintln(cwbtcAmount, "CWBTC =  ", cwbtcUnderlying, "WBTC")

}

func checkETHBalance() *big.Int {

	bal := EthBalance(Network.Vault)

	myPrintln("eth balance: ", bal)

	return (bal)

}

func checkCTokenBalance(who string, tokenName string, cTokenAddress string) *big.Int {

	cInstance := GetCTokenInstance(cTokenAddress)
	bal, err := cInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(who))

	if err != nil {
		myPrintln("who ", who)
		myPrintln("cTokenAddress ", cTokenAddress)
		log.Fatal("cInstance.BalanceOf err ", err)
	}

	//	myPrintln(tokenName, " balance: ", bal)

	return (bal)

}

/// param0 : fullRangeSize, param1: account
func Rebalance(_range int, acc int) {

	myTitle(".........Rebalance .........  ")

	_, stratInstance, _ := GetInstance4()

	//init ticklow and tickupp
	//GetSwapInfo(param[0])
	poolInstance := GetPoolInstance()
	slot0, err := poolInstance.Slot0(&bind.CallOpts{})
	if err != nil {
		myPrintln(Network.Pool)
		log.Fatal("poo slot0 : ", err)
	}
	tick := slot0.Tick

	hrange := new(big.Int).Div(big.NewInt(int64(_range)), big.NewInt(2))

	tickLower = new(big.Int).Sub(tick, hrange)
	tickUpper = new(big.Int).Add(tick, hrange)

	tickSpacing := Network.FeeTier / 50 // ie 3000/50= 60, 500/50 = 10, 10000/50 = 200
	myPrintln("rebalance - tickspacing:", tickSpacing)

	if tickSpacing < 10 {
		log.Fatal("wrong tickSpacing = ", tickSpacing)
	}
	//tickSpacing := param[1]
	tickLower.Div(tickLower, big.NewInt(tickSpacing)).Mul(tickLower, big.NewInt(tickSpacing))
	tickUpper.Div(tickUpper, big.NewInt(tickSpacing)).Mul(tickUpper, big.NewInt(tickSpacing))

	///require governance. redo auth
	ChangeAccount(int(acc))

	myPrintln("range size:", _range)
	myPrintln("ticklower, TICK,  tickupper in...", tickLower, tick, tickUpper)
	myPrintln("in range? ", tick.Cmp(tickLower) > 0 && tick.Cmp(tickUpper) < 0)

	// set ticklower and tickupper
	//setRange(param)
	NonceGen()

	var rebalParam VaultStrategy2.RebalanceParam
	rebalParam.NewHigh = tickUpper
	rebalParam.NewLow = tickLower
	rebalParam.SqthPercent = uint32(100)
	rebalParam.UniPortionRate = uint32(2665)
	rebalParam.SqthPortionRate = uint32(4675)
	rebalParam.ShortPortionRate = uint32(2659)
	rebalParam.ShortLever = uint32(1)

	tx, err := stratInstance.Rebalance(Auth, rebalParam)

	if err != nil {
		log.Fatal("Rebalance err:", err)
	}

	TxConfirm(tx.Hash())

	///require governance. redo auth
	ChangeAccount(Account)

	fmt.Println()

	//Readstring("Rebalance by Strategy1 sent....... ")

}

func Stat2Str(stat uint64) string {
	if stat == 0 {
		return ("FAIL")
	} else {
		return ("SUCCESS")
	}

}

func getBestAmounts(p float64, a float64, b float64, x float64, y float64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {

	sp := math.Sqrt(p) //p * *0.5
	sa := math.Sqrt(a) //a * *0.5
	sb := math.Sqrt(b) //b * *0.5
	// calculate the initial liquidity
	L := get_liquidity(x, y, sp, sa, sb)

	P1 := p
	sp1 := math.Sqrt(P1) // P1 * *0.5
	x1 := calculate_x(L, sp1, sb)
	y1 := calculate_y(L, sp1, sa)

	//fmt.Printf("x1={%.4f}\ny1={%.4f}\n", x1, y1)

	x1r := x1 / (x1 + y1/p)
	y1r := y1 / (y1 + x1*p)
	myPrintln(x1r, y1r)

	xr := x / (x + y/p)
	yr := y / (y + x*p)
	myPrintln(xr, yr)
	// 20/2000, 0.9908
	// 20 * 0.9908
	if x*p > y { // trade x for y
		zeroForOne = true

		r := xr - x1r

		swapAmount = math.Abs(x * r)

		amount0 = x - swapAmount

		amount1 = y + swapAmount*p

		//myPrintln("newX=", amount0)
		//myPrintln("newY=", amount1)

	} else if x*p < y { // trade y for x
		zeroForOne = false

		r := xr - x1r

		swapAmount = math.Abs(y * r)

		amount0 = x + swapAmount/p

		amount1 = y - swapAmount

		//myPrintln("newX=", amount0)
		//myPrintln("newY=", amount1)
	}

	//fmt.Printf("newX={%.18f}, newY={%.6f},swapamount={%.18f},zeroForOne={%t}\n", amount0, amount1, swapAmount, zeroForOne)

	return amount0, amount1, swapAmount, zeroForOne
}

func OraclePrice() (twapPrice *big.Int, spotPrice *big.Int) {

	vaultInstance := GetVaultInstance()

	var twapDuration = uint32(2)

	poolAddress := Network.Pool

	twap, _ := vaultInstance.GetTwap(&bind.CallOpts{}, common.HexToAddress(poolAddress), twapDuration)

	Sleep(100)

	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	Sleep(100)
	//	tick := slot0.Tick

	myPrintln(",slot.tick:", slot0.Tick)
	myPrintln(",slot0.SqrtPriceX96:", slot0.SqrtPriceX96)

	spotPrice, pricefloat64 := getPrice(slot0.SqrtPriceX96, slot0.Tick)

	fmt.Println("Spot price bigint:", spotPrice)
	fmt.Println("Spot price :", pricefloat64)

	if twap != nil {
		//	twap = big.NewInt(-192874)
		baseAmount := big.NewInt(1e8)
		baseToken := common.HexToAddress(Network.TokenA)
		quoteToken := common.HexToAddress(Network.TokenB)

		//myPrintln(" twap, baseAmount, baseToken, quoteToken", twap, baseAmount, baseToken, quoteToken)

		twapPrice, _ = vaultInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, baseAmount, baseToken, quoteToken)
		Sleep(200)

		fmt.Println("twap Price:", twapPrice)
	}

	myPrintln("twap:", twap)

	if twap == nil {
		twapPrice = spotPrice
	}

	return
}

/// protocol fees, my earned value, APY
func CheckFees() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Check Fees .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("vaultAddress: ", common.HexToAddress(Network.Vault))

	vaultInstance := GetVaultInstance()

	Fees, _ := vaultInstance.Fees(&bind.CallOpts{})

	myPrintln("U3Fees0, U3Fees1, LcFees0, LcFees1, AccruedProtocla Fees0,1 : {", Fees, "}")

	totalFees0 := new(big.Int).Add(Fees.U3Fees0, Fees.LcFees0)
	totalFees1 := new(big.Int).Add(Fees.U3Fees1, Fees.LcFees1)
	fmt.Println("totalFees0, totalFees1, prev0, prev1, diff0, diff1: {",
		totalFees0, totalFees1, "}",
		"{", prevFees0, prevFees1, "}",
		"{", new(big.Int).Sub(totalFees0, prevFees0), new(big.Int).Sub(totalFees1, prevFees1), "}")

	prevFees0 = totalFees0
	prevFees1 = totalFees1

	// accumulative protocol fees
	myPrintln("accruedProtocolFees0, accruedProtocolFees1 : {", Fees.AccruedProtocolFees0, Fees.AccruedProtocolFees1, "}")

	// // check team share
	// Team, _ := vaultInstance.Team(&bind.CallOpts{})
	// tokenIns, _, _, _, _ := GetTokenInstance(Network.Vault)
	// teamShare, _ := tokenIns.BalanceOf(&bind.CallOpts{}, Team)
	// totalShare, _ := tokenIns.TotalSupply(&bind.CallOpts{})
	// myPrintln("Team address: ", Team)

	myPrintln()

	return

	_, xPrice := OraclePrice()

	// crashed on forked mainnet
	for j, _ := range Network.PrivateKey {
		Sleep(2000)
		storedAccount, _ := vaultInstance.Accounts(&bind.CallOpts{}, big.NewInt(int64(j)))
		Sleep(1000)
		if storedAccount.String() != "0x0000000000000000000000000000000000000000" {

			fmt.Println("\n*My address:", storedAccount)

			myAddress := common.HexToAddress(storedAccount.String())

			myshare, totalshare := CalcShares(myAddress)

			if myshare.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			Assets, _ := vaultInstance.Assetholder(&bind.CallOpts{}, storedAccount)
			Sleep(1000)
			myPrintln("*Assetsholder: ", Assets)

			if totalshare.Cmp(big.NewInt(0)) > 0 && myshare.Cmp(big.NewInt(0)) > 0 {

				//my earned value
				myFees0 := new(big.Int).Mul(totalFees0, myshare)
				myFees0.Div(myFees0, totalshare)

				myFees1 := new(big.Int).Mul(totalFees1, myshare)
				myFees1.Div(myFees1, totalshare)

				fmt.Println("my Fees0: {", myFees0, "}")
				fmt.Println("my Fees1: {", myFees1, "}")

				myFeesInToken1 := new(big.Int).Mul(myFees0, xPrice)
				myFeesInToken1 = myFeesInToken1.Add(myFeesInToken1, myFees1)

				ListAccounts, _ := vaultInstance.Assetholder(&bind.CallOpts{}, storedAccount)
				Sleep(1000)
				myPrintln("*Assetsholder: ", ListAccounts)

				// calc APY  below
				blockNumber := Assets.Block

				block, err := EthClient.BlockByNumber(context.Background(), blockNumber)

				if err != nil {
					log.Fatal("block ", err)
				}
				Sleep(1000)

				// get block info

				timestamp := block.Time()

				myPrintln("deposit block info:", blockNumber, block.Time()) // 1527211625

				header, err := EthClient.HeaderByNumber(context.Background(), nil)
				if err != nil {
					log.Fatal("block header ", err)
				}
				Sleep(1000)

				headerblock, err := EthClient.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					log.Fatal("block ", err)
				}
				Sleep(1000)

				htimestamp := headerblock.Time()

				timediff := htimestamp - timestamp // diff in seconds

				myPrintln("timediff from now: {", timediff, htimestamp, timestamp, "}")

				fmt.Println("Period (Days):", timediff/60/60/24, "{sec:}", timediff)

				oneyearINsec := big.NewInt(365 * 24 * 60 * 60)
				myPrintln("oneyearINsec ", oneyearINsec)

				deposit0 := Assets.Deposit0
				deposit1 := Assets.Deposit1

				totals := GetTVL()
				Sleep(1000)

				myPrintln("totalTVL ", totals)

				//my value locked: mytvl0, mytvl1
				mytvl0 := new(big.Int).Mul(totals.Total0, myshare)
				mytvl0 = new(big.Int).Div(mytvl0, totalshare)

				mytvl1 := new(big.Int).Mul(totals.Total1, myshare)
				mytvl1 = new(big.Int).Div(mytvl1, totalshare)

				fmt.Println("deposit0,deposit1 {", deposit0, deposit1, "}")
				fmt.Println("mytvl0, mytvl1 {", mytvl0, mytvl1, "}")

				fd0 := BigIntToFloat64(deposit0)

				fd1 := BigIntToFloat64(deposit1) * 1e18 / BigIntToFloat64(xPrice)

				fm0 := BigIntToFloat64(mytvl0)

				fm1 := BigIntToFloat64(mytvl1) * 1e18 / BigIntToFloat64(xPrice)

				myPrintln("fm0:", fm0)
				myPrintln("fm1:", fm1)

				fdd := fd0 + fd1
				fmm := fm0 + fm1

				myPrintln("fdd, fmm:", fdd, fmm)

				APY := (fmm - fdd) / float64(timediff) * BigIntToFloat64(oneyearINsec) / fdd

				fmt.Println("APY:", APY)

				myDepositInToken1 := new(big.Int).Mul(deposit0, xPrice)
				myDepositInToken1 = myDepositInToken1.Add(myDepositInToken1, deposit1)

				fAPY := BigIntToFloat64(myFeesInToken1) / float64(timediff) * BigIntToFloat64(oneyearINsec) / BigIntToFloat64(myDepositInToken1) * 100
				fmt.Println("APY by fees/Deposit ", fAPY, "%")

				// timediff, myshare, totalshare, tvl0, tvl1, deposit0, deposit1, usPrice0, usPrice1
				// mytvl0 = tvl0 * myshare/totalshare
				// mytvl1 = tvl1 * myshare/totalshare
				//APY =  ( (mytvl0 - deposit0) + (mytvl1 -deposit1) ) / blocktimediff * oneyearINsec

				//			myPrintln(block.Difficulty().Uint64())       // 3217000136609065
				//			myPrintln("block hash:", block.Hash().Hex()) // 0x9e8751ebb5069389b855bba72d949
				//			blockHashHex := block.Hash().Hex()

			}

		}

		myPrintln()

	}

}

func GetPriceFromAavePool() {

	_, stratInstance, _ := GetInstance4()

	aavePrice, _ := stratInstance.GetPriceFromAavePool(&bind.CallOpts{})
	myPrintln("aavePrice from aave eth/usdc Pool :", aavePrice)

}

/// formula:
///
// 1. price = pow(1.0001,tick) * (1e(18-6)
///2.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) >> (96*2)
///3.  (sqrtPricex96^2 * 1e(decimal0)/1e(decimal1) / 2^(96*2)
// 4. javascript: JSBI.BigInt(sqrtPriceX96 *sqrtPriceX96* (1e(decimals_token_0))/(1e(decimals_token_1))/JSBI.BigInt(2) ** (JSBI.BigInt(192));
///5 solc: uint(sqrtPriceX96).mul(uint(sqrtPriceX96)).mul(1e(decimalsDiff)) >> (96 * 2);
/// returns priceInBigInt, priceReadable
func getPrice(SqrtPriceX96 *big.Int, tick *big.Int) (*big.Int, float64) {

	myPrintln("decimals0:", Token[0].Decimals)
	myPrintln("decimals1:", Token[1].Decimals)

	decimalDiff := int(Token[0].Decimals) - int(Token[1].Decimals)

	if int(Token[0].Decimals) < int(Token[1].Decimals) {
		decimalDiff = int(Token[1].Decimals) - int(Token[0].Decimals)
	}
	myPrintln("decimals diff:", decimalDiff)

	tick24 := float64(tick.Int64())
	//myPrintln("tick24 ", tick24)

	powTick := math.Pow(1.0001, tick24)

	//pow(1.0001,197510)/pow(10,(18-6))

	tickPrice := powTick / float64(math.Pow10(int(decimalDiff)))
	tickPriceReadable := 1 / tickPrice

	PriceBigInt := Float64ToBigInt(tickPriceReadable * math.Pow10(int(Token[1].Decimals)))
	myPrintln("pricebigint", PriceBigInt)
	myPrintln("price readable", tickPriceReadable)

	//myPrintln("counter check price with sqrtPx96 ^ 2 >> 192 = ", sqrtPx962Price)
	return PriceBigInt, tickPriceReadable
}

func getPrice2(SqrtPriceX96 *big.Int, tick *big.Int, decimal0 int, decimal1 int) (*big.Int, float64) {

	myPrintln("decimals0:", decimal0)
	myPrintln("decimals1:", decimal1)

	decimalDiff := decimal0 - decimal1

	if decimalDiff < 0 {
		decimalDiff = -decimalDiff
	}
	myPrintln("decimals diff:", decimalDiff)

	tick24 := float64(tick.Int64())
	//myPrintln("tick24 ", tick24)

	powTick := math.Pow(1.0001, tick24)

	//pow(1.0001,197510)/pow(10,(18-6))

	tickPrice := powTick / float64(math.Pow10(int(decimalDiff)))
	tickPriceReadable := 1 / tickPrice

	PriceBigInt := Float64ToBigInt(tickPriceReadable * math.Pow10(int(Token[1].Decimals)))
	myPrintln("pricebigint", PriceBigInt)
	myPrintln("price readable", tickPriceReadable)

	//myPrintln("counter check price with sqrtPx96 ^ 2 >> 192 = ", sqrtPx962Price)
	return PriceBigInt, tickPriceReadable
}

func GetVaults(n int) {
	fin, _, _ := GetInstance3()
	vaults, _ := fin.Vaults(&bind.CallOpts{}, big.NewInt(int64(n)))
	myPrintln(vaults)

}

func SetVaults() {
	fin, _, _ := GetInstance3()
	vaults, _ := fin.Vaults(&bind.CallOpts{}, big.NewInt(0))
	myPrintln(vaults)

}

func Alloc(accId int) {

	myTitle(" alloc  ")

	_, stratInstance, _ := GetInstance4()

	NonceGen()

	tx, err := stratInstance.Alloc(Auth)

	if err != nil {
		log.Fatal("alloc err: ", err)
	}

	myPrintln("alloc tx: %s", tx.Hash().Hex())

	//Readstring("alloc sent...wait for pending..next .. ")
	TxConfirm(tx.Hash())

}
func GetTotalSupply() {
	_, _, vaultInstance := GetInstance3()

	totalsupply, _ := vaultInstance.TotalSupply(&bind.CallOpts{})
	myPrintln("totalsupply:", totalsupply)

}

func GetTVL() *struct {
	Total0 *big.Int
	Total1 *big.Int
} {

	myTitle("GetTVL")

	_, stratInstance, _ := GetInstance4()
	token0Ins, _, _, _, _ := GetTokenInstance(Network.TokenA)
	token1Ins, _, _, _, _ := GetTokenInstance(Network.TokenB)

	//implement gettvl
	cHigh, _ := stratInstance.CHigh(&bind.CallOpts{})
	cLow, _ := stratInstance.CLow(&bind.CallOpts{})

	myTitle(" Uniswap ")
	uniLnA, _ := stratInstance.GetUniAmounts(&bind.CallOpts{}, cLow, cHigh)
	myPrintln("Token0: ", uniLnA.Amount0)
	myPrintln("Token1: ", uniLnA.Amount1)
	myPrintln("Liquidity: ", uniLnA.Myliquidity)
	myPrintln("cHigh, cLow:", cHigh, cLow)

	// clending0, clending1 := stratInstance.GetCAmounts(&bind.CallOpts{})
	// myPrintln("C Amounts in lending: ", clending0, clending1)

	myTitle("in strategy ")
	symbol, squeeth := GetBalance(Network.LendingContracts.OSQTH, Network.VaultStrat)
	myPrintln(symbol, ": ", squeeth)

	symbol, aToken := GetBalance(Network.LendingContracts.ATOKEN_USDC, Network.VaultStrat)
	myPrintln(symbol, ": ", aToken)

	// symbol, aaveWETH := GetBalance(Network.LendingContracts.AAVE_ETH, Network.VaultStrat)
	// myPrintln(symbol, ": ", aaveWETH)
	// symbol, aaveUSDC := GetBalance(Network.LendingContracts.AAVE_USDC, Network.VaultStrat)
	// myPrintln(symbol, ": ", aaveUSDC)

	sbalance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.VaultStrat))
	sbalance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.VaultStrat))
	myPrintln("(usdc  weth):  ", sbalance0, sbalance1)

	myTitle("in viaVault ")
	symbol, squeeth = GetBalance(Network.LendingContracts.OSQTH, Network.Vault)
	myPrintln(symbol, ": ", squeeth)

	vbalance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	vbalance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	myPrintln("(usdc  weth) :  ", vbalance0, vbalance1)

	myTitle("Total")
	Totals := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})

	amount0, amount1, err := stratInstance.GetTotalAmounts(&bind.CallOpts{})
	if err != nil {

	} else {
		myPrintln("Total (usdc , weth) in Strat :  ", amount0, amount1)

		Totals.Total0 = amount0.Add(amount0, vbalance0)
		Totals.Total1 = amount1.Add(amount1, vbalance1)

		myPrintln("TVL (usdc , weth) in All:  ", Totals.Total0, Totals.Total1)
	}
	return Totals
}

func GetLendingAmounts(vaultAddr string) (*big.Int, *big.Int, *big.Int, *big.Int) {

	cInstance0 := GetCTokenInstance(Network.CTOKEN0)
	cInstance1 := GetCTokenInstance(Network.CTOKEN1)

	//implement gettvl
	exchangeRate0, _ := cInstance0.ExchangeRateStored(&bind.CallOpts{})
	exchangeRate1, _ := cInstance1.ExchangeRateStored(&bind.CallOpts{})

	CAmount0 := checkCTokenBalance(vaultAddr, "CToken0", Network.CTOKEN0)
	CAmount1 := checkCTokenBalance(vaultAddr, "CToken1", Network.CTOKEN1)

	pow1018 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	//	pow106 := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	underlying0 := CAmount0.Mul(CAmount0, exchangeRate0).Div(CAmount0, pow1018) //.Div(CAmount0, pow1018)
	underlying1 := CAmount1.Mul(CAmount1, exchangeRate1).Div(CAmount1, pow1018) //.Div(CAmount1, pow106)

	return underlying0, underlying1, exchangeRate0, exchangeRate1

}

func ApproveToken(tokenAddress common.Address, maxAmount *big.Int, toAddress string) {

	tokenInstance, err := token.NewApi(tokenAddress, EthClient)
	if err != nil {
		log.Fatal("tokenInstance,", err)
	}

	//check allowance
	allow, _ := tokenInstance.Allowance(&bind.CallOpts{}, FromAddress, common.HexToAddress(toAddress))

	if allow.Cmp(big.NewInt(0)) > 0 {
		return
	}

	myPrintln("APPROVE:")
	myPrintln("to address be approved: ", toAddress)
	myPrintln("fromAddress: ", FromAddress)
	myPrintln("Allowance amount:", allow)

	NonceGen()

	tx, err := tokenInstance.Approve(Auth, common.HexToAddress(toAddress), maxAmount)

	if err != nil {
		log.Fatal("token approve tx, ", err)
	}

	Sleep(2000)
	//Readstring("Approve sent... wait for pending..next .. ")
	TxConfirm(tx.Hash())

}

func Approve(account int) {

	if account < 0 {
		return
	}

	myPrintln("----------------------------------------------")
	myPrintln(".........Approve.........  ")
	myPrintln("----------------------------------------------")

	ChangeAccount(account)
	NonceGen()

	poolInstance := GetPoolInstance()
	TokenA, _ := poolInstance.Token0(&bind.CallOpts{})

	TokenB, _ := poolInstance.Token1(&bind.CallOpts{})

	myPrintln("tokenA: ", TokenA)
	myPrintln("tokenB: ", TokenB)

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(TokenA, maxToken0, Network.Vault)
	ApproveToken(TokenB, maxToken1, Network.Vault)

	ChangeAccount(Account)

}

func AccountInfo() {

	myPrintln("----------------------------------------------")
	myPrintln(".........Account Info.........  ")
	myPrintln("----------------------------------------------")

	for i, _ := range Network.PrivateKey {

		MyAccountInfo(i)
	}

}

func MyAccountInfo(accId int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........My account Info.........  ")
	myPrintln("----------------------------------------------")

	accountAddress := GetAddress(accId)

	myPrintln("Account  ----", accId)
	myPrintln("Account address ", accountAddress)

	myPrintln("Eth balance:", EthBalance(accountAddress.String()))

	tokenAInstance, _, symbolA, _, _ := GetTokenInstance(Network.TokenA)

	bal, _ := tokenAInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln(symbolA, " in Wallet ", bal)

	tokenBInstance, _, symbolB, _, _ := GetTokenInstance(Network.TokenB)

	bal, _ = tokenBInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	myPrintln(symbolB, " in Wallet ", bal)

	///----------- token in vault

	// vaultInstance, err := vault.NewApi(vaultAddress, EthClient)
	// if err != nil {
	// 	log.Fatal("vault.NewApi ", err)
	// }

	mybal, totalbal := CalcShares(accountAddress)

	if totalbal.Cmp(big.NewInt(0)) > 0 {

		myPrintln("my share / totalSupply ", mybal.Mul(mybal, big.NewInt(100)).Div(mybal, totalbal), "%")
	}

	vaultInstance := GetVaultInstance()
	Assets, _ := vaultInstance.Assetholder(&bind.CallOpts{}, accountAddress)

	myPrintln("*Assetsholder: ", Assets)

	myPrintln("------- In My Wallet ---------- ")

	myPrintln(GetBalance(Network.LendingContracts.DAI, accountAddress.Hex()))
	myPrintln(GetBalance(Network.LendingContracts.USDC, accountAddress.Hex()))
	myPrintln(GetBalance(Network.LendingContracts.WETH, accountAddress.Hex()))
	myPrintln(GetBalance(Network.LendingContracts.OSQTH, accountAddress.Hex()))

	myPrintln()

}

func CalcShares(myAddress common.Address) (mybal *big.Int, totalbal *big.Int) {

	/// vault as token
	vaultTokenInstance, err := token.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Fatal("vaultTokenInstance,", err)
	}

	mybal, _ = vaultTokenInstance.BalanceOf(&bind.CallOpts{}, myAddress)
	Sleep(1000)

	//if mybal.Cmp(big.NewInt(0)) > 0 {
	myPrintln("myShares in vault ", mybal)

	//}

	totalbal, err = vaultTokenInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal("vault token balance,", err)
	}
	Sleep(1000)
	//if totalbal.Cmp(big.NewInt(0)) > 0 {
	myPrintln("totalSupply in vault ", totalbal)
	//}

	return mybal, totalbal
}

func VaultInfo() {
	VaultInfo2(Network.Vault)
}

func VaultInfo2(vaultAddr string) {

	myTitle(".........Vault Info.........  ")

	//vaultInstance := GetVaultInstance2(vaultAddr)
	_, vaultInstance, _ := GetInstance4()

	//myPrintln("Vault Address:  ", vaultAddr)

	//poolAddress := Network.Pool
	//get ctoken address
	// _CToken0Addr, _ := vaultInstance.CToken0(&bind.CallOpts{})
	// _CToken1Addr, _ := vaultInstance.CToken1(&bind.CallOpts{})
	// myPrintln("Ctoken0 address:", _CToken0Addr)
	// myPrintln("Ctoken1 address:", _CToken1Addr)

	// totalSupply, err := vaultInstance.TotalSupply(&bind.CallOpts{})

	// myPrintln("totalSupply (total shares in vault) :", totalSupply)
	// if err != nil {
	// 	log.Fatal("totalsupply ", err)
	// }

	Sleep(100)
	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	Sleep(100)
	tick := slot0.Tick

	qTickLower, _ := vaultInstance.CLow(&bind.CallOpts{})
	Sleep(100)
	qTickUpper, _ := vaultInstance.CHigh(&bind.CallOpts{})
	Sleep(100)

	myTitle("tick range: ")
	myPrintln("cLow, tick, cHigh  :", qTickLower, tick, qTickUpper)

	myTitle("in range?")
	if qTickLower != nil && qTickUpper != nil {
		fmt.Println(tick.Cmp(qTickLower) > 0 && tick.Cmp(qTickUpper) < 0)
	}

	// liquidity, err := vaultInstance.GetSSLiquidity(&bind.CallOpts{}, qTickLower, qTickUpper)
	// Sleep(100)
	// myPrintln("Current liquidity in pool :", liquidity)

	// if err != nil {
	// 	log.Fatal("getssliquidity  ", err)
	// }

	// xy, err := vaultInstance.GetPositionAmounts(&bind.CallOpts{}, qTickLower, qTickUpper)
	// Sleep(100)
	// myPrintln("tokenA (x) in pool ", xy.Amount0)
	// myPrintln("tokenB (y) in pool ", xy.Amount1)

	///----------- 分别返回两个toten 的总数, also print tokens in vault
	// totals := GetTVL2(vaultAddr)

	// fmt.Printf("TVL token0=%d\n", totals.Total0)
	// fmt.Printf("TVL token1=%d\n", totals.Total1)
	// myPrintln("decimals0:", int(Token[0].Decimals))
	// myPrintln("decimals1:", int(Token[1].Decimals))

	// uniPortion, _ := vaultInstance.UniPortion(&bind.CallOpts{})
	// Sleep(100)
	// myPrintln("uniPortionRate:", uniPortion)

	// protocolFeeRate, _ := vaultInstance.ProtocolFee(&bind.CallOpts{})
	// Sleep(100)
	// myPrintln("ProtocolFeeRate:", protocolFeeRate)

	//	sqrtPriceX96 := slot0.SqrtPriceX96
	// uniswapPriceBySqrtP, _ := vaultInstance.GetPriceBySQRTP(&bind.CallOpts{}, sqrtPriceX96)
	// myPrintln("GetPriceBySQRTP:", uniswapPriceBySqrtP)

	myTitle("------ tokens in Vault-----------")
	myPrintln(GetBalance(Network.LendingContracts.DAI, Network.Vault))
	myPrintln(GetBalance(Network.LendingContracts.USDC, Network.Vault))
	myPrintln(GetBalance(Network.LendingContracts.WETH, Network.Vault))
	myPrintln(GetBalance(Network.LendingContracts.OSQTH, Network.Vault))
	myPrintln(GetBalance(Network.LendingContracts.ATOKEN_USDC, Network.Vault))

	myTitle("------ tokens in Strategy-----------")
	myPrintln(GetBalance(Network.LendingContracts.DAI, Network.VaultStrat))
	myPrintln(GetBalance(Network.LendingContracts.USDC, Network.VaultStrat))
	myPrintln(GetBalance(Network.LendingContracts.WETH, Network.VaultStrat))
	myPrintln(GetBalance(Network.LendingContracts.OSQTH, Network.VaultStrat))
	myPrintln(GetBalance(Network.LendingContracts.ATOKEN_USDC, Network.VaultStrat))

	GetTVL()

}

func VaultInfo3(accId int) {
	/*
		prepare:

			TotalValueLocked (tvl0, tvl1, tvlInUsdc )
			_usdcPrice
			_sqthPrice
			myShare
			myDeposits  <-- for my gain/loss
			totalSupply
			sqthAmount
			aaUsdcAmount
			tick, cLow, cHigh <- for in range


		implementation:
		   My Total Value:
			[$usdc]				<- tvlInUSDC  * myShare / totalSupply

		   Estimated Gains: [$usdc]		<-  todo
		   APY: [%]						<-  34%  (hard coded for now)

		   My Liquidity: 				<-  amount0, amount1, liquidity = GetUniAmounts(cLow, cHigh)
			[$usdc]						<-  (amount0 + amount1 * _usdcPrice / 1e18) * myShare / totalSupply
		   	usdc: [amount]				<-	 usdcAmount = amount0 *myShare / totalSupply / 1e6
		   	weth:  [amount]				<-   ethAmount = amount1 *myShare / totalSupply / 1e18

		   My Hedge : 				sqthBalance= sqth.balanceOf(stratContract),
									aTokenBalance= aTokenUsdc.balanceOf(stratContract)
			[$usdc]				<-	(sqthBalance * _sqthPrice /1e18 *_usdcPrice/1e18  + aTokenBalance) *myshare/totalSupply
		   	squeeth position: [eth]	 <-sqthBalance * _sqthPrice /1e18 * myshare/totalSupply
		   	short position:	[$eth]	 <- aTokenBalance *_usdcPrice/1e18  * myshare/totalSupply
		   	squeeth factor:	[0.789]	 <-  0.789

			My Income:
			[$usdc]					 <- myTotalInUSDC - myDepositInUSDC)
		   	usdc: [amount]				<-
		   	weth: [amount]				<-

		   Total Amounts in uni:  $usdc		uniAmount0 + (uniAmount1 * _usdcPrice /1e18)
		   Total value of ETH^2:   $usdc	sqthBalance * _sqthPrice /1e18 *_usdcPrice/1e18
		   Total Value of SHORT:   $usdc	aTokenBalance
		   Projected Income(this epoch): 	$usdc		(amount0 + amount1 * _usdcPrice / 1e18) * APY / 52
		   Projected APY(this epoch): X%				34%
	*/
	myTitle(".........Vault Info 3.........  ")

	_, stratInstance, vaultInstance := GetInstance4()

	//# prepare totalsupply
	totalSupply, err := vaultInstance.TotalSupply(&bind.CallOpts{})

	if err != nil {
		log.Println("totalsupply ", err)
	}

	//# prepare  myshare

	vaultTokenInstance, err := token.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Fatal("vaultinfo3-vaultTokenInstance,", err)
	}

	_, myAddress := GetSignatureByKey(Network.PrivateKey[accId])
	myShare, err := vaultTokenInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(myAddress))
	if err != nil {
		log.Fatal("vaultinfo3-myshare,", err)
	}

	assetHolder, err := vaultInstance.Assetholder(&bind.CallOpts{}, common.HexToAddress(myAddress))
	if err != nil {
		log.Fatal("vaultinfo3-myDeposits,", err)
	}

	//# prepare prices

	osqthPrice := GetPricePer(Network.LendingContracts.OSQTH, 3000)
	usdcPrice := GetPricePer(Network.LendingContracts.USDC, 500)

	//# Prepare tvl

	// tvl in vault
	token0Ins, _, _, _, _ := GetTokenInstance(Network.TokenA)
	token1Ins, _, _, _, _ := GetTokenInstance(Network.TokenB)
	vbalance0, _ := token0Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	vbalance1, _ := token1Ins.BalanceOf(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	// tvl in strategy
	amount0, amount1, err := stratInstance.GetTotalAmounts(&bind.CallOpts{})

	// calculate total tvl
	var tvlInUSDC *big.Int
	var tvl0 *big.Int
	var tvl1 *big.Int
	if err != nil {
		log.Fatal("GetTotalAmounts Error:", err)
	} else {
		tvl0 = new(big.Int).Add(vbalance0, amount0)
		tvl1 = new(big.Int).Add(vbalance1, amount1)
		tvlInUSDC = CalcValue(1, tvl0, tvl1, usdcPrice, 6, 18)
	}

	//# prepare uniswap position:

	poolInstance := GetPoolInstance()
	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})
	tick := slot0.Tick
	cLow, _ := stratInstance.CLow(&bind.CallOpts{})
	cHigh, _ := stratInstance.CHigh(&bind.CallOpts{})

	uniAmounts, _ := stratInstance.GetUniAmounts(&bind.CallOpts{}, cLow, cHigh)
	uniTotalInUSDC := CalcValue(1, uniAmounts.Amount0, uniAmounts.Amount1, usdcPrice, 6, 18)

	InRange := tick.Cmp(cLow) > 0 && tick.Cmp(cHigh) < 0

	//# prepare osqth amount:

	osqthSymbol, osqthAmount := GetBalance(Network.LendingContracts.OSQTH, Network.VaultStrat)

	osqthValueETH := CalcValue(-1, big.NewInt(0), osqthAmount, osqthPrice, 18, 18)
	osqthValueUSDC := CalcValue(1, big.NewInt(0), osqthValueETH, usdcPrice, 6, 18)

	//# prepare aToken amount:  1 aToken = 1 usdc

	aTokenSymbol, aTokenAmount := GetBalance(Network.LendingContracts.ATOKEN_USDC, Network.VaultStrat)

	//# prepare 输出：
	myTitle("Prepare ...")
	myPrintln("usdc Price:", usdcPrice)
	myPrintln("osqth Price:", osqthPrice)
	myPrintln("totalSupply:", totalSupply)
	myPrintln("myShare:", myShare)
	myPrintln("myDeposits:(usdc, weth)", assetHolder.Deposit0, assetHolder.Deposit1)
	myPrintln("TotalValueLocked(usdc, weth): ", tvl0, tvl1)
	myPrintln("TVL in usdc:  ", tvlInUSDC, ", $", Readable(tvlInUSDC, 6))
	myPrintln("uni USDC amount: ", uniAmounts.Amount0)
	myPrintln("uni WETH amount: ", uniAmounts.Amount1)
	myPrintln(osqthSymbol, " amount(osqth): ", osqthAmount)
	myPrintln(osqthSymbol, " value(eth): ", osqthValueETH)
	myPrintln(osqthSymbol, " value(usdc): ", osqthValueUSDC)
	myPrintln(aTokenSymbol, "amount(aTokenUsdc): ", aTokenAmount)

	//# calculation 输出：

	myTitle("My Total Value ...")

	MyDepositInUSDC := CalcValue(1, assetHolder.Deposit0, assetHolder.Deposit1, usdcPrice, 6, 18)
	MyTotalValue := CalcMyValue(tvlInUSDC, myShare, totalSupply)
	myTvl0 := CalcMyValue(tvl0, myShare, totalSupply)
	myTvl1 := CalcMyValue(tvl1, myShare, totalSupply)

	MyLiquidity := Readable(CalcMyValue(uniTotalInUSDC, myShare, totalSupply), 6)

	MyLiquidity_USDC := Readable(CalcMyValue(uniAmounts.Amount0, myShare, totalSupply), 6)
	MyLiquidity_WETH := Readable(CalcMyValue(uniAmounts.Amount1, myShare, totalSupply), 18)

	// my hedge, my income, my gain, my apy
	TotalHedge := new(big.Int).Add(osqthValueUSDC, aTokenAmount)
	MyIncome := new(big.Int).Sub(MyTotalValue, MyDepositInUSDC)
	myIncome_USDC := new(big.Int).Sub(myTvl0, assetHolder.Deposit0)
	myIncome_ETH := new(big.Int).Sub(myTvl1, assetHolder.Deposit1)

	EstimatedGains := MyIncome
	APY := BigIntToFloat64(MyIncome) / BigIntToFloat64(MyDepositInUSDC) * 365.0

	myPrintln("My Total Value: $", Readable(MyTotalValue, 6))
	myPrintln("EstimatedGains: $", Readable(EstimatedGains, 6))
	myPrintln("APY:", APY, "%")

	myTitle("My Liquidity ...")

	myPrintln("My Liquidity: $", MyLiquidity)
	myPrintln("My Liquidity USDC:", MyLiquidity_USDC)
	myPrintln("My Liquidity WETH:", MyLiquidity_WETH)

	myTitle("My Hedge")

	myPrintln("Total: $", Readable(CalcMyValue(TotalHedge, myShare, totalSupply), 6))
	myPrintln("USDC:", Readable(CalcMyValue(aTokenAmount, myShare, totalSupply), 6))
	myPrintln("ETH:", Readable(CalcMyValue(osqthValueETH, myShare, totalSupply), 18))

	myTitle("My Gain/Loss:")
	myPrintln("$", Readable(MyIncome, 6))
	myPrintln("USDC:", Readable(myIncome_USDC, 6))
	myPrintln("ETH:", Readable(myIncome_ETH, 18))
	// 		[$usdc]					 <- (ufees0 + ufees1 * _usdcPrice/ 1e18) * myShare/totalSupply
	// 	   	usdc: [amount]				<-	ufee0 * myShare/totalSupply
	// 	   	weth: [amount]				<- ufee1 * myShare/totalSupply

	myTitle("Total Amounts(in middle bottom)")
	myPrintln("Total Liquidity(USDC+ETH): $", Readable(uniTotalInUSDC, 6))
	myPrintln("Total Value of ETH ^2: $", Readable(osqthValueUSDC, 6))
	myPrintln("Total Value of Short: $", Readable(aTokenAmount, 6))
	//projectedIncome := CalcValue(1, uFees0, uFees1, usdcPrice, 6, 18)
	//myPrintln("Projected Income:", Readable(projectedIncome, 6))
	//myPrintln("Projected APY", projectedIncome / MyLiquidity_USDC *365)

	myTitle("in range? ...")
	myPrintln(InRange)
	myPrintln("cLow, tick, cHigh", cLow, tick, cHigh)
	priceHigh := 1 / (Powerf3(1.0001, int(cLow.Int64())) / Powerf3(10.0, (18-6)))
	priceNow := 1 / (Powerf3(1.0001, int(tick.Int64())) / Powerf3(10.0, (18-6)))
	priceLow := 1 / (Powerf3(1.0001, int(cHigh.Int64())) / Powerf3(10.0, (18-6)))
	myPrintln("priceNow", priceNow)
	myPrintln("priceLow", priceLow)
	myPrintln("priceHigh", priceHigh)

}

func Powerf3(x float64, n int) float64 {
	ans := 1.0
	for n != 0 {
		ans *= x
		n--
	}
	return ans
}
func GetFees() (*big.Int, *big.Int) {

	var uFees0 *big.Int
	var uFees1 *big.Int
	return uFees0, uFees1
}
func CalcMyValue(value *big.Int, myShare *big.Int, totalSupply *big.Int) *big.Int {

	if totalSupply.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	} else {
		return value.Mul(value, myShare).Div(value, totalSupply)
	}
}
func CalcValue(dir int, quoteAmount *big.Int, perAmount *big.Int, quotePrice *big.Int, quoteDecimal int, perDecimal int) *big.Int {

	var bigValue *big.Int
	if dir < 0 {
		// calc other for eth
		// perAmount / quotePrice * 10 ^ perDecimal + quoteAmount
		part1 := new(big.Int).Mul(perAmount, PowX(10, perDecimal))
		part2 := new(big.Int).Div(part1, quotePrice)
		bigValue = new(big.Int).Add(part2, quoteAmount)
	} else {
		// calc eth for other
		// perAmount * quotePrice / 10 ^ perDecimal + quoteAmount
		part1 := new(big.Int).Mul(perAmount, quotePrice)
		part2 := new(big.Int).Div(part1, PowX(10, perDecimal))
		bigValue = new(big.Int).Add(part2, quoteAmount)
	}
	return bigValue

}
func Readable(bigv *big.Int, decimal int) float64 {

	readableValue := BigIntToFloat64(bigv) / BigIntToFloat64(PowX(10, decimal))

	//return readableValue
	//return math.Ceil(readableValue*100) / 100  // 会出现科学计数
	//return math.Trunc(readableValue*1e2+0.5) * 1e-2    //
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", readableValue), 64)
	return value

}

func get_liquidity_0(x float64, sa float64, sb float64) float64 {
	return x * sa * sb / (sb - sa)
}

func get_liquidity_1(y float64, sa float64, sb float64) float64 {
	return y / (sb - sa)
}

func get_liquidity(x float64, y float64, sp float64, sa float64, sb float64) float64 {
	var liquidity, liquidity0, liquidity1 float64
	if sp <= sa {
		liquidity = get_liquidity_0(x, sa, sb)
	} else if sp < sb {
		liquidity0 = get_liquidity_0(x, sp, sb)
		liquidity1 = get_liquidity_1(y, sa, sp)
		liquidity = math.Min(liquidity0, liquidity1)
	} else {
		liquidity = get_liquidity_1(y, sa, sb)
	}
	return liquidity
}

func calculate_x(L float64, sp float64, sb float64) float64 {
	return L * (sb - sp) / (sp * sb)
}
func calculate_y(L float64, sp float64, sa float64) float64 {
	return L * (sp - sa)
}

func getTicks(p float64, a float64, b float64, xDecimals float64, yDecimals float64) (float64, float64) {

	//calc tick  p(i) = 1.0001i

	diffDecimals := math.Pow(10, xDecimals-yDecimals)

	// log(p , 1.0001)  ==  log(p)/ log(1.0001)
	tick := math.Log(p/diffDecimals) / math.Log(1.0001)
	tickLower := math.Log(a/diffDecimals) / math.Log(1.0001)
	tickUpper := math.Log(b/diffDecimals) / math.Log(1.0001)

	fmt.Printf("tick={%.f}\n", tick)
	fmt.Printf("tickLower={%.f}\n", tickLower)
	fmt.Printf("tickUpper={%.f}\n", tickUpper)
	fmt.Printf("\n")

	return tickLower, tickUpper

}

func SetVaultAddress(_address string, ind int64) {

	myTitle("Set Address in VaultBridge")
	instance, err := bridge.NewApi(common.HexToAddress(Network.VaultBridge), EthClient)

	if err != nil {
		log.Fatal("vaultbridgeInstance err:", err)
	}

	tx, err := instance.SetAddress(Auth, common.HexToAddress(_address), big.NewInt(ind))

	if err != nil {
		log.Fatal("setVaultBridge err: ", err)
	}

	TxConfirm(tx.Hash())

}

func GetVaultAddress(ind int64) {
	//bridgeAddress := "0xF065F185Cd6fE0feD4448A2Ef56f1017e7359277"
	bridgeAddress := Network.VaultBridge

	myPrintln("vault bridge address: ", bridgeAddress)

	instance, err := bridge.NewApi(common.HexToAddress(bridgeAddress), EthClient)

	if err != nil {
		log.Fatal("vaultBridgeInstance err:", err)
	}

	vaultAddress, err := instance.GetAddress(&bind.CallOpts{}, big.NewInt(ind))

	if err != nil {
		log.Fatal("getVaultBridge err: ", err)
	}

	fmt.Println("vaultAddress ", ind, ": ", vaultAddress)

}

func AuthAdmin(vaultAddr string, _admin string) {

	instance, err := admin.NewApi(common.HexToAddress(vaultAddr), EthClient)

	if err != nil {
		log.Fatal("vaultAdmin Instance err:", err)
	}

	exists, err := instance.AuthAdmin(&bind.CallOpts{}, common.HexToAddress(_admin))

	if err != nil {
		log.Fatal("auth admin  err: ", err)
	}

	fmt.Println("auth? ", exists)

}

func SetPermit(_admin string, level int) {

	bridgeInstance, err := bridge.NewApi(common.HexToAddress(Network.VaultBridge), EthClient)

	if err != nil {
		log.Fatal("vaultbridge Instance err:", err)
	}

	tx, err := bridgeInstance.SetPermit(Auth, common.HexToAddress(_admin), big.NewInt(int64(level)))

	if err != nil {
		log.Fatal("setpermit  err: ", err)
	}

	TxConfirm(tx.Hash())

}

func GetPermit(_admin string) {

	bridgeInstance, err := bridge.NewApi(common.HexToAddress(Network.VaultBridge), EthClient)

	if err != nil {
		log.Fatal("vaultbridge Instance err:", err)
	}

	level, err := bridgeInstance.GetPermit(&bind.CallOpts{}, common.HexToAddress(_admin))

	if err != nil {
		log.Fatal("setpermit  err: ", err)
	}

	myPrintln("permit level:", level)

}

func GetArbInstance() *arb.Api {

	arbAddress := "0xa9712b5e7C1537936Ba383B2632455A02D9d49B6"

	instance, err := arb.NewApi(common.HexToAddress(arbAddress), EthClient)
	if err != nil {
		log.Fatal("arb Instance err -- :", err)
	}
	return instance

}

func EthBalanceArb(_addr string) *big.Int {

	instance := GetArbInstance()

	bal, err := instance.EthBalance(&bind.CallOpts{}, common.HexToAddress(_addr))
	if err != nil {
		log.Fatal("eth balance vis Arb instance err:    ", err)
	}

	return bal

}

func ERC20Balance(_erc20 string, _owner string) {

	erc20Instance, name, symbol, decimals, _ := GetTokenInstance(_erc20)
	_ = name
	_ = decimals
	balance, _ := erc20Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(_owner))
	myPrintln(symbol, " in Wallet ", balance)

}

func InitSqueeth() {

	squeethInstance := GetSqueethInstance(Network.LendingContracts.OSQTH)

	tx, err := squeethInstance.Init(Auth, FromAddress)
	if err != nil {
		log.Fatal("squeeth init ", err)
	}

	TxConfirm(tx.Hash())

}

func MintSqueeth(to string, amount *big.Int) {

	squeethInstance := GetSqueethInstance(Network.LendingContracts.OSQTH)

	tx, err := squeethInstance.Mint(Auth, common.HexToAddress(to), amount)
	if err != nil {
		log.Fatal("squeeth mint ", err)
	}

	TxConfirm(tx.Hash())

}
