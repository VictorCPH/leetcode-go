/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
*/
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	combination(&res, candidates, target, []int{}, 0)
	return res
}

func combination(res *[][]int, candidates []int, target int, cur []int, idx int) {
	if target == 0 {
		match := make([]int, len(cur))
		copy(match, cur)
		*res = append(*res, match)
		return
	}

	for i := idx; i < len(candidates) && candidates[i] <= target; i++ {
		if i != idx && i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		combination(res, candidates, target-candidates[i], append(cur, candidates[i]), i+1)
	}
}
