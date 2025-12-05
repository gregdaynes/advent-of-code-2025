package main

import (
	"fmt"
	"testing"
)

var input = "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."

func TestDay04(t *testing.T) {
	t.Run("Day", func(t *testing.T) {
		got := Day04(input)
		want := 13

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("Roll Access", func(t *testing.T) {
		cases := []struct {
			x    int
			y    int
			want bool
		}{
			{
				x:    2,
				y:    0,
				want: true,
			},
			{
				x:    3,
				y:    0,
				want: true,
			},
			{
				x:    5,
				y:    0,
				want: true,
			},
			{
				x:    6,
				y:    0,
				want: true,
			},
			{
				x:    8,
				y:    0,
				want: true,
			},
			{
				x:    0,
				y:    1,
				want: true,
			},
			{
				x:    6,
				y:    2,
				want: true,
			},
			{
				x:    0,
				y:    4,
				want: true,
			},
			{
				x:    9,
				y:    4,
				want: true,
			},
			{
				x:    0,
				y:    7,
				want: true,
			},
			{
				x:    0,
				y:    9,
				want: true,
			},
			{
				x:    2,
				y:    9,
				want: true,
			},
			{
				x:    8,
				y:    9,
				want: true,
			},
		}

		for n, test := range cases {
			table, rows, cols := parseInput(input)

			t.Run(fmt.Sprintf("case %d", n), func(t *testing.T) {
				got := checkRollAccess(table, rows, cols, test.x, test.y)
				want := true

				if got != want {
					t.Errorf("wanted %v, got %v", want, got)
				}
			})
		}
	})
}
