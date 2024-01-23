package main

import (
	"fmt"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	numStr := []byte(fmt.Sprint(num))

	sort.Slice(numStr, func(i, j int) bool { return numStr[i] < numStr[j] })

	var num1, num2 string

	for i, char := range numStr {
		if i%2 == 0 {
			num1 += string(char)
		} else {
			num2 += string(char)
		}
	}

	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)

	return num1Int + num2Int
}
