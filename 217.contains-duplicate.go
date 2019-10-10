/*
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	ret := containsDuplicate(nums)
	fmt.Println(ret)
}
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}
