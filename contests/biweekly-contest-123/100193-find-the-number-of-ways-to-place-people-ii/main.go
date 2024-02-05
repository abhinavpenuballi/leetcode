package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][][]int{
		{{1, 1}, {2, 2}, {3, 3}},
		{{6, 2}, {4, 4}, {2, 6}},
		{{3, 1}, {1, 3}, {1, 1}},
		{{1, 0}, {22, 14}, {4, 10}, {9, 6}, {15, 4}, {14, 8}, {10, 23}, {10, 14}, {2, 12}, {19, 7}, {21, 12}, {22, 7}, {16, 22}},
	}

	for _, test := range tests {
		fmt.Println(numberOfPairs(test))
	}
}

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] > points[j][1]
	})

	fmt.Println(points)

	ways := 0

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] <= points[j][0] && points[i][1] >= points[j][1] && !anyoneElse(points, i, j) {
				ways++
			}
		}
	}

	return ways
}

func anyoneElse(points [][]int, i, j int) bool {
	for k := 0; k < len(points); k++ {
		if k == i || k == j {
			continue
		}

		if (points[i][0] <= points[k][0] && points[k][0] <= points[j][0]) &&
			(points[i][1] >= points[k][1] && points[k][1] >= points[j][1]) {
			return true
		}
	}

	return false
}
