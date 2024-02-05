package main

import (
	"testing"
	"time"
)

func TestSortVowelse(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type sortVowelseTest struct {
			arg, expected string
		}

		sortVowelseTests := []sortVowelseTest{
			{"lEetcOde", "lEOtcede"},
			{"lYmpH", "lYmpH"},
		}

		for _, test := range sortVowelseTests {
			if output := sortVowels(test.arg); output != test.expected {
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
