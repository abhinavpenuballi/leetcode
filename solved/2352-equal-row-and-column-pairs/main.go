package main

import "fmt"

func main() {
	grids := [][][]int{
		{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
		{{3, 1, 2, 2}, {1, 4, 4, 4}, {2, 4, 2, 2}, {2, 4, 2, 2}},
	}

	for _, grid := range grids {
		fmt.Println(equalPairs(grid))
	}
}

func equalPairs(grid [][]int) int {
	rows := map[string]int{}
	cols := map[string]int{}

	numOfRows := len(grid)
	numOfCols := len(grid[0])

	for i := 0; i < numOfRows; i++ {
		row := ""

		for j := 0; j < numOfCols; j++ {
			row += fmt.Sprint(grid[i][j]) + ","
		}

		rows[row]++
	}

	for i := 0; i < numOfCols; i++ {
		col := ""

		for j := 0; j < numOfRows; j++ {
			col += fmt.Sprint(grid[j][i]) + ","
		}

		cols[col]++
	}

	matches := 0

	for i, j := range rows {
		matches += j * cols[i]
	}

	return matches
}
