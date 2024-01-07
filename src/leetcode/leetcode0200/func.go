package leetcode200

func numIslands(grid [][]byte) int {
	ret := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ret++
				dfs(grid, i, j)
			}
		}
	}
	return ret
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
