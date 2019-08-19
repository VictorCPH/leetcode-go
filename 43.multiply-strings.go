/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("11", "721"))
}
*/
func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')
			res[i+j+1] += product
		}
	}

	carry := 0
	sum := ""
	for i := len(res) - 1; i >= 0; i-- {
		res[i] += carry
		carry = res[i] / 10
		sum = strconv.Itoa(res[i]%10) + sum
	}

	idx := 0
	for ; idx < len(res)-1 && res[idx] == 0; idx++ {
	}
	return sum[idx:]
}

/*
func multiply(num1 string, num2 string) string {
	if (len(num1) == 1 && int(num1[0]-'0') == 0) ||
		(len(num2) == 1 && int(num2[0]-'0') == 0) {
		return "0"
	}

	res := []string{}

	for i := 0; i < len(num2); i++ {
		add := 0
		res = append(res, "")
		for j := len(num1) - 1; j >= 0; j-- {
			product := int(num2[i]-'0')*int(num1[j]-'0') + add
			res[i] = strconv.Itoa(product%10) + res[i]
			add = product / 10
		}
		for k := i + 1; k < len(num2); k++ {
			res[i] += "0"
		}
		if add != 0 {
			res[i] = strconv.Itoa(add) + res[i]
		}
	}

	result := ""
	add := 0
	for i := 0; i < len(res[0]); i++ {
		sum := 0
		for j := 0; j < len(res); j++ {
			if len(res[j])-1 >= i {
				sum += int(res[j][len(res[j])-1-i] - '0')
			}
		}
		sum += add
		result = strconv.Itoa(sum%10) + result
		add = sum / 10
	}
	if add != 0 {
		result = strconv.Itoa(add) + result
	}
	return result
}
*/
