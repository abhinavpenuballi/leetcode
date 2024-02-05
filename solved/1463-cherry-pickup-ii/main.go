package main

import "fmt"

func cherryPickup(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	return next(grid, rows, cols, 0, 0, cols-1, 0, &map[string]int{})
}

func next(grid [][]int, rows, cols, row, c1, c2, sum int, dp *map[string]int) int {
	state := fmt.Sprint(row, c1, c2)

	if val, ok := (*dp)[state]; ok {
		return val
	}

	if (c1 >= c2) || (c1 < 0 || cols <= c1) || (c2 < 0 || cols <= c2) {
		(*dp)[state] = 0
		return 0
	}

	if row == rows-1 {
		(*dp)[state] = grid[row][c1] + grid[row][c2]
		return grid[row][c1] + grid[row][c2]
	}

	row++

	(*dp)[state] = grid[row-1][c1] + grid[row-1][c2] +
		max(
			next(grid, rows, cols, row, c1-1, c2-1, sum, dp),
			next(grid, rows, cols, row, c1-1, c2, sum, dp),
			next(grid, rows, cols, row, c1-1, c2+1, sum, dp),

			next(grid, rows, cols, row, c1, c2-1, sum, dp),
			next(grid, rows, cols, row, c1, c2, sum, dp),
			next(grid, rows, cols, row, c1, c2+1, sum, dp),

			next(grid, rows, cols, row, c1+1, c2-1, sum, dp),
			next(grid, rows, cols, row, c1+1, c2, sum, dp),
			next(grid, rows, cols, row, c1+1, c2+1, sum, dp),
		)

	return (*dp)[state]
}
