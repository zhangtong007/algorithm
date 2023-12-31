package leetcode684

// 冗余连接
func findRedundantConnection(edges [][]int) []int {
	var n int
	// 求解有多少个节点
	for _, arr := range edges {
		n = max(n, max(arr[0], arr[1]))
	}
	to := make([][]int, n+1)
	// 最直观做法，构造出边数组，每加一条边做一次dfs，如果找到环，则为答案
	for _, arr := range edges {
		to[arr[0]] = append(to[arr[0]], arr[1])
		to[arr[1]] = append(to[arr[1]], arr[0])
		// 从arr{0}出发，作bfs，看是否有环
		if dfs(to, make([]bool, n+1), arr[0], 0) {
			return arr
		}
	}
	// 题目必有环，作占位返回
	return []int{}
}

func dfs(to [][]int, visited []bool, node, pre int) bool {
	if visited[node] {
		return true
	}
	visited[node] = true
	// 访问node的出边，作bfs
	for _, v := range to[node] {
		// 前一个点没必要遍历
		if v == pre {
			continue
		}
		if dfs(to, visited, v, node) {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
