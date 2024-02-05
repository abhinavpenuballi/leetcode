package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head

	for fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	firstHalf, secondHalf := head, slow.Next

	firstHalfRev := &ListNode{}
	firstHalfRev = nil

	for ; firstHalf != slow.Next; firstHalf = firstHalf.Next {
		firstHalfRev = &ListNode{firstHalf.Val, firstHalfRev}
	}

	maximum := 0

	for ; firstHalfRev != nil; firstHalfRev, secondHalf = firstHalfRev.Next, secondHalf.Next {
		sum := firstHalfRev.Val + secondHalf.Val

		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}
