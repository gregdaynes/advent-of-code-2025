package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

type Range struct {
	low  int
	high int
}

func main() {
	fmt.Println("Running Day 05")

	input := utils.ReadInput("./cmd/day-05/input.txt")
	results := Day05(input)

	fmt.Printf("results: %+v\n", results)
}

func Day05(input string) int {
	ranges := make([]Range, 0, 100)
	freshIDs := map[int]bool{}

	// 1. Prepare the data by taking the first part (before \n\n)
	// 2. Turning it into a low/high range and storing it in a slice
	// 3. iterate over the available ids (after \n\n)
	// 4. check through each slice if its within or equal the range
	var check bool
	for v := range strings.SplitSeq(input, "\n") {
		if v == "" {
			check = true
			continue
		}

		if !check {
			parts := strings.Split(v, "-")
			low, _ := strconv.Atoi(parts[0])
			high, _ := strconv.Atoi(parts[1])

			ranges = append(ranges, Range{
				low:  low,
				high: high,
			})
		}

		if check {
			vint, _ := strconv.Atoi(v)

			for _, r := range ranges {
				if vint >= r.low && vint <= r.high {
					freshIDs[vint] = true
				}
			}
		}
	}

	return len(freshIDs)
}
