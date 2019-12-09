package main

import "fmt"

func main() {
	fmt.Println(combine(4, 3))
}

func combine(n, k int) [][]int {
	if n <= 0 || k <= 0 || n < k {
		return [][]int{}
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return helper(nums, k)
}

func helper(nums []int, k int) [][]int {
	res := [][]int{}
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			res = append(res, []int{nums[i]})
		}
		return res
	}
	for i := 0; i < len(nums)-1; i++ {
		s := helper(nums[i+1:], k-1)
		for _, v := range s {
			res = append(res, append([]int{nums[i]}, v...))
		}
	}
	return res
}
