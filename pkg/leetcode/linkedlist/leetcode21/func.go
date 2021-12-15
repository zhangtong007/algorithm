package leetcode21

import . "core/leetcode/datastruct"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	merge(l1, l2, head)
	return head
}

func merge(l1 *ListNode, l2 *ListNode, head *ListNode) {
	for {
		if l1 == nil {
			head.Next = l2
			return
		}
		if l2 == nil {
			head.Next = l1
			return
		}
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
			head = head.Next
		} else {
			head.Next, l2 = l2, l2.Next
			head = head.Next
		}
	}
}
