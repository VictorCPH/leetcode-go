package main

import "fmt"

type land struct {
	i int
	j int
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))

	grid1 := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid1))

	grid2 := [][]byte{
		[]byte{'1', '1', '1'},
		[]byte{'1', '1', '0'},
		[]byte{'1', '1', '1'},
	}
	fmt.Println(numIslands(grid2))
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '2'
				list := []land{land{i, j}}
				for len(list) != 0 {
					f := list[0]
					list = list[1:]
					if f.i+1 < len(grid) && grid[f.i+1][f.j] == '1' {
						grid[f.i+1][f.j] = '2'
						list = append(list, land{f.i + 1, f.j})
					}
					if f.j+1 < len(grid[0]) && grid[f.i][f.j+1] == '1' {
						grid[f.i][f.j+1] = '2'
						list = append(list, land{f.i, f.j + 1})
					}
					if f.i-1 >= 0 && grid[f.i-1][f.j] == '1' {
						grid[f.i-1][f.j] = '2'
						list = append(list, land{f.i - 1, f.j})
					}
					if f.j-1 >= 0 && grid[f.i][f.j-1] == '1' {
						grid[f.i][f.j-1] = '2'
						list = append(list, land{f.i, f.j - 1})
					}
				}
				count += 1
			}
		}
	}
	return count
}
