package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	input := utils.ReadInput("./cmd/day-06/input.txt")
	results := Day06(input)

	fmt.Printf("results: %+v\n", results)
}

type question struct {
	operator string
	parts    []string
}

func Day06(input string) (total int) {
	lines := strings.Split(input, "\n")
	maxI := len(lines[0]) - 1
	maxJ := len(lines)

	qs := make([]question, maxI)
	n := 0

QUESTIONS:
	for i := maxI; i >= 0; i-- {
		q := qs[n]
		part := ""

		for j := range maxJ {
			char := string(lines[j][i])

			// flush the parts, then set the operator
			if char == "+" || char == "*" {
				parts := append(q.parts, part)
				q.parts = parts

				q.operator = char
				qs[n] = q
				i -= 1
				n += 1
				continue QUESTIONS
			}

			part += char
		}

		parts := append(q.parts, part)
		qs[n].parts = parts
	}

	//
	for _, q := range qs {
		fmt.Println(q)
		switch q.operator {
		case "+":
			v := 0
			for _, pt := range q.parts {
				pt, _ := strconv.Atoi(strings.Trim(pt, " "))
				v = v + pt
			}
			fmt.Println(v)
			total += v
		case "*":
			v := 1
			for _, pt := range q.parts {
				pt, _ := strconv.Atoi(strings.Trim(pt, " "))
				v = v * pt
			}
			fmt.Println(v)
			total += v
		}
	}

	return total
}
