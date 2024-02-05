package main

import (
	"testing"
	"time"
)

func TestFindPaths(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type findPathsTest struct {
			arg1, arg2, arg3, arg4, arg5, expected int
		}

		findPathsTests := []findPathsTest{
			{2, 2, 2, 0, 0, 6},
			{1, 3, 3, 0, 1, 12},
		}

		for _, test := range findPathsTests {
			if output := findPaths(test.arg1, test.arg2, test.arg3, test.arg4, test.arg5); output != test.expected {
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
