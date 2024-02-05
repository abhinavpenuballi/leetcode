package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	paths := 0

	next(root, 0, &paths)

	return paths
}

func next(root *TreeNode, mask int, paths *int) {
	if root == nil {
		return
	}

	mask ^= 1 << root.Val

	if root.Left == nil && root.Right == nil {
		if mask&(mask-1) == 0 {
			*paths++
		}
		return
	}

	next(root.Left, mask, paths)
	next(root.Right, mask, paths)
}
