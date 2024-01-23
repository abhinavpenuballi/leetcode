package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNextGreaterElement(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type nextGreaterElementTest struct {
			arg1, arg2, expected []int
		}

		nextGreaterElementTests := []nextGreaterElementTest{
			{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
			{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		}

		for _, test := range nextGreaterElementTests {
			if output := nextGreaterElement(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
