package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s))
	return helper(s, wordDict, dp, 0)
}

func helper(s string, wordDict []string, dp []int, start int) bool {
	if start == len(s) {
		return true
	}

	if dp[start] == 1 {
		return true
	} else if dp[start] == -1 {
		return false
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s[start:], word) {
			if helper(s, wordDict, dp, start+len(word)) {
				dp[start] = 1
				return true
			}
		}
	}
	dp[start] = -1
	return false
}
