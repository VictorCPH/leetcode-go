package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	str := []byte(s)
	i, j := 0, len(str)-1
	for i < j {
		if !isVowels(str[i]) {
			i++
		} else if !isVowels(str[j]) {
			j--
		} else {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}

func isVowels(p byte) bool {
	for _, v := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		if v == p {
			return true
		}
	}
	return false
}
