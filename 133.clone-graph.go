package main

import "fmt"

func main() {
	list := []*Node{&Node{1, []*Node{}}, &Node{2, []*Node{}}, &Node{3, []*Node{}}, &Node{4, []*Node{}}}
	list[0].Neighbors = []*Node{list[1], list[3]}
	list[1].Neighbors = []*Node{list[0], list[2]}
	list[2].Neighbors = []*Node{list[1], list[3]}
	list[3].Neighbors = []*Node{list[0], list[2]}
	clone := cloneGraph(list[0])
	travel(clone)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	arr := make([]*Node, 101)
	q1 := []*Node{node}
	for len(q1) != 0 {
		first := q1[0]
		q1 = q1[1:]
		if arr[first.Val] == nil {
			arr[first.Val] = &Node{first.Val, []*Node{}}
		}
		for _, v := range first.Neighbors {
			if arr[v.Val] == nil {
				q1 = append(q1, v)
			}
		}
	}

	q2 := []*Node{node}
	travel := make([]bool, 101)
	travel[1] = true
	for len(q2) != 0 {
		first := q2[0]
		q2 = q2[1:]
		for _, v := range first.Neighbors {
			arr[first.Val].Neighbors = append(arr[first.Val].Neighbors, arr[v.Val])
			if travel[v.Val] == false {
				q2 = append(q2, v)
				travel[v.Val] = true
			}
		}
	}
	return arr[1]
}

func travel(node *Node) {
	if node == nil {
		return
	}
	travel := make([]bool, 101)
	travel[1] = true
	q := []*Node{node}
	for len(q) != 0 {
		first := q[0]
		q = q[1:]
		for _, v := range first.Neighbors {
			fmt.Printf("%d ", v.Val)
			if travel[v.Val] == false {
				q = append(q, v)
				travel[v.Val] = true
			}
		}
		fmt.Println()
	}
}
