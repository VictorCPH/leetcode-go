/*
package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := mergeTwoLists(l1, l2)
	for ; l3 != nil; l3 = l3.Next {
		fmt.Println(l3.Val)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		l1 = l2
	}
	for ; l1 != nil; l1 = l1.Next {
		cur.Next = l1
		cur = cur.Next
	}

	return head.Next
}
