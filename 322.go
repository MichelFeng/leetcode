package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(coinChangeWithRecursive([]int{1, 2, 5}, 11))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(coinChangeWithMemo([]int{1, 2, 5}, 11))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(coinChangeWithDpTable([]int{1, 2, 5}, 11))
	fmt.Println(time.Since(start))

}

/*
1. 确定状态：目标金额 amount
2. 确定dp函数的定义：目标金额是 n，至少需要 dp(n) 个硬币
3. 确定选择：从硬币列表中选择一个，目标金额减少
4. 明确 base case：目标金额为0是，dp(0) = 0，目标金额小于0时，无解
*/

func coinChangeWithDpTable(coins []int, amount int) int {
	// dp[i] = x 表示目标金额为 i 时，至少需要 x 枚硬币
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0

	for i := 0; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 时间复杂度：O(kn)
// 子问题数目：O(n)
// 处理一个问题的时间复杂度：O(k)，k是硬币面值的个数
func coinChangeWithMemo(coins []int, amount int) int {
	memo := map[int]int{}
	var dp func(n int) int
	dp = func(n int) int {

		// 备忘录中已经存在，直接返回，无需计算
		if ans, exist := memo[n]; exist {
			return ans
		}
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxInt64
		for _, coin := range coins {

			subProblem := dp(n - coin)

			if subProblem == -1 {
				continue
			}
			res = min(subProblem+1, res)
		}

		// 将结果存入备忘录
		if res == math.MaxInt64 {
			memo[n] = -1
			return -1
		}
		memo[n] = res
		return memo[n]

	}
	return dp(amount)
}

// 时间复杂度：O(n^k)
// 子问题数目：O(n)
// 处理一个问题的时间复杂度：O(k*n^k)，k是硬币面值的个数
func coinChangeWithRecursive(coins []int, amount int) int {
	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxInt64
		for _, coin := range coins {
			subProblem := dp(n - coin)

			if subProblem == -1 {
				continue
			}
			res = min(subProblem+1, res)
		}

		if res == math.MaxInt64 {
			return -1
		}
		return res
	}

	return dp(amount)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
