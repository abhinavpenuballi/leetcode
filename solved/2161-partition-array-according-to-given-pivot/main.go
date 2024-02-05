package main

func pivotArray(nums []int, pivot int) []int {
	var putNextLessAt, putNextPivotAt int

	for i, val := range nums {
		switch {
		case val < pivot:
			nums = removeAndRotateRight(nums, i, putNextLessAt)
			putNextLessAt++
			putNextPivotAt++
		case val == pivot:
			nums = removeAndRotateRight(nums, i, putNextPivotAt)
			putNextPivotAt++
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
