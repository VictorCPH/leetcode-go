package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1000}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, v := range nums {
		max = maxInt(max, v)
	}

	bucket := make([][]int, 10)
	for base := 10; base/10 <= max; base *= 10 {
		for _, v := range nums {
			idx := (v % base) / (base / 10)
			bucket[idx] = append(bucket[idx], v)
		}
		idx := 0
		for i := 0; i < 10; i++ {
			for _, v := range bucket[i] {
				nums[idx] = v
				idx += 1
			}
			bucket[i] = []int{}
		}
	}
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		maxGap = maxInt(maxGap, nums[i]-nums[i-1])
	}
	return maxGap
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
