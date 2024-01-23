package main

import (
	"testing"
	"time"
)

func TestUniquePaths(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type uniquePathsTest struct {
			arg1, arg2, expected int
		}

		uniquePathsTests := []uniquePathsTest{
			{3, 7, 28},
			{3, 2, 3},
		}

		for _, test := range uniquePathsTests {
			if output := uniquePaths(test.arg1, test.arg2); output != test.expected {
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
