package main

import (
	"testing"
	"time"
)

func TestMinWindow(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type minWindowTest struct {
			arg1, arg2, expected string
		}

		minWindowTests := []minWindowTest{
			{"ADOBECODEBANC", "ABC", "BANC"},
			{"a", "a", "a"},
			{"a", "aa", ""},
		}

		for _, test := range minWindowTests {
			if output := minWindow(test.arg1, test.arg2); output != test.expected {
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
