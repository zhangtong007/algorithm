package leetcode2368

func reachableNodes(n int, edges [][]int, restricted []int) int {
	// 1. 建图, 出边数组
	toArr := make([][]int, n)
	for _, edge := range edges {
		toArr[edge[0]] = append(toArr[edge[0]], edge[1])
		toArr[edge[1]] = append(toArr[edge[1]], edge[0])
	}
	// 2. 记录受限节点
	restrictedBool := make([]bool, n)
	for _, v := range restricted {
		restrictedBool[v] = true
	}
	// 3. 广搜，直到不可达
	length := 0
	queue := []int{0}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		length++
		node := queue[0]
		visited[node] = true
		queue = queue[1:]
		for _, v := range toArr[node] {
			if restrictedBool[v] || visited[v] {
				continue
			}
			queue = append(queue, v)
		}
	}
	return length
}
