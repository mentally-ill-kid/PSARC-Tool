package main

import (
	"os"

	"github.com/mentally-ill-kid/PSARC-Tool/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}
}
