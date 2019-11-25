package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := BalanceAndDept(root)
	return balanced
}

func BalanceAndDept(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	balancedLeft, deptLeft := BalanceAndDept(root.Left)
	balancedRight, deptRight := BalanceAndDept(root.Right)
	if balancedLeft && balancedRight {
		if deptLeft-deptRight >= -1 && deptLeft-deptRight <= 1 {
			return true, max(deptLeft, deptRight) + 1
		}
	}
	return false, max(deptLeft, deptRight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
