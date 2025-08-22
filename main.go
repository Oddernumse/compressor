package main

import (
	"compression/compress"
	"compression/fileHandling"
	"compression/tree"
	"os"
)

func main() {
	uncompressed := fileHandling.OpenFile(os.Args[1])
	frequencies := fileHandling.RuneFreq(uncompressed)
	huff := tree.BuildTree(frequencies)

	codeMap := make(map[string]string)
	tree.BuildCode(huff, "", codeMap)
	compressed := compress.Compress(codeMap, uncompressed)

	fileHandling.WriteBin(compressed)
	fh, _ := os.ReadFile("./compressed.bin")
	// The compressed file technically IS missing a header section
	// for someone else to be able to decompress this
	compress.Decompress(fh, huff)
}
