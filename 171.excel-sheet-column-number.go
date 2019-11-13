package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	res := 0
	tmp := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i]-'A'+1) * tmp
		tmp *= 26
	}
	return res
}
