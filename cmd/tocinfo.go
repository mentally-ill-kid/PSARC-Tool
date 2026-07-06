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
	tocCmd.Flags().BoolP("write", "w", false, "write ToC into file")
}

func tocInfo(cmd *cobra.Command, args []string) {
	pathToFile, _ := cmd.Flags().GetString("path")
	isWritable, _ := cmd.Flags().GetBool("write")
	absPath, err := psarc.FindArchive(pathToFile)
	if err != nil {
		fmt.Printf("Could not find file: %s", err)
		return
	}
	headerBytes, err := psarc.ReadHeader(absPath)
	if err != nil {
		fmt.Printf("Could not read file: %s", err)
		return
	}
	tocSize, tocEntrySize, tocEntries := psarc.GetTOCSize(headerBytes)
	TOCbytes, err := psarc.ReadTOC(absPath, tocSize)
	if err != nil {
		fmt.Printf("Could not read TOC: %s", err)
		return
	}
	psarc.GetTOCInfo(TOCbytes, tocEntrySize, tocEntries, isWritable)
}
