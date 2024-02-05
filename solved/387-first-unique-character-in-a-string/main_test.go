package main

import (
	"testing"
	"time"
)

func TestFirstUniqChar(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type firstUniqCharTest struct {
			arg      string
			expected int
		}

		firstUniqCharTests := []firstUniqCharTest{
			{"leetcode", 0},
			{"loveleetcode", 2},
			{"aabb", -1},
			{"aadadaad", -1},
			{"dddccdbba", 8},
		}

		for _, test := range firstUniqCharTests {
			if output := firstUniqChar(test.arg); output != test.expected {
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
