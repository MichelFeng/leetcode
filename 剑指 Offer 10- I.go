package main

import "fmt"

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(45))
	fmt.Println(fib(100))
}

// leetcode 的int 默认是32位的，会有溢出问题
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var (
		a int64 = 0
		b int64 = 1
	)

	for i := 2; i < n; i++ {
		t := a%(1e9+7) + b%(1e9+7)
		a = b
		b = t
	}
	return int((a + b) % (1e9 + 7))
}
