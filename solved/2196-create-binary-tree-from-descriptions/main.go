package main

func main() {
	descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	createBinaryTree(descriptions)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	parents, children := map[int]struct{}{}, map[int]struct{}{}
	descMap := map[int]map[int]int{}

	for _, desc := range descriptions {
		parents[desc[0]] = struct{}{}
		children[desc[1]] = struct{}{}

		if _, ok := descMap[desc[0]]; !ok {
			descMap[desc[0]] = make(map[int]int)
		}
		descMap[desc[0]][desc[2]] = desc[1]
	}

	root := -1

	for parent := range parents {
		if _, ok := children[parent]; !ok {
			root = parent
		}
	}

	return create(descMap, root)
}

func create(descMap map[int]map[int]int, num int) *TreeNode {
	node := &TreeNode{
		Val: num,
	}

	if _, ok := descMap[num][1]; ok {
		node.Left = create(descMap, descMap[num][1])
	}

	if _, ok := descMap[num][0]; ok {
		node.Right = create(descMap, descMap[num][0])
	}

	return node
}
