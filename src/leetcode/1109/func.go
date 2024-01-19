package leetcode1109

// 使用差分
/* 差分思想：
对一个数组A而言：
计算A的差分数组B
计算B的前缀和数组C  C即是A

应用:
B[i] = A[i] - A[i-1]
如果对一个数组的第i到j 元素同时加一个数N或者减一个数N 时间复杂度(O(j-i+1))
如果利用差分数组，我们可以推出(以+为例)
比如 		A： {1, 2, 4, 2, 1}    ----> 2-4分别加4 ==> {1, 2, 8, 6, 5}
差分数组	B： {1, 1, 2, -2, -1}
2-4分别加4：B:  {1, 1, 6, -2, -1 | -5}
计算前缀和：B： {1, 2, 8, 6, 5}
最终可以只计算 B[i] +/- N  和 B[j+1] -/+ N
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	// 初始化都为0，所以差分数组全为0
	diff := make([]int, n)
	for _, arr := range bookings {
		start := arr[0] - 1
		end := arr[1] - 1
		num := arr[2]
		diff[start] += num
		if end+1 < n {
			diff[end+1] -= num
		}
	}
	// 对差分数组做前缀和，转化为前缀和数组
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}
