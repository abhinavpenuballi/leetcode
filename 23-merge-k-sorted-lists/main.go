package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lists := []*ListNode{
		{Val: 1,
			Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 5}}},

		{Val: 1,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4}}},

		{Val: 2,
			Next: &ListNode{Val: 6}},
	}

	mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	var list *ListNode

	for len(lists) != 0 {
		fmt.Println(next(lists))
		time.Sleep(500 * time.Millisecond)
	}

	return list
}

func next(lists []*ListNode) ([]*ListNode, *ListNode) {
	if len(lists) == 0 {
		return lists, nil
	}

	lowNode, lowIndex := lists[0], 0

	for i, node := range lists {
		if node != nil && lowNode != nil && node.Val < lowNode.Val {
			lowNode = node
			lowIndex = i
		}
	}

	if lists[lowIndex] != nil {
		lists[lowIndex] = lists[lowIndex].Next
	}

	lists = removeNils(lists)

	return lists, lowNode
}

func removeNils(lists []*ListNode) []*ListNode {
	newLists := []*ListNode{}

	for _, list := range lists {
		if list != nil {
			newLists = append(newLists, list)
		}
	}

	return newLists
}
