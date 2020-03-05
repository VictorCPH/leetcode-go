package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	var stack Stack
	for _, val := range s {
		if val == ']' {
			str := ""
			for c := stack.Pop(); c != '['; c = stack.Pop() {
				str = string(c) + str
			}
			digist, d := "", rune('0')
			for d = stack.Pop(); d >= '0' && d <= '9'; d = stack.Pop() {
				digist = string(d) + digist
			}
			if d != '#' {
				stack.Push(d)
			}
			num, _ := strconv.Atoi(digist)
			for i := 0; i < num; i++ {
				for _, c := range str {
					stack.Push(c)
				}
			}
		} else {
			stack.Push(val)
		}
	}
	return stack.String()
}

type Stack struct {
	data []rune
}

func (s *Stack) Push(x rune) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() rune {
	if !s.Empty() {
		out := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return out
	}
	return rune('#')
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) String() string {
	str := ""
	for _, c := range s.data {
		str += string(c)
	}
	return str
}
