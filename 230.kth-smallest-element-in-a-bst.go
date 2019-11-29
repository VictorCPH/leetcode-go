package main

import "fmt"

func main() {
	root := &TreeNode{6,
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{10, &TreeNode{7, nil, nil}, &TreeNode{11, nil, nil}},
	}
	fmt.Println(kthSmallest(root, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	cur := root
	stack := []*TreeNode{root}
	for cur == root || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k -= 1
		if k == 0 {
			return node.Val
		}
		cur = node.Right
	}
	return -1
}
