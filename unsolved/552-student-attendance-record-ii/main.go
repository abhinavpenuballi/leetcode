package main

import "fmt"

func checkRecord(n int) int {
	return next(n, 2, 1, &map[string]int{})
}

func next(n, lates, absents int, dp *map[string]int) int {
	state := fmt.Sprintf("%d,%d,%d", n, lates, absents)

	if val, ok := (*dp)[state]; ok {
		return val
	}

	if n == 0 {
		(*dp)[state] = 1
		return 1
	}

	n--

	possibilities := next(n, 2, absents, dp) % 1000000007

	if absents > 0 {
		possibilities += next(n, 2, absents-1, dp) % 1000000007
	}

	if lates > 0 {
		possibilities += next(n, lates-1, absents, dp) % 1000000007
	}

	(*dp)[state] = possibilities % 1000000007
	return (*dp)[state]
}
