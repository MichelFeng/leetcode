package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))

	fmt.Println(checkInclusion("b", "eibdaooo"))

	fmt.Println(checkInclusion("hello", "ooolleooolleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)

	if n > m {
		return false
	}

	var cnt1, cnt2 [26]int
	for i := 0; i < n; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

// // 超时了
// func checkInclusion(s1 string, s2 string) bool {
// 	m1 := map[byte]int{}
// 	m2 := map[byte]int{}
// 	lenS1 := len(s1)
// 	for i := 0; i < lenS1; i++ {
// 		m1[s1[i]]++
// 		m2[s1[i]]++
// 	}

// 	n := len(s2)

// 	for i := 0; i < n; i++ {
// 		rk := i
// 		for rk < n && m1[s2[rk]] > 0 {

// 			m1[s2[rk]]--
// 			if m1[s2[rk]] == 0 {
// 				delete(m1, s2[rk])
// 			}
// 			rk++
// 		}

// 		if len(m1) == 0 {
// 			return true
// 		}
// 		for k, v := range m2 {
// 			m1[k] = v
// 		}

// 	}
// 	return false
// }
