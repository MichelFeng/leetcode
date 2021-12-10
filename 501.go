package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Left: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println(findMode(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := map[int]int{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		m[root.Val]++
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)

	maxCount := 0
	res := []int{}
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}
	for k, v := range m {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}
