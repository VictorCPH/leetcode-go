package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
}

func addBinary(a, b string) string {
	res := ""
	for len(a) > len(b) {
		b = "0" + b
	}
	for len(a) < len(b) {
		a = "0" + a
	}

	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		m := int(a[i] - '0')
		n := int(b[i] - '0')
		sum := m + n + carry
		if sum > 1 {
			res = string(byte(sum%2+'0')) + res
			carry = 1
		} else {
			res = string(byte(sum+'0')) + res
			carry = 0
		}
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}
