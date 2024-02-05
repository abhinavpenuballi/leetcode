package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	return next(m, n, maxMove, startRow, startColumn, &map[int]map[int]map[int]int{})
}

func next(rows, cols, moves, row, col int, dp *map[int]map[int]map[int]int) int {
	if val, ok := (*dp)[row][col][moves]; ok {
		return val
	}

	if (row < 0 || rows <= row) || (col < 0 || cols <= col) {
		return 1
	}

	if moves <= 0 {
		return 0
	}

	if _, ok := (*dp)[row]; !ok {
		(*dp)[row] = make(map[int]map[int]int)
	}

	if _, ok := (*dp)[row][col]; !ok {
		(*dp)[row][col] = make(map[int]int)
	}

	(*dp)[row][col][moves] = (next(rows, cols, moves-1, row-1, col, dp) +
		next(rows, cols, moves-1, row+1, col, dp) +
		next(rows, cols, moves-1, row, col-1, dp) +
		next(rows, cols, moves-1, row, col+1, dp)) % 1000000007

	return (*dp)[row][col][moves]
}
