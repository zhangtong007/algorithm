package leetcode97

// dp[i][j]表示s1的前i个和s2的前j个能否与s3的前i+j个交错
// dp[i][j] =
//
//	dp[i-1][j] && s1[i-1] == s3[i+j-1] ||
//	dp[i][j-1] && s2[j-1] == s3[i+j-1]
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[n][m]
}
