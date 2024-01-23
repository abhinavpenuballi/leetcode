package main

import (
	"testing"
	"time"
)

func TestMaxScoreWords(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type maxScoreWordsTest struct {
			arg1     []string
			arg2     []byte
			arg3     []int
			expected int
		}

		maxScoreWordsTests := []maxScoreWordsTest{
			{[]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 23},
			{[]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}, 27},
			{[]string{"leetcode"}, []byte{'l', 'e', 't', 'c', 'o', 'd'}, []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, 0},
			{[]string{"azb", "ax", "awb", "ayb", "bpppp"}, []byte{'z', 'a', 'w', 'x', 'y', 'b', 'p', 'p', 'p'}, []int{10, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 2, 3, 3}, 14},
		}

		for _, test := range maxScoreWordsTests {
			if output := maxScoreWords(test.arg1, test.arg2, test.arg3); output != test.expected {
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
