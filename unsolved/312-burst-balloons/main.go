package main

func maxCoins(nums []int) int {
	return next(nums, 0)
}

func next(nums []int, sum int) int {
	if len(nums) == 0 {
		return sum
	}

	maximum := 0

	for i := 0; i < len(nums); i++ {
		numsCopy := copySlice(nums)

		val := next(append(numsCopy[:i], numsCopy[i+1:]...), sum+pick(nums, i-1)*pick(nums, i)*pick(nums, i+1))

		if val > maximum {
			maximum = val
		}
	}

	return maximum
}

func pick(nums []int, i int) int {
	if i < 0 || len(nums) <= i {
		return 1
	}

	return nums[i]
}

func copySlice(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}
