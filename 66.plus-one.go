package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{0}))
}

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+plus >= 10 {
			res[i+1] = (digits[i] + 1) % 10
		} else {
			res[i+1] = digits[i] + plus
			plus = 0
		}
	}
	if plus == 1 {
		res[0] += 1
		return res
	}
	return res[1:]
}
