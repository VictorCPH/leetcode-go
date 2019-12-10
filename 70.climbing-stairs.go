package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	x, y := 1, 1
	for i := 0; i < n-1; i++ {
		x, y = y, x+y
	}
	return y
}
