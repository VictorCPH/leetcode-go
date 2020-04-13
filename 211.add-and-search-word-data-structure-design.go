package main

import "fmt"

func main() {
	s := Constructor()
	s.AddWord("a")
	s.AddWord("ab")
	fmt.Println(s.Search("a"))
	fmt.Println(s.Search("a."))
	fmt.Println(s.Search("ab"))
	fmt.Println(s.Search(".a"))
	fmt.Println(s.Search(".b"))
	fmt.Println(s.Search("ab."))
	fmt.Println(s.Search("."))
	fmt.Println(s.Search(".."))
}

type WordDictionary struct {
	root *Node
}

type Node struct {
	Leaf     bool
	Adjacent map[byte]*Node
}

func Constructor() WordDictionary {
	return WordDictionary{&Node{false, make(map[byte]*Node)}}
}

func (s *WordDictionary) AddWord(w string) {
	node := s.root
	leaf := false
	for i := 0; i < len(w); i++ {
		if _, ok := node.Adjacent[w[i]]; !ok {
			if i == len(w)-1 {
				leaf = true
			}
			node.Adjacent[w[i]] = &Node{leaf, make(map[byte]*Node)}
		}
		node = node.Adjacent[w[i]]
	}
}

func (s *WordDictionary) Search(w string) bool {
	return dfs(s.root, w)
}

func dfs(root *Node, w string) bool {
	if root == nil {
		return false
	}
	if root.Leaf && len(w) == 0 {
		return true
	}
	if !root.Leaf && len(w) == 0 {
		return false
	}
	b := w[0]
	if b != '.' {
		return dfs(root.Adjacent[b], w[1:])
	} else {
		for _, v := range root.Adjacent {
			if dfs(v, w[1:]) {
				return true
			}
		}
		return false
	}
}
