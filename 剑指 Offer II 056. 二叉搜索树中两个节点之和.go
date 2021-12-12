package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 11,
			},
		},
	}

	fmt.Println(findTarget(root, 12))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}

	exist := false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := m[root.Val]; ok {
			exist = true
			return
		}
		diff := k - root.Val
		m[diff] = true
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return exist
}
