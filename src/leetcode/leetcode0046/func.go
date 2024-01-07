package leetcode46

func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	arrangement := []int{}
	used := make([]bool, n)
	recur(&ans, arrangement, nums, used, n, 0)
	return ans
}

func recur(ans *[][]int, arrangement, nums []int, used []bool, n, pos int) {
	// 到达递归边界
	if pos == n {
		bak := make([]int, len(arrangement))
		copy(bak, arrangement)
		*ans = append(*ans, bak)
		return
	}
	// process
	for i := 0; i < n; i++ {
		// 如果i未被使用
		if !used[i] {
			// 如果i未被使用，则当前位置可选择i
			used[i] = true
			recur(ans, append(arrangement, nums[i]), nums, used, n, pos+1)
			// 还原
			used[i] = false
		}
	}
}
