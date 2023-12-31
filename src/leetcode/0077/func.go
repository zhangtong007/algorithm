package leetcode77

func combine(n int, k int) [][]int {
	// n里面选k个，每个数有两个选择，选或者不选，递归
	ans := [][]int{}
	chosen := []int{}
	recur(&ans, chosen, n, k, 1)
	return ans
}

func recur(ans *[][]int, chosen []int, n, k, i int) {
	// 减枝，如果当前的选择已经大于K, 或者剩余的选择已经不够，就不用再继续了
	if len(chosen) > k || (n-i+1)+len(chosen) < k {
		return
	}
	// i 作为变化量
	if i == n+1 {
		if len(chosen) == k {
			bak := make([]int, len(chosen))
			copy(bak, chosen)
			*ans = append(*ans, bak)
		}
		return
	}
	// process
	recur(ans, chosen, n, k, i+1)
	recur(ans, append(chosen, i), n, k, i+1)
}
