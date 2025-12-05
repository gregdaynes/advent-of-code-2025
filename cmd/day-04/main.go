package main

import (
	"fmt"
	"strings"

	"github.com/gregdaynes/advent-of-code-2025/internal/utils"
)

func main() {
	fmt.Println("Running Day 04")

	input := utils.ReadInput("./cmd/day-04/input.txt")
	results := Day04(input)

	fmt.Printf("results: %+v\n", results)
}

func Day04(input string) (rolls int) {
	table, rows, cols := parseInput(input)

	for y := range rows {
		for x := range cols {
			if checkRollAccess(table, cols, rows, x, y) {
				rolls += 1
			}
		}
	}

	return rolls
}

func checkRollAccess(table [][]bool, cols, rows, x, y int) bool {
	cell := table[y][x]
	if cell != true {
		return false
	}

	// 3. find 8 adjacent cells
	//    doesn't have to remember location, just value
	//    we will need to check the cell exists before checking for value
	adj := countAdjacent(table, x, y, rows, cols)

	// 4. check if there are less than  4 trues around coords

	return adj < 4
}

func countAdjacent(table [][]bool, x, y, width, height int) (adj int) {
	if !validPos(x, y, width, height) {
		return 0
	}

	// TL
	if validPos(x-1, y-1, width, height) {
		if table[y-1][x-1] == true {
			adj += 1
		}
	}

	// TC
	if validPos(x, y-1, width, height) {
		if table[y-1][x] == true {
			adj += 1
		}
	}

	// TR
	if validPos(x+1, y-1, width, height) {
		if table[y-1][x+1] == true {
			adj += 1
		}
	}

	// ML
	if validPos(x-1, y, width, height) {
		if table[y][x-1] == true {
			adj += 1
		}
	}

	// MR
	if validPos(x+1, y, width, height) {
		if table[y][x+1] == true {
			adj += 1
		}
	}

	// BL
	if validPos(x-1, y+1, width, height) {
		if table[y+1][x-1] == true {
			adj += 1
		}
	}

	// BC
	if validPos(x, y+1, width, height) {
		if table[y+1][x] == true {
			adj += 1
		}
	}

	// BR
	if validPos(x+1, y+1, width, height) {
		if table[y+1][x+1] == true {
			adj += 1
		}
	}

	return adj
}

func validPos(x, y, xx, yy int) bool {
	if x < 0 || x >= xx {
		return false
	}

	if y < 0 || y >= yy {
		return false
	}

	return true
}

func parseInput(input string) (table [][]bool, countRows int, countCols int) {
	rows := strings.Split(input, "\n")
	countRows = len(rows)
	countCols = len(rows[0])

	table = make([][]bool, countRows)
	for i := range len(table) {
		table[i] = make([]bool, countCols)

		for pos, char := range rows[i] {
			if string(char) == "@" {
				table[i][pos] = true
			}
		}
	}

	return table, countRows, countCols
}
