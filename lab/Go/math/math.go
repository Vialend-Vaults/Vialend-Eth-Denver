package main

/*

maybeprint()  with quiet variable

apy()

cantor  pair function
// https://en.wikipedia.org/wiki/Cantor_function
// https://en.wikipedia.org/wiki/Pairing_function
// https://gist.github.com/hannesl/8031402


*/

import (
	_ "flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
)

var quiet bool

func main() {
	//	cantor_pair_reverse2(cantor_pair_calculate2(big.NewInt(1), big.NewInt(23)))
	//	fmt.Println()
	// quiet = true
	// myPrintln("abc", 1, 2.3)

	cnt := 361

	fmt.Println(cnt%180 == 0)
	return

	fmt.Println(PowX(2, 6))
	return

	// quiet = true // normally you wouldn't change this manually you'd let flag.Parse do it

	// maybePrintf("Goodbye\n", 3)
	// return

	// apy()
	// return

	// power(1.0001, -192934)
	// return

	a1 := cantor_pair_calculate(1, 9)

	a2 := cantor_pair_calculate(6, 8)

	a3 := cantor_pair_calculate(a1, a2)

	x1, x2 := cantor_pair_reverse(a3)
	cantor_pair_reverse(x1)
	cantor_pair_reverse(x2)
	// cantor_pair_reverse(cantor_pair_calculate(33, 1))

	return

	x := int64(1)
	y := int64(33)
	z := int64(1)

	if x > y {
		cantor_pair_reverse(cantor_pair_calculate(y, x))

	} else {
		cantor_pair_reverse(cantor_pair_calculate(x, y))
		z = -1 * z
	}
	fmt.Println("z ", z)

	// x, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000", 10)
	// fmt.Println(X1E18X(x))

	//	convert()
	//bigInt()

	//bigFloat()

	//	test()

}
func myPrintln(a ...interface{}) {
	if !quiet {
		fmt.Println(a)
	}
}

func maybePrintf(format string, a ...interface{}) {
	if !quiet {
		fmt.Printf(format, a...)
		fmt.Println()
	}

}

func apy() {
	/*
			*My address: 0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
		myShares in vault  2307744692237727414
		totalSupply in vault  18722457719334674299
		deposit block info: 5746880 1635375579
		timediff from now: { 2730 1635378309 1635375579 }
		oneyearINsec  31536000
		totalTVL  &{2994806194084426860 3626238001}
		deposit0,deposit1 { 1000000000000000000 300000000 }
		mytvl0, mytvl1 { 369142139471452104 446972914 }
		APY: -1022.9096660623015

		*My address: 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
		myShares in vault  16413390391638327568
		totalSupply in vault  18722457719334674299
		deposit block info: 5746883 1635375624
		timediff from now: { 2685 1635378309 1635375624 }
		oneyearINsec  31536000
		totalTVL  &{2994806194084426860 3626238001}
		deposit0,deposit1 { 2000000000000000000 3322000000 }
		mytvl0, mytvl1 { 2625452488539579473 3179008913 }
		APY: 169.43076334253118
	*/

	// get block info

	timestamp := int(1635375624)

	htimestamp := int(1635378309)

	timediff := htimestamp - timestamp // diff in seconds

	oneyearINsec := big.NewInt(365 * 24 * 60 * 60)

	deposit0 := big.NewInt(1000000000000000000)
	deposit1 := big.NewInt(300000000)

	mytvl0 := big.NewInt(369142139471452104)
	mytvl1 := big.NewInt(446972914)

	usPrice0 := 226.00  // eth/usd
	usPrice1 := 0.99999 // usdc/usd

	fd0 := BigIntToFloat64(deposit0) * usPrice0

	fd1 := BigIntToFloat64(deposit1) * usPrice1 * 1e12

	fm0 := BigIntToFloat64(mytvl0) * usPrice0

	fm1 := BigIntToFloat64(mytvl1) * usPrice1 * 1e12

	fdd := fd0 + fd1
	fmm := fm0 + fm1

	APY := (fmm - fdd) / float64(timediff) * BigIntToFloat64(oneyearINsec) / fdd

	fmt.Println("APY:", APY)
}

func BigIntToFloat64(bigN *big.Int) float64 {

	bigF, _ := new(big.Float).SetString(bigN.String())

	f, _ := bigF.Float64()

	return float64(f)
}

/*
实际上，只要是乘以或除以一个整数，均可以用移位的方法得到结果，如：
1.0001 ** b = (10001 ** b) * ( 1/10000 ** b ) =


a=a*9
分析a*9可以拆分成a*(8+1)即a*8+a*1, 因此可以改为：    a=(a<<3)+a

a=a*7
分析a*7可以拆分成a*(8-1)即a*8-a*1, 因此可以改为：    a=(a<<3)-a

关于除法读者可以类推, 此略.
*/
func power(x float64, n int64) float64 {

	var result = 1.0

	for n > 0 {
		if n&1 == 1 { // n是 奇数 则 1， 偶数则0
			result = x * result
		}

		n = n >> 1

		x = x * x
	}

	fmt.Println(result)
	return result

}

func cantor_pair_calculate2(x *big.Int, y *big.Int) *big.Int {
	//result := ((x+y)*(x+y+1))/2 + y

	x.Add(x, y).Mul(x, new(big.Int).Add(x, big.NewInt(1))).Div(x, big.NewInt(2)).Add(x, y)
	fmt.Println(x)

	return x
}

/**
 * Return the source integers from a cantor pair integer.
 */
func cantor_pair_reverse2(z *big.Int) {

	negOne := big.NewFloat(-1)
	eight := big.NewFloat(8)
	one := big.NewFloat(1)
	two := big.NewFloat(2)
	three := big.NewFloat(3)
	zFloat := new(big.Float)
	zFloat.SetString(z.String())

	//t := math.floor((-1 + math.sqrt(1+8*z)) / 2)
	zFloat8 := new(big.Float).Mul(zFloat, eight)
	z81 := new(big.Float).Add(one, zFloat8)
	sqrtZ81 := new(big.Float).Sqrt(z81)
	sqrtZ81S1 := new(big.Float).Add(negOne, sqrtZ81)
	fmt.Println("z81.Sqrt(z81).Quo(z81, two)=", sqrtZ81S1)
	t := new(big.Float).Quo(sqrtZ81S1, two)

	fmt.Println("t:", t)

	t = big.NewFloat(24)

	//x := t*(t+3)/2 - float64(z)
	tx := new(big.Float)
	tx.SetString(t.String())

	t3 := new(big.Float).Add(tx, three)

	tx.Mul(tx, t3).Quo(tx, two)
	fmt.Println("t*(t+3)/2:", tx)

	tx.Sub(tx, zFloat)

	x := tx.Mul(tx, new(big.Float).Add(tx, three).Quo(tx, two))
	x.Sub(x, zFloat)

	//y := z - t*(t+1)/2
	ty := new(big.Float)
	ty.SetString(t.String())
	y := ty.Add(zFloat, ty.Mul(new(big.Float).Add(ty, one), t).Quo(ty, two))
	fmt.Println(x, y)
}

/**
Cantor pairing functions in PHP. Pass any two positive integers and get a unique integer back. Feed the unique integer back into the reverse function and get the original integers back.
https://gist.github.com/hannesl/8031402

 * Calculate a unique integer based on two integers (cantor pairing).
*/
func cantor_pair_calculate(x int64, y int64) int64 {

	fmt.Println(x, y)

	result := ((x+y)*(x+y+1))/2 + y

	fmt.Println(result)

	return result
}

/**
 * Return the source integers from a cantor pair integer.
 */
func cantor_pair_reverse(z int64) (int64, int64) {

	t := math.Floor((-1 + math.Sqrt(1+8*float64(z))) / 2)
	x := t*(t+3)/2 - float64(z)
	y := float64(z) - t*(t+1)/2

	fmt.Println(x, y)

	return int64(x), int64(y)

}

func test() {
	var price float64
	var sqrtPriceX96 float64

	fmt.Println("----------------------------------------------")
	fmt.Println(".........移位 左移， 右移 .........  ")
	fmt.Println("----------------------------------------------")
	/// bit shift
	// i << j  => i* (1<<j) => i * pow(2,j)
	fmt.Println("左移:")
	fmt.Println("i <<j  = ", 3<<2)
	fmt.Println("i * (1<<j) =", 3.0*(1<<2))
	fmt.Println("i * pow(2,j) =", 3.0*(math.Pow(2, 2)))

	fmt.Println("右移:")
	fmt.Println("i>>j ", 164>>4)
	fmt.Println("i/pow(2,j) ", math.Floor(164/math.Pow(2, 4)))

	/// example
	price = 1e18 / 2000e6

	fmt.Println("price:", price)

	/// price to sqrtPx96   sqrt(price) << 96
	sqrtPriceX96 = math.Floor(math.Sqrt(price) * (1 << 96))
	fmt.Println("sqrPriceX96:", sqrtPriceX96)

	/// sqrtPx96 to price  sqrtPriceX96 ** 2  >> 96*2
	sqrtPx962Price := math.Floor((sqrtPriceX96 * sqrtPriceX96) / math.Pow(2, 2*96))
	fmt.Println("sqrtP*96 -> Price ", sqrtPx962Price, " == ", price)

	/// convert float64 to big.int
	s := fmt.Sprintf("%.0f", sqrtPriceX96)
	bigSqrtPX96, _ := new(big.Int).SetString(s, 10)

	fmt.Println("float to bigint:", bigSqrtPX96, "==", sqrtPriceX96)

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println(".........深度copy 大数运算 big.Int 一次赋值引发的血案 .........  ")
	fmt.Println("----------------------------------------------")
	///------------ big int value test ， 深度copy bigint variable
	fmt.Println("\nbig int value test....")
	// 初始化两个变量: a = 1, b = 2
	a := big.NewInt(1)
	b := big.NewInt(2)

	// 打印交换前的数值
	fmt.Printf("a = %v   b = %v\n", a, b)
	// 使用中間變量法進行交換
	tmp := big.NewInt(0)
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)

	// 交換完成, 對中間變量加100
	tmp.Add(tmp, big.NewInt(100))

	// 打印交換後的結果
	fmt.Printf("a = %v    b = %v   tmp = %v\n", a, b, tmp)

	fmt.Println(X1E18(5000000000000005))

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println("......... big float Quo 除法 ， big int Div.........  ")
	fmt.Println("----------------------------------------------")
	//var bv = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(25))
	var bv, _ = new(big.Int).SetString("25729324269165216041", 10)
	fmt.Println("big int 整数", bv)
	fbalance := new(big.Float)
	fbalance.SetString(bv.String())
	fValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(fValue) // 25.729324269165216041

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big.float 百分比计算 .........  ")
	fmt.Println("----------------------------------------------")
	// calculate percentable   a * 100 / B = x %
	bal := big.NewInt(125)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	fmt.Println(bal, "*", 100, "/", fbal)
	fmt.Println(new(big.Float).Quo(fbal, big.NewFloat(1200)))
	x := new(big.Float).Quo(new(big.Float).Mul(fbal, big.NewFloat(100)), big.NewFloat(1200))
	fmt.Println(x, "%")

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println(".........向下取整， 向上取整， 四舍五入 .........  ")
	fmt.Println("----------------------------------------------")

	var f = float64(4.5)
	fmt.Println(f, "向上取整= ", math.Ceil(f))           // 5
	fmt.Println(f, "向下取整= ", math.Floor(f))          // 4
	fmt.Println(f, "四舍五入= ", int(math.Floor(f+0.5))) //

}

func convert() {

	tick, _ := new(big.Int).SetString("-293422343333334", 10)

	fmt.Println(tick.Int64())
	fmt.Println(tick.IsInt64())

	fmt.Println(float64(tick.Int64()))

	tick, _ = new(big.Int).SetString("-29342223442332333333333333333333333334", 10)

	fmt.Println(tick.Int64())

	tick, _ = new(big.Int).SetString("-29342223423423423423423434333332333333333333333333333333333334", 10)

	fmt.Println(tick.Int64())
	fmt.Println(tick.IsInt64())
	fmt.Println(float64(tick.Int64())) //wrong because int overflow

	tick = big.NewInt(-23748)
	fmt.Println(tick.Quo(tick, big.NewInt(60)))
	tick = big.NewInt(-23748)
	fmt.Println(tick.Div(tick, big.NewInt(60)))

	tick = big.NewInt(23748)
	fmt.Println(tick.Quo(tick, big.NewInt(60)))
	tick = big.NewInt(23748)
	tickSpacing := big.NewInt(60)
	fmt.Println(tick.Div(tick, tickSpacing).Mul(tick, tickSpacing))

	tick, _ = new(big.Int).SetString("-23", 10)
	fmt.Println(new(big.Int).Quo(tick, big.NewInt(6)))
	tick, _ = new(big.Int).SetString("-23", 10)
	fmt.Println(tick.Div(tick, big.NewInt(6)))

}
func bigFloat() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big float test .........  ")
	fmt.Println("----------------------------------------------")

	bigf, _ := new(big.Float).SetString("234923490230942903402349")

	fmt.Println(bigf)

	f1, _ := new(big.Float).SetString("2.32342343333333333433333333")

	fmt.Println(f1)

	f2, _ := new(big.Float).SetString("2323423433333333334333333332342342342342423433333333333333333333333333333333333333333333333333")
	fmt.Println(f2)

	a1, _ := new(big.Float).SetString("3")
	a2, _ := new(big.Float).SetString("2")
	fmt.Println(a1.Quo(a1, a2))

	fmt.Println(a1.Quo(bigf, f2))

	bigd, _ := new(big.Int).SetString("234923490230942903402349", 10)
	d1, _ := new(big.Int).SetString("23234234333333333343333333323423423423424234333333333333333333333333333333333333333333333333335", 10)

	fmt.Println(d1)

	fmt.Println(bigd.Div(bigd, d1))
}

func bigInt() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big int test .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("12345642342342341234564234234234123456423423423412345642342342341234564234234234")
	fmt.Println(FloatToBigInt(float64(12345642342342341234564234234234123456423423423412345642342342341234564234234234)))
	a, _ := new(big.Int).SetString("12345642342342341234564234234234123456423423423412345642342342341234564234234234", 10)
	b, _ := new(big.Int).SetString("4234123456423423423412345642342342341234564234234234", 10)
	c, _ := new(big.Int).SetString("28", 10)
	d, _ := new(big.Int).SetString("2", 10)
	e, _ := new(big.Int).SetString("3", 10)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("a*b", new(big.Int).Mul(a, b))
	fmt.Println("a/b", new(big.Int).Div(a, b))
	fmt.Println("c*d/e", c.Mul(c, d).Div(c, e))

	fmt.Println("c", c)

	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)

	fmt.Println("c+d+e", c.Add(c, d).Add(c, e))
	fmt.Println("(c+d)/e", c.Add(c, d).Div(c, e))
	fmt.Println("(c+d)/e", c.Add(c, d).Div(c, e))

	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)

	fmt.Println("c+d/e", d.Div(d, e).Add(d, c))
	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)
	fmt.Println("c+d/e", c.Add(c, d.Div(d, e)))

}

func FloatToBigInt(val float64) *big.Int {

	//price = int(sqrt(price) * (1 << 96))
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(val)
	bigf := newNum.FloatString(0)

	//fmt.Println("val, bigf:", val,  bigf)

	bigI, ok := new(big.Int).SetString(bigf, 10)
	if !ok {
		log.Fatal("float64 to bigInt err ", val, bigI)
	}

	return bigI

}

func X1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}

func X1E18X(bigx *big.Int) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)

	return bigx.Mul(bigx, e18)
}

func PowX(x int64, d int) *big.Int {

	dd := strconv.Itoa(d)
	s := fmt.Sprintf("1%0"+dd+"d", 0)
	ex, _ := new(big.Int).SetString(s, 10)

	bigx := big.NewInt(x)

	return bigx.Mul(bigx, ex)

}
