package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	fmt.Println(obj.Top())
}

type MinStack struct {
	s   []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	if x < this.min {
		this.min = x
	}
	this.s = append(this.s, x)
}

func (this *MinStack) Pop() {
	if len(this.s) > 0 {
		if this.Top() == this.min {
			this.min = math.MaxInt32
			for _, v := range this.s[:len(this.s)-1] {
				if v < this.min {
					this.min = v
				}
			}
		}
		this.s = this.s[:len(this.s)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.s) == 0 {
		return 0
	}
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
