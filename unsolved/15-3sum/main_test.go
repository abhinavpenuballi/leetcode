package main

import (
	"reflect"
	"testing"
	"time"
)

func TestThreeSum(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type threeSumTest struct {
			arg      []int
			expected [][]int
		}

		threeSumTests := []threeSumTest{
			{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
			{[]int{0, 1, 1}, [][]int{}},
			{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		}

		for _, test := range threeSumTests {
			if output := threeSum(test.arg); !reflect.DeepEqual(output, test.expected) {
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
