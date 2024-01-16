package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return next(root, 0, 100000)
}

func next(root *TreeNode, maxSeen, minSeen int) int {
	if root == nil {
		return 0
	}

	if root.Val > maxSeen {
		maxSeen = root.Val
	}

	if root.Val < minSeen {
		minSeen = root.Val
	}

	diff := maxSeen - minSeen

	return max(diff, next(root.Left, maxSeen, minSeen), next(root.Right, maxSeen, minSeen))
}
