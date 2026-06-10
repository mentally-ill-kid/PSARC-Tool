package cmd

import (
	"fmt"

	"github.com/mentally-ill-kid/PSARC-Tool/internal/psarc"
	"github.com/spf13/cobra"
)

var tocCmd = &cobra.Command{
	Use:   "toc",
	Short: "Prints information about Table of Content into terminal",
	Run:   tocInfo,
}

func init() {
	rootCmd.AddCommand(tocCmd)
	tocCmd.Flags().StringP("path", "p", "", "path to directory containing PSARC archive")
}

func tocInfo(cmd *cobra.Command, args []string) {
	pathToFile, _ := cmd.Flags().GetString("path")
	absPath, err := psarc.FindArchive(pathToFile)
	if err != nil {
		fmt.Printf("Could not find file: %s", err)
	}
	headerBytes, err := psarc.ReadHeader(absPath)
	if err != nil {
		fmt.Printf("Could not read file")
	}
	toc_size, toc_entry_size := psarc.GetTOCSize(headerBytes)
	TOCbytes, err := psarc.ReadTOC(absPath, toc_size)
	psarc.GetTOCInfo(TOCbytes, toc_entry_size)
}
