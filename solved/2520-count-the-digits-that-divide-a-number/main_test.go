package main

import (
	"testing"
	"time"
)

func TestCountDigits(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type countDigitsTest struct {
			arg, expected int
		}

		countDigitsTests := []countDigitsTest{
			{7, 1},
			{121, 2},
			{1248, 4},
		}

		for _, test := range countDigitsTests {
			if output := countDigits(test.arg); output != test.expected {
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
