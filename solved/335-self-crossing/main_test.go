package main

import (
	"testing"
	"time"
)

func TestIsSelfCrossing(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type isSelfCrossingTest struct {
			arg      []int
			expected bool
		}

		isSelfCrossingTests := []isSelfCrossingTest{
			{[]int{2, 1, 1, 2}, true},
			{[]int{1, 2, 3, 4}, false},
			{[]int{1, 1, 1, 2, 1}, true},
			{[]int{1, 1, 2, 2, 1, 1}, true},
		}

		for _, test := range isSelfCrossingTests {
			if output := isSelfCrossing(test.arg); output != test.expected {
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
