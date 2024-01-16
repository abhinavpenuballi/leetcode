package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return createNodes(nums)
}

func createNodes(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, index := max(nums)
	node := &TreeNode{Val: max}

	node.Left = createNodes(nums[:index])
	node.Right = createNodes(nums[index+1:])

	return node
}

func max(nums []int) (int, int) {
	var max, index int

	for i, n := range nums {
		if n > max {
			max = n
			index = i
		}
	}

	return max, index
}
