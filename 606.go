package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Left: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(tree2str(root))
}

func tree2str(root *TreeNode) string {
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		tmp := []string{}
		tmp = append(tmp, fmt.Sprintf("%v", root.Val))
		if root.Left != nil || root.Right != nil {

			if root.Left != nil {
				tmp = append(tmp, "(")
				tmp = append(tmp, dfs(root.Left))
				tmp = append(tmp, ")")
			} else if root.Right != nil {
				tmp = append(tmp, "(")
				tmp = append(tmp, ")")
			}
			if root.Right != nil {
				tmp = append(tmp, "(")
				tmp = append(tmp, dfs(root.Right))
				tmp = append(tmp, ")")
			}

		}
		return strings.Join(tmp, "")
	}

	return dfs(root)
}
