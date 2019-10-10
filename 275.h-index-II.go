/*
package main

import "fmt"

func main() {
	citations := []int{0, 1, 3, 4, 5, 6}
	hIdx := hIndex(citations)
	fmt.Println(hIdx)
}
*/
func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	low := 0
	high := len(citations) - 1
	hIdx := 0
	for low < high {
		hIdx = (low + high) / 2
		if citations[hIdx] > len(citations)-hIdx {
			high = hIdx
		} else if citations[hIdx] < len(citations)-hIdx {
			low = hIdx + 1
		} else {
			return len(citations) - hIdx
		}
	}
	if low < len(citations) && citations[low] >= len(citations)-low {
		return len(citations) - low
	}
	return 0
}
