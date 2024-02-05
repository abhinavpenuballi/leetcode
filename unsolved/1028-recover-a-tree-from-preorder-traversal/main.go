package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7"))
	fmt.Println(recoverFromPreorder("1-2--3---4-5--6---7"))
}

func recoverFromPreorder(traversal string) *TreeNode {
	return nil
}
