package main

import (
	"reflect"
	"testing"
	"time"
)

func TestDailyTemperatures(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type dailyTemperaturesTest struct {
			arg, expected []int
		}

		dailyTemperaturesTests := []dailyTemperaturesTest{
			{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
			{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
			{[]int{30, 60, 90}, []int{1, 1, 0}},
		}

		for _, test := range dailyTemperaturesTests {
			if output := dailyTemperatures(test.arg); !reflect.DeepEqual(output, test.expected) {
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
