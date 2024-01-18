package main

func climbStairs(n int) int {
	return next(n, &map[int]int{0: 1, 1: 1})
}

func next(n int, dp *map[int]int) int {
	if val, ok := (*dp)[n]; ok {
		return val
	}

	(*dp)[n] = next(n-2, dp) + next(n-1, dp)

	return (*dp)[n]
}
