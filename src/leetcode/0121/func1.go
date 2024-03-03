package leetcode121

// 贪心+前缀最小值
func maxProfit1(prices []int) int {
	// 贪心
	// min: 当前已遍历过的最小价格
	// max: 当前的价格，减去已遍历过的最小价格所得利润的最大
	min, max := prices[0], 0
	for i := 0; i < len(prices); i++ {
		min = minFunc(min, prices[i])
		max = maxFunc(max, prices[i]-min)
	}
	return max
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
