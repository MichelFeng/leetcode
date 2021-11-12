package main

// func main() {
// 	// 704.go
// 	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
// 	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
// }

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	i := 0
	j := l - 1
	for i <= j {
		mid := (j-i)/2 + i
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
