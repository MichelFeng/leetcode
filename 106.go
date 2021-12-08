package main

import "fmt"

func main() {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	root := &TreeNode{}
	root.Val = postorder[len(postorder)-1]
	idx := search(inorder, root.Val)
	if idx != -1 {
		left := inorder[:idx]
		right := inorder[idx+1:]
		if len(left) > 0 {
			root.Left = buildTree(left, postorder[0:len(left)])
		}
		if len(right) > 0 {
			root.Right = buildTree(right, postorder[len(left):len(postorder)-1])
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
