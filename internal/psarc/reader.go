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
		fmt.Printf("cound not open the file: %s\n", err)
	} else {
		// fmt.Printf("opened the file\n")
	}
	header, err := file.ReadAt(bytes, 0)
	if err != nil {
		fmt.Printf("could not read the file: %s\n", err)
	} else {
		fmt.Printf("Header: %d\n", header)
		// fmt.Printf("Header: %x\n", bytes)
		return bytes, nil
	}
	return nil, err
}
