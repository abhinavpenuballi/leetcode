package main

import (
	"testing"
	"time"
)

func TestCountRoutes(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type countRoutesTest struct {
			arg1                       []int
			arg2, arg3, arg4, expected int
		}

		countRoutesTests := []countRoutesTest{
			{[]int{2, 3, 6, 8, 4}, 1, 3, 5, 4},
			{[]int{4, 3, 1}, 1, 0, 6, 5},
			{[]int{5, 2, 1}, 0, 2, 3, 0},
			{[]int{1, 2, 3}, 0, 2, 40, 615088286},
		}

		for _, test := range countRoutesTests {
			if output := countRoutes(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
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
