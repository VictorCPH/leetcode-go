/*
package main

import "fmt"

func main() {
	s := "mississippi"
	p := "mis*is*p*."
	fmt.Println(isMatch(s, p))
}
*/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
