/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
}
*/
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	combination(candidates, target, &res, []int{}, 0)
	return res
}

func combination(candidates []int, target int, res *[][]int, cur []int, idx int) {
	if target == 0 {
		match := make([]int, len(cur))
		copy(match, cur)
		*res = append(*res, match)
		return
	}

	for i := idx; i < len(candidates) && candidates[i] <= target; i++ {
		cur = append(cur, candidates[i])
		combination(candidates, target-candidates[i], res, cur, i)
		cur = cur[:len(cur)-1]
	}
}
