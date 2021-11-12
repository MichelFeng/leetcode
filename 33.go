package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		n := nums[mid]
		if n == target {
			return mid
		}
		if nums[l] <= n {
			if nums[l] <= target && target <= n {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if n < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
