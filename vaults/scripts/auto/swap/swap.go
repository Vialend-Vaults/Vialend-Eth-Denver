package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	_ "time"

	project "viaroot/scripts/project"
)

const (
	// -1  0forexact1
	// -2  exact0for1
	// 1   exact1for0
	// 2   1forexact0
	Exact1for0  = 1
	X1forexact0 = 2
	Exact0for1  = -2
	X0forexact1 = -1
)

/*
	usage:

	// get token0  , give token1
		swap exact1for0 amountInToken1
		swap 1forexact0 amountInToken0

	// get token1  , give token0
		swap exact0for1 amountInToken0
		swap 0forexact1	amountInToken1

*/

func main() {

	project.ConfigParser()
	project.Init(-1, -1)
	project.Quiet = false

	accoundid := 0     // define default swap account
	amount := int64(1) // default swap amount

	var op = 1 // -1,-2, sell; 1, 2, buy

	//做price move 和 genfee 要 检查钱包是否有钱 否则FAIL

	if len(os.Args) > 1 {

		if os.Args[1] == "buyeth" || os.Args[1] == "selleth" {
			if os.Args[1] == "buyeth" {
				op = 1
			} else {
				op = 0
			}
		} else if os.Args[1] == "checkprice" || os.Args[1] == "price" || os.Args[1] == "p" {
			project.PrintPrice()
			os.Exit(3)

		} else {
			log.Fatal("operation not support. ", os.Args[1])
		}

		if len(os.Args[2]) > 0 {
			amt, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
			if err != nil {
				log.Fatal("argument err ", err)
			}
			amount = int64(amt)

		} else {
			log.Fatal("missing amount")
		}

		if len(os.Args) > 3 && len(os.Args[3]) > 0 {
			acc, err := strconv.Atoi(strings.TrimSpace(os.Args[3]))
			if err != nil {
				log.Fatal("acc argument err ", err)
			}
			accoundid = int(acc)

		}

	} else {
		log.Fatal("missing argument, usage: swap exact1for0 amount")
	}

	if op == 1 {
		//buy eth
		project.TokenSwap(0, project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, project.PowX(amount*10, 8))
	} else {
		// sell eth
		project.TokenSwap(accoundid, project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, project.PowX(amount*10, 18))

	}

	project.PrintPrice()

}

// func main() {

// 	project.Init(-1, -1)

// 	project.Quiet = false

// 	accoundid := 0     // define default swap account
// 	amount := int64(1) // default swap amount

// 	var op = 1 // -1,-2, sell; 1, 2, buy

// 	var tn = 1

// 	//做price move 和 genfee 要 检查钱包是否有钱 否则FAIL

// 	if len(os.Args) > 1 {

// 		if os.Args[1] == "exact0for1" || os.Args[1] == "0forexact1" {
// 			if os.Args[1] == "exact0for1" {
// 				op = Exact0for1 //-2
// 			} else {
// 				op = X0forexact1 //-1
// 			}
// 		} else if os.Args[1] == "exact1for0" || os.Args[1] == "1forexact0" {
// 			if os.Args[1] == "exact1for0" {
// 				op = Exact1for0 //1
// 			} else {
// 				op = X1forexact0 //2
// 			}

// 		} else if os.Args[1] == "checkprice" || os.Args[1] == "price" || os.Args[1] == "p" {
// 			project.PrintPrice()
// 			os.Exit(3)

// 		} else {
// 			log.Fatal("operation not support. ", os.Args[1])
// 		}

// 		if len(os.Args[2]) > 0 {
// 			amt, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
// 			if err != nil {
// 				log.Fatal("argument err ", err)
// 			}
// 			amount = int64(amt)

// 		} else {
// 			log.Fatal("missing amount")
// 		}

// 		if len(os.Args) > 3 && len(os.Args[3]) > 0 {
// 			acc, err := strconv.Atoi(strings.TrimSpace(os.Args[3]))
// 			if err != nil {
// 				log.Fatal("acc argument err ", err)
// 			}
// 			accoundid = int(acc)

// 		}
// 		project.Swap(accoundid, amount, op, tn)

// 	} else {
// 		log.Fatal("missing argument, usage: swap exact1for0 amount")
// 	}

// }
