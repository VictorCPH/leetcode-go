package main

import "fmt"

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(rob(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := make(map[*TreeNode]int)
	return helper(root, result)
}

func helper(root *TreeNode, result map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := result[root]; ok {
		return v
	}
	sum := 0
	if root.Left != nil {
		sum += helper(root.Left.Left, result) + helper(root.Left.Right, result)
	}
	if root.Right != nil {
		sum += helper(root.Right.Left, result) + helper(root.Right.Right, result)
	}
	sum = max(root.Val+sum, helper(root.Left, result)+helper(root.Right, result))
	result[root] = sum
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
