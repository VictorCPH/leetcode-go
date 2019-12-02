package main

import (
	"fmt"
	"strconv"
)

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(inorder, postorder)
	fmt.Println(serialize(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	idx := find(inorder, val)
	leftInorder := inorder[:idx]
	RightInorder := inorder[idx+1:]
	leftPostorder := postorder[:len(leftInorder)]
	RightPostorder := postorder[len(leftInorder) : len(postorder)-1]
	return &TreeNode{
		val,
		buildTree(leftInorder, leftPostorder),
		buildTree(RightInorder, RightPostorder),
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
