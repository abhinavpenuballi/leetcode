package main

func rangeBitwiseAnd(left int, right int) int {
	and := left

	for i := left + 1; i <= right; i++ {
		and &= i

		if and == 0 {
			break
		}
	}

	return and
}
