package cmd

import (
	"fmt"

	"github.com/mentally-ill-kid/PSARC-Tool/internal/psarc"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract contents into folder",
	Run:   extract,
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().StringP("path", "p", "", "path to the file or directory")
	extractCmd.Flags().StringP("output", "o", "", "directory to extract into")
}

func extract(cmd *cobra.Command, args []string) {
	pathToFile, _ := cmd.Flags().GetString("path")
	outputDir, _ := cmd.Flags().GetString("output")
	absPath, err := psarc.FindArchive(pathToFile)
	if err != nil {
		fmt.Printf("Could not find file: %s", err)
		return
	}
	if err := psarc.ExtractArchive(absPath, outputDir); err != nil {
		fmt.Printf("Could not extract archive: %s\n", err)
	}

}
