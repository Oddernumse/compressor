package main

import (
	"bufio"
	"compression/compress"
	"compression/fileHandling"
	"compression/tree"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	uncompressed := fileHandling.OpenFile("./test.txt")

	frequencies := fileHandling.RuneFreq(uncompressed)
	huff := tree.BuildTree(frequencies)

	codeMap := make(map[string]string)

	tree.BuildCode(huff, "", codeMap)

	compress.Compress(codeMap, uncompressed)

	//fileHandling.WriteFile(compressed)

	f, err := os.Open("./compressed.bin")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		println(binary.ReadVarint(buf))
	}
}
