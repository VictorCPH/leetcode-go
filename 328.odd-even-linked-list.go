package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	list := oddEvenList(head)
	for ; list != nil; list = list.Next {
		fmt.Printf("%d ", list.Val)
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	cur := head
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
		}
		if cur.Next != nil {
			cur = cur.Next
		}
	}
	cur.Next = even
	return odd
}
