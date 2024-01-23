package main

func uniquePathsIII(grid [][]int) int {
	var start, end [2]int

	mustCover := 1

	for i, val := range grid {
		for j, val := range val {
			switch val {
			case 1:
				start = [2]int{i, j}
				grid[i][j] = 0
			case 2:
				end = [2]int{i, j}
			case 0:
				mustCover++
			}
		}
	}

	rows, cols := len(grid), len(grid[0])

	return next(grid, rows, cols, mustCover, 0, start, end)
}

func next(grid [][]int, rows, cols, mustCover, covered int, currCell, end [2]int) int {
	if currCell == end {
		if covered == mustCover {
			return 1
		} else {
			return 0
		}
	}

	gridCopy := copyGrid(grid)

	if gridCopy[currCell[0]][currCell[1]] != 0 {
		return 0
	} else {
		gridCopy[currCell[0]][currCell[1]] = 3
	}

	covered++

	paths := 0

	if currCell[0] > 0 {
		nextCell := currCell
		nextCell[0]--
		paths += next(gridCopy, rows, cols, mustCover, covered, nextCell, end)
	}

	if currCell[0] < rows-1 {
		nextCell := currCell
		nextCell[0]++
		paths += next(gridCopy, rows, cols, mustCover, covered, nextCell, end)
	}

	if currCell[1] > 0 {
		nextCell := currCell
		nextCell[1]--
		paths += next(gridCopy, rows, cols, mustCover, covered, nextCell, end)
	}

	if currCell[1] < cols-1 {
		nextCell := currCell
		nextCell[1]++
		paths += next(gridCopy, rows, cols, mustCover, covered, nextCell, end)
	}

	return paths
}

func copyGrid(grid [][]int) [][]int {
	gridCopy := [][]int{}

	for _, row := range grid {
		gridCopy = append(gridCopy, append([]int{}, row...))
	}

	return gridCopy
}
