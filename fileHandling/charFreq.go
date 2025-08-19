package filehandling

func RuneFreq(fileContent string) map[rune]int {
	frequency := make(map[rune]int)

	for _, rune := range fileContent {
		frequency[rune] = frequency[rune] + 1
	}

	return frequency
}
