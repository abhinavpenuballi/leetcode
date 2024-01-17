package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	type uniqueOccurrencesTest struct {
		arg      []int
		expected bool
	}

	uniqueOccurrencesTests := []uniqueOccurrencesTest{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range uniqueOccurrencesTests {
		if output := uniqueOccurrences(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
