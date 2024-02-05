package main

import "sort"

type MyCalendarThree struct {
	events   [][2]int
	overlaps map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{overlaps: map[int]int{}}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.events = append(this.events, [2]int{startTime, endTime})

	sort.Slice(this.events, func(i, j int) bool {
		return this.events[i][0] < this.events[j][0]
	})

	for i := startTime; i < endTime; i++ {
		this.overlaps[i]++
	}

	maxSeen := 0

	for _, val := range this.overlaps {
		if val > maxSeen {
			maxSeen = val
		}
	}

	return maxSeen
}

// ["MyCalendarThree","book"]
// [[],[0,1000000000]]
