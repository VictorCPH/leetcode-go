package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum4([]int{3, 1, 2}, 35))
}

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	return helper(nums, target, dp)
}

func helper(nums []int, target int, dp []int) int {
	if target < 0 {
		return 0
	}
	if dp[target] != -1 {
		return dp[target]
	}

	res := 0
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if target == nums[i] {
			res += 1
		} else {
			res += helper(nums, target-nums[i], dp)
		}
	}
	dp[target] = res
	return res
}
