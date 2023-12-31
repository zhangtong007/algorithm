package bfs

/*
	BFS方式：
	作拓扑排序，从入度为0的点出发，作bfs
*/
var (
	m, n     int
	to       [][]int
	degree   []int
	distance []int
)

var (
	direct_x = []int{-1, 0, 0, 1}
	direct_y = []int{0, -1, 1, 0}
)

func addEdge(out, into int) {
	to[out] = append(to[out], into)
	degree[into]++
}

func numsNode(i, j int) int {
	return i*m + j
}

//test
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func valid(i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func longestIncreasingPath(matrix [][]int) int {
	// 建图
	m, n = len(matrix), len(matrix[0])
	to = make([][]int, m*n)
	degree = make([]int, m*n)
	distance = make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 遍历四个方向
			for k := 0; k < 4; k++ {
				x := i + direct_x[k]
				y := j + direct_y[k]
				if valid(x, y) && matrix[i][j] < matrix[x][y] {
					addEdge(numsNode(i, j), numsNode(x, y))
				}
			}
		}
	}

	// 找到所有度为0的点，作深搜起点
	queue := []int{}
	for i, v := range degree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		out := queue[0]
		queue = queue[1:]
		for _, into := range to[out] {
			degree[into]--
			distance[into] = max(distance[into], distance[out]+1)
			if degree[into] == 0 {
				queue = append(queue, into)
			}
		}
	}
	var max int
	for _, v := range distance {
		if max < v {
			max = v
		}
	}
	return max
}
