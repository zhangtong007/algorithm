package leetcode322

func Test(coins []int, amount int) int {
	return coinChange(coins, amount)
}

// 动态规划
// dp[i] 表示从总金额为i的最少硬币数
// dp[i] = min(dp[i-coin]..., coin ∈ coins && i > coin)
// dp[0] = 1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i > coin {
				dp[i] = min(dp[i], dp[i-coin])
			}
		}
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
