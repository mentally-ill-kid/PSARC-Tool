# PSARC Tool

A simple pet project for learning file handling in Go.
Command-line unpacker for `.psarc` (PlayStation Archive) files, built with [Cobra](https://github.com/spf13/cobra).

## Current features

- Find `.psarc` files in a given directory
- Display detailed archive info (magic, version, compression, file count, etc.)
- Extract files from a `.psarc` archive (supports zlib-compressed archives)

## Build and run

`go build -o psarc-tool .`

## Future plans

- Add ability to create `.psarc` archives
- Dump ToC entries to a .txt file 

## Contributing

Contributions are **very** welcome! Feel free to fork and submit your PRs

## Tested on

PSARC-Tool was tested **ONLY** on [InFamous](https://en.wikipedia.org/wiki/Infamous_(video_game)) dump.
COMPATIBILITY ON OTHER TITLES NOT GUARANTEED