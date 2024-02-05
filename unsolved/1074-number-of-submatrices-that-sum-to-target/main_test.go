package main

import (
	"testing"
	"time"
)

func TestNumSubmatrixSumTarget(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type numSubmatrixSumTargetTest struct {
			arg1           [][]int
			arg2, expected int
		}

		numSubmatrixSumTargetTests := []numSubmatrixSumTargetTest{
			{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
			{[][]int{{1, -1}, {-1, 1}}, 0, 5},
			{[][]int{{904}}, 0, 0},
		}

		for _, test := range numSubmatrixSumTargetTests {
			if output := numSubmatrixSumTarget(test.arg1, test.arg2); output != test.expected {
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
