package compress

import "strings"

func Compress(codeMap map[string]string, inputText string) string {
	var encodedText strings.Builder

	for _, char := range inputText {
		encodedText.WriteString(codeMap[string(char)])
	}

	return encodedText.String()
}
