package main

import "fmt"

func main() {
	fmt.Println(validAnagram("anagram", "nagaram"))
	fmt.Println(validAnagram("rat", "cat"))
}

func validAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] += 1
		m[t[i]-'a'] -= 1
	}
	for i := 0; i < len(m); i++ {
		if m[i] != 0 {
			return false
		}
	}
	return true
}
