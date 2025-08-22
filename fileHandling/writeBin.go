package fileHandling

import (
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

// Slightly less ugly implementation of the cleverest solution
func BinStringToNumber(binaryString []string) uint8 {
	var num uint8

	for index, value := range binaryString {
		if value == "1" {
			num += uint8(math.Pow(2, float64(index)))
		}
	}

	return num
}

func WriteBin(compressed string) {
	file, err := os.Create("./compressed.bin")
	if err != nil {
		log.Fatal(err)
	}

	splitString := strings.Split(compressed, "")
	chunks := slices.Chunk(splitString, 8)

	for stringByte := range chunks {
		file.Write([]byte{BinStringToNumber(stringByte)})
	}
}
