package main

import (
	"fmt"
	"testing"
)

func TestDay02(t *testing.T) {
	cases := []struct {
		IDRange string
		Want    []int
	}{
		{"11-22", []int{11, 22}},
		{"95-115", []int{99, 111}},
		{"998-1012", []int{999, 1010}},
		{"1188511880-1188511890", []int{1188511885}},
		{"222220-222224", []int{222222}},
		{"1698522-1698528", []int{}},
		{"446443-446449", []int{446446}},
		{"38593856-38593862", []int{38593859}},
		{"565653-565659", []int{565656}},
		{"824824821-824824827", []int{824824824}},
		{"2121212118-2121212124", []int{2121212121}},
	}

	for n, test := range cases {
		t.Run(fmt.Sprintf("test %d", n), func(t *testing.T) {
			invalidIDs := findInvalidID(test.IDRange)

			if len(invalidIDs) != len(test.Want) {
				t.Errorf("wanted %v ids, got %v", len(test.Want), len(invalidIDs))
			}

			lenWant := len(test.Want)
			for i, got := range invalidIDs {
				if lenWant >= i {
					continue
				}

				want := test.Want[i]
				if got != want {
					t.Errorf("wanted %v, got %v", want, got)
				}

			}

		})
	}

	t.Run("testing sum over all ids", func(t *testing.T) {
		idGroups := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
		got := Day02(idGroups)
		want := 4174379265

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
