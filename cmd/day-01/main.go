package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	fmt.Println("Running Day 01")

	input := utils.ReadInput("./cmd/day-01/input.txt")

	results := Day01(50, input)
	fmt.Printf("results: %+v\n", results)
}

// a sequence end is marked by rotating to 0
func Day01(state int, input string) (seqCount int) {
	for v := range strings.SplitSeq(input, "\n") {
		dir := 1
		if d := v[:1] == "L"; d == true {
			dir = -1
		}

		count, err := strconv.Atoi(v[1:])
		if err != nil {
			fmt.Printf("count was not an int, %v received\n", v)
		}

		for range count {
			state = state + 1*dir

			if state < 0 {
				state = 99
			} else if state > 99 {
				state = 0
			}

			if state == 0 {
				seqCount++
			}
		}
	}

	return seqCount
}
