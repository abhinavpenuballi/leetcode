package main

import (
	"testing"
	"time"
)

func TestSumOfMultiples(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type sumOfMultiplesTest struct {
			arg, expected int
		}

		sumOfMultiplesTests := []sumOfMultiplesTest{
			{7, 21},
			{10, 40},
			{9, 30},
		}

		for _, test := range sumOfMultiplesTests {
			if output := sumOfMultiples(test.arg); output != test.expected {
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
