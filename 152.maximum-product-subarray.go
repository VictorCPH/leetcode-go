package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	max, min, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax := max
		max = maxInt(maxInt(max*nums[i], min*nums[i]), nums[i])
		min = minInt(minInt(tmpMax*nums[i], min*nums[i]), nums[i])
		res = maxInt(max, res)
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
