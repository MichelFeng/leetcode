package main

import "fmt"

func main() {
	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))

	fmt.Println(sortedSquares([]int{-5, -3, -2, 0, 1}))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1

	ans := make([]int, n)
	for pos := n - 1; pos >= 0; pos-- {
		ls := nums[l] * nums[l]
		rs := nums[r] * nums[r]
		if ls > rs {
			ans[pos] = ls
			l++
		} else {
			ans[pos] = rs
			r--
		}
	}
	return ans
}
