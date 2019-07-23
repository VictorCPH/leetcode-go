/*
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(countAndSay(i))
	}
}
*/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	count := 1
	i := 1
	var buf bytes.Buffer
	for i = 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			buf.WriteString(strconv.Itoa(count))
			buf.WriteByte(s[i-1])
			count = 1
		}
	}
	buf.WriteString(strconv.Itoa(count))
	buf.WriteByte(s[i-1])

	return buf.String()
}
