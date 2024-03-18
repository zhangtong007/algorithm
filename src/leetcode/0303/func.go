package leetcode303

// 前缀和
type NumArray struct {
	nums      []int
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return NumArray{
		nums:      nums,
		prefixSum: prefixSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right] - this.prefixSum[left] + this.nums[left]
}
