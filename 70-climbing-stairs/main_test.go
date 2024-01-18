package main

import "testing"

func TestClimbStairs(t *testing.T) {
	type climbStairsTest struct {
		arg      int
		expected int
	}

	climbStairsTests := []climbStairsTest{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{40, 165580141},
		{41, 267914296},
		{42, 433494437},
		{43, 701408733},
		{44, 1134903170},
		{100, 1298777728820984005},
	}

	for _, test := range climbStairsTests {
		if output := climbStairs(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
