package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindows("ADOBECODEBANC", "ABC"))
}

func minWindows(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for idx := range t {
		need[t[idx]]++
		window[t[idx]] = 0
	}

	left, right := 0, 0
	valid := 0
	start := 0
	length := math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]

}
