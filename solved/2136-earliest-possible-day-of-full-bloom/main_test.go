package main

import (
	"testing"
	"time"
)

func TestEarliestFullBloom(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type earliestFullBloomTest struct {
			arg1, arg2 []int
			expected   int
		}

		earliestFullBloomTests := []earliestFullBloomTest{
			{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
			{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9},
			{[]int{1}, []int{1}, 2},
			{[]int{1, 2}, []int{1000, 3}, 1001},
			{[]int{27, 5, 24, 17, 27, 4, 23, 16, 6, 26, 13, 17, 21, 3, 9, 10, 28, 26, 4, 10, 28, 2}, []int{26, 9, 14, 17, 6, 14, 23, 24, 11, 6, 27, 14, 13, 1, 15, 5, 12, 15, 23, 27, 28, 12}, 348},
		}

		for _, test := range earliestFullBloomTests {
			if output := earliestFullBloom(test.arg1, test.arg2); output != test.expected {
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
