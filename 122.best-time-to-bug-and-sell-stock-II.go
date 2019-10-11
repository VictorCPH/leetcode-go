/*
package main

import "fmt"

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1))
	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices2))
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices3))
}
*/
func maxProfit(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			total += profit
		}
	}
	return total
}
