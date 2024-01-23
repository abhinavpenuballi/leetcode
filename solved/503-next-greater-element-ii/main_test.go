package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNextGreaterElements(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type nextGreaterElementsTest struct {
			arg1, expected []int
		}

		nextGreaterElementsTests := []nextGreaterElementsTest{
			{[]int{1, 2, 1}, []int{2, -1, 2}},
			{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		}

		for _, test := range nextGreaterElementsTests {
			if output := nextGreaterElements(test.arg1); !reflect.DeepEqual(output, test.expected) {
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
