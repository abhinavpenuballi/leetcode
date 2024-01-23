package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSolveNQueens(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type solveNQueensTest struct {
			arg      int
			expected [][]string
		}

		solveNQueensTests := []solveNQueensTest{
			{4, nil},
		}

		for _, test := range solveNQueensTests {
			if output := solveNQueens(test.arg); reflect.DeepEqual(output, test.expected) {
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
