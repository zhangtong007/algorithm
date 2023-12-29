package leetcode11

// 双指针
func maxArea(height []int) int {
	n := len(height)
	// width := r-l
	ans := 0
	l, r := 0, n-1
	for l < r {
		ans = max(ans, (r-l)*min(height[r], height[l]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
