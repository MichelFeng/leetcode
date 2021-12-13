package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	depth := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			tmp := queue[1:]

			for _, n := range cur.Children {
				tmp = append(tmp, n)
			}
			queue = tmp
		}
		depth++
	}

	return depth
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	fmt.Println(maxDepth(root))
}
