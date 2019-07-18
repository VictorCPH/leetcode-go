/*
package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	i, j := 0, 0
	for ; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] {
			j++
		} else {
			i = i - j
			j = 0
		}
	}
	if j != len(needle) {
		return -1
	}
	return i - j
}
