package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(sumNumbers(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	for _, v := range sumNumberHelper(root) {
		val, _ := strconv.Atoi(v)
		sum += val
	}
	return sum
}

func sumNumberHelper(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
	} else {
		for _, v := range sumNumberHelper(root.Left) {
			res = append(res, strconv.Itoa(root.Val)+v)
		}
		for _, v := range sumNumberHelper(root.Right) {
			res = append(res, strconv.Itoa(root.Val)+v)
		}
	}
	return res
}
