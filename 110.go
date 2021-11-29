package main

import (
	"fmt"
)

func main() {
	// root := &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 9,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 20,
	// 		Left: &TreeNode{
	// 			Val: 15,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,

				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(root *TreeNode, depth int) int
	dfs = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}

		left := dfs(root.Left, depth+1)

		right := dfs(root.Right, depth+1)
		if left > right {
			return left
		}
		return right
	}

	left := dfs(root.Left, 0)
	right := dfs(root.Right, 0)
	if left > right {
		return left-right <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
	}
	return right-left <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
