package main

import (
	"testing"
	"time"
)

func TestCountGoodTriplets(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type countGoodTripletsTest struct {
			arg1                       []int
			arg2, arg3, arg4, expected int
		}

		countGoodTripletsTests := []countGoodTripletsTest{
			{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
			{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		}

		for _, test := range countGoodTripletsTests {
			if output := countGoodTriplets(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
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
