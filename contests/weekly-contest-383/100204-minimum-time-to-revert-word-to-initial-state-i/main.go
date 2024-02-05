package main

import (
	"strings"
)

func minimumTimeToInitialState(word string, k int) int {
	time := 1

	for initialWord := word; len(word) > k; time++ {
		word = word[k:]
		if strings.Index(initialWord, word) == 0 {
			return time
		}
	}

	return time
}
