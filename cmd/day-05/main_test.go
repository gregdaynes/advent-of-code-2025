package main

import "testing"

var input = "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"

func TestDay05(t *testing.T) {
	t.Run("Day05", func(t *testing.T) {
		got := Day05(input)
		want := 3
		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
