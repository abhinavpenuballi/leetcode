package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSequentialDigits(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type sequentialDigitsTest struct {
			arg1, arg2 int
			expected   []int
		}

		sequentialDigitsTests := []sequentialDigitsTest{
			{100, 300, []int{123, 234}},
			{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
		}

		for _, test := range sequentialDigitsTests {
			if output := sequentialDigits(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
