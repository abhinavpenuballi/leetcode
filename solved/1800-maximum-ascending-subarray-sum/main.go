package main

func maxAscendingSum(nums []int) int {
	var max, lastSeen, sum int

	for _, num := range nums {
		if num > lastSeen {
			sum += num
		} else {
			if sum > max {
				max = sum
			}

			sum = num
		}

		lastSeen = num
	}

	if sum > max {
		max = sum
	}

	return max
}
