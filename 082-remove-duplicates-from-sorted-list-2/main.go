package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	head = nonDuplicatedNode(head)
	current := head
	for current != nil {
		next := nonDuplicatedNode(current.Next)
		current.Next = next
		current = next
	}

	return head
}

func nonDuplicatedNode(node *ListNode) *ListNode {
	for node != nil && node.Next != nil && node.Val == node.Next.Val {
		next := node.Next
		for next.Next != nil && next.Val == next.Next.Val {
			next = next.Next
		}
		node = next.Next
	}
	return node
}
