/*
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123321))
}
*/

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	right := 0
	for x > right {
		right = right*10 + x%10
		x = x / 10
	}

	return x == right || x == right/10
}
