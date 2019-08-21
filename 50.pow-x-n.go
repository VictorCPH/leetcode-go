/*
package main

import "fmt"

func main() {
	fmt.Println(myPow(-1, -100))
	fmt.Println(myPow(0.4331, 0))
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.3, -3))
	fmt.Println(myPow(34.00515, 3))
	fmt.Println(myPow(34.00515, -3))
}
*/
func myPow(x float64, n int) float64 {
	if x == 0. || x == 1. {
		return 1.
	}
	if x == -1. {
		if n%2 == 0 {
			return 1.
		} else {
			return -1.
		}
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.
	}

	if n < 0 {
		return 1 / pow(x, -n)
	}

	res := pow(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return x * res * res
	}
}
