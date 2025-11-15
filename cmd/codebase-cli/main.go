package main

import (
	"os"

	"github.com/codebase-interface/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
