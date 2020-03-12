package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{[]int{0, 1}}))
	fmt.Println(canFinish(2, [][]int{[]int{0, 1}, []int{1, 0}}))
	fmt.Println(canFinish(4, [][]int{[]int{0, 1}, []int{0, 2}}))
}

type Node struct {
	Degree int
	Edge   [][]int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make([]Node, numCourses)
	for _, pre := range prerequisites {
		nodes[pre[1]].Degree += 1
		nodes[pre[0]].Edge = append(nodes[pre[0]].Edge, pre)
	}

	queue := []*Node{}
	for idx, node := range nodes {
		if node.Degree == 0 {
			queue = append(queue, &nodes[idx])
		}
	}
	count := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		count += 1
		for _, edge := range node.Edge {
			nodes[edge[1]].Degree -= 1
			if nodes[edge[1]].Degree == 0 {
				queue = append(queue, &nodes[edge[1]])
			}
		}
	}
	if count < len(nodes) {
		return false
	}
	return true
}
