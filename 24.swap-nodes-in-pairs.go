/*
package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	l2 := swapPairs(l1)
	for ; l2 != nil; l2 = l2.Next {
		fmt.Println(l2.Val)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	first := pre
	for ; pre.Next != nil && pre.Next.Next != nil; pre = pre.Next.Next {
		cur := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur
	}
	return first.Next
}
