package compress

import "strings"

// This could also just use the binary tree, but using a map makes this function simpler ¯\_(ツ)_/¯
func Compress(codeMap map[string]string, inputText string) string {
	var encodedText strings.Builder

	for _, char := range inputText {
		encodedText.WriteString(codeMap[string(char)])
	}

	return encodedText.String()
}
