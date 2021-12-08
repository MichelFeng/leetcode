package main

import (
	"fmt"
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
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, min *TreeNode, max *TreeNode) bool

	dfs = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
		if root == nil {
			return true
		}

		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return dfs(root.Left, min, root) && dfs(root.Right, root, max)
	}

	return dfs(root, nil, nil)
}
