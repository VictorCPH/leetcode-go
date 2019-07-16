/*
package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
*/

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	i, j := 1, len(height)
	maxArea := 0
	for i < j {
		if (j-i)*min(height[i-1], height[j-1]) > maxArea {
			maxArea = (j - i) * min(height[i-1], height[j-1])
		}
		if height[i-1] > height[j-1] {
			j -= 1
		} else {
			i += 1
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
