package main

import (
	"testing"
	"time"
)

func TestTruncateSentence(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type truncateSentenceTest struct {
			arg1     string
			arg2     int
			expected string
		}

		truncateSentenceTests := []truncateSentenceTest{
			{"Hello how are you Contestant", 4, "Hello how are you"},
			{"What is the solution to this problem", 4, "What is the solution"},
			{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
		}

		for _, test := range truncateSentenceTests {
			if output := truncateSentence(test.arg1, test.arg2); output != test.expected {
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
