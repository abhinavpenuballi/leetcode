package main

import "fmt"

func uniquePaths(m int, n int) int {
	return next(m, n, 0, 0, &map[string]int{})
}

func next(rows, cols, row, col int, dp *map[string]int) int {
	cell := fmt.Sprint(row, col)

	if val, ok := (*dp)[cell]; ok {
		return val
	}

	if row == rows-1 && col == cols-1 {
		(*dp)[cell] = 1
		return 1
	}

	paths := 0

	if row < rows-1 {
		paths += next(rows, cols, row+1, col, dp)
	}

	if col < cols-1 {
		paths += next(rows, cols, row, col+1, dp)
	}

	(*dp)[cell] = paths
	return paths
}
