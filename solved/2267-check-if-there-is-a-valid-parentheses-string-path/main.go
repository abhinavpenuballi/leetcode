package main

import "fmt"

func hasValidPath(grid [][]byte) bool {
	return next(grid, len(grid)-1, len(grid[0])-1, 0, 0, 0, &map[string]bool{})
}

func next(grid [][]byte, rows, cols, row, col, stack int, dp *map[string]bool) bool {
	state := fmt.Sprintf("%d,%d,%d", row, col, stack)

	if val, ok := (*dp)[state]; ok {
		return val
	}

	if grid[row][col] == '(' {
		stack++
	} else {
		stack--
	}

	if stack < 0 {
		(*dp)[state] = false

		return false
	}

	if row < rows {
		if next(grid, rows, cols, row+1, col, stack, dp) {
			return true
		}

		(*dp)[state] = false
	}

	if col < cols {
		if next(grid, rows, cols, row, col+1, stack, dp) {
			return true
		}

		(*dp)[state] = false
	}

	if row == rows && col == cols && stack == 0 {
		fmt.Println(dp)
		return true
	} else {
		(*dp)[state] = false

		return false
	}
}
