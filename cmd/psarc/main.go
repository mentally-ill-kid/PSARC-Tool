package main

import (
	"fmt"
	"os"

	"github.com/mentally-ill-kid/PSARC-Tool/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "pack":
		Pack(os.Args[:2])
	case "unpack":
		Unpack(os.Args[:2])
	case "-h", "--help":
		cli.PrintUsage()
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		cli.PrintUsage()
		os.Exit(1)
	}
}
