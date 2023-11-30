package main

import (
	"os"

	"github.com/thisdotEric/gitxtend/internal"
)

func main() {
	stat, _ := os.Stdin.Stat()

	// Check if the app is used to receive pipe outputs
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		internal.HandlePipeInput("")
	} else {
		internal.HandleCLIInput(os.Args)
	}
}
