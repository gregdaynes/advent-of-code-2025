package utils

import (
	"log"
	"os"
	"strings"
)

func ReadInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error reading input", err)
	}

	return strings.Trim(string(bytes), "\n")
}
