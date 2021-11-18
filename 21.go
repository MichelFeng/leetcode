package main

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	mergeTwoLists(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		head *ListNode
		node *ListNode
	)
	for l1 != nil && l2 != nil {
		var val int
		if l1.Val < l2.Val {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}
		if node == nil {
			node = &ListNode{Val: val}
			head = node
		} else {
			newNode := &ListNode{Val: val}
			node.Next = newNode
			node = node.Next
		}
	}
	if l1 == nil {
		node.Next = l2
	} else {
		node.Next = l1
	}
	return head
}
