package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 1, 5, 2, 100, 330}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	maxCount := 1
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]-1] {
			count := 1
			curVal := nums[i]
			for m[curVal+1] {
				count += 1
				curVal += 1
			}
			maxCount = maxInt(maxCount, count)
		}
	}
	return maxCount
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
