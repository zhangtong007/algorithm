package leetcode1793

func maximumScore(nums []int, k int) int {
	var (
		n    = len(nums)
		ans  = nums[k]
		l, r = k - 1, k + 1
	)
	for i := nums[k]; ; i-- {
		// 左端点移动
		for l >= 0 && nums[l] >= i {
			l--
		}
		// 右端点移动
		for r < n && nums[r] >= i {
			r++
		}
		// 更新答案
		// 如果i不在nums里面，l和r是不会发生移动的
		ans = max(ans, (r-l-1)*i)
		// 遍历结束则退出循环
		if l == -1 && r == n {
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
