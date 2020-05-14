package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(numbers); i++ {
		if idx, ok := m[numbers[i]]; ok {
			return []int{idx + 1, i + 1}
		} else {
			m[target-numbers[i]] = i
		}
	}
	return []int{-1, -1}
}
