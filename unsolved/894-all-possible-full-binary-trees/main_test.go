package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAllPossibleFBT(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type allPossibleFBTTest struct {
			arg      int
			expected []*TreeNode
		}

		allPossibleFBTTests := []allPossibleFBTTest{
			/* {7, []*TreeNode{

			}}, */
			{3, []*TreeNode{
				&TreeNode{Left: &TreeNode{}, Right: &TreeNode{}},
			}},
		}

		for _, test := range allPossibleFBTTests {
			if output := allPossibleFBT(test.arg); !reflect.DeepEqual(output, test.expected) {
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
