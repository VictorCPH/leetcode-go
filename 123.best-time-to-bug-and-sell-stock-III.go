/*
package main

import "fmt"

func main() {
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices1))
	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices2))
	prices3 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(maxProfit(prices3))
	prices4 := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfit(prices4))
}
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	leftProfit := make([]int, len(prices))
	rightProfit := make([]int, len(prices))

	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit = maxInt(maxProfit, prices[i]-min)
		min = minInt(min, prices[i])
		leftProfit[i] = maxProfit
	}

	max := prices[len(prices)-1]
	maxProfit = 0
	for i := len(prices) - 2; i >= 0; i-- {
		maxProfit = maxInt(maxProfit, max-prices[i])
		max = maxInt(max, prices[i])
		rightProfit[i] = maxProfit
	}

	maxProfit = 0
	for i := 0; i < len(prices); i++ {
		maxProfit = maxInt(maxProfit, rightProfit[i]+leftProfit[i])
	}
	return maxProfit
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
