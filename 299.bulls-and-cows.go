/*
package main

import "fmt"

func main() {
	secret := "1123"
	guess := "0111"
	res := getHint(secret, guess)
	fmt.Println(res)
}
*/
func getHint(secret string, guess string) string {
	secret_arr := [10]int{0}
	guess_arr := [10]int{0}

	bull := 0
	cow := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
		} else {
			secret_arr[secret[i]-'0']++
			guess_arr[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cow += min(secret_arr[i], guess_arr[i])
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
