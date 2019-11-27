package main

import "fmt"

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(rightSideView(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MyNode struct {
	Node  *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*MyNode{&MyNode{root, 0}}
	res := []int{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if len(res) == node.Level {
			res = append(res, node.Node.Val)
		}
		if node.Node.Right != nil {
			queue = append(queue, &MyNode{node.Node.Right, node.Level + 1})
		}
		if node.Node.Left != nil {
			queue = append(queue, &MyNode{node.Node.Left, node.Level + 1})
		}
	}
	return res
}
