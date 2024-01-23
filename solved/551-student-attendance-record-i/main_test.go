package main

import (
	"testing"
	"time"
)

func TestCheckRecord(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type checkRecordTest struct {
			arg      string
			expected bool
		}

		checkRecordTests := []checkRecordTest{
			{"PPALLP", true},
			{"PPALLL", false},
		}

		for _, test := range checkRecordTests {
			if output := checkRecord(test.arg); output != test.expected {
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
