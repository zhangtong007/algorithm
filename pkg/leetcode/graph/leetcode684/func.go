package leetcode684

// 冗余连接

func findRedundantConnection(edges [][]int) []int {
	// 初始化出边数组
	// 找到n
	n := 0
	for _, arr := range edges {
		n = max(n, max(arr[0], arr[1]))
	}
	// 初始化出边数组
	to := make([][]int, n+1)
	// 构造出边数组
	// 深度优先遍历出边数组，找环
	for _, arr := range edges {
		visited := make([]bool, n+1)
		to[arr[0]] = append(to[arr[0]], arr[1])
		to[arr[1]] = append(to[arr[1]], arr[0])
		// 每次加边的时候做一次判断是否有环，有环则该边为答案
		if dfs(to, visited, arr[0], 0) {
			return arr
		}
	}
	return nil
}

func dfs(to [][]int, visited []bool, val, parent int) bool {
	if visited[val] {
		return true
	}
	visited[val] = true
	for _, v := range to[val] {
		if v == parent {
			continue
		}
		if dfs(to, visited, v, val) {
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
