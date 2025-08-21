package compress

import (
	"compression/tree"
	"strings"
)

func Decompress(compressed string, root *tree.Leaf) string {
	node := root
	var decompressed strings.Builder

	for _, bit := range compressed {
		if string(bit) == "0" {
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
