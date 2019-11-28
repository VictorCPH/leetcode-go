package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{1,
		&TreeNode{1, nil, nil},
		nil,
	}
	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

func helper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}
	if root.Val > high || root.Val < low {
		return false
	}
	return helper(root.Left, low, root.Val-1) && helper(root.Right, root.Val+1, high)
}
