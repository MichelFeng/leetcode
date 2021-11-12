package main

import "fmt"

// func main() {
// 01.06.go
// fmt.Println(compressString("aabcccccaa"))
// fmt.Println(compressString("aabc"))
// }

func compressString(S string) string {
	if len(S) < 2 {
		return S
	}

	res := ""
	c := S[0]
	n := 1

	for i := 1; i < len(S); i++ {
		if S[i] == c {
			n++
		} else {
			res += fmt.Sprintf("%c%v", c, n)
			c = S[i]
			n = 1
		}
	}
	res += fmt.Sprintf("%c%v", c, n)

	if len(S) > len(res) {
		return res
	}
	return S
}
