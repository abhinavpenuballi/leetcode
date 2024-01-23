package main

import (
	"fmt"
	"strconv"
)

func subtractProductAndSum(n int) int {
	digits := []int{}

	for _, digit := range fmt.Sprint(n) {
		num, _ := strconv.Atoi(string(digit))
		digits = append(digits, num)
	}

	product, sum := 1, 0

	for _, digit := range digits {
		product *= digit
		sum += digit
	}

	return product - sum
}
