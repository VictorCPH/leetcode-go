package main

import "fmt"

func main() {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{6, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}},
	}
	fmt.Println(pathSum(tree, 9))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			res = append(res, []int{root.Val})
		}
	} else {
		for _, v := range pathSum(root.Left, sum-root.Val) {
			res = append(res, append([]int{root.Val}, v...))
		}
		for _, v := range pathSum(root.Right, sum-root.Val) {
			res = append(res, append([]int{root.Val}, v...))
		}
	}
	return res
}
