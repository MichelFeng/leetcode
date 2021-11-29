package main

import "fmt"

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	n := len(nums)
	if n == 0 {
		return nil
	}
	var generate func(left, right int) *TreeNode

	generate = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2

		root := &TreeNode{}
		root.Val = nums[mid]
		root.Left = generate(left, mid-1)
		root.Right = generate(mid+1, right)

		return root
	}

	return generate(0, n-1)
}
