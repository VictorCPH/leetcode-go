/*
package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}
*/
func generateParenthesis(n int) []string {
	res := []string{}
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, s string, open int, closed int, n int) {
	if len(s) == n*2 {
		*res = append(*res, s)
		return
	}

	if open < n {
		backtrack(res, s+"(", open+1, closed, n)
	}
	if closed < open {
		backtrack(res, s+")", open, closed+1, n)
	}
}
