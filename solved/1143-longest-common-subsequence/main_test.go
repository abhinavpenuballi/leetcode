package main

import (
	"testing"
	"time"
)

func TestLongestCommonSubsequence(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type longestCommonSubsequenceTest struct {
			arg1, arg2 string
			expected   int
		}

		longestCommonSubsequenceTests := []longestCommonSubsequenceTest{
			{"abcde", "ace", 3},
			{"abcde", "acea", 3},
			{"abc", "abc", 3},
			{"abc", "def", 0},
			{"abcde", "abaaedcde", 5},
			{"ezupkr", "ubmrapg", 2},
			{"bl", "yby", 1},
			{"oxcpqrsvwf", "shmtulqrypy", 2},
		}

		for _, test := range longestCommonSubsequenceTests {
			if output := longestCommonSubsequence(test.arg1, test.arg2); output != test.expected {
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
