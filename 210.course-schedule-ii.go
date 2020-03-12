package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{[]int{0, 1}}))
	fmt.Println(findOrder(2, [][]int{[]int{0, 1}, []int{1, 0}}))
	fmt.Println(findOrder(4, [][]int{[]int{0, 1}, []int{0, 2}}))
}

type Node struct {
	Idx    int
	Degree int
	Edge   []int
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	nodes := make([]Node, numCourses)
	for idx, _ := range nodes {
		nodes[idx].Idx = idx
	}
	for _, pre := range prerequisites {
		nodes[pre[0]].Degree += 1
		nodes[pre[1]].Edge = append(nodes[pre[1]].Edge, pre[0])
	}

	queue := []*Node{}
	for idx, node := range nodes {
		if node.Degree == 0 {
			queue = append(queue, &nodes[idx])
		}
	}

	res := []int{}
	count := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		count += 1
		res = append(res, node.Idx)
		for _, edge := range node.Edge {
			nodes[edge].Degree -= 1
			if nodes[edge].Degree == 0 {
				queue = append(queue, &nodes[edge])
			}
		}
	}
	if count < len(nodes) {
		return []int{}
	}
	return res
}
