/*
package main

import "fmt"

func main() {
	str := "abbcbbcba"
	subStr := longestPalindrome(str)
	fmt.Println(subStr)
}
*/

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i, _ := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
