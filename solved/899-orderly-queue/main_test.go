package main

import (
	"testing"
	"time"
)

func TestOrderlyQueue(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type orderlyQueueTest struct {
			arg1     string
			arg2     int
			expected string
		}

		orderlyQueueTests := []orderlyQueueTest{
			{"cba", 1, "acb"},
			{"baaca", 3, "aaabc"},
		}

		for _, test := range orderlyQueueTests {
			if output := orderlyQueue(test.arg1, test.arg2); output != test.expected {
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
