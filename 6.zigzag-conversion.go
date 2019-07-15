/*
package main

import (
	"bytes"
	"fmt"
)

func main() {
	p := convert("AB", 1)
	fmt.Println(p)
}
*/
func convert(s string, numRows int) string {
	if len(s) == 0 || len(s) == 1 || numRows <= 1 {
		return s
	}

	var buffer bytes.Buffer
	for i := 0; i < len(s); i += 2 * (numRows - 1) {
		buffer.WriteByte(s[i])
	}

	for i := 1; i < numRows-1; i++ {
		k := i
		for j := numRows - 1; k < len(s); j += numRows - 1 {
			buffer.WriteByte(s[k])
			k = (j-k)*2 + k
		}
	}

	for i := numRows - 1; i < len(s); i += 2 * (numRows - 1) {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}
