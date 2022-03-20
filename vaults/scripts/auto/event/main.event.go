package main

import (
	"fmt"
	"os"
	_ "os"
	"strconv"
	"strings"
	_ "time"

	p "viaroot/scripts/project"
)

func main() {

	p.Init(-1, -1)

	contract := p.Network.Vault

	//contract := "0x3acA8F82e5fca1f437D71BEc742F472b843451d0"

	if len(os.Args) > 1 {
		if len(os.Args[1]) > 0 {
			comd := os.Args[1]

			//event -l Deposit, or -l
			if comd == "-l" || comd == "-listen" {
				if len(os.Args) > 2 {
					p.Event(contract, os.Args[2], 0, 0)
				} else {
					p.Event(contract, "", 0, 0)
				}

			}

			if comd == "-f" || comd == "-find" {

				//event -f Widthdraw 6046718 6046718

				_from, _ := strconv.Atoi(strings.TrimSpace(os.Args[3]))
				_to, _ := strconv.Atoi(strings.TrimSpace(os.Args[4]))

				p.Event(contract, os.Args[2], _from, _to)
			}

		}
	} else {

		fmt.Println("invalid paramters. e.g.: event -l Deposit")
	}

}
