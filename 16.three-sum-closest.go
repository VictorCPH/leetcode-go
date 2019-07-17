/*
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
*/
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	closest := math.MaxInt32
	res := 0
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum-target == 0 {
				return target
			} else if sum-target > 0 {
				k--
				if sum-target < closest {
					closest = sum - target
					res = sum
				}
			} else {
				j++
				if target-sum < closest {
					closest = target - sum
					res = sum
				}
			}
		}
	}
	return res
}
