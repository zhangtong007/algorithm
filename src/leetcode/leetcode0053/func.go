package leetcode53

// 最大子数组和

// 方法1： 前缀和思想
func maxSubArray1(nums []int) int {
	n := len(nums)
	// 求取前缀和数组
	sums := make([]int, n+1)
	sums[0] = 0
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// 字段和最大：sums[i] - sums[i-l] = max
	// 由表达式分析，要使得max最大，只需要sums[i-l]最小
	// 题目给定返回 -10^5 --- 10^5
	max := -100000
	min := sums[0]
	for i := 1; i <= n; i++ {
		max = maxNum(max, sums[i]-min)
		min = minNum(min, sums[i])
	}
	return max
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法二：贪心
/* nums[i] = {
		nums[i] + nums[i-1], nums[i-1] > 0;  // 解释，如果第i个数前面的数带来正向收益，则接受，否则拒绝。
		nums[i] = nums[i]
}
dp[i] 表示0-i的子数组最大子序和
*/
func maxSubArray2(nums []int) int {
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		// 如果是递增的，相加
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		max = maxNum(max, nums[i])
	}
	return max
}
