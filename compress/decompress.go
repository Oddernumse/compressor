package compress

import (
	"compression/tree"
	"os"
	"strings"
)

func ByteToBinaryString(byteNumber byte) string {
	v := byte(byteNumber)
	binaryString := ""

	for i := 0; i < 8; i++ {
		if (v>>i)&1 == 1 {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}

	return binaryString
}

func Decompress(compressed []byte, root *tree.Leaf) string {
	node := root
	var decompressed strings.Builder

	for _, byte := range compressed {

		binString := ByteToBinaryString(byte)

		for _, num := range binString {
			if string(num) == "0" {
				node = node.Left
			} else {
				node = node.Right
			}

			if node.Char != 0 {
				decompressed.WriteRune(node.Char)
				node = root
			}
		}
	}

	os.WriteFile("./uncompressed.txt", []byte(decompressed.String()), 0644)

	return decompressed.String()
}
