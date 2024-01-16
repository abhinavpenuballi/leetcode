package main

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	answer := []int{}

	for _, query := range queries {
		count := 0

		for _, point := range points {
			distance := math.Sqrt(math.Pow(float64(query[0]-point[0]), 2) + math.Pow(float64(query[1]-point[1]), 2))

			if distance <= float64(query[2]) {
				count++
			}
		}

		answer = append(answer, count)
	}

	return answer
}
