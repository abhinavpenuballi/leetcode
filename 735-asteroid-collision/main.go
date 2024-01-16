package main

import (
	"fmt"
	"math"
)

func main() {
	asteroidss := [][]int{
		{5, 10, -5},
		{8, -8},
		{10, 2, -5},
		{-2, -1, 1, 2},
		{-2, -2, -2, 1},
		{1, -1, -2, -2},
		{-2, -2},
	}

	for _, asteroids := range asteroidss {
		fmt.Println(asteroidCollision(asteroids))
	}
}

func asteroidCollision(asteroids []int) []int {
	for {
		anyExplosion := false

		for i := 1; i < len(asteroids); i++ {
			if opposite(asteroids[i-1], asteroids[i]) {
				winner := winner(asteroids[i-1], asteroids[i])

				if winner == -1 {
					asteroids = append(asteroids[:i], asteroids[i+1:]...)
					i--
				} else if winner == 1 {
					asteroids = append(asteroids[:i-1], asteroids[i:]...)
					i--
				} else {
					asteroids = append(asteroids[:i-1], asteroids[i+1:]...)
					i--
				}

				anyExplosion = true
			}
		}

		if !anyExplosion {
			break
		}
	}

	return asteroids
}

func winner(a1, a2 int) int {
	if math.Abs(float64(a1)) > math.Abs(float64(a2)) {
		return -1
	}

	if math.Abs(float64(a1)) < math.Abs(float64(a2)) {
		return 1
	}

	return 0
}

func opposite(a1, a2 int) bool {
	return a1 > 0 && 0 > a2
}
