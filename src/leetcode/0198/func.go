package leetcode198

// dp[i][j] 表示前i间房，在第i间房状态j(0:没被偷，1:被偷)时的最大收益
/*
   决断：对于每间房i来说，要么偷，要么不偷
   // 偷：
   dp[i][1] = dp[i-1][0] + nums[i]
   // 不偷:
   dp[i][0] = max(dp[i-1][0], dp[i-1][1])
*/
func rob(nums []int) int {
	n := len(nums)
	// 处理下标为1
	nums = append(make([]int, 1, n+1), nums...)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			// 求最大值，错误状态定义最小
			dp[i][j] = -1e9
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][1] = dp[i-1][0] + nums[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[n][0], dp[n][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
