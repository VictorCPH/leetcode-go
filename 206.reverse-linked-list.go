package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	for list := reverseList(head); list != nil; list = list.Next {
		fmt.Printf("%d ", list.Val)
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := &ListNode{Next: head}
	for head.Next != nil {
		tmp := head.Next.Next
		list.Next, head.Next.Next = head.Next, list.Next
		head.Next = tmp
	}
	return list.Next
}
