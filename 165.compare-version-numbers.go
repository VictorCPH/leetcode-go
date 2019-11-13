package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.1.0", "1.1.1"))
	fmt.Println(compareVersion("1.01", "1.0001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}

func compareVersion(version1 string, version2 string) int {
	verArr1 := strings.Split(version1, ".")
	verArr2 := strings.Split(version2, ".")

	maxLen := max(len(verArr1), len(verArr2))
	for len(verArr1) < len(verArr2) {
		verArr1 = append(verArr1, "0")
	}
	for len(verArr1) > len(verArr2) {
		verArr2 = append(verArr2, "0")
	}

	for i := 0; i < maxLen; i++ {
		num1, _ := strconv.Atoi(verArr1[i])
		num2, _ := strconv.Atoi(verArr2[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
