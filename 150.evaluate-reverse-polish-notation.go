package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	stack := NewStack()
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			second, _ := strconv.Atoi(stack.Pop())
			first, _ := strconv.Atoi(stack.Pop())
			switch tokens[i] {
			case "+":
				stack.Push(strconv.Itoa(first + second))
			case "-":
				stack.Push(strconv.Itoa(first - second))
			case "*":
				stack.Push(strconv.Itoa(first * second))
			case "/":
				stack.Push(strconv.Itoa(first / second))
			}
		} else {
			stack.Push(tokens[i])
		}
	}
	res, _ := strconv.Atoi(stack.Pop())
	return res
}

type Stack struct {
	data []string
}

func NewStack() Stack {
	return Stack{[]string{}}
}

func (s *Stack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() string {
	if len(s.data) > 0 {
		elm := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return elm
	}
	return ""
}
