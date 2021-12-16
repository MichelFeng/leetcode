package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	fmt.Println(rangeSumBST(root, 7, 15))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low, high int) int {

	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// 没有使用到二叉搜索树的特性
func rangeSumBST_slow(root *TreeNode, low int, high int) int {
	var sum int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)

	return sum
}
