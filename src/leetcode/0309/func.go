package leetcode309

// 动态规划 122题的基础上增加冷冻期
// 定义状态：dp[i][j][state]表示第i天持有j股(0,1),state(0 不在冷冻期, 1 冷冻期)时的最大利润
/*
    买：未持有股票且未在冷冻期，才能买。
	dp[i][1][0] <- dp[i-1][0][0] - prices[i]
	卖: 持有股票才能卖，持有时不可能在冷冻期
	dp[i][0][1] <- dp[i-1][1][0] + prices[i]
	不买不卖: 冷冻期会接触，所以state只会是0
	dp[i][j][0] <- dp[i][j][state]
*/
func maxProfit(prices []int) int {
	n := len(prices)
	// 处理prices
	prices = append([]int{0}, prices...)
	// 初始化状态
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			for state := range dp[i][j] {
				dp[i][j][state] = -1e9
			}
		}
	}
	dp[0][0][0] = 0
	// 决断：状态转移
	for i := 1; i < len(dp); i++ {
		// 买
		dp[i][1][0] = max(dp[i][1][0], dp[i-1][0][0]-prices[i])
		// 卖
		dp[i][0][1] = max(dp[i][0][1], dp[i-1][1][0]+prices[i])
		for j := range dp[i] {
			for state := range dp[i][j] {
				// 不买不卖
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][state])
			}
		}
	}
	return max(dp[n][0][0], dp[n][0][1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
