package main

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum, found := int64(-9223372036854775808), false

	for i := 0; i < len(nums)-1; i++ {
		for j, sum := i+1, int64(nums[i]); j < len(nums); j++ {
			sum += int64(nums[j])
			if abs(nums[i]-nums[j]) == k {
				found = true
				maxSum = max(maxSum, sum)
			}
		}
	}

	if !found {
		return 0
	}

	return maxSum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
