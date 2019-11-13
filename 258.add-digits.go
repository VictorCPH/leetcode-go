package main

import "fmt"

func main() {
	fmt.Println(addDigits(123456))
}

func addDigits(num int) int {
	for num > 9 {
		sum := 0
		for num > 0 {
			mod := num % 10
			num = num / 10
			sum += mod
		}
		num = sum
	}
	return num
}
