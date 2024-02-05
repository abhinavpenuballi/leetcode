package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFinalPrices(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type finalPricesTest struct {
			arg, expected []int
		}

		finalPricesTests := []finalPricesTest{
			{[]int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
			{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			{[]int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
		}

		for _, test := range finalPricesTests {
			if output := finalPrices(test.arg); !reflect.DeepEqual(output, test.expected) {
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
