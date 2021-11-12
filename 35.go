package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1}, 0))
	fmt.Println(searchInsert([]int{1, 3}, 2))
}

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums) - 1
	ans := len(nums)
	for i <= j {
		mid := (j-i)/2 + i
		n := nums[mid]
		if n >= target {
			ans = mid
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return ans
}
