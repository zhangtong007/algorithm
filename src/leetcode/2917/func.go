package leetcode2917

// 枚举每一个二进制位即可
func findKOr(nums []int, k int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			ans |= 1 << i
		}
	}
	return ans
}
