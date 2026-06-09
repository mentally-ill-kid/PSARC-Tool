package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints information about archive into terminal",

	Run: runInfo,
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().StringP("path", "p", "", "path to the file")
}

func runInfo(cmd *cobra.Command, args []string) {
	pathToFile, _ := cmd.Flags().GetString("path")

	if pathToFile == "" {
		_, err := filepath.Glob("*.psarc??")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Could not file psarc file in current directory")
		} else {
			fmt.Printf("file has been located")
		}
	} else {
		_, err := filepath.Glob(filepath.Join(pathToFile, "*psarc??"))
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Count not find psarc file in %s", pathToFile)
		} else {
			fmt.Printf("file has been located in %s", pathToFile)
		}
	}
}
