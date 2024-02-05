package main

import (
	"reflect"
	"testing"
	"time"
)

func TestDivideArray(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type divideArrayTest struct {
			arg1     []int
			arg2     int
			expected [][]int
		}

		divideArrayTests := []divideArrayTest{
			{[]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2, [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}}},
			{[]int{1, 3, 3, 2, 7, 3}, 3, [][]int{}},
		}

		for _, test := range divideArrayTests {
			if output := divideArray(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
