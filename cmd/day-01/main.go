package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running Day 01")

	start := 50
	input := ""

	bytes, err := os.ReadFile("./cmd/day-01/input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return
	}
	input = strings.Trim(string(bytes), "\n")

	results := Day01(start, input)
	fmt.Printf("results: %+v", results)
}

// a sequence end is marked by rotating to 0
func Day01(start int, input string) int {
	seqCount := 0
	state := start

	for v := range strings.SplitSeq(input, "\n") {
		mult := 1
		if dir := v[:1] == "L"; dir == true {
			mult = -1
		}

		count, err := strconv.Atoi(v[1:])
		if err != nil {
			fmt.Printf("count was not an int, %v received", v)
		}

		for range count {
			state = state + 1*mult

			if state < 0 {
				state = 99
			} else if state > 99 {
				state = 0
			}
		}

		if state == 0 {
			seqCount = seqCount + 1
		}
	}

	return seqCount
}
