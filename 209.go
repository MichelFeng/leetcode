package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

/*
由于nums数组中的所有数字都是正数，所以前n个数字的和是有序递增的
因此可以使用前序和来构造有序数组，然后通过二分查找来得到区间
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sums := make([]int, n+1) // sums[i] 表示 nums 中前 i-1 个数的和
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	minLen := math.MaxInt64
	for i := 1; i <= n; i++ {
		newTarget := target + sums[i-1]
		leftBound := binarySearch(newTarget, sums)
		if leftBound < 0 {
			leftBound = -leftBound - 1
		}
		if leftBound <= n {
			length := (leftBound - i + 1)
			if minLen > length {
				minLen = length
			}
		}
	}
	//
	if minLen == math.MaxInt64 {
		return 0
	}
	return minLen
}

// 通过二分查找，寻找前缀和大于等于target左边界
func binarySearch(target int, sums []int) int {

	left := 0
	right := len(sums) - 1
	for left <= right { // 左闭右闭区间，终止条件是 left = right+1
		mid := left + (right-left)/2
		if target == sums[mid] {
			// 向左收缩
			right = mid - 1
		} else if target > sums[mid] {
			// 区间调整为 [mid+1, right]
			left = mid + 1
		} else if target < sums[mid] {
			// 区间调整为 [left, mid - 1]
			right = mid - 1
		}
	}

	if left >= len(sums) || sums[left] >= target {
		return left
	}
	return -1
}
