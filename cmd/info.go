package cmd

import (
	"fmt"

	"github.com/mentally-ill-kid/PSARC-Tool/internal/psarc"
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
	absPath, err := psarc.FindArchive(pathToFile)
	if err != nil {
		fmt.Printf("Encountered error: %s", err)
	}
	bytes, err := psarc.ReadHeader(absPath)
	psarc.GetHeaderInfo(bytes)
}
