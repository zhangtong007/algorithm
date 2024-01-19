package leetcode304

// 二维前缀和
type NumMatrix struct {
	prefixes [][]int
}

func calPrefixes(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	prefixes := make([][]int, 0, row+1)
	prefixes = append(prefixes, make([]int, col+1))
	for i := 1; i <= row; i++ {
		prefixes = append(prefixes, make([]int, col+1))
		for j := 1; j <= col; j++ {
			prefixes[i][j] = prefixes[i-1][j] + prefixes[i][j-1] - prefixes[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return prefixes
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		prefixes: calPrefixes(matrix),
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 前缀和设计 第i个从i-1开始  所以题目的第i应该从i+1开始，需要做转化
	row1, col1, row2, col2 = row1+1, col1+1, row2+1, col2+1
	return this.prefixes[row2][col2] - this.prefixes[row2][col1-1] - this.prefixes[row1-1][col2] + this.prefixes[row1-1][col1-1]
}
