package main

import (
	"testing"
	"time"
)

func TestMinPairSum(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type minPairSumTest struct {
			arg      []int
			expected int
		}

		minPairSumTests := []minPairSumTest{
			{[]int{3, 5, 2, 3}, 7},
			{[]int{3, 5, 4, 2, 4, 6}, 8},
		}

		for _, test := range minPairSumTests {
			if output := minPairSum(test.arg); output != test.expected {
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
