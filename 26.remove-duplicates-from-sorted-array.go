/*
package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 5}
	l := removeDuplicates(nums)
	for i := 0; i < l; i++ {
		fmt.Println(nums[i])
	}
}
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
