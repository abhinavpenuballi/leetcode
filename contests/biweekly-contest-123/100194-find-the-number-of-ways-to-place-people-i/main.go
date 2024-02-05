package main

import "fmt"

func main() {
	tests := [][][]int{
		{{1, 1}, {2, 2}, {3, 3}},
		{{6, 2}, {4, 4}, {2, 6}},
		{{3, 1}, {1, 3}, {1, 1}},
	}

	for _, test := range tests {
		fmt.Println(numberOfPairs(test))
	}
}

func numberOfPairs(points [][]int) int {
	ways := 0

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

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
