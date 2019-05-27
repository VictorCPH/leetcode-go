/*
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 := &ListNode{5, &ListNode{6, nil}}
	l3 := addTwoNumbers(l1, l2)
	for ; l3 != nil; l3 = l3.Next {
		fmt.Println(l3.Val)
	}
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	last := res
	m, n, sum := 0, 0, 0
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + m
		m = sum / 10
		n = sum % 10
		cur := &ListNode{Val: n, Next: nil}
		last.Next = cur
		last = cur
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		l1 = l2
	}
	for l1 != nil {
		sum = l1.Val + m
		m = sum / 10
		n = sum % 10
		cur := &ListNode{Val: n, Next: nil}
		last.Next = cur
		last = cur
		l1 = l1.Next
	}
	if m != 0 {
		last.Next = &ListNode{Val: m, Next: nil}
	}
	return res.Next
}
