package leetcode130

var (
	direct_x = []int{1, 0, -1, 0}
	direct_y = []int{0, -1, 0, 1}
)

func solve(board [][]byte) {
	// 从边界上的O作深搜，记忆化不能被X的节点
	for i := range board {
		for j := range board[i] {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) &&
				board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	// 修改
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'N' {
				continue
			}
			board[i][j] = 'X'
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'N' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'N'
	for idx := 0; idx < 4; idx++ {
		dfs(board, i+direct_x[idx], j+direct_y[idx])
	}
}
