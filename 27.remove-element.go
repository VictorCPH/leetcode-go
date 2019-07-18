/*
package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 5}
	l := removeElement(nums, 1)
	for i := 0; i < l; i++ {
		fmt.Println(nums[i])
	}
}
*/
func removeElement(nums []int, val int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
