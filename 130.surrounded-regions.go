package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(board)
	printResult(board)

	board1 := [][]byte{
		[]byte{'O', 'O'},
		[]byte{'O', 'O'},
	}
	solve(board1)
	printResult(board1)
}

func printResult(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

type area struct {
	i int
	j int
}

func solve(board [][]byte) {
	border := []area{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
				rollback, region := helper(i, j, board)
				if rollback {
					border = append(border, region...)
				}
			}
		}
	}
	for _, b := range border {
		board[b.i][b.j] = 'O'
	}
}

func helper(i, j int, board [][]byte) (bool, []area) {
	list := []area{area{i, j}}
	res := []area{}
	rollback := false
	for len(list) != 0 {
		f := list[0]
		list = list[1:]
		res = append(res, f)
		if f.i == 0 || f.i == len(board)-1 || f.j == 0 || f.j == len(board[0])-1 {
			rollback = true
		}
		if f.i+1 < len(board) && board[f.i+1][f.j] == 'O' {
			board[f.i+1][f.j] = 'X'
			list = append(list, area{f.i + 1, f.j})
		}
		if f.j+1 < len(board[0]) && board[f.i][f.j+1] == 'O' {
			board[f.i][f.j+1] = 'X'
			list = append(list, area{f.i, f.j + 1})
		}
		if f.i-1 >= 0 && board[f.i-1][f.j] == 'O' {
			board[f.i-1][f.j] = 'X'
			list = append(list, area{f.i - 1, f.j})
		}
		if f.j-1 >= 0 && board[f.i][f.j-1] == 'O' {
			board[f.i][f.j-1] = 'X'
			list = append(list, area{f.i, f.j - 1})
		}
	}
	return rollback, res
}
