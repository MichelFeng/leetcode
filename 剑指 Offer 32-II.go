package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
			Val: 20,
		},
	}
	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
			queue = tmp
		}
		res = append(res, level)
		depth++
	}

	return res
}
