/*
package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
*/
func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	max := 0
	maxIdx := 0
	for i, v := range height {
		if v > max {
			max = v
			maxIdx = i
		}
		leftMax[i] = maxIdx
	}

	max = 0
	maxIdx = len(height) - 1
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
			maxIdx = i
		}
		rightMax[i] = maxIdx
	}

	res := 0
	for i, v := range height {
		if height[leftMax[i]] > v && height[rightMax[i]] > v {
			ht := min(height[leftMax[i]], height[rightMax[i]]) - v
			res += ht
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
