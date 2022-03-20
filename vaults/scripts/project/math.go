package include

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
)

func Float64ToBigInt(val float64) *big.Int {

	//price = int(sqrt(price) * (1 << 96))
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(val)
	bigf := newNum.FloatString(0)

	//fmt.Println("val, bigf:", val,  bigf)
	//os.Exit(3)

	bigI, ok := new(big.Int).SetString(bigf, 10)
	if !ok {
		log.Fatal("float64 to bigInt err ", val, bigI)
	}

	return bigI

}

func BigFloatToBigInt(bigval *big.Float) *big.Int {

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func BigIntToFloat64(bigN *big.Int) float64 {

	bigF, _ := new(big.Float).SetString(bigN.String())

	f, _ := bigF.Float64()

	return float64(f)
}

func Pricef(priceInWei *big.Int, decimal int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(priceInWei.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimal)))

	//	fmt.Println(value) // 25.729324269165216041
	return value
}

func Str2BigInt(sn string) *big.Int {

	bsn, _ := new(big.Int).SetString(sn, 10)
	return bsn

}

func X1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}

// x ^ d , e.g. x=10, d=18 = 10^18
func PowX(x int64, d int) *big.Int {

	dd := strconv.Itoa(d - 1)

	s := fmt.Sprintf("1%0"+dd+"d", 0)

	ex, _ := new(big.Int).SetString(s, 10)

	bigx := big.NewInt(x)

	return bigx.Mul(bigx, ex)

}

func FloorFloat64ToBigInt(f64 float64) *big.Int {

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.")
	}

	return big.NewInt(int64(math.Floor(f64)))
}
func RoundFloat64ToBigInt(f64 float64) *big.Int {

	if f64 >= math.MaxInt64 || f64 <= math.MinInt64 {
		log.Fatal("f64 is out of int64 range.")
	}

	return big.NewInt(int64(math.Round(f64)))
}
