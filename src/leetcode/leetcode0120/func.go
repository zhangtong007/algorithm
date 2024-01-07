package leetcode120

import "math"

// 动态规划
// 找寻状态、每一个点的位置作为移动状态
// dp[i][j] 表示点{i,j}的最小路径和
// 		易得：dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
//		特殊状态：
//			1. 第一列的来源只能是上一层的第一列，即dp[i][0] = dp[i-1][0] + triangle[i][j]
//			2. 最后一列的来源只能是上一层的最后一个，即dp[i][last] = dp[i-1][last-1]+ triangle[i][j], last为第i行的最后一列
// 			3. dp[0][0] = triangle[0][0]
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	dp[0] = []int{triangle[0][0]}
	for i := 1; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			switch j {
			case 0:
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			case len(triangle[i]) - 1:
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			default:
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	// 遍历最后一层，返回最小路径和
	ret := math.MaxInt32
	for _, v := range dp[m-1] {
		ret = min(ret, v)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
