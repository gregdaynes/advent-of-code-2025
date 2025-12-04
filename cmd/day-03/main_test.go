package main

import (
	"fmt"
	"testing"
)

func TestDay03(t *testing.T) {
	t.Run("find high pair in bank", func(t *testing.T) {
		cases := []struct {
			Bank string
			High int
		}{
			{"987654321111111", 98},
			{"811111111111119", 89},
			{"234234234234278", 78},
			{"818181911112111", 92},
		}

		for n, test := range cases {
			t.Run(fmt.Sprintf("case %d", n), func(t *testing.T) {
				got := findHighPair(test.Bank)
				want := test.High

				if got != want {
					t.Errorf("wanted %v, got %v", want, got)
				}
			})
		}
	})

	t.Run("day03", func(t *testing.T) {
		input := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
		got := Day03(input)
		want := 357

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
