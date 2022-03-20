package main

import (
	_ "os"
	_ "time"

	p "viaroot/scripts/project"
)

func main() {

	//p.ViaStratUniCompPublicList()
	//return

	//emergency
	//p.ViaStratUniCompPublicFunc(1, int64(1), int64(1))

	//rebalance
	//p.ViaStratUniCompPublicFunc(2, int64(1), int64(1))

	//getPrice
	p.ViaStratUniCompPublicFunc(3)

}
