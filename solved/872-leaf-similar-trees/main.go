package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var wg sync.WaitGroup

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := "", ""

	wg.Add(2)
	go func() {
		defer wg.Done()
		getLeaves(root1, &leaves1)
	}()
	go func() {
		defer wg.Done()
		getLeaves(root2, &leaves2)
	}()
	wg.Wait()

	return leaves1 == leaves2
}

func getLeaves(root *TreeNode, leaves *string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		(*leaves) += fmt.Sprint(root.Val) + "#"
		return
	}

	getLeaves(root.Left, leaves)
	getLeaves(root.Right, leaves)
}
