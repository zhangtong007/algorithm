package leetcode1

// 利用map来解决该题
func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i := range nums {
		other := target - nums[i]
		if idx, exist := indexes[other]; exist {
			return []int{idx, i}
		}
		indexes[nums[i]] = i
	}
	return nil
}
