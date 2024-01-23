package main

import "slices"

func maxSatisfaction(satisfaction []int) int {
	slices.Sort(satisfaction)

	if satisfaction[len(satisfaction)-1] <= 0 {
		return 0
	}

	start := 0

	for i := 1; i < len(satisfaction); i++ {
		if satisfaction[i-1] < 0 && satisfaction[i] >= 0 {
			start = i
			break
		}
	}

	max := 0

	for i := start; i >= 0; i-- {
		mul, sum := 1, 0

		for j := i; j < len(satisfaction); j++ {
			sum += satisfaction[j] * mul
			mul++
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
