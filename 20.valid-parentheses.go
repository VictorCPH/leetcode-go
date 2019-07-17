/*
package main

import "fmt"

func main() {
	fmt.Println(isValid("{{{{[][][[]]}}}}"))
}
*/
func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}
	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			stack = append(stack, v)
		} else if v == '}' || v == ')' || v == ']' {
			if len(stack) > 0 && m[v] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
