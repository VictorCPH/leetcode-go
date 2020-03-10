package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 2, 1}))
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
