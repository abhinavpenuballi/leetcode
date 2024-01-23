package main

import "fmt"

func totalNQueens(n int) int {
	board := [][]string{}

	for i := 0; i < n; i++ {
		board = append(board, make([]string, n))
	}

	fmt.Println(board)

	return 0
}
