package main

import "fmt"

func main() {
	ns := []int{0, 1, 2, 3, 4}

	for _, n := range ns {
		fmt.Println(n, generateParenthesis(n))
	}
}

func generateParenthesis(n int) []string {
	combinations := []string{}

	open("", n, 0, 0, &combinations)

	return combinations
}

func open(str string, n, addedOpens, closedOpens int, combinations *[]string) {
	if n == addedOpens {
		return
	}

	str += "("
	addedOpens++

	if addedOpens > closedOpens {
		close(str, n, addedOpens, closedOpens, combinations)
	}
	open(str, n, addedOpens, closedOpens, combinations)
}

func close(str string, n, addedOpens, closedOpens int, combinations *[]string) {
	str += ")"
	closedOpens++

	if n == closedOpens {
		*combinations = append(*combinations, str)
		return
	}

	if addedOpens > closedOpens {
		close(str, n, addedOpens, closedOpens, combinations)
	}
	open(str, n, addedOpens, closedOpens, combinations)
}
