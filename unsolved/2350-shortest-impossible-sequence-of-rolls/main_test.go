package main

import (
	"testing"
	"time"
)

func TestShortestSequence(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type shortestSequenceTest struct {
			arg1           []int
			arg2, expected int
		}

		shortestSequenceTests := []shortestSequenceTest{
			{[]int{4, 2, 1, 2, 3, 3, 2, 4, 1}, 4, 3},
			{[]int{1, 1, 2, 2}, 2, 2},
			{[]int{1, 1, 3, 2, 2, 2, 3, 3}, 4, 1},
		}

		for _, test := range shortestSequenceTests {
			if output := shortestSequence(test.arg1, test.arg2); output != test.expected {
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
