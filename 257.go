package main

import (
	"fmt"
	"strings"
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
	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	slice := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {

		slice = append(slice, fmt.Sprintf("%v", root.Val))
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, strings.Join(slice, "->"))
		}
		slice = slice[:len(slice)-1]

	}
	dfs(root)
	return res
}
