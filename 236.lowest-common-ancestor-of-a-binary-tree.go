package main

import "fmt"

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, nil, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(lowestCommonAncestor(root, &TreeNode{7, nil, nil}, &TreeNode{0, nil, nil}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	helper(root, p, q, &result)
	return result
}

func helper(root, p, q *TreeNode, result **TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, p, q, result)
	right := helper(root.Right, p, q, result)

	val := left + right
	if root.Val == p.Val || root.Val == q.Val {
		val += 1
	}
	if *result == nil && val == 2 {
		*result = root
	}
	return val
}
