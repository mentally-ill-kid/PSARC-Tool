package psarc

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type TOCEntry struct {
	Digest [16]byte
	ZIndex uint32
	Length uint64
	Offset uint64
}

func readUint40BE(data []byte) uint64 {
	return uint64(data[0])<<32 |
		uint64(data[1])<<24 |
		uint64(data[2])<<16 |
		uint64(data[3])<<8 |
		uint64(data[4])
}

func ParseTOCEntries(hexTOC []byte, tocEntrySize int, tocEntries int) ([]TOCEntry, error) {
	if tocEntrySize <= 0 {
		return nil, errors.New("TOC entry size is invalid")
	}

	entries := tocEntries
	if entries <= 0 {
		entries = len(hexTOC) / tocEntrySize
	}
	if maxEntries := len(hexTOC) / tocEntrySize; entries > maxEntries {
		entries = maxEntries
	}

	parsed := make([]TOCEntry, 0, entries)
	for i := 0; i < entries; i++ {
		start := i * tocEntrySize
		entry := hexTOC[start : start+tocEntrySize]
		if len(entry) < 30 {
			continue
		}

		var digest [16]byte
		copy(digest[:], entry[:16])
		parsed = append(parsed, TOCEntry{
			Digest: digest,
			ZIndex: binary.BigEndian.Uint32(entry[16:20]),
			Length: readUint40BE(entry[20:25]),
			Offset: readUint40BE(entry[25:30]),
		})
	}

	return parsed, nil
}

func GetTOCInfo(hexTOC []byte, tocEntrySize int, tocEntries int) {
	entries, err := ParseTOCEntries(hexTOC, tocEntrySize, tocEntries)
	if err != nil {
		fmt.Println(err)
		return
	}

	fixedBytes := len(entries) * tocEntrySize
	if fixedBytes > len(hexTOC) {
		fixedBytes = len(hexTOC)
	}
	remainder := len(hexTOC) - fixedBytes

	fmt.Printf("TOC size: %d bytes\n", len(hexTOC))
	fmt.Printf("TOC entry size: %d bytes\n", tocEntrySize)
	fmt.Printf("Parsed entries: %d\n", len(entries))
	if remainder != 0 {
		fmt.Printf("Trailing block table bytes: %d\n", remainder)
	}

	if len(entries) == 0 {
		return
	}

	fmt.Println("Entries:")
	for i, entry := range entries {
		fmt.Printf("%03d: md5=%x zIndex=%d length=%d offset=%d\n", i, entry.Digest, entry.ZIndex, entry.Length, entry.Offset)
	}
}
