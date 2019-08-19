/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
*/
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	arr := permuteUnique(nums[1:])
	res := [][]int{}
	set := make(map[string]int)
	for _, val := range arr {
		tmp := append([]int{nums[0]}, val...)
		tmpStr := sliceToStr(tmp)
		if _, ok := set[tmpStr]; !ok {
			fmt.Println(tmpStr)
			set[tmpStr] = 1
			res = append(res, tmp)
		}

		for i := 0; i < len(val); i++ {
			tmp = make([]int, len(val)+1)
			copy(tmp, val[:i+1])
			tmp[i+1] = nums[0]
			copy(tmp[i+2:], val[i+1:])
			tmpStr = sliceToStr(tmp)
			if _, ok := set[tmpStr]; !ok {
				fmt.Println(tmpStr)
				set[tmpStr] = 1
				res = append(res, tmp)
			}
		}
	}
	return res
}

func sliceToStr(s []int) string {
	res := ""
	for _, v := range s {
		res += strconv.Itoa(v)
	}
	return res
}
