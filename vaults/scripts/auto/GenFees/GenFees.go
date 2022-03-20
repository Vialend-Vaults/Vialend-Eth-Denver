package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	_ "time"

	project "viaroot/scripts/project"
)

func main() {

	project.Init(-1, -1)

	project.Quiet = true

	swapTimes := 4
	inSec := time.Duration(10)
	//genFees 4
	if len(os.Args) > 1 {

		if os.Args[1] == "checkprice" || os.Args[1] == "price" || os.Args[1] == "p" {
			project.PrintPrice()
			os.Exit(3)

		} else {
			st, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
			if err != nil {
				log.Fatal("argument err ", err)
			}
			swapTimes = st

		}

	}

	project.GenFees(swapTimes, inSec) // swap times

}
