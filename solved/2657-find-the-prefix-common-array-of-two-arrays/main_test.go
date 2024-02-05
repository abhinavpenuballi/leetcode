package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type findThePrefixCommonArrayTest struct {
			arg1, arg2, expected []int
		}

		findThePrefixCommonArrayTests := []findThePrefixCommonArrayTest{
			{[]int{1, 3, 2, 4}, []int{3, 1, 2, 4}, []int{0, 2, 3, 4}},
			{[]int{2, 3, 1}, []int{3, 1, 2}, []int{0, 1, 3}},
		}

		for _, test := range findThePrefixCommonArrayTests {
			if output := findThePrefixCommonArray(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
