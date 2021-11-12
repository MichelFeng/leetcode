package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		nl := numbers[l]
		nr := numbers[r]
		if nl+nr > target {
			r--
		} else if nl+nr == target {
			return []int{l + 1, r + 1}
		} else {
			l++
		}
	}
	return nil
}
