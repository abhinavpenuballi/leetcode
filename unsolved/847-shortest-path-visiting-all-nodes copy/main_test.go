package main

import (
	"testing"
	"time"
)

func TestShortestPathLength(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type shortestPathLengthTest struct {
			arg      [][]int
			expected int
		}

		shortestPathLengthTests := []shortestPathLengthTest{
			{[][]int{{1, 2, 3}, {0}, {0}, {0}}, 4},
			{[][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}, 4},
		}

		for _, test := range shortestPathLengthTests {
			if output := shortestPathLength(test.arg); output != test.expected {
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
