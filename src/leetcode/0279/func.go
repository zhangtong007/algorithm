package leetcode279

import "math"

// sqt(n)种数里面选择数，每种可以无限选择，和为n的最少数量是多少
// 完全背包问题
func numSquares(n int) int {
	// 转化以下模型，target变为目标，n变为可选数数量，nums为备选数类
	target := n
	n = int(math.Sqrt(float64(n)))
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = i * i
	}
	f := make([]int, target+1)
	for j := 1; j <= target; j++ {
		f[j] = 1e9
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := nums[i]; j <= target; j++ {
			f[j] = min(f[j], f[j-nums[i]]+1)
		}
	}
	return f[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
