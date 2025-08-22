package fileHandling

import (
	"log"
	"os"
	"slices"
	"strings"
)

// Ugliest implementation of the cleverest solution
func BinStringToNumber(binaryString []string) uint8 {
	var num uint8

	for index, value := range binaryString {
		if value == "1" {
			switch index {
			case 0:
				num += 1
			case 1:
				num += 2
			case 2:
				num += 4
			case 3:
				num += 8
			case 4:
				num += 16
			case 5:
				num += 32
			case 6:
				num += 64
			case 7:
				num += 128
			}
		}
	}

	return num
}

func WriteFile(compressed string) {
	file, err := os.Create("./compressed.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitString := strings.Split(compressed, "")
	chunks := slices.Chunk(splitString, 8)

	for stringByte := range chunks {
		file.Write([]byte{BinStringToNumber(stringByte)})
	}
}
