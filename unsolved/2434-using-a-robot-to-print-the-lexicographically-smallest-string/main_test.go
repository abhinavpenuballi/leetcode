package main

import "testing"

func TestRobotWithString(t *testing.T) {
	type robotWithStringTest struct {
		arg, expected string
	}

	robotWithStringTests := []robotWithStringTest{
		{"zza", "azz"},
		{"bac", "abc"},
		{"bdda", "addb"},
		{"abc", "abc"},
		{"mmuqezwmomeplrtskz", "eekstrlpmomwzqummz"},
	}

	for _, test := range robotWithStringTests {
		if output := robotWithString(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
