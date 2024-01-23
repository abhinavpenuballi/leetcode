package main

import "fmt"

func main() {
	ns := []int{1, 2, 33, 50, 100, 1000}

	for _, n := range ns {
		fmt.Println(countVowelStrings(n))
	}
}

func countVowelStrings(n int) int {
	return next(1, n, &map[int]map[int]int{})
}

func next(start, n int, dp *map[int]map[int]int) int {
	if n == 0 {
		return 1
	}

	if (*dp)[start][n] != 0 {
		return (*dp)[start][n]
	} else {
		(*dp)[start] = map[int]int{}
	}

	ways := 0

	for i := start; i <= 5; i++ {
		ways += next(i, n-1, dp)
	}

	(*dp)[start][n] = ways

	return (*dp)[start][n]
}
