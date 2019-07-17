/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{0, 0, 0, 0, 0}, 0))
}
*/

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j-1 >= i+1 && nums[j] == nums[j-1] {
				continue
			}

			m, n := j+1, len(nums)-1
			for m < n {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if sum < target {
					m++
				} else if sum > target {
					n--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
					m++
					for ; m < n && nums[m] == nums[m-1]; m++ {
					}
				}
			}
		}
	}
	return res
}
