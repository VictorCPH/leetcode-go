package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		obstacleGrid[0][0] = 0
	} else {
		obstacleGrid[0][0] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] != 1 {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for j := 1; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] != 1 {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		} else {
			obstacleGrid[0][j] = 0
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] != 1 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
