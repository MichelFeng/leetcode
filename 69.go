package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(6))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(0))
}

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (r-l)/2 + l
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
