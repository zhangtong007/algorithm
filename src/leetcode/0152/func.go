package leetcode0152

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max(max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
		dpMin[i] = min(min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
	}
	m := dpMax[0]
	for i := range dpMax {
		m = max(m, dpMax[i])
	}
	return m
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
