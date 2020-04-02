package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{
		[]int{0, 3}, []int{1, 3}, []int{2, 3}, []int{4, 3}, []int{5, 4},
	}
	fmt.Println(findMinHeightTrees(n, edges))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	m := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		m[edges[i][0]] = append(m[edges[i][0]], edges[i][1])
		m[edges[i][1]] = append(m[edges[i][1]], edges[i][0])
	}

	deg := make([]int, n)
	queue := []int{}
	for i := 0; i < n; i++ {
		deg[i] = len(m[i])
		if deg[i] == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		layerSize := len(queue)
		n -= layerSize
		for i := 0; i < layerSize; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, v := range m[node] {
				deg[v] -= 1
				if deg[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return queue
}
