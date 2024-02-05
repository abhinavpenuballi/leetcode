package main

import (
	"testing"
	"time"
)

func TestMinInsertions(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type minInsertionsTest struct {
			arg      string
			expected int
		}

		minInsertionsTests := []minInsertionsTest{
			{"zzazz", 0},
			{"mbadm", 2},
			{"leetcode", 5},
			{"vsrgaxxpgfiqdnwvrlpddcz", 17},
		}

		for _, test := range minInsertionsTests {
			if output := minInsertions(test.arg); output != test.expected {
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
