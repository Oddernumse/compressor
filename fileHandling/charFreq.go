package filehandling

func RuneFreq(fileContent string) map[string]int {
	frequency := make(map[string]int)

	for _, rune := range fileContent {
		frequency[string(rune)] = frequency[string(rune)] + 1
	}

	return frequency
}
