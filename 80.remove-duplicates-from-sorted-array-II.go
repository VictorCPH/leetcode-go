/*
package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	length := removeDuplicates(nums)
	for i := 0; i < length; i++ {
		fmt.Println(nums[i])
	}
}
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	idx := 2
	last := -1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-2] || last == i-2 {
			if nums[i] != nums[idx] {
				last = idx
			}
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
