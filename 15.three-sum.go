/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 2, 2, 3, 0, 1, 2, 3, -1, -4, 2}))
}
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i < 1 || nums[i] != nums[i-1] {
			tmp := twoSum(nums[i+1:], 0-nums[i])
			for _, v := range tmp {
				res = append(res, append(v, nums[i]))
			}
		}
	}
	return res
}

func twoSum(nums []int, sum int) [][]int {
	m := map[int]int{}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			if v != -1 {
				res = append(res, []int{nums[v], nums[i]})
				m[nums[i]] = -1
			}
		} else {
			m[sum-nums[i]] = i
		}
	}
	return res
}
