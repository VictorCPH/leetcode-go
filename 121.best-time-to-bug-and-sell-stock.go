/*
package main

import "fmt"

func main() {
	prices := []int{7, 4, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices1 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices1))
}
*/
func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[minPrice]
		if profit > maxProfit {
			maxProfit = profit
		} else if profit < 0 {
			minPrice = i
		}
	}
	return maxProfit
}
