package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(isCousins(root, 4, 5))
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := make([]*TreeNode, 0)
	if root.Val == x || root.Val == y {
		return false
	}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		size := len(queue)
		xOk, yOk := false, false
		for i := 0; i < size; {
			left := queue[0]
			right := queue[1]
			tmp := queue[2:]
			leftOk, rightOk := false, false
			if left != nil {
				if left.Val == x {
					xOk = true
					leftOk = true
				}
				if left.Val == y {
					yOk = true
					leftOk = true
				}
				tmp = append(tmp, left.Left)
				tmp = append(tmp, left.Right)
			}
			if right != nil {
				if right.Val == x {
					xOk = true
					rightOk = true
				}
				if right.Val == y {
					yOk = true
					rightOk = true
				}
				tmp = append(tmp, right.Left)
				tmp = append(tmp, right.Right)
			}
			if leftOk && rightOk {
				return false
			}
			i += 2
			queue = tmp
		}
		if xOk && yOk {
			return true
		}
	}
	return false
}
