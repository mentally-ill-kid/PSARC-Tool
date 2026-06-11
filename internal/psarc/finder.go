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

	var pattern string
	if pathToFile == "" {
		pattern = "*.psarc"
	} else {
		pattern = filepath.Join(pathToFile, "*.psarc")
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	if len(matches) == 0 {
		if pathToFile == "" {
			fmt.Println("Could not find psarc file in current directory")
		} else {
			fmt.Printf("Could not find psarc file in %s\n", pathToFile)
		}
		return "", err
	}

	for _, match := range matches {
		absPath, err := filepath.Abs(match)
		if err != nil {
			fmt.Printf("Error resolving path %s: %v\n", match, err)
			continue
		}
		fmt.Printf("File has been located: %s\n", absPath)
		return absPath, nil
	}
	return "", err
}
