package main

import (
	filehandling "compression/fileHandling"
	"compression/tree"
)

func main() {
	frequencies := filehandling.RuneFreq(filehandling.OpenFile("./test.txt"))
	huff := tree.BuildTree(frequencies)

	tree.BuildCode(huff, "")

	tree.PrintTree(huff)
}
