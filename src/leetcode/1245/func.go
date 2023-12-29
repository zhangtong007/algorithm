package leetcode1245

// 树的直径

/*
树的直径定义： 数上任意两个节点之间的距离(边数)的最大值
求解步骤：
1. 任意一个点A出发，寻找距离点A最远的点P  (P为直径的一个端点)
2. 从P出发，寻找距离点P最远的点Q  经过的距离即为直径
*/
func treeDiameter(edges [][]int) int {
	if len(edges) == 0 {
		return 0
	}
	// map 记录所有点的到达情况 // 也可考虑出边数组
	to := map[int][]int{}
	for _, edge := range edges {
		to[edge[0]] = append(to[edge[0]], edge[1])
		to[edge[1]] = append(to[edge[1]], edge[0])
	}

	// BFS  确定出发点 edges[0][0]，寻找点P
	p, _ := bfs(to, edges[0][0])
	// BFS 确定出发点 P，寻找点Q
	_, path := bfs(to, p)
	return path
}

// node 为最远端点(可能不唯一，返回其一)， path为路径长度
func bfs(to map[int][]int, start int) (node, path int) {
	queue := []int{start}
	path = -1
	// 记忆化搜索，防止路径重复
	used := map[int]bool{}
	for len(queue) > 0 {
		node = queue[0]
		path++
		newQueue := []int{}
		for _, start := range queue {
			used[start] = true
			for _, toNode := range to[start] {
				if used[toNode] {
					continue
				}
				newQueue = append(newQueue, toNode)
			}
		}
		queue = newQueue
	}
	return
}
