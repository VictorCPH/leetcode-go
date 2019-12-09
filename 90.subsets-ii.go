package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{[]int{}}
	lastL, l := 0, 1
	start := 0
	for i := 0; i < len(nums); i++ {
		lastL = l
		l = len(res)
		if i > 0 && nums[i] == nums[i-1] {
			start = lastL
		} else {
			start = 0
		}
		for j := start; j < l; j++ {
			tmp := make([]int, len(res[j]))
			copy(tmp, res[j])
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
	}
	return res
}
