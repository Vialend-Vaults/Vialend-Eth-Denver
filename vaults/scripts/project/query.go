package include

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	tester "viaroot/deploy/tester"
	pool "viaroot/deploy/uniswap/v3/deploy/UniswapV3Pool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckPrice(do bool, interval time.Duration) {

	if !do {
		return
	}

	//var network = Networks[Networkid]

	poolAddress := common.HexToAddress(Network.Pool)
	poolInstance, err := pool.NewApi(poolAddress, EthClient)
	if err != nil {
		log.Fatal("pool.NewApi ", err)
	}

	SECONDS_AGO := []uint32{10}

	for i := 0; i < 500; i++ {
		fmt.Println(i, "/500")

		//TickCumulatives                    []*big.Int
		//SecondsPerLiquidityCumulativeX128s []*big.Int
		tickPrices, _ := poolInstance.Observe(&bind.CallOpts{}, SECONDS_AGO)
		before := tickPrices.TickCumulatives
		after := tickPrices.SecondsPerLiquidityCumulativeX128s

		fmt.Println("tickPrices.TickCumulatives:", before)
		fmt.Println("tickPrices.SecondsPerLiquidityCumulativeX128s:", after)

		a1 := after[0].Sub(after[0], before[0])
		twap := a1.Div(a1, big.NewInt(int64(SECONDS_AGO[0])))
		slot0, err := poolInstance.Slot0(&bind.CallOpts{})
		if err != nil {
			log.Fatal("slot0 ", err)
		}

		last := slot0.Tick
		fmt.Println("twap=", twap)
		fmt.Println("slot0.tick=", last)
		//f := new(big.Float).SetInt(twap)
		//f64 := Float64(f, 0)
		//f := 10001 / 10000
		//**twap
		//x := 1.0001

		//one2 := big.NewInt(twap).Exp(1.0001, f)

		//fmt.Println("twap\t{twap}\t 1.0001**twap:" 1.0001**twap, twap)
		//print(f"last\t{last}\t{1.0001**last}")
		//print(f"trend\t{last-twap}")
		//print()

		time.Sleep(interval * time.Second)
	}
}

func PoolInfo2(do bool) {

	//	poolAddress, err := newInstance.GetPool(&bind.CallOpts{}, Network.TokenA, Network.TokenB, fee)

	fmt.Println("----------------------------------------------")
	fmt.Println(".............. Pool Info2 ....... ")
	fmt.Println("----------------------------------------------")

	poolAddress := common.HexToAddress("0xaa16E934A327D500fdE1493302CeB394Ff6Ff0b2")
	fmt.Println("pool address:", poolAddress)

	poolInstance, err := pool.NewApi(common.HexToAddress(Network.Pool), EthClient)

	if err != nil {
		log.Fatal(err)
	}

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	fmt.Println("slot0.SqrtPriceX96:", slot0.SqrtPriceX96)
	fmt.Println("slot0.Tick:", slot0.Tick)

	_, price := getPrice(slot0.SqrtPriceX96, slot0.Tick)
	fmt.Println("getPrice ", price)

	sqrtPf := new(big.Float)
	///convert big Int to big Float
	sqrtPf.SetString(slot0.SqrtPriceX96.String())
	///convert big float to float64
	sp64, _ := sqrtPf.Float64()
	// operate float64
	sqrtPx962Price := math.Floor((sp64 * sp64) / math.Pow(2, 2*96))
	///verify the price
	fmt.Println("veryfy sqrtP*96 = ", sqrtPx962Price, ",  getPrice = ", price)

	fmt.Println("Price in decimals: ", price)

	feeRate, _ := poolInstance.Fee(&bind.CallOpts{})
	fmt.Println("Fee tier:", feeRate)

	liquidity, _ := poolInstance.Liquidity(&bind.CallOpts{})
	fmt.Println("liquidity in pool:", liquidity)

}

func Equation(do bool, doSet bool) {

	if !do {
		return
	}

	//var network = Networks[Networkid]

	poolAddress := common.HexToAddress("0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c")
	testerAddress := common.HexToAddress("0xa9c7752485080511344A90212C84c868628DB7A3")
	LPaddress := common.HexToAddress("0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E")

	poolInstance, err := pool.NewApi(poolAddress, EthClient)
	if err != nil {
		log.Fatal("pool.NewApi ", err)
	}

	slot0, err := poolInstance.Slot0(&bind.CallOpts{})
	if err != nil {
		log.Fatal("slot0 ", err)
	}

	fmt.Println("slot0=", slot0)

	testerInstance, err := tester.NewApi(testerAddress, EthClient)
	if err != nil {
		log.Fatal("isntance err: ", err)
	}

	tickLower := big.NewInt(16080)
	tickUpper := big.NewInt(27060)

	if doSet {
		NonceGen()
		tx, err := testerInstance.SetPosition(Auth, poolAddress, LPaddress, tickLower, tickUpper)
		if err != nil {
			log.Fatal("setposition err:", err)
		}
		_ = tx

		Readstring("press any key to continue")
	} /*
		p1, p2, p3, p4, p5, p6, _ := testerInstance.GetPosition(&bind.CallOpts{})
		if err != nil {
			log.Fatal("info err:", err)
		}
		fmt.Println(p1, p2, p3, p4, p5, p6)
	*/

	NonceGen()
	tx, err := testerInstance.CallPosition(Auth, poolAddress, LPaddress, tickLower, tickUpper)
	if err != nil {
		log.Fatal("call trans err:", err)
	}
	_ = tx

	Readstring("press any key to continue")

	//	p1, p2, p3, p4, p5,amount0,amount1 _ := testerInstance.CallPosition(&bind.CallOpts{}, poolAddress, LPaddress, tickLower, tickUpper)
	p1, p2, p3, p4, p5, p6, amount0, amount1, _ := testerInstance.GetPosition(&bind.CallOpts{})

	if err != nil {
		log.Fatal("info err:", err)
	}
	fmt.Println(p1, p2, p3, p4, p5, p6, amount0, amount1)

}
