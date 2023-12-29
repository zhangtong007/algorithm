package leetcode20

var initMap = map[uint8]uint8{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	// 初始化栈
	stack := make([]uint8, len(s))
	// pop 表示栈顶索引
	pop := -1
	for i := range s {
		// 当前元素是有效左括号，入栈
		if isValidValue(s[i]) {
			pop++
			stack[pop] = s[i]
			continue
		}
		// 如果栈顶为空，或者栈顶与当前元素不匹配，return false
		if pop < 0 || !match(stack[pop], s[i]) {
			return false
		}
		// 如果栈顶元素与当前元素匹配，栈顶出栈
		pop--
	}
	// 元素遍历完毕，判断栈是否空 非空则无效，空则有效
	return pop == -1
}

func match(pop, v uint8) bool {
	return initMap[pop] == v
}

func isValidValue(v uint8) bool {
	_, ok := initMap[v]
	return ok
}
