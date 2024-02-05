package main

import (
	"testing"
	"time"
)

func TestPairSum(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type pairSumTest struct {
			arg      *ListNode
			expected int
		}

		pairSumTests := []pairSumTest{
			{&ListNode{5, &ListNode{4, &ListNode{2, &ListNode{1, nil}}}}, 6},
			{&ListNode{4, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}, 7},
			{&ListNode{1, &ListNode{100000, nil}}, 100001},
		}

		for _, test := range pairSumTests {
			if output := pairSum(test.arg); output != test.expected {
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
