package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{5,
		&TreeNode{3, &TreeNode{7, nil, nil}, &TreeNode{0, &TreeNode{0, nil, nil}, nil}},
		&TreeNode{2, &TreeNode{6, nil, nil}, nil},
	}
	fmt.Println(serialize(root))
	fmt.Println(serialize(deserialize(serialize(root))))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")
	num, _ := strconv.Atoi(strs[0])
	root := &TreeNode{num, nil, nil}
	curLevel := []*TreeNode{root}

	for i := 1; i < len(strs); {
		curLen := len(curLevel)
		for j := 0; j < curLen; j++ {
			node := curLevel[j]
			if strs[i] != "null" {
				tmp, _ := strconv.Atoi(strs[i])
				node.Left = &TreeNode{tmp, nil, nil}
				curLevel = append(curLevel, node.Left)
			}
			i += 1

			if strs[i] != "null" {
				tmp, _ := strconv.Atoi(strs[i])
				node.Right = &TreeNode{tmp, nil, nil}
				curLevel = append(curLevel, node.Right)
			}
			i += 1
		}
		curLevel = curLevel[curLen:]
	}
	return root
}
