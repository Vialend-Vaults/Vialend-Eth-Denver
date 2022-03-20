package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	_ "time"

	project "viaroot/scripts/project"
)

func main() {

	project.Quiet = true

	iteration := -1 //-1 infinite

	nid := project.Networkid
	acc := 1

	if len(os.Args) > 1 {
		if len(os.Args[1]) > 0 {
			_nid, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
			if err != nil {
				log.Fatal("argument 1 err ", err)
			}
			nid = _nid

		}
	}

	if len(os.Args) > 2 {
		_acc, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
		if err != nil {
			log.Fatal("argument 2 err ", err)
		}

		acc = _acc

	}

	if nid != 3 && nid != 4 && nid != 5 {
		log.Fatal("Wrong networkid ", nid)
	}

	project.Init(nid, acc)

	// s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	// s.Color("red")                                              // Set the spinner color to red
	// s.Start()                                                   // Start the spinner

	project.MonitorVault(nid, acc, iteration)

}
