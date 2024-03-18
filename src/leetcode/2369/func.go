package leetcode2369

// 动态规划
/*
	dp[i] = (dp[i-2] && nums[i-2] == nums[i-1]) ||
	(dp[i-3] &&
	((nums[i-3] == nums[i-2] && nums[i-3] == nums[i-1]) ||
	nums[i-3]+1 == nums[i-2] && nums[i-2]+1 == nums[i-1])))dp下标从0开始，return dp[n]
*/
func validPartition(nums []int) bool {
	n := len(nums)
	// 定义dp
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		if i >= 2 {
			dp[i] = dp[i-2] && nums[i-2] == nums[i-1]
		}
		if i >= 3 {
			dp[i] = dp[i] ||
				(dp[i-3] &&
					((nums[i-3] == nums[i-2] && nums[i-2] == nums[i-1]) ||
						(nums[i-3]+1 == nums[i-2] && nums[i-2]+1 == nums[i-1])))
		}
	}
	return dp[n]
}
