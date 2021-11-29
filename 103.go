package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,

				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	var bfs func(root *TreeNode) [][]int
	bfs = func(root *TreeNode) [][]int {
		res := make([][]int, 0)
		if root == nil {
			return res
		}

		idx := 0
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
			if idx%2 == 1 {
				for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
					level[i], level[j] = level[j], level[i]
				}
			}

			res = append(res, level)
			idx++
		}
		return res
	}
	return bfs(root)
}
