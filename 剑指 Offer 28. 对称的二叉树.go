package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	res := isSymmetric(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	mirror := mirrorTree(root)

	var dfs func(root, mirror *TreeNode) bool

	dfs = func(root, mirror *TreeNode) bool {
		if root == nil && mirror == nil {
			return true
		} else if root == nil || mirror == nil {
			return false
		}
		if root.Val != mirror.Val {
			return false
		}

		return dfs(root.Left, mirror.Left) && dfs(root.Right, mirror.Right)
	}
	return dfs(root, mirror)
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{
		Val: root.Val,
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)

	newRoot.Left = right
	newRoot.Right = left
	return newRoot
}
