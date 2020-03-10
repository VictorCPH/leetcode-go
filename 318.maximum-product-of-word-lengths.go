package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct([]string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}))
}

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	bit := make([]int, len(words))
	for i, word := range words {
		for _, w := range word {
			bit[i] = bit[i] | (0x1 << (w - 'a'))
		}
	}
	maxVal := 0
	for i := 0; i < len(words)-1; i++ {
		if maxVal >= len(words[i])*len(words[i]) {
			break
		}
		for j := i + 1; j < len(words); j++ {
			if len(words[i])*len(words[j]) <= maxVal {
				continue
			} else {
				if bit[i]&bit[j] == 0 {
					maxVal = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return maxVal
}
