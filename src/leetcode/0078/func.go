package leetcode78

import "sort"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	chosen := []int{}
	// 从0开始递归
	recur(&ans, chosen, nums, 0)
	for i := range ans {
		sort.Slice(ans[i], func(m, n int) bool {
			return ans[i][m] < ans[i][n]
		})
	}
	return ans
}

/*
ans, chosen 递归修改相关值，传指针，因为值是全局相关
nums 是源数组，i表示第i个数据
递归子集抉择策略：
每个数，选或者不选，是一个策略
*/
func recur(ans *[][]int, chosen []int, nums []int, i int) {
	// 到达递归边界
	if i == len(nums) {
		// 更新答案
		// 做拷贝
		arr := make([]int, len(chosen))
		copy(arr, chosen)
		*ans = append(*ans, arr)
		return
	}
	// 具体操作
	// 如果i不选择
	recur(ans, chosen, nums, i+1)
	// 如果i选择，将i记录到chosen里
	chosen = append(chosen, nums[i])
	recur(ans, chosen, nums, i+1)

	// 还原现场
	chosen = chosen[:len(chosen)-1]
}

func recur1(ans *[][]int, chosen []int, nums []int, i int) {
	// 到达递归边界
	if i == len(nums) {
		// 更新答案
		// 坑点！！！做笔记！！！
		arr := make([]int, len(chosen))
		copy(arr, chosen)
		*ans = append(*ans, arr)
		return
	}
	// 具体操作
	// 如果i不选择
	recur(ans, chosen, nums, i+1)
	// 将选择直接带入，无需再次还原
	recur(ans, append(chosen, nums[i]), nums, i+1)
}
