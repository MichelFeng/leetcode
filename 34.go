package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false) - 1

	if start <= end && end < len(nums) && nums[start] == target && nums[end] == target {
		return []int{start, end}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}
