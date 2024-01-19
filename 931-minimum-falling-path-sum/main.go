package main

func minFallingPathSum(matrix [][]int) int {
	size, minimimum := len(matrix), 10001

	for row := 1; row < size; row++ {
		for col := 0; col < size; col++ {
			matrix[row][col] += min(
				pick(matrix, size, row-1, col-1),
				pick(matrix, size, row-1, col),
				pick(matrix, size, row-1, col+1),
			)
		}
	}

	for _, val := range matrix[size-1] {
		if val < minimimum {
			minimimum = val
		}
	}

	return minimimum
}

func pick(matrix [][]int, size, row, col int) int {
	if col < 0 || size <= col {
		return 10001
	}

	return matrix[row][col]
}
