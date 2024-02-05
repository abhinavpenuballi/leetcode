package main

import "sort"

type MyCalendar struct {
	events [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, event := range this.events {
		if (event[0] <= start && start < event[1]) || (event[0] < end && end <= event[1]) || (start <= event[0] && event[1] <= end) {
			return false
		}
	}

	this.events = append(this.events, [2]int{start, end})

	sort.Slice(this.events, func(i, j int) bool {
		return this.events[i][0] < this.events[j][0]
	})

	return true
}
