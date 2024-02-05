package main

func maximumScore(nums []int, k int) int {
	mins := make([]int, len(nums))

	minimum := nums[k]

	for i := k; i >= 0; i-- {
		if nums[i] < minimum {
			minimum = nums[i]
		}
		mins[i] = minimum
	}

	minimum = nums[k]

	for i := k; i < len(nums); i++ {
		if nums[i] < minimum {
			minimum = nums[i]
		}
		mins[i] = minimum
	}

	maxScore := 0
	l, r := k, k

	for l >= 0 && r < len(nums) {
		score := min(mins[l], mins[r]) * (r - l + 1)

		if score > maxScore {
			maxScore = score
		}

		if l-1 >= 0 && r+1 < len(nums) {
			if mins[l-1] < mins[r+1] {
				r++
				continue
			}
			l--
			continue
		}

		if l-1 >= 0 {
			l--
			continue
		}

		r++
	}

	return maxScore
}
