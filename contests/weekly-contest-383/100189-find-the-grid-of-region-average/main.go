package main

import "fmt"

func main() {
	fmt.Println(resultGrid([][]int{{5, 6, 7, 10}, {8, 9, 10, 10}, {11, 12, 13, 10}}, 3))
	fmt.Println(resultGrid([][]int{{10, 20, 30}, {15, 25, 35}, {20, 30, 40}, {25, 35, 45}}, 12))
	fmt.Println(resultGrid([][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}, 1))
}

func resultGrid(image [][]int, threshold int) [][]int {
	rows, cols := len(image), len(image[0])

	affected := map[int]map[int]int{}

	for i := 0; i < rows; i++ {
		affected[i] = make(map[int]int)
	}

	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}

	for rowStart := 0; rowStart <= rows-3; rowStart++ {
		for colStart := 0; colStart <= cols-3; colStart++ {
			if isRegion, val := isRegion(image, rowStart, colStart, threshold); isRegion {
				for row := rowStart; row < rowStart+3; row++ {
					for col := colStart; col < colStart+3; col++ {
						result[row][col] = (result[row][col]*affected[row][col] + val) / (affected[row][col] + 1)
						affected[row][col]++
					}
				}
			} else {
				for row := rowStart; row < rowStart+3; row++ {
					for col := colStart; col < colStart+3; col++ {
						result[row][col] = image[row][col]
					}
				}
			}
		}
	}

	return result
}

func isRegion(image [][]int, rowStart, colStart, threshold int) (bool, int) {
	sum := 0

	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			sum += image[row][col]

			for testRow := rowStart; testRow < rowStart+3; testRow++ {
				for testCol := colStart; testCol < colStart+3; testCol++ {
					if row == testRow && col == testCol {
						continue
					}
					if abs(row-testRow)+abs(col-testCol) == 1 && abs(image[row][col]-image[testRow][testCol]) > threshold {
						return false, 0
					}
				}
			}
		}
	}

	return true, sum / 9
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
