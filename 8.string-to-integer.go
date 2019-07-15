/*
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-2147483649"))
}
*/

func myAtoi(str string) int {
	res := 0
	symbol := 1
	idx := 0
	for idx < len(str) && str[idx] == ' ' {
		idx++
	}

	if idx < len(str) && (str[idx] == '+' || str[idx] == '-' || (str[idx] >= '0' && str[idx] <= '9')) {
		if str[idx] == '+' || str[idx] == '-' {
			if str[idx] == '-' {
				symbol = -1
			}
			idx += 1
		}
		for idx < len(str) && str[idx] >= '0' && str[idx] <= '9' {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && str[idx]-'0' > 7) {
				if symbol < 0 {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
			res = res*10 + int(str[idx]-'0')
			idx += 1
		}
		return res * symbol
	} else {
		return 0
	}
}
