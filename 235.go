package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 4}

	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val >= root.Val && q.Val <= root.Val || p.Val <= root.Val && q.Val >= root.Val {
		return root
	} else if p.Val <= root.Val && q.Val <= root.Val {
		// 左子树
		if root.Left != nil {
			root = lowestCommonAncestor(root.Left, p, q)
		}

	} else {
		// 右子树
		if root.Right != nil {
			root = lowestCommonAncestor(root.Right, p, q)
		}

	}
	return root
}
