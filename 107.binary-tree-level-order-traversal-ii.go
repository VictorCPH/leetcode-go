package main

import "fmt"

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(levelOrderBottom(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeAndLevel struct {
	Node  *TreeNode
	Level int
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*NodeAndLevel{&NodeAndLevel{root, 0}}
	tmp := map[int][]int{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		tmp[node.Level] = append(tmp[node.Level], node.Node.Val)
		if node.Node.Left != nil {
			queue = append(queue, &NodeAndLevel{node.Node.Left, node.Level + 1})
		}
		if node.Node.Right != nil {
			queue = append(queue, &NodeAndLevel{node.Node.Right, node.Level + 1})
		}
	}
	res := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}
