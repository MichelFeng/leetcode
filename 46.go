package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
}

func permute(nums []int) [][]int {
	var res = make([][]int, 0)

	var backtrack func(nums []int, track []int)
	backtrack = func(nums, track []int) {
		if len(track) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			contain := false
			for j := 0; j < len(track); j++ {
				if nums[i] == track[j] {
					contain = true
				}
			}
			if contain {
				continue
			}
	
			track = append(track, nums[i])
			backtrack(nums, track)
			track = track[:len(track)-1]
		}
	}

	track := make([]int, 0, len(nums))
	backtrack(nums, track)
	return res
}
