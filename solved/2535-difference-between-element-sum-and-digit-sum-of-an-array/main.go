package main

func differenceOfSum(nums []int) int {
	var elementSum, digitSum int

	for _, num := range nums {
		elementSum += num

		for ; num > 0; num /= 10 {
			digitSum += num % 10
		}
	}

	diff := elementSum - digitSum

	if diff < 0 {
		diff = -diff
	}

	return diff
}
