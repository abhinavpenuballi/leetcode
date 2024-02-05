package main

import (
	"testing"
	"time"
)

func TestKInversePairs(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type kInversePairsTest struct {
			arg1, arg2, expected int
		}

		kInversePairsTests := []kInversePairsTest{
			{3, 0, 1},
			{3, 1, 2},
			{10, 3, 155},
		}

		for _, test := range kInversePairsTests {
			if output := kInversePairs(test.arg1, test.arg2); output != test.expected {
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
