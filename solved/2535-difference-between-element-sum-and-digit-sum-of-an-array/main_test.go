package main

import (
	"testing"
	"time"
)

func TestDifferenceOfSum(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type differenceOfSumTest struct {
			arg      []int
			expected int
		}

		differenceOfSumTests := []differenceOfSumTest{
			{[]int{1, 15, 6, 3}, 9},
			{[]int{1, 2, 3, 4}, 0},
		}

		for _, test := range differenceOfSumTests {
			if output := differenceOfSum(test.arg); output != test.expected {
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
