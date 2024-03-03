package leetcode123

// 套用188题通用模板，将最大次数k替换为2即可
func maxProfit(prices []int) int {
	return maxProfitGeneral(2, prices)
}

// 动态规划
// 分析可变状态：dp[i][j][m] 表示前i支股票在第i处位置是持有股票j(0, 1股)，且交易了m笔交易的最大利润
// 每个状态的决断点：买股票、卖股票、不买也不卖
// 分析三个状态的转移易得：
/*
    买股票：dp[i][1][m] <- dp[i-1][0][m-1] - prices[i]
	买股票: dp[i][0][m] <- dp[i-1][1][m] + prices[i]
	啥也不干: dp[i][j][m] <- dp[i-1][j][m]
*/
// 初始状态:dp[0][0][0] = 0, else: min
func maxProfitGeneral(k int, prices []int) int {
	n := len(prices)
	// prices 处理
	prices = append([]int{0}, prices...)
	// 初始化
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for m := range dp[i][j] {
				dp[i][j][m] = -1e9
			}
		}
	}
	dp[0][0][0] = 0
	// 决断
	for i := 1; i < len(dp); i++ {
		for j := range dp[i] {
			for k := range dp[i][j] {
				//
				if k > 0 {
					dp[i][1][k] = max(dp[i][1][k], dp[i-1][0][k-1]-prices[i])
				}
				// 卖
				dp[i][0][k] = max(dp[i][0][k], dp[i-1][1][k]+prices[i])
				// 啥也不干
				dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k])
			}
		}
	}
	// 返回最终状态，题意最多卖k次，所以返回0-k次的最大值
	ret := int(-1e9)
	for m := 0; m <= k; m++ {
		if dp[n][0][m] > ret {
			ret = dp[n][0][m]
		}
	}
	return int(ret)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
