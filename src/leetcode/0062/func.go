package leetcode62

// 线性动态规划
// 分析：
//     1. 机器人每次只能向下或者向右移动=>那么每个位置只能来源于上方或者左方
//     2. dp[i][j]表示机器人到达i，j位置时的总路径数，易得到 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//     3. m行n列，最后一个位置返回dp[m-1][n-1]，下标以0开始
//     4. 初始状态：处于0,0位置，有且只有一种方法 dp[0][0]=1

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
