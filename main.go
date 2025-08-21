package main

import (
	"compression/compress"
	"compression/fileHandling"
	"compression/tree"
)

func main() {
	uncompressed := fileHandling.OpenFile("./test.txt")

	frequencies := fileHandling.RuneFreq(uncompressed)
	huff := tree.BuildTree(frequencies)

	codeMap := make(map[string]string)

	tree.BuildCode(huff, "", codeMap)

	compressed := compress.Compress(codeMap, uncompressed)
	compress.Decompress(compressed, huff)

	fileHandling.WriteFile(compressed)

	//test := byte("11111111")

	//fmt.Printf("%08b", test)

	//os.WriteFile("./compressed.txt", test, 0644)

	//testing, _ := os.ReadFile("compressed.bin")
	//println(string([]byte(testing)))
}
