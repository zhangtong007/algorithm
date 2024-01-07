package leetcode673

// 动态规划
// dp[i] 存储(maxLen, cnt)，maxLen表示以i结尾的最长子序列的长度， cnt表示以i结尾的最长子序列的个数
type Pair struct {
	maxLen, cnt int
}

func findNumberOfLIS(nums []int) int {
	dp := make([]Pair, len(nums))
	maxPair := Pair{maxLen: 0, cnt: 0}
	for i := 0; i < len(nums); i++ {
		dp[i] = Pair{1, 1}
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 构成升序，优先判断新的长度，是否大于当前状态里面维护的长度，若大于，则更新
				if dp[j].maxLen+1 > dp[i].maxLen {
					dp[i].maxLen = dp[j].maxLen + 1
					dp[i].cnt = dp[j].cnt
					continue
				}
				// 构成升序，若长度相等，则方案叠加。更新
				if dp[j].maxLen+1 == dp[i].maxLen {
					dp[i].cnt += dp[j].cnt
					continue
				}
			}
		}
		if dp[i].maxLen > maxPair.maxLen {
			maxPair.maxLen = dp[i].maxLen
			maxPair.cnt = dp[i].cnt
			continue
		}
		if dp[i].maxLen == maxPair.maxLen {
			maxPair.cnt += dp[i].cnt
		}
	}
	return maxPair.cnt
}
