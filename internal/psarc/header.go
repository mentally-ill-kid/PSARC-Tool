package psarc

import (
	"encoding/binary"
	"fmt"
)

func GetHeaderInfo(hexHeader []byte) {
	magic := hexHeader[0:4]
	versionMajor := hexHeader[4:6]
	versionMinor := hexHeader[6:8]
	compression_type := hexHeader[8:12]
	toc_length := binary.BigEndian.Uint32(hexHeader[12:16])
	toc_entry_size := binary.BigEndian.Uint32(hexHeader[16:20])
	toc_entries := binary.BigEndian.Uint32(hexHeader[20:24])
	block_size := binary.BigEndian.Uint32(hexHeader[24:28])
	archive_flags := binary.BigEndian.Uint32(hexHeader[28:32])

	fmt.Printf(
		`Magic: %s
Version: v%d.%d
Compression type: %s
TOC length: %d
TOC entry size: %d bytes
TOC entries: 1 manifest + %d files
Block size: %d
archive flags: %d`,
		string(magic),
		versionMajor[1],
		versionMinor[1],
		string(compression_type),
		toc_length,
		toc_entry_size,
		toc_entries-1,
		block_size,
		archive_flags)
}

func GetTOCSize(hexHeader []byte) (int, int, int) {
	toc_length := binary.BigEndian.Uint32(hexHeader[12:16])
	toc_entry_size := binary.BigEndian.Uint32(hexHeader[16:20])
	toc_entries := binary.BigEndian.Uint32(hexHeader[20:24])

	return int(toc_length), int(toc_entry_size), int(toc_entries)
}
