package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("ab", "aa"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[byte]byte{}
	mp := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			if _, ok2 := mp[t[i]]; !ok2 {
				m[s[i]] = t[i]
				mp[t[i]] = true
			} else {
				return false
			}
		} else {
			if m[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
