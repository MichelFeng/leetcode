package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	fmt.Println(getMinimumDifference(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	min := math.MaxInt32
	pre := -1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 && root.Val-pre < min {
			min = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return min
}
