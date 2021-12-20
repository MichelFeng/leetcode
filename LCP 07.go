package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	nums := make([][]int, n)
	for _, r := range relation {
		row, col := r[0], r[1]
		nums[row] = append(nums[row], col)
	}
	count := 0
	var dfs func(col int, step int)
	dfs = func(col int, step int) {
		if step == k {
			if col == n-1 {
				count++
			}
			return
		}
		for _, y := range nums[col] {
			dfs(y, step+1)
		}
	}
	dfs(0, 0)

	return count
}

func main() {
	fmt.Println(numWays(5, [][]int{
		{0, 2},
		{2, 1},
		{3, 4},
		{2, 3},
		{1, 4},
		{2, 0},
		{0, 4},
	}, 3))
	// fmt.Println(numWays(3, [][]int{
	// 	{0, 2},
	// 	{0, 1},
	// 	{1, 2},
	// 	{2, 1},
	// 	{2, 0},
	// 	{1, 0},
	// }, 1))
}
