package main

import (
	"testing"
	"time"
)

func TestFindErrorNums(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type findErrorNumsTest struct {
			arg, expected []int
		}

		findErrorNumsTests := []findErrorNumsTest{
			{[]int{1, 2, 2, 4}, []int{2, 3}},
			{[]int{1, 1}, []int{1, 2}},
		}

		for _, test := range findErrorNumsTests {
			if output := findErrorNums(test.arg); !correctOutput(output, test.expected) {
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

func correctOutput(output, expected []int) bool {
	if len(output) != len(expected) {
		return false
	}

	for i := 0; i < len(output); i++ {
		if output[i] != expected[i] {
			return false
		}
	}

	return true
}
