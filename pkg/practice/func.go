package practice

var (
	direction = [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
)

func shortestBridge(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	arr := make([][]int, 0, 16)
	ret := -1
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	flag := true
	for i := 0; i < row && flag; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, &arr, visited)
				flag = false
				break
			}
		}
	}
	for len(arr) != 0 {
		num := len(arr)
		ret++
		for i := 0; i < num; i++ {
			nodes := arr[i]
			for j := 0; j < 4; j++ {
				nx := nodes[0] + direction[j][0]
				ny := nodes[1] + direction[j][1]
				if nx < 0 || nx >= row || ny < 0 || ny >= col || visited[nx][ny] {
					continue
				}
				if grid[nx][ny] == 1 {
					return ret
				}
				visited[nx][ny] = true
				arr = append(arr, []int{nx, ny})
			}
		}
	}
	return ret
}

func dfs(grid [][]int, i, j int, arr *[][]int, visited [][]bool) {
	if i < 0 || i > len(grid) || j < 0 || j > len(grid[0]) || visited[i][j] || grid[i][j] != 1 {
		return
	}
	visited[i][j] = true
	*arr = append(*arr, []int{i, j})
	dfs(grid, i-1, j, arr, visited)
	dfs(grid, i+1, j, arr, visited)
	dfs(grid, i, j-1, arr, visited)
	dfs(grid, i, j+1, arr, visited)
}
