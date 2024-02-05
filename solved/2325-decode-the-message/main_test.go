package main

import (
	"testing"
	"time"
)

func TestDecodeMessage(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type decodeMessageTest struct {
			arg1, arg2, expected string
		}

		decodeMessageTests := []decodeMessageTest{
			{"the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv", "this is a secret"},
			{"eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb", "the five boxing wizards jump quickly"},
		}

		for _, test := range decodeMessageTests {
			if output := decodeMessage(test.arg1, test.arg2); output != test.expected {
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
