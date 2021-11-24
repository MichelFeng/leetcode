package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	return dp[n]
}
