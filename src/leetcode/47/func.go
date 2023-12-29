package leetcode47

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	arrangemant := []int{}
	used := make([]bool, len(nums))
	recur(&ans, arrangemant, nums, used, 0)
	return ans
}

func recur(ans *[][]int, arrangemant, nums []int, used []bool, pos int) {
	// 递归边界
	if pos == len(nums) {
		bak := make([]int, len(arrangemant))
		copy(bak, arrangemant)
		*ans = append(*ans, bak)
		return
	}

	// process
	for i := range nums {
		if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] {
			continue
		}
		used[i] = true
		arrangemant = append(arrangemant, nums[i])
		recur(ans, arrangemant, nums, used, pos+1)
		arrangemant = arrangemant[:len(arrangemant)-1]
		used[i] = false
	}
}
