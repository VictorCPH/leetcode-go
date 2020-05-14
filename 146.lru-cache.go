package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	size int
	m    map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, head, nil}
	head.Next = tail
	return LRUCache{
		size: capacity,
		m:    make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		tail := this.tail
		tmp := tail.Prev
		node.Next = tail
		tail.Prev = node
		tmp.Next = node
		node.Prev = tmp
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.Val = value

		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		tail := this.tail
		tmp := tail.Prev
		node.Next = tail
		tail.Prev = node
		tmp.Next = node
		node.Prev = tmp
	} else {
		this.m[key] = &Node{key, value, this.tail.Prev, this.tail}
		this.tail.Prev.Next = this.m[key]
		this.tail.Prev = this.m[key]
		if len(this.m) == this.size+1 {
			tmp := this.head.Next
			this.head.Next = tmp.Next
			tmp.Next.Prev = this.head
			delete(this.m, tmp.Key)
		}
	}
}
