package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if _, ok := seen[node]; ok {
			return node
		}
		seen[node] = true
		node = node.Next
	}
	return nil
}
