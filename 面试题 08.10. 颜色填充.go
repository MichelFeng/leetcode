package main

import "fmt"

func main() {
	// image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	fmt.Println(floodFill(image, 1, 1, 1))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var di = [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	rows := len(image)
	cols := len(image[0])

	visited := make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		visited = append(visited, make([]int, cols))
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})

	visited[sr][sc] = 1
	oldColor := image[sr][sc]
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]
			visited[cur[0]][cur[1]] = 1
			for j := 0; j < 4; j++ {

				nr, nc := di[j][0]+cur[0], di[j][1]+cur[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && image[nr][nc] == oldColor && visited[nr][nc] == 0 {
					tmp = append(tmp, []int{nr, nc})
				}
			}
			image[cur[0]][cur[1]] = newColor
			queue = tmp
		}
	}
	return image
}
