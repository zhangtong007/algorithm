package leetcode70

func climbStairs(n int) int {
	// dp[i] 表示前i+1阶台阶能到达的方法数
	// dp[i] = dp[i-2] + dp[i-1]
	dp := make([]int, n+1) // 这里的n+1没有特殊含义，只是为了在n=1时初始状态dp[1]不发生越界
	// 初始状态，跳到第1阶和跳到第二阶分别有，1和2种方法
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n-1]
}
