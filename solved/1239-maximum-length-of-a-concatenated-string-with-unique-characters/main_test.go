package main

import (
	"testing"
	"time"
)

func TestMaxLength(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type maxLengthTest struct {
			arg      []string
			expected int
		}

		maxLengthTests := []maxLengthTest{
			{[]string{"un", "iq", "ue"}, 4},
			{[]string{"cha", "r", "act", "ers"}, 6},
			{[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
			{[]string{"aa", "bb"}, 0},
			{[]string{"jnfbyktlrqumowxd", "mvhgcpxnjzrdei"}, 16},
		}

		for _, test := range maxLengthTests {
			if output := maxLength(test.arg); output != test.expected {
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
