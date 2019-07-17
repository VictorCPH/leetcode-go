/*
package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list := removeNthFromEnd(head, 5)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}

	if size == n {
		return head.Next
	}

	cur := head
	for i := 1; i < size-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}
