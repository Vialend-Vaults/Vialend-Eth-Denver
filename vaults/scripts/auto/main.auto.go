package main

import (
	"fmt"
	_ "time"

	project "viaroot/scripts/project"
)

func main() {

	fmt.Println("Env: NetworkId=", project.Networkid, ",client=", project.Network.ProviderUrl[project.ProviderSortId])
	project.Init(-1, -1)

	project.Quiet = true

	// project.MovePrice(project.Sell)
	// project.MovePrice(project.Buy)

	project.GenFees(4, 5) // swap times

	//project.Rebal(20) //interval seconds

	//AutoBoth() // Rebal(30) + GenFees

}

func AutoBoth() {

	go project.GenFees(10, 10) // {swap num of times, account, amount , sleepSeconds}

	project.Rebal(10, 20) // max rebalance run ,  interval seconds between each rebalance

}
