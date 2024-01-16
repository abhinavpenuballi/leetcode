package main

import "fmt"

func main() {
	longest([]int{0, -1, 1})
	longest([]int{0, -1, -1, 1, -1, -1})
	longest([]int{0, 1, -1, 1, -1})
	longest([]int{0, 1, -1, -1})
	longest([]int{0, -1})
}

func longest(path []int) int {
	longest, currLongest := 0, 0

	for i := 1; i < len(path); i++ {
		if path[i] != path[i-1] {
			currLongest++

			if currLongest > longest {
				longest = currLongest
			}
		} else {
			currLongest = 1
		}
	}

	fmt.Println(path, longest)

	return longest
}
