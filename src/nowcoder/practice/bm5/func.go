package bm5

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoList(lists[0], lists[1])
	}
	return mergeTwoList(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoList(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoList(list1, list2.Next)
	return list2
}
