package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2Indexes := map[int]int{}

	for i, val := range nums2 {
		nums2Indexes[val] = i
	}

	out := []int{}

	for _, val := range nums1 {
		found := -1

		for i := nums2Indexes[val] + 1; i < len(nums2); i++ {
			if val < nums2[i] {
				found = nums2[i]
				break
			}
		}

		out = append(out, found)
	}

	return out
}
