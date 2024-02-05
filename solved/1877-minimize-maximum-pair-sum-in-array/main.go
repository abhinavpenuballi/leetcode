package main

import (
	"slices"
)

func minPairSum(nums []int) int {
	slices.Sort(nums)

	l, r := 0, len(nums)-1
	max := 0

	for l < r {
		sum := nums[l] + nums[r]

		if sum > max {
			max = sum
		}

		l++
		r--
	}

	return max
}
