package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]struct{}{}
	for _, n := range nums1 {
		m[n] = struct{}{}
	}
	inter := map[int]struct{}{}
	for _, n := range nums2 {
		if _, exist := m[n]; exist {
			inter[n] = struct{}{}
		}
	}
	res := make([]int, 0)
	for k := range inter {
		res = append(res, k)
	}

	return res
}
