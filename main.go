package main

import (
	filehandling "compression/fileHandling"
)

func main() {
	frequencies := filehandling.RuneFreq(filehandling.OpenFile("./test.txt"))

	for rune, value := range frequencies {
		println(rune, " ", value)
	}
}
