package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}

	mid := 0
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] < nums[right] {
		return nums[left]
	}
	return nums[right]
}
