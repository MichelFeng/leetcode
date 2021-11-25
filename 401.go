package main

import "fmt"

func main() {
	// fmt.Println(trans(0))
	fmt.Println(readBinaryWatch(1))
	fmt.Println(readBinaryWatch(9))
}

/*
1. 路径：表盘，board[10]byte
2. 选择列表：10
3. 结束条件：board 中1个个数==turnedOn

*/
func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)

	board := make([]byte, 10)
	var backtrack func(board []byte, track int)
	backtrack = func(board []byte, track int) {

		if track >= 10 {
			return
		}
		// 结束条件
		if count(board) == turnedOn {
			total := 0
			for i := 0; i < len(board); i++ {
				if board[i] == 1 {
					total += trans(i)
				}
			}
			res = append(res, toString(total))
			return
		}

		for i := 0; i < 10; i++ {
			// 冲突检测
			if isConflict(board, i) {
				continue
			}

			// 选择
			board[i] = 1
			// 递归
			backtrack(board, track+1)
			// 取消选择
			board[i] = 0
		}

	}

	backtrack(board, 0)
	return res
}

func isConflict(board []byte, i int) bool {
	if i >= 6 {
		total := 0
		for j := 0; j < 4; j++ {
			total += 2 << board[6+j]
		}
		if total > 11 {
			return true
		}
	}
	total := 0
	for j := 0; j < 6; j++ {
		total += 2 << board[j]
	}
	if total > 60 {
		return true
	}

	return false
}

func trans(n int) int {
	if n >= 6 {
		return 60 * (1 << (n - 6))
	}
	return 1 << n
}

func toString(n int) string {
	hours := n / 60
	minutes := n % 60
	return fmt.Sprintf("%d:%02d", hours, minutes)
}

func count(board []byte) int {
	c := 0
	for i := range board {
		if board[i] == 1 {
			c++
		}
	}
	return c
}
