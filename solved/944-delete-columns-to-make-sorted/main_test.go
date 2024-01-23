package main

import (
	"testing"
	"time"
)

func TestMinDeletionSize(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type minDeletionSizeTest struct {
			arg      []string
			expected int
		}

		minDeletionSizeTests := []minDeletionSizeTest{
			{[]string{"cba", "daf", "ghi"}, 1},
			{[]string{"a", "b"}, 0},
			{[]string{"zyx", "wvu", "tsr"}, 3},
			{[]string{"rrjk", "furt", "guzm"}, 2},
		}

		for _, test := range minDeletionSizeTests {
			if output := minDeletionSize(test.arg); output != test.expected {
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
