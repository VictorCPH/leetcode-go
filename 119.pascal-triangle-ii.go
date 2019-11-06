/*
package main

import "fmt"

func main() {
	fmt.Println(getRow(10))
}
*/
func getRow(rowIndex int) []int {
	nums := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		nums[0] = 1
		nums[i] = 1
		for j := i - 1; j > 0; j-- {
			nums[j] = nums[j-1] + nums[j]
		}
	}
	return nums
}
