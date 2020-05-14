package main

import "fmt"

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
	fmt.Println(minSubArrayLen(3, []int{1, 1}))
}

func minSubArrayLen(s int, nums []int) int {
	start, end := 0, len(nums)
	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		if sum >= s {
			if j-i < end-start {
				start, end = i, j
			}
			for sum >= s {
				sum -= nums[i]
				i += 1
				if sum >= s && j-i < end-start {
					start, end = i, j
				}
			}
		}
	}
	if end == len(nums) {
		return 0
	}
	return end - start + 1
}
