/*
package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
*/

// O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := []int{}
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] <= nums2[idx2] {
			nums3 = append(nums3, nums1[idx1])
			idx1 += 1
		} else {
			nums3 = append(nums3, nums2[idx2])
			idx2 += 1
		}
	}
	if idx1 == len(nums1) {
		nums3 = append(nums3, nums2[idx2:]...)
	} else {
		nums3 = append(nums3, nums1[idx1:]...)
	}
	return float64((nums3[(len(nums3)-1)/2] + nums3[len(nums3)/2])) / 2
}

