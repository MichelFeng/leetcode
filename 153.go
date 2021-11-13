package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	ans := math.MaxInt64
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < ans {
			ans = nums[mid]
		}
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] < ans {
		ans = nums[l]
	}
	return ans
}
