package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	return helper(nums, k, n)
}

func helper(nums []int, k, n int) [][]int {
	if n <= 0 || k <= 0 {
		return [][]int{}
	}
	if k == 1 {
		if len(nums) > 0 && n >= nums[0] && n <= nums[len(nums)-1] {
			return [][]int{[]int{n}}
		} else {
			return [][]int{}
		}
	}

	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		s := helper(nums[i+1:], k-1, n-nums[i])
		for _, v := range s {
			res = append(res, append([]int{nums[i]}, v...))
		}
	}
	return res
}
