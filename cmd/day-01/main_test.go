package main

import "testing"

var input = "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
var start = 50

func TestDay01(t *testing.T) {
	t.Run("count number of sequences", func(t *testing.T) {
		want := 6
		got := Day01(start, input)

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
