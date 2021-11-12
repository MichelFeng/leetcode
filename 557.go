package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseStr(s []byte) string {
	for l, r := 0, len(s)-1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}

func reverseWords(s string) string {
	l := strings.Split(s, " ")
	res := []string{}
	for _, i := range l {
		res = append(res, reverseStr([]byte(i)))
	}
	return strings.Join(res, " ")
}
