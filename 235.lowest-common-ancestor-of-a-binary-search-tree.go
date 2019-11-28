package main

import "fmt"

func main() {
	root := &TreeNode{5,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		&TreeNode{10, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(lowestCommonAncestor(root, &TreeNode{4, nil, nil}, &TreeNode{1, nil, nil}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		return helper(root, q.Val, p.Val)
	}
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, min, max int) *TreeNode {
	if root.Val >= min && root.Val <= max {
		return root
	}
	if root.Val > min && root.Val > max {
		return helper(root.Left, min, max)
	}
	return helper(root.Right, min, max)
}
