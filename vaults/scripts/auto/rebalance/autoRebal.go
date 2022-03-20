package main

import (
	_ "time"

	project "viaroot/scripts/project"
)

func main() {

	project.Init(-1, -1)

	project.Quiet = true

	AutoRebalance(0) // 1 + GenFees

}

func AutoRebalance(_genfee int) {

	if _genfee > 0 {
		go project.GenFees(10, 10) // {swap num of times, account, amount , sleepSeconds}
	}

	project.Rebal(-1, 5) // max rebalance run ,  interval seconds between each rebalance

}
