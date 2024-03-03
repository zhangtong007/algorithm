package leetcode122

// 解法一：贪心
func maxProfit1(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	ret := 0
	// cur，记录买入的初始值
	cur := 0
	// flag，true表示买入，false表示未买入
	flag := false
	for i := 0; i < len(prices); i++ {
		// 如果是最后一个了
		if i == len(prices)-1 {
			// 判断是否有买入，有买入就将他卖出
			if flag {
				ret += prices[i] - cur
			}
			// 退出循环
			break
		}
		// 如果有买入，判断下一个是否小于当前，如果小于，说明会跌，赶紧卖出
		if flag && prices[i] > prices[i+1] {
			ret += prices[i] - cur
			flag = false
		}
		// 如果没有买入，在下一次涨的时候买入，即next > cur
		if !flag && prices[i] < prices[i+1] {
			cur = prices[i]
			flag = true
		}
	}
	return ret
}
