package main

import (
	"fmt"
)

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))
}

func updateMatrix(mat [][]int) [][]int {
	var (
		dx = []int{0, 1, 0, -1}
		dy = []int{-1, 0, 1, 0}
	)
	row, column := len(mat), len(mat[0])
	ans := make([][]int, row)
	visited := make([][]int, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]int, column)
		visited[i] = make([]int, column)
	}
	queue := [][]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if mat[i][j] == 0 {
				visited[i][j] = 1
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]
		for d := 0; d < 4; d++ {
			mx, my := cell[0]+dx[d], cell[1]+dy[d]
			if mx >= 0 && mx < row && my >= 0 && my < column && visited[mx][my] == 0 {
				queue = append(queue, []int{mx, my})
				ans[mx][my] = ans[cell[0]][cell[1]] + 1
				visited[mx][my] = 1
			}
		}
	}
	return ans
}
