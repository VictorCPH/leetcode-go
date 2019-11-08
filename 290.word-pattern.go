package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	m := map[byte]string{}
	mp := map[string]bool{}
	for i := 0; i < len(pattern); i++ {
		if _, ok := m[pattern[i]]; !ok {
			if _, ok2 := mp[strs[i]]; !ok2 {
				m[pattern[i]] = strs[i]
				mp[strs[i]] = true
			} else {
				return false
			}
		} else {
			if m[pattern[i]] != strs[i] {
				return false
			}
		}
	}
	return true
}
