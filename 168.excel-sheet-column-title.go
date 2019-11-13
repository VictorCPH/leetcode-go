package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(701))
}

func convertToTitle(n int) string {
	s := ""
	for n > 0 {
		mod := n % 26
		n = n / 26
		if mod == 0 {
			mod = 26
			n -= 1
		}
		s = string(mod-1+'A') + s
	}
	return s
}
