package main

import (
	"testing"
	"time"
)

func TestNextGreaterElement(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type nextGreaterElementTest struct {
			arg1, expected int
		}

		nextGreaterElementTests := []nextGreaterElementTest{
			{12, 21},
			{21, -1},
			{1234, 1243},
			{1324, 1342},
			{1243, 1324},
			{230241, 230412},
		}

		for _, test := range nextGreaterElementTests {
			if output := nextGreaterElement(test.arg1); output != test.expected {
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
