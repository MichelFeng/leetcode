package main

import "fmt"

func main() {
	fmt.Println(guessNumber(8))
}

func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := (r-l)/2 + l
		switch guess(mid) {
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		case 0:
			return mid
		}
	}
	return -1
}

func guess(n int) int {
	if n > 6 {
		return -1
	} else if n == 6 {
		return 0
	}
	return 1
}
