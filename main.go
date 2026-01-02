package main

import (
	"os"

	"github.com/seu-usuario/stress-test/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
