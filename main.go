package main

import (
	"compression/compress"
	"compression/fileHandling"
	"compression/tree"
	"os"
)

func main() {
	uncompressed := fileHandling.OpenFile("./test.txt")

	frequencies := fileHandling.RuneFreq(uncompressed)
	huff := tree.BuildTree(frequencies)

	codeMap := make(map[string]string)

	tree.BuildCode(huff, "", codeMap)

	compressed := compress.Compress(codeMap, uncompressed)

	fileHandling.WriteBin(compressed)

	fh, _ := os.ReadFile("./compressed.bin")
	println(compress.Decompress(fh, huff))
}
