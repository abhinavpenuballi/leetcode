package main

import (
	"fmt"
	"strconv"
)

func nextGreaterElement(n int) int {
	nStr := fmt.Sprint(n)

	last, minIndex := nStr[len(nStr)-1], len(nStr)-1

	for i := len(nStr) - 2; i >= 0; i-- {
		if nStr[i] < last {
			minIndex = i
			break
		}
	}

	if minIndex == len(nStr)-1 {
		return -1
	}

	nStr = nStr[:minIndex] + string(last) + nStr[minIndex:len(nStr)-1]

	n, _ = strconv.Atoi(nStr)

	if n == int(int32(n)) {
		return n
	} else {
		return -1
	}
}
