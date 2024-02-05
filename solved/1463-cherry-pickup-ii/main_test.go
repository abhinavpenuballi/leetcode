package main

import (
	"testing"
	"time"
)

func TestCherryPickup(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type cherryPickupTest struct {
			arg      [][]int
			expected int
		}

		cherryPickupTests := []cherryPickupTest{
			{[][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}, 24},
			{[][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}, 28},
			{[][]int{{16, 1, 9, 16, 5, 16, 3, 16, 6, 3, 1, 7}, {10, 0, 0, 12, 9, 12, 19, 3, 9, 3, 18, 9}, {18, 6, 6, 13, 2, 1, 9, 8, 12, 2, 10, 6}, {16, 7, 6, 10, 5, 17, 16, 0, 12, 7, 5, 15}, {20, 11, 17, 15, 2, 8, 12, 2, 17, 13, 12, 0}, {11, 1, 10, 11, 13, 9, 16, 7, 1, 12, 13, 8}, {12, 19, 19, 3, 13, 0, 7, 1, 1, 3, 1, 16}, {4, 9, 1, 4, 16, 19, 11, 11, 3, 9, 2, 7}, {10, 13, 1, 4, 3, 7, 14, 3, 20, 7, 6, 11}, {8, 17, 14, 13, 2, 5, 9, 11, 11, 18, 19, 15}, {16, 4, 3, 13, 18, 17, 14, 16, 15, 12, 20, 13}, {20, 0, 19, 16, 0, 3, 16, 16, 1, 15, 2, 20}, {20, 18, 8, 11, 0, 13, 8, 7, 13, 6, 18, 19}, {9, 2, 9, 18, 10, 16, 0, 5, 9, 11, 4, 14}, {11, 4, 18, 8, 18, 18, 10, 12, 11, 0, 10, 13}, {0, 7, 9, 2, 1, 17, 4, 1, 6, 9, 7, 9}}, 465},
		}

		for _, test := range cherryPickupTests {
			if output := cherryPickup(test.arg); output != test.expected {
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
