package leetcode84

type Rect struct {
	width  int
	height int
}

func largestRectangleArea(heights []int) int {
	ret := 0
	heights = append(heights, 0)
	stack := make([]Rect, 0)
	for _, height := range heights {
		sumWidth := 0
		// 如果栈顶之前高度大于当前高度，单调性破坏，确定栈顶扩展范围，需要更新
		for len(stack) > 0 && stack[len(stack)-1].height >= height {
			sumWidth += stack[len(stack)-1].width
			ret = max(ret, stack[len(stack)-1].height*sumWidth)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Rect{
			width:  sumWidth + 1,
			height: height,
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
