package main

import (
	"testing"
	"time"
)

func TestUniquePathsIII(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type uniquePathsIIITest struct {
			arg      [][]int
			expected int
		}

		uniquePathsIIITests := []uniquePathsIIITest{
			{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}, 2},
			{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}, 4},
			{[][]int{{0, 1}, {2, 0}}, 0},
		}

		for _, test := range uniquePathsIIITests {
			if output := uniquePathsIII(test.arg); output != test.expected {
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
