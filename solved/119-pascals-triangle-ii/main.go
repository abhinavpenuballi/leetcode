package main

import "fmt"

func main() {
	ns := []int{3, 0, 1, 20}

	for _, n := range ns {
		fmt.Println(getRow(n))
	}
}

func getRow(rowIndex int) []int {
	dp := map[int]int{0: 1, 1: 1}

	row := []int{}

	for i := 0; i <= rowIndex; i++ {
		elem := fact(rowIndex, &dp) / (fact(i, &dp) * fact(rowIndex-i, &dp))
		row = append(row, elem)
	}

	return row
}

func fact(n int, dp *map[int]int) int {
	if val, ok := (*dp)[n]; ok {
		return val
	}

	(*dp)[n] = n * fact(n-1, dp)

	return (*dp)[n]
}
