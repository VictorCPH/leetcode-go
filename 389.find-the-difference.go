package main

import "fmt"

func main() {
	fmt.Printf("%c\n", findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	bm := [26]int{}
	for i := 0; i < len(s); i++ {
		bm[s[i]-'a'] -= 1
		bm[t[i]-'a'] += 1
	}
	bm[t[len(t)-1]-'a'] += 1
	for i := 0; i < 26; i++ {
		if bm[i] == 1 {
			return byte(i + 'a')
		}
	}
	return byte('a')
}
