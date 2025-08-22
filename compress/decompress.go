package compress

import (
	"compression/tree"
	"strings"
)

func ByteToBinaryString(byteNumber byte) {

}

func Decompress(compressed []byte, root *tree.Leaf) string {
	node := root
	var decompressed strings.Builder

	for _, bit := range compressed {

		if bit == 0 {
			node = node.Left
		} else {
			node = node.Right
		}

		if node.Char != 0 {
			decompressed.WriteRune(node.Char)
			node = root
		}
	}

	return decompressed.String()
}
