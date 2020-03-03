package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}

type MyStack struct {
	queue Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{NewQueue()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	l := this.queue.Size()
	last := 0
	for i := 0; i < l; i++ {
		last = this.queue.Peek()
		if i != l-1 {
			this.queue.Push(last)
		}
		this.queue.Pop()
	}
	return last
}

/** Get the top element. */
func (this *MyStack) Top() int {
	l := this.queue.Size()
	last := 0
	for i := 0; i < l; i++ {
		last = this.queue.Peek()
		this.queue.Push(last)
		this.queue.Pop()
	}
	return last
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

type Queue struct {
	s []int
}

func NewQueue() Queue {
	return Queue{[]int{}}
}

func (q *Queue) Push(x int) {
	q.s = append(q.s, x)
}

func (q *Queue) Pop() {
	if len(q.s) > 0 {
		q.s = q.s[1:]
	}
}

func (q *Queue) Peek() int {
	if len(q.s) > 0 {
		return q.s[0]
	}
	return 0
}

func (q *Queue) Size() int {
	return len(q.s)
}

func (q *Queue) Empty() bool {
	return len(q.s) == 0
}
