package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error reading input", err)
	}

	return strings.Trim(string(bytes), "\n")
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not convert %v to int", s)
	}
	return i
}

func LenI(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
