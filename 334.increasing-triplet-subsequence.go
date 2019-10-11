package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{5, 4, 3, 1, 1, 5, 1, 1, 7}
	fmt.Println(increasingTriplet(nums))
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums2))
	nums3 := []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums3))
	nums4 := []int{2, 4, -2, -3}
	fmt.Println(increasingTriplet(nums4))
	nums5 := []int{0, 4, 2, 1, 0, -1, -3}
	fmt.Println(increasingTriplet(nums5))
	nums6 := []int{2, 5, 3, 4, 5}
	fmt.Println(increasingTriplet(nums6))
}

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
