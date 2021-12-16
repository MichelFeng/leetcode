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
	fmt.Println(minDiffInBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var min = math.MaxInt32
	var pre *TreeNode
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}
		if pre != nil {
			diff := root.Val - pre.Val
			if diff < min {
				min = diff
			}
		}
		pre = root
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return min
}
