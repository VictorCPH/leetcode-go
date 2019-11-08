package main

import "fmt"

func main() {
	s := []byte{'a', 'b', 'c', 'd', 'e'}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
