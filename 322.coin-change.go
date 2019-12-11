package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	res := helper(coins, amount, dp)
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func helper(coins []int, amount int, dp []int) int {
	if dp[amount] != 0 {
		return dp[amount]
	}
	if amount == 0 {
		return 0
	}
	res := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		if amount >= coins[i] {
			res = min(res, helper(coins, amount-coins[i], dp)+1)
		}
	}
	dp[amount] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
