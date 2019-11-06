/*
package main

import "fmt"

func main() {
	nums := generate(10)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Printf("%d ", nums[i][j])
		}
		fmt.Println()
	}
}
*/
func generate(numRows int) [][]int {
	nums := [][]int{}
	for i := 0; i < numRows; i++ {
		nums = append(nums, make([]int, i+1))
	}

	for i := 0; i < numRows; i++ {
		nums[i][0] = 1
		nums[i][i] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < len(nums[i])-1; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}

	return nums
}
