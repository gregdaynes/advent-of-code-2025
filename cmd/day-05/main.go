package main

import (
	"fmt"
	"slices"
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

func Day05(input string) (count int) {
	acc := make([]Range, 0, 200)

CURRENT:
	for v := range strings.SplitSeq(input, "\n") {
		if v == "" {
			break CURRENT
		}

		parts := strings.Split(v, "-")
		slices.SortFunc(parts, func(a, b string) int {
			low, _ := strconv.Atoi(b)
			high, _ := strconv.Atoi(a)
			return high - low
		})

		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])

		// EXISTING:
		for i, e := range acc {
			// covered already
			if low >= e.low && high <= e.high {
				continue CURRENT
			}

			// covers existing
			if low <= e.high && high >= e.low {
				acc[i] = Range{}
			}

			// lower than low
			if low <= e.low && high >= e.low && high <= e.high {
				high = e.high
				acc[i] = Range{}
			}

			// higher than high
			if high >= e.high && low <= e.high && low >= e.low {
				low = e.low
				acc[i] = Range{}
			}
		}

		acc = append(acc, Range{
			low:  low,
			high: high,
		})
	}

	for _, v := range acc {
		if v.high != 0 {
			count += v.high - v.low + 1
		}
	}

	return count
}
