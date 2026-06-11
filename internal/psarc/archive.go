package psarc

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func isZlibCompressed(hexHeader []byte) bool {
	return len(hexHeader) >= 12 && string(hexHeader[8:12]) == "zlib"
}

func decompressIfNeeded(data []byte, compressed bool) ([]byte, error) {
	if !compressed {
		return data, nil
	}

	var output bytes.Buffer
	remaining := data
	for len(remaining) > 0 {
		streamReader := bytes.NewReader(remaining)
		reader, err := zlib.NewReader(streamReader)
		if err != nil {
			return nil, err
		}

		chunk, readErr := io.ReadAll(reader)
		closeErr := reader.Close()
		if readErr != nil {
			return nil, readErr
		}
		if closeErr != nil {
			return nil, closeErr
		}

		output.Write(chunk)
		consumed := len(remaining) - streamReader.Len()
		if consumed <= 0 {
			break
		}
		remaining = remaining[consumed:]
	}

	return output.Bytes(), nil
}

func readArchiveEntry(pathToFile string, entry TOCEntry, nextOffset uint64, compressed bool) ([]byte, error) {
	if nextOffset < entry.Offset {
		return nil, fmt.Errorf("invalid entry offsets for archive data")
	}

	raw, err := ReadBytesAt(pathToFile, int64(entry.Offset), int64(nextOffset-entry.Offset))
	if err != nil {
		return nil, err
	}

	return decompressIfNeeded(raw, compressed)
}

func sanitizeArchivePath(name string) (string, bool) {
	cleaned := filepath.Clean(strings.TrimPrefix(name, "/"))
	if cleaned == "." || cleaned == "" || strings.HasPrefix(cleaned, "..") {
		return "", false
	}
	return cleaned, true
}

func ExtractArchive(pathToFile string, outputDir string) error {
	headerBytes, err := ReadHeader(pathToFile)
	if err != nil {
		return err
	}

	tocSize, tocEntrySize, tocEntries := GetTOCSize(headerBytes)
	tocBytes, err := ReadTOC(pathToFile, tocSize)
	if err != nil {
		return err
	}

	entries, err := ParseTOCEntries(tocBytes, tocEntrySize, tocEntries)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return fmt.Errorf("archive does not contain any TOC entries")
	}
	if len(entries) < 2 {
		return fmt.Errorf("archive does not contain any extractable files")
	}

	compressed := isZlibCompressed(headerBytes)
	manifestData, err := readArchiveEntry(pathToFile, entries[0], entries[1].Offset, compressed)
	if err != nil {
		return err
	}

	manifestLines := strings.Split(strings.ReplaceAll(string(manifestData), "\r\n", "\n"), "\n")
	if len(manifestLines) > 0 && manifestLines[len(manifestLines)-1] == "" {
		manifestLines = manifestLines[:len(manifestLines)-1]
	}

	if outputDir == "" {
		outputDir = "extracted"
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}

	manifestPath := filepath.Join(outputDir, "manifest.txt")
	if err := os.WriteFile(manifestPath, manifestData, 0o644); err != nil {
		return err
	}

	filesExtracted := 0
	for index := 1; index < len(entries); index++ {
		entry := entries[index]
		var nextOffset uint64
		if index+1 < len(entries) {
			nextOffset = entries[index+1].Offset
		} else {
			info, statErr := os.Stat(pathToFile)
			if statErr != nil {
				return statErr
			}
			nextOffset = uint64(info.Size())
		}

		data, err := readArchiveEntry(pathToFile, entry, nextOffset, compressed)
		if err != nil {
			return fmt.Errorf("failed to read entry %d: %w", index, err)
		}

		fileName := fmt.Sprintf("entry-%03d.bin", index)
		if index-1 < len(manifestLines) {
			if sanitized, ok := sanitizeArchivePath(manifestLines[index-1]); ok {
				fileName = sanitized
			}
		}

		outputPath := filepath.Join(outputDir, fileName)
		if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(outputPath, data, 0o644); err != nil {
			return err
		}
		filesExtracted++
	}

	fmt.Printf("Extracted %d files to %s\n", filesExtracted, outputDir)
	return nil
}
