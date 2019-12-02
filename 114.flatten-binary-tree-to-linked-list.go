package main

import "fmt"

func main() {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, nil, &TreeNode{6, nil, nil}},
	}
	flatten(root)
	for ; root != nil; root = root.Right {
		fmt.Println(root.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	pre := &TreeNode{0, nil, nil}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre.Left = nil
		pre.Right = node
		pre = node
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}
