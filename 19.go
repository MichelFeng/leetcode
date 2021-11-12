package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := ListNode{Val: 1}
	// n2 := ListNode{Val: 2}
	// n3 := ListNode{Val: 3}
	// n4 := ListNode{Val: 4}
	// n5 := ListNode{Val: 5}
	// n1.Next = &n2
	// n2.Next = &n3
	// n3.Next = &n4
	// n4.Next = &n5

	res := removeNthFromEnd(&n1, 1)
	fmt.Println(res.Val)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
    first, second := head, dummy
    for i := 0; i < n; i++ {
        first = first.Next
    }
    for ; first != nil; first = first.Next {
        second = second.Next
    }
    second.Next = second.Next.Next
    return dummy.Next
}
