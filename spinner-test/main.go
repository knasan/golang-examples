package main

import "github.com/theckman/yacspin"
import "time"

func main() {

	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Suffix:          "\t system call",
		SuffixAutoColon: true,
		Message:         "starting ...",
		StopCharacter:   "✓",
		StopColors:      []string{"fgGreen"},
	}

	spinner, _ := yacspin.New(cfg)
	// handle the error

	spinner.Start()

	// doing some work
	time.Sleep(2 * time.Second)

	spinner.Message("booting system ...")

	// upload...
	time.Sleep(2 * time.Second)
	spinner.Message("booting finished")
	time.Sleep(2 * time.Second)

	spinner.Stop()
}