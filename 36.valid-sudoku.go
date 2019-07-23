/*
package main

import "fmt"

func main() {
	fmt.Println(isValidSudoku([][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}
*/
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if _, ok := m[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' {
				m[board[i][j]] = true
			}
		}
	}

	for j := 0; j < 9; j++ {
		n := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if _, ok := n[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' {
				n[board[i][j]] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			p := make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if _, ok := p[board[i+k][j+l]]; ok {
						return false
					} else if board[i+k][j+l] != '.' {
						p[board[i+k][j+l]] = true
					}
				}
			}
		}
	}
	return true
}
