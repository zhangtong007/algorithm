package leetcode697

// 题目转化为求出现最多元素的最左最右的距离
type element struct {
	// 数据出现次数
	num int
	// 数据最左索引
	left int
	// 数据最右索引
	right int
}

func findShortestSubArray(nums []int) int {
	numMap := make(map[int]*element)
	for i, v := range nums {
		if _, exist := numMap[v]; exist {
			numMap[v].num++
			numMap[v].right = i
			continue
		}
		numMap[v] = &element{
			num:  1,
			left: i,
		}
	}
	max := &element{num: 0}
	for _, e := range numMap {
		if max.num < e.num || (max.num != 1 && max.num == e.num && max.right-max.left > e.right-e.left) {
			max = e
		}
	}
	if max.num == 1 {
		return 1
	}
	return max.right - max.left + 1
}
