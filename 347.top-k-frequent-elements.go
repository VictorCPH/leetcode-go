package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type KV struct {
	Key int
	Val int
}

type ByVal []KV

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i].Val < a[j].Val }

func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num] += 1
	}
	kv := []KV{}
	for k, v := range maps {
		kv = append(kv, KV{k, v})
	}
	sort.Sort(ByVal(kv))
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, kv[len(kv)-1-i].Key)
	}
	return res
}
