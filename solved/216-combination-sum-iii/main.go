package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(4, 24))
}

func combinationSum3(k int, n int) [][]int {
	valid := [][]int{}

	next([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}, k, n, &valid)

	return valid
}

func next(pickable, picked []int, k, n int, valid *[][]int) {
	lenPicked := len(picked)

	if lenPicked > k {
		return
	}

	currSum := sum(picked)

	if currSum > n {
		return
	}

	if lenPicked == k && currSum == n {
		fmt.Println(picked)
		tempPicked := make([]int, len(picked))
		copy(tempPicked, picked)

		*valid = append(*valid, tempPicked)

		return
	}

	for len(pickable) > 0 {
		tempPicked := append(picked, pickable[0])

		pickable = pickable[1:]

		next(pickable, tempPicked, k, n, valid)
	}
}

func sum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}
