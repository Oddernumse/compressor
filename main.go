package main

import (
	filehandling "compression/fileHandling"
	"compression/tree"
)

func main() {
	bst := tree.BST{}

	frequencies := filehandling.RuneFreq(filehandling.OpenFile("./test.txt"))
	huff := tree.BuildTree(frequencies)

	bst.BuildCode(huff)
}
