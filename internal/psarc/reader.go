package psarc

import (
	"fmt"
	"os"
)

func ReadHeader(pathToFile string) ([]byte, error) {
	headerSize := 32

	bytes := make([]byte, headerSize)
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Printf("could not open the file: %s\n", err)
	} else {
		// fmt.Printf("opened the file\n")
	}
	header, err := file.ReadAt(bytes, 0)
	if err != nil || header != headerSize {
		fmt.Printf("could not read the file: %s\n", err)
	} else {
		fmt.Printf("Header: %d\n", header)
		// fmt.Printf("Header: %x\n", bytes)
		return bytes, nil
	}
	return nil, err
}

func ReadTOC(pathToFile string, TOCsize int) ([]byte, error) {
	bytes := make([]byte, TOCsize)
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Printf("Could not open the file: %s", err)
	}
	TOC, err := file.ReadAt(bytes, 32) //32 -TOC beginning
	if err != nil || TOC != TOCsize {
		fmt.Printf("Could not read the file: %s", err)
	} else {
		return bytes, err
	}
	return nil, err
}
