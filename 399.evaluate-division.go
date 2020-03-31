package main

import "fmt"

func main() {
	equations := [][]string{
		[]string{"a", "b"},
		[]string{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		[]string{"a", "c"},
		[]string{"b", "a"},
		[]string{"a", "e"},
		[]string{"a", "a"},
		[]string{"x", "x"},
	}
	fmt.Println(calcEquation(equations, values, queries))
}

type Pair struct {
	Str string
	Val float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string][]Pair)
	for i := 0; i < len(equations); i++ {
		m[equations[i][0]] = append(m[equations[i][0]], Pair{equations[i][1], values[i]})
		m[equations[i][1]] = append(m[equations[i][1]], Pair{equations[i][0], 1 / values[i]})
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = dfs(queries[i][0], queries[i][1], m)
	}
	return res
}

type Node struct {
	Str  string
	Next []Pair
	Val  float64
}

func dfs(s1, s2 string, m map[string][]Pair) float64 {
	if _, ok := m[s1]; !ok {
		return -1.0
	}
	if s1 == s2 {
		return 1.0
	}

	queue := []Node{Node{s1, m[s1], 1.0}}
	travel := make(map[string]bool)
	for len(queue) != 0 {
		f := queue[0]
		queue = queue[1:]
		travel[f.Str] = true
		for _, v := range f.Next {
			if v.Str == s2 {
				return v.Val * f.Val
			}
			if travel[v.Str] == false {
				queue = append(queue, Node{v.Str, m[v.Str], f.Val * v.Val})
			}
		}
	}
	return -1.0
}
