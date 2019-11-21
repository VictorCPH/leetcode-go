package main

import "fmt"

func main() {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, nil, &TreeNode{5, nil, nil}}}
	fmt.Println(inorderTraversal(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}
