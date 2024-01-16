package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 != nil {
			return list1
		}

		if list2 != nil {
			return list2
		}

		return nil
	}

	list := &ListNode{}
	head := list

	if list1.Val < list2.Val {
		list = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		list = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}

	for list1.Next != nil && list2.Next != nil {
		if list1.Val < list2.Val {
			list.Next = list1
			list = list.Next
			list1 = list1.Next
		} else if list1.Val == list2.Val {
			list.Next = list1
			list = list.Next
			list1 = list1.Next

			list.Next = list2
			list = list.Next
			list2 = list2.Next
		} else {
			list.Next = list2
			list = list.Next
			list2 = list2.Next
		}
	}

	if list1.Next != nil {
		list.Next = list1
	}

	if list2.Next != nil {
		list.Next = list2
	}

	return head
}
