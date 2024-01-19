package leetcode300

func TestDemo(nums []int) int {
	return lengthOfLIS(nums)
}

// 动态规划
// dp[i] 表示以第i个元素结尾的最长子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxVal := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
