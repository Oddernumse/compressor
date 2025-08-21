package fileHandling

import (
	"log"
	"os"
	"slices"
	"strings"
)

func binaryStringToByte(s string) byte {
	var result byte
	for i := 0; i < len(s); i++ {
		result <<= 1
		if s[i] == '1' {
			result |= 1
		}
	}
	return result
}

func binaryStringsToBytes(strings []string) []byte {
	bytes := make([]byte, len(strings))
	for i, s := range strings {
		bytes[i] = binaryStringToByte(s)
	}

	println(len(bytes))
	return bytes
}

func WriteFile(compressed string) {
	file, err := os.Create("./compressed.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitString := strings.Split(compressed, "")
	chunks := slices.Chunk(splitString, 8)

	for byte := range chunks {
		file.Write(binaryStringsToBytes(byte))
	}
}
