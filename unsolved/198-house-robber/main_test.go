package main

import (
	"testing"
	"time"
)

func TestRob(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type robTest struct {
			arg      []int
			expected int
		}

		robTests := []robTest{
			{[]int{1, 2, 3, 1}, 4},
			{[]int{2, 7, 9, 3, 1}, 12},
			{[]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}, 4173},
		}

		for _, test := range robTests {
			if output := rob(test.arg); output != test.expected {
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
