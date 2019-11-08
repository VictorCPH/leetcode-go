package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	if i < 0 {
		return 0
	}

	for j := i - 1; j >= 0; j-- {
		if s[j] == ' ' {
			return i - j
		}
	}
	return i + 1
}
