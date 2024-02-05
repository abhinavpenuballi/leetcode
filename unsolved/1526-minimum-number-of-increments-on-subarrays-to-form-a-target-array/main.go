package main

import "fmt"

func minNumberOperations(target []int) int {
	return next(target, &map[string]int{})
}

func next(target []int, dp *map[string]int) int {
	state := fmt.Sprint(target)

	if val, ok := (*dp)[state]; ok {
		return val
	}

	if len(target) == 0 {
		return 0
	}

	targets := breakOnZeroes(target)

	if len(targets) == 0 {
		return 0
	}

	ops := 0

	for _, target := range targets {
		for i := 0; i < len(target); i++ {
			target[i]--
		}

		ops++

		ops += next(target, dp)
	}

	(*dp)[state] = ops
	return ops
}

func breakOnZeroes(target []int) [][]int {
	targets, nums := [][]int{}, []int{}

	for _, val := range target {
		if val != 0 {
			nums = append(nums, val)
		} else {
			if len(nums) > 0 {
				targets = append(targets, nums)
			}

			nums = []int{}
		}
	}

	if len(nums) > 0 {
		targets = append(targets, nums)
	}

	return targets
}
