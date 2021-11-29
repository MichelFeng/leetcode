package main

import "fmt"

func main() {
	fmt.Println(generateTrees(3))
	fmt.Println(generateTrees(1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	var generate func(left, right int) []*TreeNode
	generate = func(left, right int) []*TreeNode {
		res := make([]*TreeNode, 0)
		if left > right {
			res = append(res, nil)
			return res
		}

		for mid := left; mid <= right; mid++ {
			lefts := generate(left, mid-1)
			rights := generate(mid+1, right)

			for i := range lefts {
				for j := range rights {
					root := &TreeNode{
						Left:  lefts[i],
						Right: rights[j],
						Val:   mid,
					}
					res = append(res, root)
				}
			}
		}
		return res
	}
	return generate(1, n)

}
