package main

import "testing"

func TestHasValidPath(t *testing.T) {
	type hasValidPathTest struct {
		arg      [][]byte
		expected bool
	}

	hasValidPathTests := []hasValidPathTest{
		{[][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}}, true},
		{[][]byte{{')', ')'}, {'(', '('}}, false},
		{[][]byte{{'(', ')', ')', '(', '(', '(', '(', ')', ')', '(', ')', '(', ')', '(', '(', '(', '(', ')', '(', ')', '('}, {'(', '(', ')', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', '(', '(', ')'}, {')', ')', '(', ')', ')', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', '(', '(', ')', ')'}, {'(', '(', ')', '(', ')', '(', ')', ')', ')', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')'}, {'(', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', ')', ')', ')', '(', ')', '(', '('}, {')', ')', '(', '(', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', '(', ')', '(', '(', '(', '(', ')'}, {')', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', ')', '(', '(', ')', ')', '(', ')', '('}, {'(', ')', '(', '(', '(', ')', ')', ')', ')', '(', '(', ')', '(', '(', ')', ')', '(', ')', ')', ')', '('}, {'(', ')', '(', ')', '(', '(', '(', '(', ')', '(', '(', '(', '(', '(', '(', ')', '(', ')', '(', ')', ')'}, {'(', ')', '(', '(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', '(', ')', ')', '(', '(', '(', ')'}, {'(', '(', ')', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', ')', '(', ')', '(', ')', ')', ')', '('}, {')', '(', '(', '(', ')', '(', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', ')', '(', '(', ')'}, {'(', ')', '(', ')', ')', '(', '(', ')', '(', ')', '(', ')', ')', ')', '(', '(', '(', '(', ')', '(', ')'}, {'(', '(', ')', '(', ')', ')', '(', '(', '(', ')', '(', ')', '(', '(', ')', ')', '(', '(', '(', ')', ')'}, {'(', '(', '(', '(', ')', ')', '(', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', '('}, {'(', '(', '(', ')', ')', ')', '(', ')', ')', '(', ')', ')', '(', '(', ')', '(', ')', '(', '(', '(', ')'}, {')', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', '(', ')', '(', ')', '(', '(', '(', '(', ')', ')'}}, false},
	}

	for _, test := range hasValidPathTests {
		if output := hasValidPath(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}