package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := generateTrees(5)
	for _, root := range res {
		fmt.Println(serialize(root))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	nodes := []int{}
	for i := 1; i <= n; i++ {
		nodes = append(nodes, i)
	}
	return helper(nodes)
}

func helper(nodes []int) []*TreeNode {
	if len(nodes) == 0 {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	for i := 0; i < len(nodes); i++ {
		left := helper(nodes[:i])
		right := helper(nodes[i+1:])
		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{nodes[i], l, r})
			}
		}
	}
	return res
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
