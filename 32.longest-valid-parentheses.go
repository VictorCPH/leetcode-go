/*
package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")())(()))"))
	fmt.Println(longestValidParentheses("()(((()))"))
	fmt.Println(longestValidParentheses("()(())"))
}
*/
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxCount := 0
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			if maxCount < i-stack[len(stack)-1] {
				maxCount = i - stack[len(stack)-1]
			}
		}
	}
	return maxCount
}
