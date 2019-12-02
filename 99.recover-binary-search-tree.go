package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil},
	}
	fmt.Println(serialize(root))
	recoverTree(root)
	fmt.Println(serialize(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre    *TreeNode
	first  *TreeNode
	second *TreeNode
)

func recoverTree(root *TreeNode) {
	helper(root)
	first.Val, second.Val = second.Val, first.Val
}

func helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		if pre != nil && pre.Val > root.Val {
			if first == nil {
				first = pre
				second = root
			} else {
				second = root
			}
		}
		pre = root
		helper(root.Right)
	}
}

func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := []*TreeNode{root}
	tmp := []*TreeNode{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		tmp = append(tmp, node)
		for node == nil && len(queue) != 0 {
			node = queue[0]
			queue = queue[1:]
			tmp = append(tmp, node)
		}
		if node == nil {
			break
		}

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	res := ""
	for _, node := range tmp {
		if node == nil {
			res += "null,"
		} else {

			res += strconv.Itoa(node.Val) + ","
		}
	}
	return res[:len(res)-1]
}
