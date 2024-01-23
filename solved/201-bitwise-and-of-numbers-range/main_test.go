package main

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	type rangeBitwiseAndTest struct {
		arg1, arg2, expected int
	}

	rangeBitwiseAndTests := []rangeBitwiseAndTest{
		{5, 7, 4},
		{0, 0, 0},
		{1, 2147483647, 0},
		{600000000, 2147483645, 0},
	}

	for _, test := range rangeBitwiseAndTests {
		if output := rangeBitwiseAnd(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
