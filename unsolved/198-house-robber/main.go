package main

func rob(nums []int) int {
	maximum := 0

	for i := 0; i < len(nums); i++ {
		val := next(nums, i, 0)

		if val > maximum {
			maximum = val
		}
	}

	return maximum
}

func next(nums []int, robbing, sum int) int {
	if robbing >= len(nums) {
		return sum
	}

	sum += nums[robbing]

	maximum := 0

	for i := robbing + 2; i <= len(nums)+1; i++ {
		val := next(nums, i, sum)

		if val > maximum {
			maximum = val
		}
	}

	return maximum
}
