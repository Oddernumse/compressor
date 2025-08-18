package filehandling

import (
	"log"
	"os"
)

func OpenFile(fileName string) string {
	fh, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(fh)
}
