package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	sumInRangePerRow := getSumInRangePerRow(matrix)
	fmt.Println(sumInRangePerRow)

	rows, cols := len(matrix), len(matrix[0])
	selectRows, selectCols := 1, 1

	return 0
}

func getSumInRangePerRow(matrix [][]int) map[int]map[int]map[int]int {
	sumInRangePerRow := map[int]map[int]map[int]int{}

	for rowNum, row := range matrix {
		if _, ok := sumInRangePerRow[rowNum]; !ok {
			sumInRangePerRow[rowNum] = make(map[int]map[int]int)
		}

		for startCol := 0; startCol < len(row); startCol++ {
			if _, ok := sumInRangePerRow[rowNum][startCol]; !ok {
				sumInRangePerRow[rowNum][startCol] = make(map[int]int)
			}

			for sum, endCol := 0, startCol; endCol < len(row); endCol++ {
				sum += row[endCol]
				sumInRangePerRow[rowNum][startCol][endCol] = sum
			}

		}
	}

	return sumInRangePerRow
}

func next(matrix [][]int, rows, cols, selectRows, selectCols, target int) int {
	sumsToTarget := 0

	for rowNum := 0; rowNum < rows-selectRows; rowNum++ {

	}

	for selectRowsCopy := selectRows + 1; selectRowsCopy < rows; selectRowsCopy++ {
		sumsToTarget += next(matrix, rows, cols, selectRowsCopy, selectCols, target)
	}

	for selectColsCopy := selectCols + 1; selectColsCopy < cols; selectColsCopy++ {
		sumsToTarget += next(matrix, rows, cols, selectRows, selectColsCopy, target)
	}

	return sumsToTarget
}
