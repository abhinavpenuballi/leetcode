package main

import (
	"testing"
	"time"
)

func TestCheckRecord(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type checkRecordTest struct {
			arg, expected int
		}

		checkRecordTests := []checkRecordTest{
			{1, 3},
			{2, 8},
			{3, 19},
			{4, 43},
			{27, 278757449},
			{28, 530803311},
			{10101, 183236316},
			{33883, 617300274},
			{99995, 969766675},
		}

		for _, test := range checkRecordTests {
			if output := checkRecord(test.arg); output != test.expected {
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
