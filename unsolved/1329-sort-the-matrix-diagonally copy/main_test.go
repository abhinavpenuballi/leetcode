package main

import (
	"reflect"
	"testing"
	"time"
)

func TestDiagonalSort(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type diagonalSortTest struct {
			arg, expected [][]int
		}

		diagonalSortTests := []diagonalSortTest{
			{[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}, [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}},
			{[][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}}, [][]int{{5, 17, 4, 1, 52, 7}, {11, 11, 25, 45, 8, 69}, {14, 23, 25, 44, 58, 15}, {22, 27, 31, 36, 50, 66}, {84, 28, 75, 33, 55, 68}}},
		}

		for _, test := range diagonalSortTests {
			if output := diagonalSort(test.arg); !reflect.DeepEqual(output, test.expected) {
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
