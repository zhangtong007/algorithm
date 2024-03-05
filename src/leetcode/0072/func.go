package leetcode72

/*
动态规划

	定义dp[i][j]表示word1前i个字符转为为word2前j个字符需要的最少操作数
	1. 初始状态: dp[i][0] = i, dp[0][j] = j
	2. 决断：由word1[i] -> word2[j] 有插入、删除、替换、不变四类决断
	    插入： dp[i][j] = dp[i-1][j] + 1
	    删除:  dp[i][j] = dp[i][j-1] + 1
	    替换:  dp[i][j] = dp[i-1][j-1] + 1
	    不变:  dp[i][j] = dp[i-1][j-1]
	    其中替换和不变的条件为：word1[i]==word2[j]，则不变。否则替换
	3. 状态转移方程：
	    if word1[i] == word2[j] opts == 1 else opts = 0
	    dp[i][j] = min(min(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+opts)
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 1. 索引移到第一位
	word1 = " " + word1
	word2 = " " + word2
	// 2. 定义dp
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 3. 初始化dp
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	// 4. 遍历状态，最优子结构推导
	var opts int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i] == word2[j] {
				opts = 0
			} else {
				opts = 1
			}
			dp[i][j] = min(min(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+opts)
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
