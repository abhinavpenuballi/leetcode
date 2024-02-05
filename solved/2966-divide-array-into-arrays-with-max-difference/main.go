package main

import (
	"slices"
)

func divideArray(nums []int, k int) [][]int {
	out := [][]int{}

	slices.Sort(nums)

	for i := 0; i < len(nums); i += 3 {
		a, b, c := nums[i], nums[i+1], nums[i+2]

		if b-a <= k && c-a <= k && c-b <= k {
			out = append(out, []int{a, b, c})
		} else {
			return [][]int{}
		}
	}

	return out
}
