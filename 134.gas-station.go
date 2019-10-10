/*
package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)
}
*/
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		j := i
		gasSum := 0
		costSum := 0

		for {
			gasSum += gas[j]
			costSum += cost[j]
			if gasSum < costSum {
				if j >= i && j < len(gas)-1 {
					i = j
				} else {
					return -1
				}
				break
			} else {
				j = (j + 1) % len(gas)
			}
			if j == i {
				return i
			}
		}
	}
	return -1
}
