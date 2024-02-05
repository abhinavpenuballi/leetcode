package main

import (
	"testing"
	"time"
)

func TestWateringPlants(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type wateringPlantsTest struct {
			arg1           []int
			arg2, expected int
		}

		wateringPlantsTests := []wateringPlantsTest{
			{[]int{2, 2, 3, 3}, 5, 14},
			{[]int{1, 1, 1, 4, 2, 3}, 4, 30},
			{[]int{7, 7, 7, 7, 7, 7, 7}, 8, 49},
		}

		for _, test := range wateringPlantsTests {
			if output := wateringPlants(test.arg1, test.arg2); output != test.expected {
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
