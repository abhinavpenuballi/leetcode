package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	checkNodes(head, nil)

	return head
}

func checkNodes(node, parent *ListNode) int {
	if node == nil {
		return 0
	}

	maxVal := checkNodes(node.Next, node)

	if maxVal > node.Val {
		if parent != nil {
			parent.Next = node.Next
		} else {
			*node = *(node.Next)
		}

		return maxVal
	}

	return node.Val
}
