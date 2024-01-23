package main

import "sort"

func sumSubarrayMins(arr []int) int {
	newArr, sum := make([][2]int, len(arr)), 0

	for i := 0; i < len(arr); i++ {
		newArr[i] = [2]int{i, arr[i]}
		sum += arr[i]
	}

	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i][1] < newArr[j][1]
	})

	for size := 2; size <= len(arr); size++ {
		for l := 0; l <= len(arr)-size; l++ {
			sum += pick(newArr, l, l+size)
		}
	}

	return sum % 1000000007
}

func pick(arr [][2]int, l, r int) int {
	for _, elem := range arr {
		if l <= elem[0] && elem[0] < r {
			return elem[1]
		}
	}

	return 0
}
