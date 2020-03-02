package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
	fmt.Println(ladderLength(beginWord, "xxx", wordList))
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))
}

type node struct {
	Word  string
	Level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	list := []node{node{beginWord, 1}}
	maps := map[string]bool{beginWord: true}

	for len(list) != 0 {
		first := list[0]
		list = list[1:]
		for _, w := range wordList {
			count := 0
			for i, v := range w {
				if first.Word[i] != byte(v) {
					count += 1
				}
			}
			if count <= 1 {
				if w == endWord {
					return first.Level + 1
				}
				if !maps[w] {
					maps[w] = true
					list = append(list, node{w, first.Level + 1})
				}
			}
		}
	}
	return 0
}
