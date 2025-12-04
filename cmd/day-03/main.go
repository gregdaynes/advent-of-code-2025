package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	fmt.Println("Running Day 03")

	input := utils.ReadInput("./cmd/day-03/input.txt")
	results := Day03(input)

	fmt.Printf("results: %+v\n", results)
}

func Day03(input string) (total int) {
	banks := strings.SplitSeq(input, "\n")
	for bank := range banks {
		result, _ := strconv.Atoi(findHighPair(bank))
		total += result
	}

	return total
}

type High struct {
	Number   int
	Index    int
	Reserved int
}

func findHighPair(bank string) (found string) {
	// 1. find largest digit with 11 followers
	// which means we the distance reserved from end must be 11

	// 2. find the largest digit with 10 followers
	// the reserved contracts to 10

	// ...

	// repeat until 12 found

	c := High{
		Number:   0,
		Index:    -1,
		Reserved: 12,
	}

	for len(found) <= 11 {
		end := len(bank) - c.Reserved
		// fmt.Println("r", c.Reserved, "end", end)
		// find next highest number with reserved amount trailing
		for i := c.Index + 1; i <= end; i++ {
			num, _ := strconv.Atoi(string(bank[i]))
			if num > c.Number {
				// fmt.Printf("num %v is higher than %v\n", num, c.Number)

				c.Number = num
				c.Index = i
			}
		}

		found = found + string(bank[c.Index])
		// keep index, but reset high number
		c.Number = 0
		c.Reserved -= 1
	}

	return found
}
