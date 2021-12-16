package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}
	res := increasingBST(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var head, cur *TreeNode
	head = &TreeNode{}
	cur = head
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		node := &TreeNode{
			Val: root.Val,
		}
		cur.Right = node
		cur = cur.Right
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return head.Right
}
