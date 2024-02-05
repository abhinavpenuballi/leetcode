package main

import (
	"testing"
	"time"
)

func TestRestoreString(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type restoreStringTest struct {
			arg1     string
			arg2     []int
			expected string
		}

		restoreStringTests := []restoreStringTest{
			{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
			{"abc", []int{0, 1, 2}, "abc"},
		}

		for _, test := range restoreStringTests {
			if output := restoreString(test.arg1, test.arg2); output != test.expected {
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
