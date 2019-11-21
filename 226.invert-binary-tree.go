package main

import "fmt"

func main() {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}
	invertTree := invertTree(tree)
	travel(invertTree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		root.Val,
		invertTree(root.Right),
		invertTree(root.Left),
	}
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	travel(root.Left)
	travel(root.Right)
}
