package include

import (
	"time"

	"github.com/briandowns/spinner"
)

func spin(n int) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()
	if n > 0 {
		time.Sleep(4 * time.Second) // Run for some time to simulate work
		s.Stop()
	}

}
