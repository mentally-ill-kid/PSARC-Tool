package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints information about archive into terminal",
	Run:   runInfo,
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringP("path", "p", "", "path to the file or directory")
}

func runInfo(cmd *cobra.Command, args []string) {
	pathToFile, _ := cmd.Flags().GetString("path")

	var pattern string
	if pathToFile == "" {
		pattern = "*.psarc"
	} else {
		pattern = filepath.Join(pathToFile, "*.psarc")
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("Error searching for files: %v\n", err)
		return
	}

	if len(matches) == 0 {
		if pathToFile == "" {
			fmt.Println("Could not find psarc file in current directory")
		} else {
			fmt.Printf("Could not find psarc file in %s\n", pathToFile)
		}
		return
	}

	for _, match := range matches {
		absPath, err := filepath.Abs(match)
		if err != nil {
			fmt.Printf("Error resolving path %s: %v\n", match, err)
			continue
		}
		fmt.Printf("File has been located: %s\n", absPath)
		readHeader(absPath)
	}
}

func readHeader(absolutePath string) {
	fmt.Printf("Filepath is: %s\n", absolutePath)
}
