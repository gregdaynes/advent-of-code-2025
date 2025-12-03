package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
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
		id := strconv.Itoa(i)
		l := len(id)
		if l%2 > 0 {
			continue
		}

		a := id[:l/2]
		b := id[l/2:]
		if a == b {
			invalid = append(invalid, i)
		}
	}

	return invalid
}
