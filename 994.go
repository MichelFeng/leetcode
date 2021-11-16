package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}

func orangesRotting(grid [][]int) int {
	var (
		dx = []int{0, 1, 0, -1}
		dy = []int{-1, 0, 1, 0}
	)
	row, column := len(grid), len(grid[0])

	queue := [][]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		cell := queue[0]
		ans = cell[2]
		queue = queue[1:]
		for d := 0; d < 4; d++ {
			mx, my := cell[0]+dx[d], cell[1]+dy[d]
			if mx >= 0 && mx < row && my >= 0 && my < column && grid[mx][my] == 1 {
				queue = append(queue, []int{mx, my, ans + 1})
				grid[mx][my] = 2
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}
