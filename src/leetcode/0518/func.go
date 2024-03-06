package leetcode518

// 完全背包问题
// n种资源，每种资源容量为coins[n]，求能达到总容量为amount的方案数
func change(amount int, coins []int) int {
	coins = append(make([]int, 1, len(coins)), coins...)
	f := make([]int, amount+1)
	f[0] = 1
	for i := 1; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] += f[j-coins[i]]
		}
	}
	return f[amount]
}
