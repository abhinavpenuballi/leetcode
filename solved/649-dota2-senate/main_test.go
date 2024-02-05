package main

import (
	"testing"
	"time"
)

func TestPredictPartyVictory(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type predictPartyVictoryTest struct {
			arg, expected string
		}

		predictPartyVictoryTests := []predictPartyVictoryTest{
			{"RD", "Radiant"},
			{"RDD", "Dire"},
		}

		for _, test := range predictPartyVictoryTests {
			if output := predictPartyVictory(test.arg); output != test.expected {
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
