package leetcode1143

// 动态规划
// dp[i][j] 表示text1前i个元素和text2前j个元素的最长公共子序列
// if text[i] == text[j], dp[i][j] = dp[i-1][j-1] + 1
// if text[i] != text[j], dp[i][j] = max(dp[i-1][j], dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	// 优雅处理边界判断, 状态从1开始判定
	text1 = " " + text1
	text2 = " " + text2
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 状态转移
	for i := 1; i < m; i++ {
		dp = append(dp, make([]int, n))
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
