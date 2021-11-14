package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	var (
		dx = []int{1, 0, 0, -1}
		dy = []int{0, 1, -1, 0}
	)

	oldColor := image[sr][sc]

	if oldColor == newColor {
		return image
	}

	r, c := len(image), len(image[0])
	visited := [][]int{}
	visited = append(visited, []int{sr, sc})
	image[sr][sc] = newColor
	for i := 0; i < len(visited); i++ {
		cell := visited[i]
		for j := 0; j < 4; j++ {
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < r && my >= 0 && my < c && image[mx][my] == oldColor {
				visited = append(visited, []int{mx, my})
				image[mx][my] = newColor
			}
		}
	}
	return image
}
