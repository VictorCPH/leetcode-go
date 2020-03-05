package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums1, 1))
	fmt.Println(findKthLargest(nums1, 2))
	fmt.Println(findKthLargest(nums1, 3))
	fmt.Println(findKthLargest(nums1, 4))
	fmt.Println(findKthLargest(nums1, 5))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	idx := partition(nums, 0, len(nums)-1, 0)
	if idx == len(nums)-k {
		return nums[idx]
	} else if idx > len(nums)-k {
		return findKthLargest(nums[:idx], k-(len(nums)-idx))
	} else {
		return findKthLargest(nums[idx+1:], k)
	}
}

func partition(nums []int, left, right, idx int) int {
	val := nums[idx]
	nums[idx], nums[right] = nums[right], nums[idx]
	storeIdx := left
	for i := left; i < right; i++ {
		if nums[i] < val {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx += 1
		}
	}
	nums[right], nums[storeIdx] = nums[storeIdx], nums[right]
	return storeIdx
}
