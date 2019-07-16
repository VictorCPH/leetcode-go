/*
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1994))
}
*/

func intToRoman(num int) string {
	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	var buf bytes.Buffer
	for _, v := range arr {
		if num/v > 0 {
			for i := 0; i < num/v; i++ {
				buf.WriteString(m[v])
			}
			num = num % v
		}
	}
	return buf.String()
}
