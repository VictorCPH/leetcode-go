/*
package main

import "fmt"

func main() {
	var s string = "hwwekh"
	length := lengthOfLongestSubstring(s)
	fmt.Println(length)
}
*/
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	res := make(map[byte]int)
	start, end := 0, 0
	for end < len(s) {
		if t, ok := res[s[end]]; ok && t >= start {
			if end-start > maxLen {
				maxLen = end - start
			}
			start = t + 1
		} else {
			res[s[end]] = end
			end += 1
		}
	}
	if end-start > maxLen {
		maxLen = end - start
	}
	return maxLen
}
