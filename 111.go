package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(minDepth(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	depth := 1
	for len(queue) > 0 {
		size := len(queue) // 当前层的结点个数
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]

			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}

			queue = tmp
		}
		depth++
	}

	return depth
}
