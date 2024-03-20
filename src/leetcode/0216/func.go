package leetcode216

// 递归搜索即可
func combinationSum3(k int, n int) [][]int {
	if k > n {
		return nil
	}
	ans := [][]int{}
	chosen := []int{}
	var recursion func(k, n, i int)
	recursion = func(k, n, i int) {
		// 边界
		if k == len(chosen) || i == 10 {
			// 答案是否符合选择
			if k == len(chosen) && n == 0 {
				ans = append(ans, append(make([]int, 0, len(chosen)), chosen...))
			}
			return
		}
		// 选择i
		chosen = append(chosen, i)
		recursion(k, n-i, i+1)
		// 不选择i
		// 状态恢复
		chosen = chosen[:len(chosen)-1]
		recursion(k, n, i+1)
	}
	// 从1开始递归
	recursion(k, n, 1)
	return ans
}
