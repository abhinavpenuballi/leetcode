package main

import (
	"testing"
	"time"
)

func TestCheckOnesSegment(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type checkOnesSegmentTest struct {
			arg      string
			expected bool
		}

		checkOnesSegmentTests := []checkOnesSegmentTest{
			{"1001", false},
			{"110", true},
		}

		for _, test := range checkOnesSegmentTests {
			if output := checkOnesSegment(test.arg); output != test.expected {
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
