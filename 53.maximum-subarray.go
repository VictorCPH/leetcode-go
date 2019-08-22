/*
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
*/

/*
// O(n)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
*/

// O(nlogn) divide and conquer
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxSub(nums, 0, len(nums)-1)
}

func maxSub(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}

	mid := (l + r - 1) / 2
	lMax := maxSub(nums, l, mid)
	rMax := maxSub(nums, mid+1, r)

	lSum, lSumMax := 0, math.MinInt32
	for i := mid; i >= l; i-- {
		lSum += nums[i]
		if lSum > lSumMax {
			lSumMax = lSum
		}
	}

	rSum, rSumMax := 0, math.MinInt32
	for i := mid + 1; i <= r; i++ {
		rSum += nums[i]
		if rSum > rSumMax {
			rSumMax = rSum
		}
	}

	return maxInt(maxInt(lMax, rMax), lSumMax+rSumMax)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
