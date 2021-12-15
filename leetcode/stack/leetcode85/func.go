package leetcode85

// 借助84题进行解答

type Rect struct {
	width  int
	height int
}

func largestRectangleArea(heights []byte) int {
	ret := 0
	heights = append(heights, 0)
	stack := make([]Rect, 0)
	for _, height := range heights {
		sumWidth := 0
		// 如果栈顶之前高度大于当前高度，单调性破坏，确定栈顶扩展范围，需要更新
		for len(stack) > 0 && stack[len(stack)-1].height >= int(height)-'0' {
			sumWidth += stack[len(stack)-1].width
			ret = max(ret, stack[len(stack)-1].height*sumWidth)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Rect{
			width:  sumWidth + 1,
			height: int(height) - '0',
		})
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	initArr(matrix)
	retArr := perRow(matrix)
	return findMax(retArr)
}

func findMax(retArr []int) int {
	max := retArr[0]
	for i := 1; i < len(retArr); i++ {
		if max < retArr[i] {
			max = retArr[i]
		}
	}
	return max
}

func perRow(matrix [][]byte) []int {
	retArr := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		retArr[i] = largestRectangleArea(matrix[i])
	}
	return retArr
}

func initArr(matrix [][]byte) {
	for j := 0; j < len(matrix[0]); j++ {
		for i := 1; i < len(matrix); i++ {
			if matrix[i][j] == '0' {
				continue
			}
			matrix[i][j] += matrix[i-1][j] - '0'
		}
	}
}
