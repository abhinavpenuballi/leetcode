package main

import "strconv"

func sequentialDigits(low int, high int) []int {
	sequentialDigits := []int{}

	for nums, selectDigits := "123456789", 1; selectDigits <= 9; selectDigits++ {
		for i := 0; i <= len(nums)-selectDigits; i++ {
			n, _ := strconv.Atoi(nums[i : i+selectDigits])
			sequentialDigits = append(sequentialDigits, n)
		}
	}

	ans := []int{}

	for i := 0; i < len(sequentialDigits); i++ {
		if low <= sequentialDigits[i] && sequentialDigits[i] <= high {
			ans = append(ans, sequentialDigits[i])
		}
	}

	return ans
}
