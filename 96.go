package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(1))
}

func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	var count func(left, right int) int
	count = func(left, right int) int {
		if left > right {
			return 1
		}
		if memo[left][right] != 0 {
			return memo[left][right]
		}
		res := 0
		for i := left; i <= right; i++ {

			res += count(left, i-1) * count(i+1, right)

		}
		memo[left][right] = res
		return res
	}
	return count(1, n)
}
