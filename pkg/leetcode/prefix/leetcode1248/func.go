package leetcode1248

// 计算前缀和，偶数加0，奇数加1，转化为子段和为k有多少个
func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	// 求取前缀和, sums[i] 表示第1个到第i个的和
	sums := make([]int, n+1)
	sums[0] = 0
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]&1
	}
	// 求取到前缀和数组之后，做差分匹配，求取字段和为k的个数
	// 字段和求取  sum[l, r] = sum[r] - sum[l-1]
	// sum[l, r] == k =>  sum[r] - sum[l-1] = k => sum[r] - k == sum[l-1]
	// 换元 二元变量转化为一元变量
	ans := 0
	// count 用于求取另一元素s[l-1]
	count := make([]int, n+1)
	// count[sums[0]]++  sums[0] = sums[1-1] 也是一种情况
	count[0] = 1
	for i := 1; i <= n; i++ {
		// 注意数组越界
		if sums[i]-k >= 0 {
			// sums[i]-k == sums[l-1] 只要该值存在，就加
			ans += count[sums[i]-k]
		}
		count[sums[i]]++
	}
	return ans
}
