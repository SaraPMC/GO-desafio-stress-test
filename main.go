package main

import (
	"os"

	"github.com/SaraPMC/stress-test/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
