package main

import "fmt"

func main() {
	fmt.Println(minimumTimeToInitialState("abacaba", 3))
}

func minimumTimeToInitialState(word string, k int) int {
	time := 1

	for i := k; i < len(word); i, time = i+k, time+1 {
		if word[:len(word)-i] == word[i:] {
			return time
		}
	}

	return time
}
