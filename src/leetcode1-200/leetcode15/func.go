package leetcode15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	ans := [][]int{}
	// 固定左端，以左端往右，做一次两数之和(双指针)
	for i := 0; i < n; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		start := i + 1
		for j := start; j < k; j++ {
			// 去重复
			if j > start && nums[j] == nums[j-1] {
				continue
			}
			// 题目要求 三数之和为0，所以判断第三数为 0-nums[i]
			for j < k && nums[j]+nums[k] > -nums[i] {
				k--
			}
			if j < k && nums[j]+nums[k] == -nums[i] {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
