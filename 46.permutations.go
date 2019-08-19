/*
package main

import "fmt"

func main() {
	res := permute([]int{1, 2, 3, 4})
	for _, val1 := range res {
		for _, val2 := range val1 {
			fmt.Printf("%d ", val2)
		}
		fmt.Println()
	}
}
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	arr := permute(nums[1:])
	res := [][]int{}
	for _, val := range arr {
		res = append(res, append([]int{nums[0]}, val...))
		for i := 0; i < len(val); i++ {
			tmp := make([]int, len(val)+1)
			copy(tmp, val[:i+1])
			tmp[i+1] = nums[0]
			copy(tmp[i+2:], val[i+1:])
			res = append(res, tmp)
		}
	}
	return res
}
