package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 2, 4))
	fmt.Println(findClosestElements([]int{1, 2}, 1, 1))
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	left, right := 0, n-k // 起始位置取值的下标区间
	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
