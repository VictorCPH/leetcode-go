/*
package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		[]int{1, 11},
		[]int{2, 12},
		[]int{3, 13},
		[]int{4, 14},
		[]int{5, 15},
		[]int{6, 16},
		[]int{7, 17},
		[]int{8, 18},
		[]int{9, 19},
		[]int{10, 20},
	}))
}
*/
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}

	i, j := 0, 0
	h := len(matrix)
	w := len(matrix[0])

	for k := 0; k < (h+1)/2 && k < (w+1)/2; k++ {
		if j >= w-k {
			break
		}
		for j = k; j < w-k; j++ {
			fmt.Println(matrix[i][j])
			res = append(res, matrix[i][j])
		}
		j--
		if i+1 == h-k {
			break
		}
		for i = i + 1; i < h-k; i++ {
			fmt.Println(matrix[i][j])
			res = append(res, matrix[i][j])
		}
		i--
		if j-1 < k {
			break
		}
		for j = j - 1; j >= k; j-- {
			fmt.Println(matrix[i][j])
			res = append(res, matrix[i][j])
		}
		j++
		if i-1 <= k {
			break
		}
		for i = i - 1; i > k; i-- {
			fmt.Println(matrix[i][j])
			res = append(res, matrix[i][j])
		}
		i++
	}
	return res
}
