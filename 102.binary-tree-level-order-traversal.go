package main

import "fmt"

func main() {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(levelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	list := []*Element{&Element{root, 0}}
	res := [][]int{}
	for len(list) != 0 {
		node := list[0].Node
		level := list[0].Level
		list = list[1:]
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			list = append(list, &Element{node.Left, level + 1})
		}
		if node.Right != nil {
			list = append(list, &Element{node.Right, level + 1})
		}
	}
	return res
}
