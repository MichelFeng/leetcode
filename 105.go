package main

import "fmt"

func main() {
	root := buildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4})
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	root := &TreeNode{}
	root.Val = preorder[0]
	idx := search(inorder, root.Val)
	if idx != -1 {
		left := inorder[:idx]
		right := inorder[idx+1:]
		if len(left) > 0 {
			root.Left = buildTree(preorder[1:len(left)+1], left)
		}
		if len(right) > 0 {
			root.Right = buildTree(preorder[len(left)+1:], right)
		}
	} else {
		return nil
	}

	return root
}

func search(nums []int, target int) int {
	for idx, val := range nums {
		if val == target {
			return idx
		}
	}
	return -1
}
