package psarc

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindArchive(pathToFile string) (string, error) {
	if pathToFile != "" {
		if info, err := os.Stat(pathToFile); err == nil && !info.IsDir() {
			absPath, absErr := filepath.Abs(pathToFile)
			if absErr != nil {
				return "", absErr
			}
			return absPath, nil
		}
	}

	var patterns []string
	if pathToFile == "" {
		patterns = []string{"*.psarc", "*.psarc_s"}
	} else {
		patterns = []string{
			filepath.Join(pathToFile, "*.psarc"),
			filepath.Join(pathToFile, "*.psarc_s"),
		}
	}

	var allMatches []string
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return "", fmt.Errorf("error searching for pattern %s: %w", pattern, err)
		}
		allMatches = append(allMatches, matches...)
	}

	if len(allMatches) == 0 {
		if pathToFile == "" {
			return "", fmt.Errorf("could not find psarc file in current directory")
		}
		return "", fmt.Errorf("could not find psarc file in %s", pathToFile)
	}

	for _, match := range allMatches {
		absPath, err := filepath.Abs(match)
		if err != nil {
			fmt.Printf("Error resolving path %s: %v\n", match, err)
			continue
		}
		fmt.Printf("File has been located: %s\n", absPath)
		return absPath, nil
	}

	return "", fmt.Errorf("no valid files found")
}
