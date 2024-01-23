package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return next(obstacleGrid, len(obstacleGrid), len(obstacleGrid[0]), 0, 0, &map[string]int{})
}

func next(obstacleGrid [][]int, rows, cols, row, col int, dp *map[string]int) int {
	cell := fmt.Sprint(row, col)

	if val, ok := (*dp)[cell]; ok {
		return val
	}

	if obstacleGrid[row][col] == 1 {
		(*dp)[cell] = 0
		return 0
	}

	if row == rows-1 && col == cols-1 {
		(*dp)[cell] = 1
		return 1
	}

	paths := 0

	if row < rows-1 {
		paths += next(obstacleGrid, rows, cols, row+1, col, dp)
	}

	if col < cols-1 {
		paths += next(obstacleGrid, rows, cols, row, col+1, dp)
	}

	(*dp)[cell] = paths
	return paths
}
