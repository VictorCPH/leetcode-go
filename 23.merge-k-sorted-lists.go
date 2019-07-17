/*
package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l3 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l4 := mergeKLists([]*ListNode{l1, l2, l3})
	for ; l4 != nil; l4 = l4.Next {
		fmt.Println(l4.Val)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(lists[i], res)
	}
	return res
}

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
