/*
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat"}))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	res := [][]string{}
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		key := generateKey(strs[i])
		if _, ok := m[key]; !ok {
			m[key] = []string{}
		}
		m[key] = append(m[key], strs[i])
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func generateKey(s string) string {
	count := [26]int{}
	for _, v := range s {
		count[int(v-'a')] += 1
	}

	var buf bytes.Buffer
	for _, v := range count {
		buf.WriteString("#")
		buf.WriteString(strconv.Itoa(v))
	}
	return buf.String()
}
