/*
package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	product := []string{""}
	tmp := []string{}
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(product); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				tmp = append(tmp, product[j]+string([]byte{m[digits[i]][k]}))
			}
		}
		product = tmp
		tmp = []string{}
	}
	return product
}
