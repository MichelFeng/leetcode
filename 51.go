package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// fmt.Println(solveNQueens(1))
	start := time.Now()
	res := solveNQueens(8)
	fmt.Println(res, len(res))
	fmt.Println(time.Since(start))
}

/*
1. 路径：棋盘
2. 选择列表：每一行的所有列
3. 结束条件：遍历到最后一行时
*/
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	var backtrack func(board []string, row int)
	backtrack = func(board []string, row int) {
		// 结束条件
		if row == len(board) {
			tmp := make([]string, n)
			copy(tmp, board)
			res = append(res, tmp)
			return
		}

		for col := 0; col < n; col++ {

			if isConflict(board, row, col) {
				continue
			}

			// 做选择
			board[row] = strings.Repeat(".", col) + "Q" + strings.Repeat(".", n-col-1)
			backtrack(board, row+1)
			// 撤销选择
			board[row] = strings.Repeat(".", n)
		}
	}

	backtrack(board, 0)

	return res
}

func isConflict(board []string, row, col int) bool {
	// 列上是否存在
	for i := 0; i < row; i++ {
		if rune(board[i][col]) == rune('Q') {
			return true
		}
	}

	// 左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if rune(board[i][j]) == rune('Q') {
			return true
		}
		i--
		j--
	}

	// 右上
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if rune(board[i][j]) == rune('Q') {
			return true
		}
		i--
		j++
	}
	return false
}
