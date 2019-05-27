/*
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	fmt.Println(res[0], res[1])
}
*/

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)
	for i, v := range nums {
		if t, ok := res[target-v]; ok {
			return []int{t, i}
		} else {
			res[v] = i
		}
	}
	return nil
}
