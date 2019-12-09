package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	for i := 0; i < len(nums); i++ {
		l := len(res)
		for j := 0; j < l; j++ {
			tmp := make([]int, len(res[j]))
			copy(tmp, res[j])
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
	}
	return res
}
