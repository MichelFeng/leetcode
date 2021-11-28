package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 20,
		},
	}
	fmt.Println(sumOfLeftLeaves(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	idx := 0
	for idx < len(queue) {
		cur := queue[idx]
		if cur.Left != nil {
			if cur.Left.Left == nil && cur.Left.Right == nil {
				sum += cur.Left.Val
			}
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		idx++
	}

	return sum
}
