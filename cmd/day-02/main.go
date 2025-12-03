package main

import (
	"fmt"
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

func Day02(idGroups string) (total int) {
	for rangeStr := range strings.SplitSeq(idGroups, ",") {
		bounds := strings.Split(rangeStr, "-")
		start := utils.Atoi(bounds[0])
		end := utils.Atoi(bounds[1])

		for id := start; id <= end; id++ {
			length := utils.LenI(id)

			for i := 1; i <= length/2; i++ {
				compId := utils.Atoi(strings.Repeat(strconv.Itoa(id)[:i], length/i))

				if compId == id {
					total += id
					break
				}
			}
		}
	}

	return total
}
