package main

import (
	"fmt"
	"strconv"
)

func main() {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	root := buildTree(preorder, inorder)
	fmt.Println(serialize(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	idx := find(inorder, val)
	leftInorder := inorder[:idx]
	RightInorder := inorder[idx+1:]
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	RightPreorder := preorder[1+len(leftInorder):]
	return &TreeNode{
		val,
		buildTree(leftPreorder, leftInorder),
		buildTree(RightPreorder, RightInorder),
	}
}

func find(s []int, t int) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := []*TreeNode{root}
	tmp := []*TreeNode{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		tmp = append(tmp, node)
		for node == nil && len(queue) != 0 {
			node = queue[0]
			queue = queue[1:]
			tmp = append(tmp, node)
		}
		if node == nil {
			break
		}

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	res := ""
	for _, node := range tmp {
		if node == nil {
			res += "null,"
		} else {

			res += strconv.Itoa(node.Val) + ","
		}
	}
	return res[:len(res)-1]
}
