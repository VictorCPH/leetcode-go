package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}

func majorityElement(nums []int) []int {
	m := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	for k, v := range m {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}
