package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	input := utils.ReadInput("./cmd/day-06/input.txt")
	results := Day06(input)

	fmt.Printf("results: %+v\n", results)
}

// consecutive spaces
var re = regexp.MustCompile("\\s+")

type question struct {
	operator string
	parts    []int
}

func Day06(input string) (total int) {
	lines := strings.Split(input, "\n")
	opLineIdx := len(lines) - 1

	qs := make([]question, questionCount(lines[0])+1)

	for i, row := range lines {
		val := re.Split(row, -1)

		for j, v := range val {
			q := qs[j]

			if i == opLineIdx {
				q.operator = v
			} else {
				parts := q.parts
				vv, _ := strconv.Atoi(v)
				q.parts = append(parts, vv)
			}

			qs[j] = q
		}
	}

	for _, q := range qs {
		switch q.operator {
		case "+":
			v := 0
			for _, pt := range q.parts {
				v = v + pt
			}
			total += v
		case "*":
			v := 1
			for _, pt := range q.parts {
				v = v * pt
			}
			total += v
		}
	}

	return total
}

func questionCount(input string) int {
	return len(re.Split(input, -1))
}
