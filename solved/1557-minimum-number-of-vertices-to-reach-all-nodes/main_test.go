package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFindSmallestSetOfVertices(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type findSmallestSetOfVerticesTest struct {
			arg1     int
			arg2     [][]int
			expected []int
		}

		findSmallestSetOfVerticesTests := []findSmallestSetOfVerticesTest{
			{6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}, []int{0, 3}},
			{5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}, []int{0, 2, 3}},
		}

		for _, test := range findSmallestSetOfVerticesTests {
			if output := findSmallestSetOfVertices(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
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
