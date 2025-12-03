package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	fmt.Println("Running Day 02")

	input := utils.ReadInput("./cmd/day-02/input.txt")
	results := Day02(input)

	fmt.Printf("results: %+v\n", results)
}

func Day02(idGroups string) (sum int) {
	for groups := range strings.SplitSeq(idGroups, ",") {
		parts := strings.Split(groups, "-")
		start := utils.Atoi(parts[0])
		end := utils.Atoi(parts[1])

		for n := start; n <= end; n++ {
			id := strconv.Itoa(n)
			length := len(id)
			maxIters := int(math.Floor(float64(length) / 2))

			for i := 1; i <= maxIters; i++ {
				if strings.Repeat(id[:i], length/i) == id {
					sum += n
					break
				}
			}
		}
	}

	return sum
}
