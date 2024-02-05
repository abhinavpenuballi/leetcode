package main

import (
	"testing"
	"time"
)

func TestCountMatches(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type countMatchesTest struct {
			arg1       [][]string
			arg2, arg3 string
			expected   int
		}

		countMatchesTests := []countMatchesTest{
			{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver", 1},
			{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone", 2},
		}

		for _, test := range countMatchesTests {
			if output := countMatches(test.arg1, test.arg2, test.arg3); output != test.expected {
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
