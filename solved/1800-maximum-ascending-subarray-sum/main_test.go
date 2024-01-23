package main

import "testing"

func TestMaxAscendingSum(t *testing.T) {
	type maxAscendingSumTest struct {
		arg      []int
		expected int
	}

	maxAscendingSumTests := []maxAscendingSumTest{
		{[]int{10, 20, 30, 5, 10, 50}, 65},
		{[]int{10, 20, 30, 40, 50}, 150},
		{[]int{12, 17, 15, 13, 10, 11, 12}, 33},
	}

	for _, test := range maxAscendingSumTests {
		if output := maxAscendingSum(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
