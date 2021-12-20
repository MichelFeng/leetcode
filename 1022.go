package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(sumRootToLeaf(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var dfs func(root *TreeNode, bytes []int)
	dfs = func(root *TreeNode, bytes []int) {
		if root == nil {

			return
		}

		bytes = append(bytes, root.Val)
		if root.Left != nil {
			dfs(root.Left, bytes)
		}
		if root.Right != nil {
			dfs(root.Right, bytes)
		}
		if root.Left == nil && root.Right == nil {
			s := 0
			for i := 0; i < len(bytes); i++ {
				if bytes[i] == 1 {
					s += 1 << (len(bytes) - 1 - i)
				}
			}
			sum += s
		}

	}
	dfs(root, []int{})
	return sum
}
