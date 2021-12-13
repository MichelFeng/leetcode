package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	res := true
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	val := root.Val
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]
			if cur.Val != val {
				res = false
				return res
			}
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
			queue = tmp
		}
	}
	return res
}
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(isUnivalTree(root))
}
