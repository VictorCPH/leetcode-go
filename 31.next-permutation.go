/*
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 3, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
*/
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	minVal := math.MaxInt32
	minIdx := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i] && nums[j] <= minVal {
					minVal = nums[j]
					minIdx = j
				}
			}
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
			m, n := i+1, len(nums)-1
			for m < n {
				nums[m], nums[n] = nums[n], nums[m]
				m++
				n--
			}
			return
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
