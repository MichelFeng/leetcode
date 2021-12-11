package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	fmt.Println(islandPerimeter([][]int{
		{1, 0},
	}))
}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dims := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	} // 上下左右

	res := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			val := grid[row][col]
			if val == 0 {
				continue
			}

			// 判断上下左右是否为0
			for i := range dims {
				dim := dims[i]
				newRow := dim[0] + row
				newCol := dim[1] + col
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					if grid[newRow][newCol] == 0 {
						res++
					}
				} else {
					res++
				}
			}
		}
	}
	return res
}
