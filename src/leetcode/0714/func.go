package leetcode714

// 动态规划 122题的基础上增加手续费
// 定义状态：dp[i][j]表示第i天持有j股(0,1)
/*
    买：
	dp[i][1] <- dp[i-1][0] - prices[i] - fee
	卖:
	dp[i][0] <- dp[i-1][1] + prices[i]
	不买不卖:
	dp[i][j] <- dp[i-1][j]
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	// 处理prices
	prices = append([]int{0}, prices...)
	// 初始化状态
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = -1e9
		}
	}
	dp[0][0] = 0
	// 决断：状态转移
	for i := 1; i < len(dp); i++ {
		// 买
		dp[i][1] = max(dp[i][1], dp[i-1][0]-prices[i]-fee)
		// 卖
		dp[i][0] = max(dp[i][0], dp[i-1][1]+prices[i])
		for j := range dp[i] {
			// 不买不卖
			dp[i][j] = max(dp[i][j], dp[i-1][j])
		}
	}
	return dp[n][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
