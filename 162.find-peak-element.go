package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
	nums1 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums1))
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
