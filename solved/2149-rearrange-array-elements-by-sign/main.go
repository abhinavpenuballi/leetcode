package main

func rearrangeArray(nums []int) []int {
	putNextPositiveAt := 0
	putNextNegativeAt := 1

	for i, val := range nums {
		if val > 0 {
			nums = removeAndRotateRight(nums, i, putNextPositiveAt)
			putNextPositiveAt += 2
		} else {
			nums = removeAndRotateRight(nums, i, putNextNegativeAt)
			putNextNegativeAt += 2
		}
	}

	return nums
}

func removeAndRotateRight(nums []int, removeAt, putNextAt int) []int {
	for i := removeAt; i > putNextAt; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	return nums
}
