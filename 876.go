package main

import "fmt"

func main() {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	res := middleNode(&n1)
	fmt.Println(res.Val)

	n6 := ListNode{Val: 6}
	n5.Next = &n6
	res = middleNode(&n1)
	fmt.Println(res.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	front := head
	for front != nil && front.Next != nil {
		front = front.Next
		head = head.Next
		if front == nil {
			break
		}
		front = front.Next
	}
	return head
}
