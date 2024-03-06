package leetcode416

// 转化题意：即在n个数中选择小于n-1个数，使其和为sum/2。0-1背包问题
func canPartition(nums []int) bool {
	n := len(nums)
	nums = append(make([]int, 1, n+1), nums...)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
