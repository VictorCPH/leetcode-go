/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{1, 0}
	ret := hIndex(citations)
	fmt.Println(ret)
}
*/
func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for i := 0; i < len(citations); i++ {
		if i+1 > citations[i] {
			return i
		}
	}
	if citations[0] < 1 {
		return 0
	}
	return len(citations)
}
