package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running Day 02")

	bytes, err := os.ReadFile("./cmd/day-02/input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return
	}
	input := strings.Trim(string(bytes), "\n")

	sum := Day02(input)
	fmt.Printf("results: %+v\n", sum)
}

func Day02(idGroups string) (sum int) {
	for group := range strings.SplitSeq(idGroups, ",") {
		ids := findInvalidID(group)

		for _, id := range ids {
			sum += id
		}
	}

	return sum
}

func findInvalidID(rangeID string) (invalid []int) {
	rangeIDs := strings.Split(rangeID, "-")
	start, err := strconv.Atoi(rangeIDs[0])
	end, err := strconv.Atoi(rangeIDs[1])
	if err != nil {
		fmt.Println("error converting index to int", start, end, err)
	}

	for i := start; i <= end; i++ {
		if !validID(strconv.Itoa(i)) {
			invalid = append(invalid, i)
		}
	}

	return invalid
}

func validID(id string) (ok bool) {
	length := len(id)
	maxIters := int(math.Floor(float64(length) / 2))

	for i := 1; i <= maxIters; i++ {
		if strings.Repeat(id[:i], length/i) == id {
			return false
		}
	}

	return true
}
