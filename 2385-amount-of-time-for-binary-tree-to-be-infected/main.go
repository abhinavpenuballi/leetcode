package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	adjacencyList := map[int][]int{}

	checkNode(root, nil, &adjacencyList)

	fmt.Println(adjacencyList)

	return 0
}

func checkNode(node, parent *TreeNode, adjacencyList *map[int][]int) {
	if node == nil {
		return
	}

	if parent != nil {
		(*adjacencyList)[node.Val] = append((*adjacencyList)[node.Val], parent.Val)
	}

	if node.Left != nil {
		(*adjacencyList)[node.Val] = append((*adjacencyList)[node.Val], node.Left.Val)
	}

	if node.Right != nil {
		(*adjacencyList)[node.Val] = append((*adjacencyList)[node.Val], node.Right.Val)
	}

	checkNode(node.Left, node, adjacencyList)
	checkNode(node.Right, node, adjacencyList)
}

func findMaxDist(adjacencyList map[int][]int, start int, path []int) int {
	if alreadySeen(start, path) {
		return len(path)
	}

	path = append(path, start)

	maxDist := 1

	for _, next := range adjacencyList[start] {
		dist := findMaxDist(adjacencyList, next, path)

		if dist > maxDist {
			maxDist = dist
		}
	}

	return maxDist
}

func alreadySeen(start int, path []int) bool {
	for _, n := range path {
		if n == start {
			return true
		}
	}
	return false
}
