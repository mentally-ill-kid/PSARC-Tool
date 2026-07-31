package psarc

import (
	"errors"
	"io"
	"os"
)

func ReadHeader(pathToFile string) ([]byte, error) {
	headerSize := 32

	bytes := make([]byte, headerSize)
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	header, err := file.ReadAt(bytes, 0)
	if err != nil {
		return nil, err
	}
	if header != headerSize {
		return nil, io.ErrUnexpectedEOF
	}
	return bytes, nil
}

func ReadTOC(pathToFile string, TOCsize int) ([]byte, error) {
	if TOCsize < 32 {
		return nil, errors.New("invalid TOC size")
	}

	bytes := make([]byte, TOCsize-32)
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	TOC, err := file.ReadAt(bytes, 32)
	if err != nil {
		return nil, err
	}
	if TOC != len(bytes) {
		return nil, io.ErrUnexpectedEOF
	}
	return bytes, nil
}

func ReadBytesAt(pathToFile string, offset int64, length int64) ([]byte, error) {
	if length < 0 {
		return nil, errors.New("invalid length")
	}

	bytes := make([]byte, int(length))
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	read, err := file.ReadAt(bytes, offset)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	if read != int(length) {
		return nil, io.ErrUnexpectedEOF
	}
	return bytes, nil
}
