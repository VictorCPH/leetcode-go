/*
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
		fmt.Print(" ")
	}
}
*/

func rotate(nums []int, k int) {
	if len(nums) == 0 || k < 0 {
		return
	}

	count := 0
	l := len(nums)
	for start := 0; count < l; start++ {
		current := start
		last := nums[current]
		for {
			idx := (current + k) % l
			val := nums[idx]
			nums[idx] = last
			count++
			current = idx
			last = val
			if current == start {
				break
			}
		}
	}
}
