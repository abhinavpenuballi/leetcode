package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRearrangeArray(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type rearrangeArrayTest struct {
			arg, expected []int
		}

		rearrangeArrayTests := []rearrangeArrayTest{
			{[]int{3, 1, -2, -5, 2, -4}, []int{3, -2, 1, -5, 2, -4}},
			{[]int{-1, 1}, []int{1, -1}},
			{[]int{1, 2, 3, -1, -2, -3}, []int{1, -1, 2, -2, 3, -3}},
			{[]int{-1, -2, -3, 1, 2, 3}, []int{1, -1, 2, -2, 3, -3}},
			{[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, -1, 2, -2, 3, -3, 4, -4, 5, -5, 6, -6, 7, -7, 8, -8, 9, -9}},
			{[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{1, -1, 2, -2, 3, -3, 4, -4, 5, -5, 6, -6, 7, -7, 8, -8, 9, -9, 10, -10, 11, -11, 12, -12, 13, -13, 14, -14, 15, -15, 16, -16, 17, -17, 18, -18, 19, -19, 20, -20}},
		}

		for _, test := range rearrangeArrayTests {
			if output := rearrangeArray(test.arg); !reflect.DeepEqual(output, test.expected) {
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
