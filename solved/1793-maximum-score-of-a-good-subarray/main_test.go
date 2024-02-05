package main

import (
	"testing"
	"time"
)

func TestMaximumScore(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type maximumScoreTest struct {
			arg1           []int
			arg2, expected int
		}

		maximumScoreTests := []maximumScoreTest{
			{[]int{1, 4, 3, 7, 4, 5}, 3, 15},
			{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0, 20},
		}

		for _, test := range maximumScoreTests {
			if output := maximumScore(test.arg1, test.arg2); output != test.expected {
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
