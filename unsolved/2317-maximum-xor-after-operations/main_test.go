package main

import (
	"testing"
	"time"
)

func TestMaximumXOR(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type maximumXORTest struct {
			arg      []int
			expected int
		}

		maximumXORTests := []maximumXORTest{
			{[]int{3, 2, 4, 6}, 7},
			{[]int{1, 2, 3, 9, 2}, 11},
			//{[]int{957, 405, 480, 21, 73, 509, 419, 529, 207, 879, 192, 134, 979, 678, 627, 919, 107, 995, 72, 245, 940, 310, 554, 662, 738, 608, 190, 931, 458, 499, 272, 718, 175, 822, 603, 821, 453, 813, 624, 181, 955, 729, 724, 281, 880, 561, 509, 916, 265, 232, 898, 736, 476, 551, 396, 916, 34, 855, 340, 160, 835, 75, 403, 283, 631, 731, 345, 295, 175, 682, 766}, 1023},
		}

		for _, test := range maximumXORTests {
			if output := maximumXOR(test.arg); output != test.expected {
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
