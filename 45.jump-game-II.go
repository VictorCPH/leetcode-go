package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	length := len(nums)
	memo := make([]int, length)

	lastReached := 0
	for i := 0; i < length-1; i++ {
		for j := lastReached + 1; j < length && j <= i+nums[i]; j++ {
			lastReached = j
			memo[lastReached] = memo[i] + 1
		}
	}
	return memo[length-1]
}
