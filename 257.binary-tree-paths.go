package main

import (
	"fmt"
	"strconv"
)

func main() {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}},
	}
	fmt.Println(binaryTreePaths(tree))
	fmt.Println(binaryTreePaths(&TreeNode{1, nil, nil}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	valStr := strconv.Itoa(root.Val)
	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)
	for _, v := range leftPaths {
		res = append(res, valStr+"->"+v)
	}
	for _, v := range rightPaths {
		res = append(res, valStr+"->"+v)
	}
	if len(leftPaths) == 0 && len(rightPaths) == 0 {
		res = append(res, valStr)
	}
	return res
}
