package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums1))
	nums2 := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums2))
	fmt.Println(summaryRanges([]int{}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := []string{}
	start := 0
	str := ""
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if i-1 == start {
				str = fmt.Sprintf("%d", nums[start])
			} else {
				str = fmt.Sprintf("%d->%d", nums[start], nums[i-1])
			}
			res = append(res, str)
			start = i
		}
	}
	if start == len(nums)-1 {
		str = fmt.Sprintf("%d", nums[start])
	} else {
		str = fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1])
	}
	res = append(res, str)
	return res
}
