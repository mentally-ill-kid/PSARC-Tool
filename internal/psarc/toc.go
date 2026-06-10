package psarc

import (
	"encoding/binary"
	"fmt"
)

func readUint40BE(data []byte) uint64 {
	return uint64(data[0])<<32 |
		uint64(data[1])<<24 |
		uint64(data[2])<<16 |
		uint64(data[3])<<8 |
		uint64(data[4])
}

func GetTOCInfo(hexTOC []byte, tocEntrySize int, tocEntries int) {
	if tocEntrySize <= 0 {
		fmt.Println("TOC entry size is invalid")
		return
	}

	entries := tocEntries
	if entries <= 0 {
		entries = len(hexTOC) / tocEntrySize
	}
	if maxEntries := len(hexTOC) / tocEntrySize; entries > maxEntries {
		entries = maxEntries
	}
	fixedBytes := entries * tocEntrySize
	if fixedBytes > len(hexTOC) {
		fixedBytes = len(hexTOC)
	}
	remainder := len(hexTOC) - fixedBytes

	fmt.Printf("TOC size: %d bytes\n", len(hexTOC))
	fmt.Printf("TOC entry size: %d bytes\n", tocEntrySize)
	fmt.Printf("Parsed entries: %d\n", entries)
	if remainder != 0 {
		fmt.Printf("Trailing block table bytes: %d\n", remainder)
	}

	if entries == 0 {
		return
	}

	fmt.Println("Entries:")
	for i := 0; i < entries; i++ {
		start := i * tocEntrySize
		entry := hexTOC[start : start+tocEntrySize]
		if len(entry) < 30 {
			fmt.Printf("%03d: %x\n", i, entry)
			continue
		}

		digest := entry[:16]
		zIndex := binary.BigEndian.Uint32(entry[16:20])
		length := readUint40BE(entry[20:25])
		offset := readUint40BE(entry[25:30])

		fmt.Printf("%03d: md5=%x zIndex=%d length=%d offset=%d\n", i, digest, zIndex, length, offset)
	}
}
