package leetcode1074

// 不熟悉
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// 枚举上边界
	ans := 0
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		// 枚举下边界
		for _, row := range matrix[i:] {
			for c, v := range row {
				sum[c] += v
			}
			ans += subArraySum(sum, target)
		}
	}
	return ans
}

func subArraySum(sum []int, target int) int {
	mp := map[int]int{0: 1}
	ans := 0
	for i, pre := 0, 0; i < len(sum); i++ {
		pre += sum[i]
		if _, exist := mp[pre-target]; exist {
			ans += mp[pre-target]
		}
		mp[pre]++
	}
	return ans
}
