package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("z"))
}

type countAndIndex struct {
	Count int
	Index int
}

func firstUniqChar(s string) int {
	m := make([]countAndIndex, 26)

	for i := 0; i < len(s); i++ {
		m[s[i]-'a'].Count += 1
		m[s[i]-'a'].Index = i
	}

	firstIdx := len(s)
	for _, v := range m {
		if v.Count == 1 {
			firstIdx = min(v.Index, firstIdx)
		}
	}
	if firstIdx == len(s) {
		return -1
	}
	return firstIdx
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
