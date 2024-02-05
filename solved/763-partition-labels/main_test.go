package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPartitionLabels(t *testing.T) {
	timeout := time.After(2 * time.Second)

	done := make(chan bool)

	go func() {
		type partitionLabelsTest struct {
			arg      string
			expected []int
		}

		partitionLabelsTests := []partitionLabelsTest{
			{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
			{"eccbbbbdec", []int{10}},
			{"ababcbacadefegdemnohijhklij", []int{9, 7, 8}},
		}

		for _, test := range partitionLabelsTests {
			if output := partitionLabels(test.arg); !reflect.DeepEqual(output, test.expected) {
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
