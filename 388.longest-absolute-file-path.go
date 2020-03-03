package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Println(lengthLongestPath("a"))
	fmt.Println(lengthLongestPath("dir\n    file.txt"))
	fmt.Println(lengthLongestPath("dir\n        file.txt"))
	fmt.Println(lengthLongestPath("a\n\tb.txt\na2\n\tb2.txt"))
}

func lengthLongestPath(input string) int {
	stack := NewStack()
	input = strings.Replace(input, "\n    ", "\n\t", -1)
	paths := strings.Split(input, "\n")
	res := 0
	for _, path := range paths {
		idx := strings.LastIndex(path, "\t")
		if stack.Empty() || idx > stack.Top().Level {
			stack.Push(Node{path[idx+1:], idx})
		} else {
			for !stack.Empty() && idx <= stack.Top().Level {
				stack.Pop()
			}
			stack.Push(Node{path[idx+1:], idx})
		}
		if strings.Contains(path[idx+1:], ".") {
			l := stack.Length()
			if l > res {
				res = l
			}
		}
	}
	return res
}

type Node struct {
	Val   string
	Level int
}
type Stack struct {
	data []Node
}

func NewStack() Stack {
	return Stack{[]Node{}}
}

func (s *Stack) Push(x Node) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() Node {
	if len(s.data) > 0 {
		elm := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return elm
	}
	return Node{}
}

func (s *Stack) Top() Node {
	if len(s.data) == 0 {
		return Node{}
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Length() int {
	l := 0
	for _, v := range s.data {
		l += len(v.Val) + 1
	}
	return l - 1
}
