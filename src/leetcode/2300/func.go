package leetcode2300

import "sort"

// 二分potions下界即可
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	n := len(potions)
	for i, v := range spells {
		l, r := 0, n
		for l < r {
			mid := (l + r) / 2
			if v*potions[mid] >= int(success) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ans[i] = n - r
	}
	return ans
}
