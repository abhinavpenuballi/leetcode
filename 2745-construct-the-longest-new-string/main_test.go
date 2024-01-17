package main

import "testing"

func TestLongestString(t *testing.T) {
	type longestStringTest struct {
		arg1, arg2, arg3, expected int
	}

	longestStringTests := []longestStringTest{
		{2, 5, 1, 12},
		{3, 2, 2, 14},
	}

	for _, test := range longestStringTests {
		if output := longestString(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
