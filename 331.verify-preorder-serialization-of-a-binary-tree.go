package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("9,#,#,1"))
	fmt.Println(isValidSerialization("1,#"))
	fmt.Println(isValidSerialization("1,#,#,#,#"))
}

func isValidSerialization(preorder string) bool {
	strs := strings.Split(preorder, ",")
	if len(strs) == 0 {
		return false
	}
	stack := []string{strs[0]}
	for i := 1; i < len(strs); i++ {
		if len(stack) == 0 {
			return false
		}
		stack = append(stack, strs[i])
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[0] == "#"
}
