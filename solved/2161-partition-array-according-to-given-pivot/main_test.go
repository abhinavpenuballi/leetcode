package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPivotArray(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type pivotArrayTest struct {
			arg1     []int
			arg2     int
			expected []int
		}

		pivotArrayTests := []pivotArrayTest{
			{[]int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
			{[]int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
		}

		for _, test := range pivotArrayTests {
			if output := pivotArray(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
