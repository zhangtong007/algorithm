package leetcode90

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	chosen := []int{}
	recur(false, &ans, chosen, nums, 0)
	return ans
}

func recur(preIsSelected bool, ans *[][]int, chosen, nums []int, pos int) {
	if pos == len(nums) {
		*ans = append(*ans, append([]int(nil), chosen...))
		return
	}

	recur(false, ans, chosen, nums, pos+1)
	// 如果要将当前作选择，需要判断上一个，如果上一个和当前重复，且上一个未被选择，那么当前不应该被选择
	if !preIsSelected && pos > 0 && nums[pos] == nums[pos-1] {
		return
	}
	recur(true, ans, append(chosen, nums[pos]), nums, pos+1)
}
