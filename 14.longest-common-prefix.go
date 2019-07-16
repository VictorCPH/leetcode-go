/*
package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; ; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if len(strs[j]) <= i || len(strs[j+1]) <= i || strs[j][i] != strs[j+1][i] {
				return strs[j][:i]
			}
		}
	}
}
