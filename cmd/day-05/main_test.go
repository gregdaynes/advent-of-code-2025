package main

import "testing"

var input = "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"
var answer = 14

// var input = "200-300\n100-101\n1-1\n2-2\n3-3\n1-3\n1-3\n2-2\n50-70\n10-10\n98-99\n99-99\n99-99\n99-100\n1-1\n2-1\n100-100\n100-100\n100-101\n200-300\n201-300\n202-300\n250-251\n98-99\n100-100\n100-101\n1-101"
// var answer = 202

func TestDay05(t *testing.T) {
	t.Run("Day05", func(t *testing.T) {
		got := Day05(input)
		want := answer
		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
