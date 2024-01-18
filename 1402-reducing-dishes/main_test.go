package main

import "testing"

func TestMaxSatisfaction(t *testing.T) {
	type maxSatisfactionTest struct {
		arg      []int
		expected int
	}

	maxSatisfactionTests := []maxSatisfactionTest{
		{[]int{-1, -8, 0, 5, -7}, 14},
		{[]int{4, 3, 2}, 20},
		{[]int{-1, -4, -5}, 0},
	}

	for _, test := range maxSatisfactionTests {
		if output := maxSatisfaction(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
