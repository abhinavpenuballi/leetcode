package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCellsInRange(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type cellsInRangeTest struct {
			arg      string
			expected []string
		}

		cellsInRangeTests := []cellsInRangeTest{
			{"K1:L2", []string{"K1", "K2", "L1", "L2"}},
			{"A1:F1", []string{"A1", "B1", "C1", "D1", "E1", "F1"}},
		}

		for _, test := range cellsInRangeTests {
			if output := cellsInRange(test.arg); !reflect.DeepEqual(output, test.expected) {
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
