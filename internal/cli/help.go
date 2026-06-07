package cli

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Fprintf(os.Stderr, `usage: psarc-tool <command> [path]
	
	Commands:
		pack      Create an archive from files/directories
		unpack    Extract files from an archive`)
}
