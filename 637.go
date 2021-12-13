package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var sum float64 = 0.0
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]
			sum += float64(cur.Val)
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
			queue = tmp
		}
		res = append(res, sum/float64(size))
	}
	return res
}

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
	fmt.Println(averageOfLevels(root))
}
