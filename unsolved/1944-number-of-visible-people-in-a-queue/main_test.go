package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCanSeePersonsCount(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type canSeePersonsCountTest struct {
			arg, expected []int
		}

		canSeePersonsCountTests := []canSeePersonsCountTest{
			{[]int{10, 6, 8, 5, 11, 9}, []int{3, 1, 2, 1, 1, 0}},
			{[]int{5, 1, 2, 3, 10}, []int{4, 1, 1, 1, 0}},
			{[]int{3, 1, 5, 8, 6}, []int{2, 1, 1, 1, 0}},
		}

		for _, test := range canSeePersonsCountTests {
			if output := canSeePersonsCount(test.arg); !reflect.DeepEqual(output, test.expected) {
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
