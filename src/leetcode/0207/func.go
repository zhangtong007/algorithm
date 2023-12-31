package leetcode207

// 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构造出边数组
	to := make([][]int, numCourses)
	// degree 统计各点入度
	degree := make([]int, numCourses)
	for _, edge := range prerequisites {
		to[edge[1]] = append(to[edge[1]], edge[0])
		degree[edge[0]]++
	}
	// 入队处理
	// 所有度为0的点才需要入队
	queue := make([]int, 0, 6)
	// visited 记录被访问点的数量
	visited := 0
	for v := range degree {
		if degree[v] == 0 {
			queue = append(queue, v)
			visited++
		}
	}
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]
		for _, v := range to[val] {
			degree[v]--
			if degree[v] == 0 {
				queue = append(queue, v)
				visited++
			}
		}
	}
	return visited == numCourses
}
