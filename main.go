package main

import (
	"fmt"
	"os"

	"github.com/thisdotEric/gx/internal"
)

func main() {
	stat, _ := os.Stdin.Stat()

	var err error

	// Check if the app is used to receive pipe outputs
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		err = internal.HandlePipeInput(os.Args)
	} else {
		err = internal.HandleCLIInput(os.Args)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
