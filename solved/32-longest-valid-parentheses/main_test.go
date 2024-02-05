package main

import (
	"testing"
	"time"
)

func TestLongestValidParentheses(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type longestValidParenthesesTest struct {
			arg      string
			expected int
		}

		longestValidParenthesesTests := []longestValidParenthesesTest{
			{"(()", 2},
			{")()())", 4},
			{"", 0},
		}

		for _, test := range longestValidParenthesesTests {
			if output := longestValidParentheses(test.arg); output != test.expected {
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
