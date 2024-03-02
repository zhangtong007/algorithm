package leetcode56

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		// 左端点大于前一个的右端点，不包含，不作处理
		if intervals[i][0] > intervals[i-1][1] {
			continue
		}
		// 右端点小于等于前一个的右端点，被前一个包含，删除
		if intervals[i][1] <= intervals[i-1][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
			// 删除之后原地停留，扫描下一个
			i--
			continue
		}
		// 右端点大于前一个左端点，则交叉。如[1,3],[2,4] -> [1,4]
		// 合并
		intervals[i-1][1] = intervals[i][1]
		// 删除
		intervals = append(intervals[:i], intervals[i+1:]...)
		// 删除之后原地停留，扫描下一个
		i--
	}
	return intervals
}
