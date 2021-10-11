package leetcode25

import "core/leetcode/datastruct"

func reverseKGroup(head *datastruct.ListNode, k int) *datastruct.ListNode {
	protected := &datastruct.ListNode{
		Val:  -1,
		Next: head,
	}
	last := protected
	var (
		groupEnd      *datastruct.ListNode
		nextGroupHead *datastruct.ListNode
	)
	for head != nil {
		// 1. 分组(往后走k-1步, 找到1组)
		//     一组开头head，结尾groupEnd
		groupEnd = getEnd(head, k)
		if groupEnd == nil {
			break
		}
		// 2. 一组内部(head 到 end) 反转
		nextGroupHead = groupEnd.Next
		reverse(head, nextGroupHead)
		// 3. 更新每组跟前一组，后一组之间的边(组间相连)
		last.Next, head.Next = groupEnd, nextGroupHead

		// 将last，head指针更新到下一组
		last, head = head, nextGroupHead
	}
	return protected.Next
}

// 拿到每一组的最末尾值
func getEnd(head *datastruct.ListNode, k int) *datastruct.ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	// 说明没有最后一组了
	return nil
}

// 反转链表，在stop停止
func reverse(head *datastruct.ListNode, stop *datastruct.ListNode) {
	last := head
	for head != stop {
		last, head, head.Next = head, head.Next, last
	}
}
