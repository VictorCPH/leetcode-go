package main

import "fmt"

func main() {
	fmt.Println(numTrees(4))
}

func numTrees(n int) int {
	tmp := make([]int, n+1)
	tmp[0], tmp[1] = 1, 1
	helper(n, tmp)
	return tmp[n]
}

func helper(n int, tmp []int) int {
	if tmp[n] != 0 {
		return tmp[n]
	}
	for i := 1; i <= n; i++ {
		tmp[n] += helper(i-1, tmp) * helper(n-i, tmp)
	}
	return tmp[n]
}
