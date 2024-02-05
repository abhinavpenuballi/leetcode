package main

import (
	"reflect"
	"testing"
	"time"
)

func TestBuildArray(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type buildArrayTest struct {
			arg1     []int
			arg2     int
			expected []string
		}

		buildArrayTests := []buildArrayTest{
			{[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
			{[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
			{[]int{1, 2}, 4, []string{"Push", "Push"}},
		}

		for _, test := range buildArrayTests {
			if output := buildArray(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
