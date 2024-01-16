package main

import "fmt"

func main() {
	ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range ns {
		fmt.Println(fib(n))
	}
}

func fib(n int) int {
	return next(n, &map[int]int{0: 0, 1: 1})
}

func next(n int, dp *map[int]int) int {
	if val, ok := (*dp)[n]; ok {
		return val
	}

	(*dp)[n] = next(n-2, dp) + next(n-1, dp)

	return (*dp)[n]
}
