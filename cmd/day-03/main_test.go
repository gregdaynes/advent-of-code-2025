package main

import (
	"fmt"
	"testing"
)

func TestDay03(t *testing.T) {
	t.Run("find high set of 12 in bank", func(t *testing.T) {
		cases := []struct {
			Bank string
			High string
		}{
			{"987654321111111", "987654321111"},
			{"811111111111119", "811111111119"},
			{"234234234234278", "434234234278"},
			{"818181911112111", "888911112111"},
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
		want := 3121910778619

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
