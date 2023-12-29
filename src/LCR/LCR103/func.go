package lcr103

func TestDemo(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = 1e10
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] >= 1e10 {
		dp[amount] = -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
