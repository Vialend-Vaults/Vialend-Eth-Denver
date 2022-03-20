package include

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var Quiet = false
var Auto = true //auto check pending status

func MyTitle(title string) {
	myTitle(title)
}

func MyPrintln(a ...interface{}) {
	myPrintln(a)
}

func myTitle(title string) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".......................", title, "...................")
	fmt.Println("----------------------------------------------")

}

func myPrintln(a ...interface{}) {
	if !Quiet {
		fmt.Println(a)
	}
}

func Readstring2(msg string, pause bool) {

	fmt.Println(msg)

	time.Sleep(Network.PendingTime * time.Second)

	if pause {
		fmt.Println("paused. press anykey to continue")

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.TrimSuffix(strings.TrimSpace(text), " \n")
			break
		}
	}
}
func Readstring(msg string) string {

	fmt.Println(msg)

	if Auto {
		time.Sleep(Network.PendingTime * time.Second)
		return ""
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.TrimSuffix(strings.TrimSpace(text), " \n")

			return text

		}
	}
}

func Pause(msg string) string {

	fmt.Println(msg)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.TrimSuffix(strings.TrimSpace(text), " \n")
		return text
	}
}

func Sleep(tn time.Duration) {
	time.Sleep(tn * time.Millisecond)
}
