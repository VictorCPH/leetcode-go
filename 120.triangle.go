package main

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	minVal := math.MaxInt32
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
			if i == len(triangle)-1 {
				minVal = min(triangle[i][j], minVal)
			}
		}
	}
	if minVal == math.MaxInt32 {
		return triangle[0][0]
	}
	return minVal
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
