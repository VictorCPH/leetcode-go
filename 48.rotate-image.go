/*
package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
*/
func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i <= l/2; i++ {
		for j := i; j <= l-2-i; j++ {
			matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i] = matrix[l-1-j][i], matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j]
		}
	}
}
