package main

import "fmt"

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
	nums1 := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums1))
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return nums[len(nums)-1]
}
