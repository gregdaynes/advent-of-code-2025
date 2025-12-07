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
		fmt.Println("---------")
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
		fmt.Println(v, low, high)

		// EXISTING:
		for i, e := range acc {
			fmt.Println("existing check", e)
			// covered already
			if low >= e.low && high <= e.high {
				fmt.Println("skipping", e, low, high)
				continue CURRENT
			}

			// covers existin
			if low <= e.high && high >= e.low {
				acc[i] = Range{}
				fmt.Println("removing", e, low, high)
			}

			// lower than low
			if low <= e.low && high >= e.low && high <= e.high {
				high = e.high
				acc[i] = Range{}
				fmt.Println("extending low", e, low, high, acc[i])
			}

			// higher than high
			if high >= e.high && low <= e.high && low >= e.low {
				low = e.low
				acc[i] = Range{}
				fmt.Println("extending high", e, low, high, acc[i])
			}
		}

		fmt.Println("new entry", low, high)
		acc = append(acc, Range{
			low:  low,
			high: high,
		})
	}

	fmt.Println(acc)

	for _, v := range acc {
		if v.high != 0 {
			count += v.high - v.low + 1
		}
	}

	return count
}
