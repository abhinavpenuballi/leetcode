package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFindingUsersActiveMinutes(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type findingUsersActiveMinutesTest struct {
			arg1     [][]int
			arg2     int
			expected []int
		}

		findingUsersActiveMinutesTests := []findingUsersActiveMinutesTest{
			{[][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5, []int{0, 2, 0, 0, 0}},
			{[][]int{{1, 1}, {2, 2}, {2, 3}}, 4, []int{1, 1, 0, 0}},
		}

		for _, test := range findingUsersActiveMinutesTests {
			if output := findingUsersActiveMinutes(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
