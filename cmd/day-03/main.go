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
		total += findHighPair(bank)
	}

	return total
}

func findHighPair(bank string) (high int) {
	for i := range len(bank) {
		for j := i + 1; j < len(bank); j++ {
			num, _ := strconv.Atoi(bank[i:i+1] + bank[j:j+1])
			if num > high {
				high = num
			}
		}
	}

	return high
}
