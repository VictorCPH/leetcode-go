package main

import "fmt"

func main() {
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords(" "))
}

func reverseWords(s string) string {
	spaceIdx := []int{}
	spaceIdx = append(spaceIdx, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaceIdx = append(spaceIdx, i)
		}
	}
	spaceIdx = append(spaceIdx, len(s))

	res := ""
	for i := len(spaceIdx) - 2; i >= 0; i-- {
		if spaceIdx[i]+1 != spaceIdx[i+1] {
			res = res + s[spaceIdx[i]+1:spaceIdx[i+1]] + " "
		}
	}
	if len(res) == 0 {
		return ""
	}
	return res[:len(res)-1]
}
