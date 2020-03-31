package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s := Constructor(nums)
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}

type Solution struct {
	origin []int
	nums   []int
}

func Constructor(nums []int) Solution {
	origin := make([]int, len(nums))
	copy(origin, nums)
	return Solution{nums, origin}
}

func (this *Solution) Reset() []int {
	return this.origin
}

func (this *Solution) Shuffle() []int {
	lastIdx := len(this.nums)
	for i := 0; i < len(this.nums); i++ {
		r := rand.Intn(lastIdx)
		tmp := this.nums[r]
		lastIdx -= 1
		this.nums[r] = this.nums[lastIdx]
		this.nums[lastIdx] = tmp
	}
	return this.nums
}
