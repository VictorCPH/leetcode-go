package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalSquare([][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		[]byte{'0'},
	}))
	fmt.Println(maximalSquare([][]byte{}))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxVal := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			maxVal = 1
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == '1' {
			maxVal = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				matrix[i][j] += min(matrix[i][j-1], min(matrix[i-1][j-1], matrix[i-1][j])) - '0'
				maxVal = max(maxVal, int(matrix[i][j]-'0'))
			}
		}
	}
	return maxVal * maxVal
}

func min(a, b byte) byte {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
