package leetcode213

// 围成一圈，，那么在简单打家劫舍198题的基础上多了一个限制条件，第一家和最后一家收尾不能同时选
// 对于线性动态规划而言，是单向线性的，从起点至终点，所以直接解无法线性动归解该题
// 提取条件，两次动归
// 既然第一家和第n家不能够同时选。那么将情景拆分求解也可以
// 1. 求解[2-n]的最大值max1
// 2. 求解[1-n-1]的最大值max2
// 3. 注意不重不漏，如果n == 1 直接返回
// 4. return max(max1, max2)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max1 := simpleRob(nums[1:])
	max2 := simpleRob(nums[0 : len(nums)-1])
	return max(max1, max2)
}

// 198动归直接复制过来
func simpleRob(nums []int) int {
	n := len(nums)
	// 处理下标为1
	nums = append(make([]int, 1, n+1), nums...)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			// 求最大值，错误状态定义最小
			dp[i][j] = -1e9
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][1] = dp[i-1][0] + nums[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[n][0], dp[n][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
