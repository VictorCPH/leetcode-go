package main

import "fmt"

func main() {
	root := &TreeNode{6,
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{10, &TreeNode{7, nil, nil}, &TreeNode{11, nil, nil}},
	}
	iter := Constructor(root)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	list []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	if root == nil {
		return BSTIterator{[]*TreeNode{}}
	}
	cur := root
	list := []*TreeNode{}
	for cur == root || cur != nil {
		list = append(list, cur)
		cur = cur.Left
	}
	return BSTIterator{list}
}

func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return -1
	}
	next := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	cur := next.Right
	for cur != nil {
		this.list = append(this.list, cur)
		cur = cur.Left
	}
	return next.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.list) != 0
}
