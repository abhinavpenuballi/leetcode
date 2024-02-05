package main

import (
	"testing"
	"time"
)

func TestMaxCoins(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type maxCoinsTest struct {
			arg      []int
			expected int
		}

		maxCoinsTests := []maxCoinsTest{
			{[]int{3, 1, 5, 8}, 167},
			{[]int{1, 5}, 10},
			{[]int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3}, 1654},
		}

		for _, test := range maxCoinsTests {
			if output := maxCoins(test.arg); output != test.expected {
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
