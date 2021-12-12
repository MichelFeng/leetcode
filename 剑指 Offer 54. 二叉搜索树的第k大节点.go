package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	fmt.Println(kthLargest(root, 1))
}

func kthLargest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	var res int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Right != nil {
			dfs(root.Right)
		}
		k--
		if k == 0 {
			res = root.Val
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}

	}

	dfs(root)
	return res
}
