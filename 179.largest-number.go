package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{101, 23213, 23132, 42424, 112233}))
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{1}))
	fmt.Println(largestNumber([]int{}))
}

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	strs := []string{}
	for _, v := range nums {
		strs = append(strs, strconv.Itoa(v))
	}

	sort.Slice(strs, func(i int, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	res := ""
	for _, v := range strs {
		res += v
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
