package leetcode142

import . "core/pkg/leetcode/datastruct"

const max = 100000000

func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		if head.Val == max {
			return head
		}
		head.Val = max
		head = head.Next
	}
	return nil
}
