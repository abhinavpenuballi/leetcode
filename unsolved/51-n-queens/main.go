package main

import "fmt"

func solveNQueens(n int) [][]string {
	board := prepareBoard(n)

	fmt.Println(board)

	return nil
}

func prepareBoard(n int) []string {
	board, row := []string{}, ""

	for i := 0; i < n; i++ {
		row += "."
	}

	for i := 0; i < n; i++ {
		board = append(board, row)
	}

	return board
}

func fill(n, remaining int, board []string) {}
