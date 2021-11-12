package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(1))
}

func firstBadVersion(n int) int {
	i, j := 0, n-1
	v := n
	for i <= j {
		mid := (j-i)/2 + i
		if isBadVersion(mid) {
			if mid < v {
				v = mid
			}
			j = mid - 1

		} else {
			i = mid + 1
		}
	}
	return v
}

func isBadVersion(n int) bool {
	return n == 1
}
