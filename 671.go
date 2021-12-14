package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(findSecondMinimumValue(root))
}

func findSecondMinimumValue(root *TreeNode) int {
	min, second := math.MaxInt64, math.MaxInt64
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]

			if cur.Val < min {
				second = min
				min = cur.Val
			} else if cur.Val > min {
				if cur.Val < second {
					second = cur.Val
				}
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
	if second == math.MaxInt64 {
		return -1
	}
	return second
}
