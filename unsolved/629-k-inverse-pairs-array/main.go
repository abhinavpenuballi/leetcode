package main

func kInversePairs(n int, k int) int {
	pickFrom := []int{}

	for i := 1; i <= n; i++ {
		pickFrom = append(pickFrom, i)
	}

	return next(pickFrom, []int{}, k)
}

func next(pickFrom, picked []int, k int) int {
	if len(pickFrom) == 0 {
		inversePairs := 0

		for i := 0; i < len(picked)-1; i++ {
			for j := i + 1; j < len(picked); j++ {
				if picked[i] > picked[j] {
					inversePairs++
				}
			}
		}

		if inversePairs == k {
			return 1
		} else {
			return 0
		}
	}

	inversePairs := 0

	for i, val := range pickFrom {
		pickFromLeft := make([]int, len(pickFrom[:i]))
		pickFromRight := make([]int, len(pickFrom[i+1:]))

		copy(pickFromLeft, pickFrom[:i])
		copy(pickFromRight, pickFrom[i+1:])

		inversePairs += next(append(pickFromLeft, pickFromRight...), append(picked, val), k)
	}

	return inversePairs
}
