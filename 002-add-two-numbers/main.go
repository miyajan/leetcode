package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	total := ListNode{}
	current := &total
	var prev *ListNode
	carryOver := 0
	for l1 != nil || l2 != nil || carryOver == 1 {
		current.Val = carryOver
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}
		if current.Val >= 10 {
			carryOver = 1
			current.Val = current.Val - 10
		} else {
			carryOver = 0
		}
		if prev != nil {
			prev.Next = current
		}
		prev = current
		current = &ListNode{}
	}

	return &total
}
