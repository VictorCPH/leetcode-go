package main

import "fmt"

func main() {
	list := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	root := sortedListToBST(list)
	levelTravel(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var perMid *ListNode
	mid, doubleMid := head, head
	for doubleMid.Next != nil && doubleMid.Next.Next != nil {
		perMid = mid
		mid = mid.Next
		doubleMid = doubleMid.Next.Next
	}

	root := &TreeNode{mid.Val, nil, nil}
	if perMid != nil {
		perMid.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(mid.Next)
	return root
}

func levelTravel(root *TreeNode) {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
