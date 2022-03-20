package include

import (
	"bufio"
	_ "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	_ "os"
	"strings"
	"time"

	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var TickCount int

/// read file
/// parse file
/// get range
/// calc range
/// trigger Strategy1

func ReadSignal() int64 {

	// read signal file
	content, _ := ioutil.ReadFile("C:\\Program Files (x86)\\MetaTrader 4 - 1\\tester\\files\\_Signal.txt")

	// trim any space or hard cariage
	s := strings.TrimSpace(string(content))

	// parse string to int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("parse error:", err)
		return -1
	}

	return i

}
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
func ReadTick() int {

	fi := "../../file/tick"
	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	sln, e := Readln(r)

	tickArray := []string{}

	for e == nil {

		s := strings.TrimSpace(sln)

		if len(s) > 0 {
			tickArray = append(tickArray, s) // a == [3 4]
		}
		sln, e = Readln(r)
	}

	tempCount := len(tickArray)

	newTick := int(0)

	if tempCount != TickCount {

		_tick, err := strconv.Atoi(tickArray[0])

		if err != nil {
			log.Fatal("strconv.atoi err:", err)
		}

		fmt.Println("new tick:", _tick)

		newTick = _tick

		TickCount = tempCount

		return newTick

	}

	return 0
}

func TickReport(_tick *big.Int) {

	var buf []byte

	d0, err := ioutil.ReadFile("../file/tick")
	if err != nil {
		log.Fatal("Readfile err:", err)
	}

	s := strings.TrimSpace(string(d0))

	d1 := []byte(_tick.String() + "\n")

	buf = append(d1, s...)

	err = os.WriteFile("../file/tick", buf, 0644)
	if err != nil {
		log.Fatal("writefile err:", err)
	}

}

//  pass in range , run strategy1
func doRebal(rng int64) {

	//	fmt.Println("Rebalance Triggered , new range:", rng)
	Strategy1(rng, int64(1))

}

// calculate range
// max : limited number to exit loop, -1= infinite
// interval : check in x seconds
func Rebal2(max int, interval int) {

	i := 0
	oldrange := int64(0)

	for {

		newRange := ReadSignal()
		i++

		//fmt.Println("Reading range:........(", i, ")", "{", newRange, "}")

		if newRange > 0 {

			// only do rebalance when new range is detected
			if newRange != oldrange {
				oldrange = newRange
				doRebal(newRange)
			}

			// in testing mode, in case something went wrong
			if max != -1 && i > max { // 30 s per check, 1000 checks for 8 hours
				os.Exit(3)
			}

		}

		//testing mode
		time.Sleep(time.Duration(interval) * time.Second)

	}

}

func Rebal(max int, interval int) {

	i := 0
	newTick := int(0)

	for {

		// in testing mode, in case something went wrong
		if max != -1 && i > max { // 30 s per check, 1000 checks for 8 hours
			os.Exit(3)
		}

		newTick = ReadTick()
		newRange := int64(800)

		if newTick != 0 && i > 0 {
			if newRange > 0 {

				fmt.Println()
				fmt.Println("REBALANCE IS TRIGGERED...STRATEGY ====  001")
				doRebal(newRange)
				//fmt.Println("doRebalance...")

			}
		} else {
			fmt.Println("Monitoring Rebalance Signal.....", i)

		}

		//testing mode
		time.Sleep(time.Duration(interval) * time.Second)

		i++

	}

}

func check(msg string, e error) {
	if e != nil {
		log.Fatal(e, msg)
	}
}

/*
zeroforone = true,
	1. liquidity pool change: 	more eth, less usdc
	2. swapper: 				swap eth for usdc,
	3. trading action: 			sell eth , buy usdc
	4. vault change: 			more eth, less usdc
	5. price change: 			eth price go down

zeroforone = false,
	1. liquidity pool change: 	less eth, more usdc  (less 0, more 1 )
	2. swapper: 				swap usdc for eth,   (swap usdc for eth)
	3. trading action: 			sell usdc , buy eth
	4. vault change: 			less eth, more usdc
	5. price change: 			eth price go up

	//@op
	// -1  0forexact1
	// -2  exact0for1
	// 1   exact1for0
	// 2   1forexact0

*/
func Swap(accountId int, amount int64, op int, times int) {

	_pool := Network.Pool // check networkid

	var zeroForOne bool

	if op < 0 {
		// eth price go down
		zeroForOne = true
	}
	if op > 0 {
		// eth price go up
		zeroForOne = false
	}

	for i := 0; i < times; i++ {
		if op > 1 || op < -1 {
			Swap0(accountId, big.NewInt(amount), zeroForOne, _pool)
		} else {
			Swap1(accountId, big.NewInt(amount), zeroForOne, _pool)
		}
	}

}

// generate fees by swap in the uniswap pool
func GenFees(t int, sleepSeconds time.Duration) {
	accountid := Account
	amount := int64(1)

	// -1  0forexact1
	// -2  exact0for1
	// 1   exact1for0
	// 2   1forexact0

	for i := 0; i < t; i++ {

		Swap(accountid, amount, -2, 1)
		Swap(accountid, amount, 2, 1)

		time.Sleep(sleepSeconds * time.Second)
	}
}

func MonitorAll() {
	// todo all vault
	go MonitorVault(3, 1, -1)
	go MonitorVault(4, 1, -1)

}

func MonitorVault(nid int, acc int, maxt int) {

	var lasttick = big.NewInt(0)

	//temp := Networkid
	Networkid = nid
	Network = Networks[nid]

	///require governance. always use account 0 as the deployer
	Auth = GetSignature(Networkid, acc)

	fmt.Println(">>> Monitoring tick range: {vault:", Network.Vault, "} >>>>>>> ")

	i := 0
	for {
		i = i + 1
		if i < maxt || maxt == -1 {
			myPrintln(i, maxt, "Time: ", time.Now().Format("15:04:05"))

			//CheckTVL()

			inRange, tick := CheckRange(lasttick)
			if !inRange {

				//TickReport(tick)    // write new tick to file  <-- old way

				fmt.Println("**** Rebalance triggered **** ")
				go doRebal(SetStrategy())
				time.Sleep(5 * time.Second)
				fmt.Println()

			} else {

				time.Sleep(2 * time.Second)
			}

			lasttick = tick

		}
	}

}

func CheckRange(lasttick *big.Int) (bool, *big.Int) {

	vaultInstance := GetVaultInstance()
	poolInstance := GetPoolInstance()

	slot0, _ := poolInstance.Slot0(&bind.CallOpts{})

	tick := slot0.Tick

	qTickLower, _ := vaultInstance.CLow(&bind.CallOpts{})
	qTickUpper, _ := vaultInstance.CHigh(&bind.CallOpts{})

	myPrintln("cLow, tick, cHigh , lasttick :", qTickLower, tick, qTickUpper, lasttick)
	//fmt.Println("cLow, tick, cHigh , lasttick :", qTickLower, tick, qTickUpper, lasttick)

	//no new tick from last
	if tick.Cmp(lasttick) == 0 {
		return true, tick
	}

	inRange := tick.Cmp(qTickLower) > 0 && tick.Cmp(qTickUpper) < 0

	tickInfo(inRange, tick, qTickLower, qTickUpper)

	return inRange, tick

}
func tickInfo(inRange bool, tick *big.Int, tickLow *big.Int, tickHigh *big.Int) {

	if !inRange {

		if tick.Cmp(tickLow) <= 0 {

			fmt.Println("** Out range! poke down ", tick, tickLow, " {", new(big.Int).Sub(tick, tickLow), ",", new(big.Int).Sub(qTickUpper, tick), "}")
		} else if tick.Cmp(tickHigh) >= 0 {
			fmt.Println("** Out range! poke up", tick, tickHigh, " {", new(big.Int).Sub(tick, tickLow), ",", new(big.Int).Sub(tickHigh, tick), "}")
		} else {
			fmt.Println("** something wrong , tick still in range? ", tickLow, tick, tickHigh, " {", new(big.Int).Sub(tick, tickLow), ",", new(big.Int).Sub(tickHigh, tick), "}")
		}

	} else {
		fmt.Println(">>> In range ")
		fmt.Println(">>> ticklower, tick,  tickupper:", tickLow, tick, tickHigh, " {", new(big.Int).Sub(tick, tickLow), ",", new(big.Int).Sub(tickHigh, tick), "}")
	}

	PrintPrice()
	fmt.Println()

}
func SetStrategy() int64 {

	// strategy scenarios:
	// which pairs ?
	// market bull?  rng = wider
	// market bear?  rng = wider
	// market flat?  rng = tight
	// decide by which indicator?  bollinger?  tma ?  envelope ?  ma ? macd? rsi?
	// timeframe? D1,H4, H1?

	rng := int64(200)
	return rng

}
