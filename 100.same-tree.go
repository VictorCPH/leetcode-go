package main

import "fmt"

func main() {
	tree1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}
	tree2 := &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(isSameTree(tree1, tree2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}
