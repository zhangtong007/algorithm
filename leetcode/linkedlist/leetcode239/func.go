package leetcode239

// 滑动窗口最大值
// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	// 下标递增，值递减的队列
	queue := make([]int, 0, k)
	ans := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		// 删除出界的选项
		if len(queue) != 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		// 插入新选项i，维护单调性(值递减)
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 取队头，更新答案
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
