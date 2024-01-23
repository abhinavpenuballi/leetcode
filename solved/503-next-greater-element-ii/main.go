package main

func nextGreaterElements(nums []int) []int {
	out := []int{}

	for i := 0; i < len(nums); i++ {
		found := -1

		for j := i + 1; j < i+len(nums); j++ {
			if nums[i] < nums[j%len(nums)] {
				found = nums[j%len(nums)]
				break
			}
		}

		out = append(out, found)
	}

	return out
}
