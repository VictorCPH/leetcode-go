package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, &ListNode{3, head}}
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	stepOne, stepTwo := head, head
	for stepTwo != nil && stepTwo.Next != nil {
		stepOne = stepOne.Next
		stepTwo = stepTwo.Next.Next
		if stepOne == stepTwo {
			return true
		}
	}
	return false
}
