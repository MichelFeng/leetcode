package main

import "fmt"

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}

func maxAreaOfIsland(grid [][]int) int {

	var (
		dx = []int{1, 0, 0, -1}
		dy = []int{0, 1, -1, 0}
	)

	r, c := len(grid), len(grid[0])

	ans := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			count := 0
			if grid[i][j] == 0 {
				continue
			}
			visited := [][]int{{i, j}}
			count++
			grid[i][j] = 0
			for k := 0; k < len(visited); k++ {
				cell := visited[k]
				for d := 0; d < 4; d++ {
					mx, my := cell[0]+dx[d], cell[1]+dy[d]
					if mx >= 0 && mx < r && my >= 0 && my < c && grid[mx][my] == 1 {
						visited = append(visited, []int{mx, my})
						grid[mx][my] = 0
						count++
					}
				}
			}
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}
