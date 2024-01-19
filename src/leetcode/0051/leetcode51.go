package leetcode51

// N皇后问题：棋盘上可以攻击：同一行、列，同一对角线上可以互相攻击
// 问题转化：每行选择一个位置放置一个“皇后”
// 需要做记忆化，如果列被使用，对角线被使用，则不选择

/*
对角线确定：y=kx+b  正对角线: k = 1, 副对角线: k=-1；即 b=y-x, b=y+x分别确认正对角线，反对角线
*/
var (
	colsUsed             = make(map[int]bool)
	positiveDiagonalUsed = make(map[int]bool)
	inverseDiagonalUsed  = make(map[int]bool)
)

func initBytes(n int) []byte {
	ret := make([]byte, n)
	for i := range ret {
		ret[i] = '.'
	}
	return ret
}

func solveNQueens(n int) [][]string {
	// 每个答案只记录对应行选择的位置(列号)，最后再将ret进行转化
	ret := [][]int{}
	dfs(&ret, make([]int, n), 0)
	// trance to answer
	ans := make([][]string, len(ret))
	for i, arr := range ret {
		strs := make([]string, n)
		for j, v := range arr {
			bytes := initBytes(n)
			bytes[v] = 'Q'
			strs[j] = string(bytes)
		}
		ans[i] = strs
	}
	return ans
}

func dfs(ret *[][]int, cols []int, pos int) {
	if pos == len(cols) {
		*ret = append(*ret, append([]int(nil), cols...))
		return
	}

	// pos 表示选择到第几个，在这儿表示行号
	for i := 0; i < len(cols); i++ {
		// 如果没有被使用过，则可选
		if colsUsed[i] || positiveDiagonalUsed[i-pos] || inverseDiagonalUsed[i+pos] {
			continue
		}
		cols[pos] = i
		// 记忆化
		colsUsed[i] = true
		positiveDiagonalUsed[i-pos] = true
		inverseDiagonalUsed[i+pos] = true
		dfs(ret, cols, pos+1)
		// 现场还原
		delete(colsUsed, i)
		delete(positiveDiagonalUsed, i-pos)
		delete(inverseDiagonalUsed, i+pos)
	}
}
