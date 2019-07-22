/*
package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{8, 8, 8, 8, 8, 8}, 8))
}
*/
func searchRange(nums []int, target int) []int {
	leftIdx := search(nums, target, true)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return []int{-1, -1}
	}
	rightIdx := search(nums, target, false) - 1
	return []int{leftIdx, rightIdx}
}

func search(nums []int, target int, left bool) int {
	low, high := 0, len(nums)

	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target || (left && nums[mid] == target) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
