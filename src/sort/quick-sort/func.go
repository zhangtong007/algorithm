package quicksort

func TestDemo(nums []int) []int {
	return quickSort(nums)
}

// 经过排序之后，返回nums的升序数组
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 随机因子，取第一个数即可
	mid := nums[0]
	// 根据快排原则，将小于等于mid的放在mid的左边，记为left集合
	// 将大于mid的放在mid右边，记为right集合
	left, right := []int{}, []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= mid {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	// 将所得到的左右元素继续进行递归快排(分治思想)
	left = quickSort(left)
	right = quickSort(right)
	// 分治之后得到有序的左右子串，合并即可 [left, mid, right]
	return append(left, append([]int{mid}, right...)...)
}
