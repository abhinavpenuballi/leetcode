package main

import (
	"testing"
	"time"
)

func TestTotalNQueens(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type totalNQueensTest struct {
			arg, expected int
		}

		totalNQueensTests := []totalNQueensTest{
			/* {1, 1},
			{2, 0},
			{3, 0}, */
			{4, 2},
		}

		for _, test := range totalNQueensTests {
			if output := totalNQueens(test.arg); output != test.expected {
				t.Errorf("Output %v not equal to expected %v", output, test.expected)
			}
		}

		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Time Limit Exceeded")
	case <-done:
	}
}
