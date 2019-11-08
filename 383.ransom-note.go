package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(a, b string) bool {
	m := [26]int{}

	for i := 0; i < len(b); i++ {
		m[b[i]-'a'] += 1
	}
	for i := 0; i < len(a); i++ {
		m[a[i]-'a'] -= 1
		if m[a[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
