package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

type Trie struct {
	IsEnd bool
	Nodes map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, make(map[byte]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if _, ok := cur.Nodes[word[i]]; !ok {
			cur.Nodes[word[i]] = &Trie{false, make(map[byte]*Trie)}
		}
		cur = cur.Nodes[word[i]]
		if i == len(word)-1 {
			cur.IsEnd = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		if _, ok := root.Nodes[word[i]]; ok {
			root = root.Nodes[word[i]]
		} else {
			return false
		}
	}
	return root.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := root.Nodes[prefix[i]]; ok {
			root = root.Nodes[prefix[i]]
		} else {
			return false
		}
	}
	return true
}
