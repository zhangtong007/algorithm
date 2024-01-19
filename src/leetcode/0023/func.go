package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n == 2 {
		return process(lists[0], lists[1])
	}
	return mergeKLists([]*ListNode{mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:])})
}

// 将两个有序链表升序合并
func process(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = process(list1.Next, list2)
		return list1
	}
	list2.Next = process(list1, list2.Next)
	return list2
}
