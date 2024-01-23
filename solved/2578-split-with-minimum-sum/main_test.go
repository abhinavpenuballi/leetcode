package main

import (
	"testing"
	"time"
)

func TestSplitNum(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type splitNumTest struct {
			arg, expected int
		}

		splitNumTests := []splitNumTest{
			{4325, 59},
			{687, 75},
		}

		for _, test := range splitNumTests {
			if output := splitNum(test.arg); output != test.expected {
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
