/*
package main

import "fmt"
import "math"

func main() {
	fmt.Println(divide(-7, 3))
	fmt.Println(divide(math.MaxInt32, 1))
}
*/

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	same := dividend^divisor >= 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := 0
	for dividend >= divisor {
		tmp := divisor
		var n uint = 0
		for tmp <= dividend {
			tmp <<= 1
			n++
		}
		res += 1 << (n - 1)
		dividend -= tmp >> 1
	}

	if same {
		return res
	}
	return -res
}
